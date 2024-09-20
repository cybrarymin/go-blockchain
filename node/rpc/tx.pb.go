// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: tx.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TxSignReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     string `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To       string `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	Value    uint64 `protobuf:"varint,3,opt,name=Value,proto3" json:"Value,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *TxSignReq) Reset() {
	*x = TxSignReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxSignReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxSignReq) ProtoMessage() {}

func (x *TxSignReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxSignReq.ProtoReflect.Descriptor instead.
func (*TxSignReq) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{0}
}

func (x *TxSignReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TxSignReq) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TxSignReq) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TxSignReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type TxSignRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx []byte `protobuf:"bytes,1,opt,name=Tx,proto3" json:"Tx,omitempty"`
}

func (x *TxSignRes) Reset() {
	*x = TxSignRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxSignRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxSignRes) ProtoMessage() {}

func (x *TxSignRes) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxSignRes.ProtoReflect.Descriptor instead.
func (*TxSignRes) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{1}
}

func (x *TxSignRes) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

type TxSendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx []byte `protobuf:"bytes,1,opt,name=Tx,proto3" json:"Tx,omitempty"`
}

func (x *TxSendReq) Reset() {
	*x = TxSendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxSendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxSendReq) ProtoMessage() {}

func (x *TxSendReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxSendReq.ProtoReflect.Descriptor instead.
func (*TxSendReq) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{2}
}

func (x *TxSendReq) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

type TxSendRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=TxHash,proto3" json:"TxHash,omitempty"`
}

func (x *TxSendRes) Reset() {
	*x = TxSendRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxSendRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxSendRes) ProtoMessage() {}

func (x *TxSendRes) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxSendRes.ProtoReflect.Descriptor instead.
func (*TxSendRes) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{3}
}

func (x *TxSendRes) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type TxReceiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx []byte `protobuf:"bytes,1,opt,name=Tx,proto3" json:"Tx,omitempty"`
}

func (x *TxReceiveReq) Reset() {
	*x = TxReceiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxReceiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxReceiveReq) ProtoMessage() {}

func (x *TxReceiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxReceiveReq.ProtoReflect.Descriptor instead.
func (*TxReceiveReq) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{4}
}

func (x *TxReceiveReq) GetTx() []byte {
	if x != nil {
		return x.Tx
	}
	return nil
}

type TxReceiveRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TxReceiveRes) Reset() {
	*x = TxReceiveRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxReceiveRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxReceiveRes) ProtoMessage() {}

func (x *TxReceiveRes) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxReceiveRes.ProtoReflect.Descriptor instead.
func (*TxReceiveRes) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{5}
}

var File_tx_proto protoreflect.FileDescriptor

var file_tx_proto_rawDesc = []byte{
	0x0a, 0x08, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x09, 0x54, 0x78,
	0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1b, 0x0a,
	0x09, 0x54, 0x78, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x78, 0x22, 0x1b, 0x0a, 0x09, 0x54, 0x78,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x78, 0x22, 0x23, 0x0a, 0x09, 0x54, 0x78, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x1e, 0x0a, 0x0c,
	0x54, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x54, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x54, 0x78, 0x22, 0x0e, 0x0a, 0x0c,
	0x54, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x32, 0x75, 0x0a, 0x02,
	0x54, 0x78, 0x12, 0x20, 0x0a, 0x06, 0x54, 0x78, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x0a, 0x2e, 0x54,
	0x78, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x54, 0x78, 0x53, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x06, 0x54, 0x78, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0a,
	0x2e, 0x54, 0x78, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x54, 0x78, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x54, 0x78, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x12, 0x0d, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x28, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tx_proto_rawDescOnce sync.Once
	file_tx_proto_rawDescData = file_tx_proto_rawDesc
)

func file_tx_proto_rawDescGZIP() []byte {
	file_tx_proto_rawDescOnce.Do(func() {
		file_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_tx_proto_rawDescData)
	})
	return file_tx_proto_rawDescData
}

var file_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tx_proto_goTypes = []any{
	(*TxSignReq)(nil),    // 0: TxSignReq
	(*TxSignRes)(nil),    // 1: TxSignRes
	(*TxSendReq)(nil),    // 2: TxSendReq
	(*TxSendRes)(nil),    // 3: TxSendRes
	(*TxReceiveReq)(nil), // 4: TxReceiveReq
	(*TxReceiveRes)(nil), // 5: TxReceiveRes
}
var file_tx_proto_depIdxs = []int32{
	0, // 0: Tx.TxSign:input_type -> TxSignReq
	2, // 1: Tx.TxSend:input_type -> TxSendReq
	4, // 2: Tx.TxReceive:input_type -> TxReceiveReq
	1, // 3: Tx.TxSign:output_type -> TxSignRes
	3, // 4: Tx.TxSend:output_type -> TxSendRes
	5, // 5: Tx.TxReceive:output_type -> TxReceiveRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tx_proto_init() }
func file_tx_proto_init() {
	if File_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tx_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TxSignReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tx_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TxSignRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tx_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TxSendReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tx_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TxSendRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tx_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TxReceiveReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tx_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TxReceiveRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tx_proto_goTypes,
		DependencyIndexes: file_tx_proto_depIdxs,
		MessageInfos:      file_tx_proto_msgTypes,
	}.Build()
	File_tx_proto = out.File
	file_tx_proto_rawDesc = nil
	file_tx_proto_goTypes = nil
	file_tx_proto_depIdxs = nil
}
