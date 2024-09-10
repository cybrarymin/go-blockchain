package rtx

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/volodymyrprokopyuk/go-blockchain/blockchain/account"
	"github.com/volodymyrprokopyuk/go-blockchain/blockchain/chain"
	"github.com/volodymyrprokopyuk/go-blockchain/blockchain/state"
)

type TxSrv struct {
  UnimplementedTxServer
  keyStoreDir string
  state *state.State
}

func NewTxSrv(keyStoreDir string, sta *state.State) *TxSrv {
  return &TxSrv{keyStoreDir: keyStoreDir, state: sta}
}

func (s *TxSrv) TxSign(_ context.Context, req *TxSignReq,) (*TxSignRes, error) {
  path := filepath.Join(s.keyStoreDir, req.From)
  acc, err := account.Read(path, []byte(req.Password))
  if err != nil {
    return nil, err
  }
  tx := chain.Tx{
    From: chain.Address(req.From),
    To: chain.Address(req.To),
    Value: uint(req.Value),
    Nonce: s.state.Pending.Nonce(chain.Address(req.From)) + 1,
    Time: time.Now(),
  }
  stx, err := acc.Sign(tx)
  if err != nil {
    return nil, err
  }
  jsnSTx, err := json.Marshal(stx)
  if err != nil {
    return nil, err
  }
  res := &TxSignRes{SigTx: jsnSTx}
  return res, nil
}

func (s *TxSrv) TxSend(_ context.Context, req *TxSendReq) (*TxSendRes, error) {
  var stx chain.SigTx
  err := json.Unmarshal(req.SigTx, &stx)
  if err != nil {
    return nil, err
  }
  err = s.state.ApplyTx(stx)
  if err != nil {
    return nil, err
  }
  fmt.Printf("* Pending state (ApplyTx)\n%v\n", s.state)
  hash, err := stx.Hash()
  if err != nil {
    return nil, err
  }
  res := &TxSendRes{Hash: hash[:]}
  return res, nil
}
