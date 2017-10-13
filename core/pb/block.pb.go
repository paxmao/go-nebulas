// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	Account
	Transaction
	BlockHeader
	Block
	Blocks
*/
package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Balance []byte `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce   uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type Transaction struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From      []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Nonce     uint64 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Alg       uint32 `protobuf:"varint,9,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign      []byte `protobuf:"bytes,10,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockHeader struct {
	Hash       []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash []byte `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Nonce      uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Coinbase   []byte `protobuf:"bytes,4,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Timestamp  int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChainId    uint32 `protobuf:"varint,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	StateRoot  []byte `protobuf:"bytes,7,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	TxsRoot    []byte `protobuf:"bytes,8,opt,name=txs_root,json=txsRoot,proto3" json:"txs_root,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{2} }

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeader) GetTxsRoot() []byte {
	if m != nil {
		return m.TxsRoot
	}
	return nil
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{3} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Blocks struct {
	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *Blocks) Reset()                    { *m = Blocks{} }
func (m *Blocks) String() string            { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()               {}
func (*Blocks) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{4} }

func (m *Blocks) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "corepb.Account")
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
	proto.RegisterType((*BlockHeader)(nil), "corepb.BlockHeader")
	proto.RegisterType((*Block)(nil), "corepb.Block")
	proto.RegisterType((*Blocks)(nil), "corepb.Blocks")
}

func init() { proto.RegisterFile("block.proto", fileDescriptorBlock) }

var fileDescriptorBlock = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0xab, 0xd4, 0x30,
	0x10, 0xc7, 0xc9, 0x76, 0xb7, 0xdb, 0x9d, 0xee, 0x13, 0x89, 0x1e, 0xa2, 0x28, 0x96, 0x82, 0x50,
	0x10, 0x56, 0xd0, 0x83, 0x78, 0xd4, 0xd3, 0xf3, 0x1a, 0xbc, 0x2f, 0x69, 0x1a, 0x5f, 0x8b, 0x6d,
	0x52, 0x9a, 0x79, 0xf2, 0x3e, 0xb0, 0xe0, 0xd7, 0x90, 0x4c, 0xfa, 0x76, 0x5b, 0xf1, 0x36, 0xf3,
	0x9f, 0xcc, 0x0c, 0xbf, 0xc9, 0x1f, 0xf2, 0xba, 0x77, 0xfa, 0xe7, 0x69, 0x9c, 0x1c, 0x3a, 0x9e,
	0x6a, 0x37, 0x99, 0xb1, 0x2e, 0x3f, 0xc3, 0xfe, 0x8b, 0xd6, 0xee, 0xde, 0x22, 0x17, 0xb0, 0xaf,
	0x55, 0xaf, 0xac, 0x36, 0x82, 0x15, 0xac, 0x3a, 0xca, 0xc7, 0x94, 0x3f, 0x87, 0x9d, 0x75, 0x41,
	0xdf, 0x14, 0xac, 0xda, 0xca, 0x98, 0x94, 0xbf, 0x19, 0xe4, 0xdf, 0x27, 0x65, 0xbd, 0xd2, 0xd8,
	0x39, 0xcb, 0x39, 0x6c, 0x5b, 0xe5, 0xdb, 0xb9, 0x99, 0xe2, 0xa0, 0xfd, 0x98, 0xdc, 0x40, 0x8d,
	0x47, 0x49, 0x31, 0x7f, 0x02, 0x1b, 0x74, 0x22, 0x21, 0x65, 0x83, 0x2e, 0x4c, 0xff, 0xa5, 0xfa,
	0x7b, 0x23, 0xb6, 0x24, 0xc5, 0xe4, 0xba, 0x73, 0xb7, 0xd8, 0xc9, 0x5f, 0xc1, 0x01, 0xbb, 0xc1,
	0x78, 0x54, 0xc3, 0x28, 0xd2, 0x82, 0x55, 0x89, 0xbc, 0x0a, 0x61, 0x5b, 0xa3, 0x50, 0x89, 0x7d,
	0xdc, 0x16, 0x62, 0xfe, 0x02, 0x32, 0xdd, 0xaa, 0xce, 0x9e, 0xbb, 0x46, 0x64, 0x05, 0xab, 0x6e,
	0xe4, 0x9e, 0xf2, 0x6f, 0x0d, 0x7f, 0x0a, 0x89, 0xea, 0xef, 0xc4, 0x81, 0xd4, 0x10, 0x86, 0x01,
	0xbe, 0xbb, 0xb3, 0x02, 0xe2, 0x80, 0x10, 0x97, 0x7f, 0x18, 0xe4, 0x5f, 0xc3, 0xe5, 0x6e, 0x8d,
	0x6a, 0xcc, 0xf4, 0x5f, 0xcc, 0x37, 0x90, 0x8f, 0x6a, 0x32, 0x16, 0xcf, 0x54, 0x8a, 0xb4, 0x10,
	0xa5, 0xdb, 0xf0, 0xe0, 0x42, 0x93, 0x2c, 0x69, 0x5e, 0x42, 0xa6, 0x5d, 0x67, 0x6b, 0xe5, 0x1f,
	0xe1, 0x2f, 0xf9, 0x9a, 0x74, 0xf7, 0x2f, 0xe9, 0x92, 0x2a, 0x5d, 0x53, 0xbd, 0x06, 0xf0, 0xa8,
	0xd0, 0x9c, 0x27, 0xe7, 0x70, 0x3e, 0xc5, 0x81, 0x14, 0xe9, 0x1c, 0x86, 0x4e, 0x7c, 0xf0, 0xb1,
	0x98, 0xc5, 0x6f, 0xc6, 0x07, 0x1f, 0x4a, 0xe5, 0x00, 0x3b, 0x02, 0xe5, 0xef, 0x20, 0x6d, 0x09,
	0x96, 0x20, 0xf3, 0x0f, 0xcf, 0x4e, 0xd1, 0x2d, 0xa7, 0xc5, 0x1d, 0xe4, 0xfc, 0x84, 0x7f, 0x82,
	0x23, 0x5e, 0x5d, 0xe0, 0xc5, 0xa6, 0x48, 0x96, 0x2d, 0x0b, 0x87, 0xc8, 0xd5, 0xc3, 0xf2, 0x3d,
	0xa4, 0x34, 0xcf, 0xf3, 0xb7, 0x90, 0x92, 0x37, 0xbd, 0x60, 0xd4, 0x7c, 0xb3, 0xda, 0x27, 0xe7,
	0x62, 0x9d, 0x92, 0x75, 0x3f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xe8, 0xbb, 0x09, 0xc9,
	0x02, 0x00, 0x00,
}