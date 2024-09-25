package node

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/volodymyrprokopyuk/go-blockchain/chain"
	"github.com/volodymyrprokopyuk/go-blockchain/node/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcMsgRelay[Msg any] func(
  ctx context.Context, conn *grpc.ClientConn, chRelay chan Msg,
) error

var grpcTxRelay grpcMsgRelay[chain.SigTx] = func(
  ctx context.Context, conn *grpc.ClientConn, chRelay chan chain.SigTx,
) error {
  cln := rpc.NewTxClient(conn)
  stream, err := cln.TxReceive(context.Background())
  if err != nil {
    return err
  }
  defer stream.CloseAndRecv()
  for {
    select {
    case <- ctx.Done():
      return nil
    case tx, open := <- chRelay:
      if !open {
        return nil
      }
      jtx, err := json.Marshal(tx)
      if err != nil {
        fmt.Println(err)
        continue
      }
      req := &rpc.TxReceiveReq{Tx: jtx}
      err = stream.Send(req)
      if err != nil {
        fmt.Println(err)
        continue
      }
    }
  }
}

var grpcBlockRelay grpcMsgRelay[chain.SigBlock] = func(
  ctx context.Context, conn *grpc.ClientConn, chRelay chan chain.SigBlock,
) error {
  cln := rpc.NewBlockClient(conn)
  stream, err := cln.BlockReceive(context.Background())
  if err != nil {
    return err
  }
  defer stream.CloseAndRecv()
  for {
    select {
    case <- ctx.Done():
      return nil
    case blk, open := <- chRelay:
      if !open {
        return nil
      }
      jblk, err := json.Marshal(blk)
      if err != nil {
        fmt.Println(err)
        continue
      }
      req := &rpc.BlockReceiveReq{Block: jblk}
      err = stream.Send(req)
      if err != nil {
        fmt.Println(err)
        continue
      }
    }
  }
}

type msgRelay[Msg any, Relay grpcMsgRelay[Msg]] struct {
  ctx context.Context
  wg *sync.WaitGroup
  chMsg chan Msg
  grpcRelay Relay
  selfRelay bool
  peerDisc *peerDiscovery
  wgRelays *sync.WaitGroup
  chPeerAdd, chPeerRem chan string
}

func newMsgRelay[Msg any, Relay grpcMsgRelay[Msg]](
  ctx context.Context, wg *sync.WaitGroup, cap int,
  grpcRelay Relay, selfRelay bool, peerDisc *peerDiscovery,
) *msgRelay[Msg, Relay] {
  return &msgRelay[Msg, Relay]{
    ctx: ctx, wg: wg, chMsg: make(chan Msg, cap),
    grpcRelay: grpcRelay, selfRelay: selfRelay, peerDisc: peerDisc,
    wgRelays: new(sync.WaitGroup),
    chPeerAdd: make(chan string), chPeerRem: make(chan string),
  }
}

func (r *msgRelay[Msg, Relay]) RelayTx(tx Msg) {
  r.chMsg <- tx
}

func (r *msgRelay[Msg, Relay]) RelayBlock(blk Msg) {
  r.chMsg <- blk
}

func (r *msgRelay[Msg, Relay]) addPeers(period time.Duration) {
  defer r.wgRelays.Done()
  tick := time.NewTicker(6 * time.Second) // initial early add peers
  defer tick.Stop()
  reset := false
  for {
    select {
    case <- r.ctx.Done():
      return
    case <- tick.C:
      if !reset {
        tick.Reset(period)
        reset = true
      }
      var peers []string
      if r.selfRelay {
        peers = r.peerDisc.SelfPeers()
      } else {
        peers = r.peerDisc.Peers()
      }
      for _, peer := range peers {
        r.chPeerAdd <- peer
      }
    }
  }
}

func (r *msgRelay[Msg, Relay]) peerRelay(peer string) chan Msg {
  chRelay := make(chan Msg)
  r.wgRelays.Add(1)
  go func () {
    defer r.wgRelays.Done()
    conn, err := grpc.NewClient(
      peer, grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    if err != nil {
      fmt.Println(err)
      r.chPeerRem <- peer
      return
    }
    defer conn.Close()
    err = r.grpcRelay(r.ctx, conn, chRelay)
    if err != nil {
      fmt.Println(err)
      r.chPeerRem <- peer
      return
    }
  }()
  return chRelay
}

func (r *msgRelay[Msg, Relay]) relayMsgs() {
  defer r.wg.Done()
  r.wgRelays.Add(1)
  go r.addPeers(10 * time.Second)
  chRelays := make(map[string]chan Msg)
  closeRelays := func() {
    for _, chRelay := range chRelays {
      close(chRelay)
    }
  }
  for {
    select {
    case <- r.ctx.Done():
      closeRelays()
      r.wgRelays.Wait()
      return
    case peer := <- r.chPeerAdd:
      _, exist := chRelays[peer]
      if exist {
        continue
      }
      fmt.Printf("* Relay peer: %v\n", peer)
      chRelay := r.peerRelay(peer)
      chRelays[peer] = chRelay
    case peer := <- r.chPeerRem:
      _, exist := chRelays[peer]
      if !exist {
        continue
      }
      chRelay := chRelays[peer]
      close(chRelay)
      delete(chRelays, peer)
    case msg := <- r.chMsg:
      for _, chRelay := range chRelays {
        chRelay <- msg
      }
    }
  }
}
