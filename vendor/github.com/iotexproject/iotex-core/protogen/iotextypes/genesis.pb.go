// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/genesis.proto

package iotextypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Genesis struct {
	Blockchain           *GenesisBlockchain `protobuf:"bytes,1,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	Account              *GenesisAccount    `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Poll                 *GenesisPoll       `protobuf:"bytes,3,opt,name=poll,proto3" json:"poll,omitempty"`
	Rewarding            *GenesisRewarding  `protobuf:"bytes,4,opt,name=rewarding,proto3" json:"rewarding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{0}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetBlockchain() *GenesisBlockchain {
	if m != nil {
		return m.Blockchain
	}
	return nil
}

func (m *Genesis) GetAccount() *GenesisAccount {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *Genesis) GetPoll() *GenesisPoll {
	if m != nil {
		return m.Poll
	}
	return nil
}

func (m *Genesis) GetRewarding() *GenesisRewarding {
	if m != nil {
		return m.Rewarding
	}
	return nil
}

type GenesisBlockchain struct {
	Timestamp             int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BlockGasLimit         uint64   `protobuf:"varint,2,opt,name=blockGasLimit,proto3" json:"blockGasLimit,omitempty"`
	ActionGasLimit        uint64   `protobuf:"varint,3,opt,name=actionGasLimit,proto3" json:"actionGasLimit,omitempty"`
	BlockInterval         int64    `protobuf:"varint,4,opt,name=blockInterval,proto3" json:"blockInterval,omitempty"`
	NumSubEpochs          uint64   `protobuf:"varint,5,opt,name=numSubEpochs,proto3" json:"numSubEpochs,omitempty"`
	NumDelegates          uint64   `protobuf:"varint,6,opt,name=numDelegates,proto3" json:"numDelegates,omitempty"`
	NumCandidateDelegates uint64   `protobuf:"varint,7,opt,name=numCandidateDelegates,proto3" json:"numCandidateDelegates,omitempty"`
	TimeBasedRotation     bool     `protobuf:"varint,8,opt,name=timeBasedRotation,proto3" json:"timeBasedRotation,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *GenesisBlockchain) Reset()         { *m = GenesisBlockchain{} }
func (m *GenesisBlockchain) String() string { return proto.CompactTextString(m) }
func (*GenesisBlockchain) ProtoMessage()    {}
func (*GenesisBlockchain) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{1}
}

func (m *GenesisBlockchain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisBlockchain.Unmarshal(m, b)
}
func (m *GenesisBlockchain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisBlockchain.Marshal(b, m, deterministic)
}
func (m *GenesisBlockchain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisBlockchain.Merge(m, src)
}
func (m *GenesisBlockchain) XXX_Size() int {
	return xxx_messageInfo_GenesisBlockchain.Size(m)
}
func (m *GenesisBlockchain) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisBlockchain.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisBlockchain proto.InternalMessageInfo

func (m *GenesisBlockchain) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *GenesisBlockchain) GetBlockGasLimit() uint64 {
	if m != nil {
		return m.BlockGasLimit
	}
	return 0
}

func (m *GenesisBlockchain) GetActionGasLimit() uint64 {
	if m != nil {
		return m.ActionGasLimit
	}
	return 0
}

func (m *GenesisBlockchain) GetBlockInterval() int64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

func (m *GenesisBlockchain) GetNumSubEpochs() uint64 {
	if m != nil {
		return m.NumSubEpochs
	}
	return 0
}

func (m *GenesisBlockchain) GetNumDelegates() uint64 {
	if m != nil {
		return m.NumDelegates
	}
	return 0
}

func (m *GenesisBlockchain) GetNumCandidateDelegates() uint64 {
	if m != nil {
		return m.NumCandidateDelegates
	}
	return 0
}

func (m *GenesisBlockchain) GetTimeBasedRotation() bool {
	if m != nil {
		return m.TimeBasedRotation
	}
	return false
}

type GenesisAccount struct {
	InitBalanceAddrs     []string `protobuf:"bytes,1,rep,name=initBalanceAddrs,proto3" json:"initBalanceAddrs,omitempty"`
	InitBalances         []string `protobuf:"bytes,2,rep,name=initBalances,proto3" json:"initBalances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenesisAccount) Reset()         { *m = GenesisAccount{} }
func (m *GenesisAccount) String() string { return proto.CompactTextString(m) }
func (*GenesisAccount) ProtoMessage()    {}
func (*GenesisAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{2}
}

func (m *GenesisAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisAccount.Unmarshal(m, b)
}
func (m *GenesisAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisAccount.Marshal(b, m, deterministic)
}
func (m *GenesisAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccount.Merge(m, src)
}
func (m *GenesisAccount) XXX_Size() int {
	return xxx_messageInfo_GenesisAccount.Size(m)
}
func (m *GenesisAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccount proto.InternalMessageInfo

func (m *GenesisAccount) GetInitBalanceAddrs() []string {
	if m != nil {
		return m.InitBalanceAddrs
	}
	return nil
}

func (m *GenesisAccount) GetInitBalances() []string {
	if m != nil {
		return m.InitBalances
	}
	return nil
}

type GenesisPoll struct {
	EnableGravityChainVoting bool               `protobuf:"varint,1,opt,name=enableGravityChainVoting,proto3" json:"enableGravityChainVoting,omitempty"`
	GravityChainStartHeight  uint64             `protobuf:"varint,2,opt,name=gravityChainStartHeight,proto3" json:"gravityChainStartHeight,omitempty"`
	RegisterContractAddress  string             `protobuf:"bytes,3,opt,name=registerContractAddress,proto3" json:"registerContractAddress,omitempty"`
	StakingContractAddress   string             `protobuf:"bytes,4,opt,name=stakingContractAddress,proto3" json:"stakingContractAddress,omitempty"`
	VoteThreshold            string             `protobuf:"bytes,5,opt,name=voteThreshold,proto3" json:"voteThreshold,omitempty"`
	ScoreThreshold           string             `protobuf:"bytes,6,opt,name=scoreThreshold,proto3" json:"scoreThreshold,omitempty"`
	SelfStakingThreshold     string             `protobuf:"bytes,7,opt,name=selfStakingThreshold,proto3" json:"selfStakingThreshold,omitempty"`
	Delegates                []*GenesisDelegate `protobuf:"bytes,8,rep,name=delegates,proto3" json:"delegates,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}           `json:"-"`
	XXX_unrecognized         []byte             `json:"-"`
	XXX_sizecache            int32              `json:"-"`
}

func (m *GenesisPoll) Reset()         { *m = GenesisPoll{} }
func (m *GenesisPoll) String() string { return proto.CompactTextString(m) }
func (*GenesisPoll) ProtoMessage()    {}
func (*GenesisPoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{3}
}

func (m *GenesisPoll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisPoll.Unmarshal(m, b)
}
func (m *GenesisPoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisPoll.Marshal(b, m, deterministic)
}
func (m *GenesisPoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisPoll.Merge(m, src)
}
func (m *GenesisPoll) XXX_Size() int {
	return xxx_messageInfo_GenesisPoll.Size(m)
}
func (m *GenesisPoll) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisPoll.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisPoll proto.InternalMessageInfo

func (m *GenesisPoll) GetEnableGravityChainVoting() bool {
	if m != nil {
		return m.EnableGravityChainVoting
	}
	return false
}

func (m *GenesisPoll) GetGravityChainStartHeight() uint64 {
	if m != nil {
		return m.GravityChainStartHeight
	}
	return 0
}

func (m *GenesisPoll) GetRegisterContractAddress() string {
	if m != nil {
		return m.RegisterContractAddress
	}
	return ""
}

func (m *GenesisPoll) GetStakingContractAddress() string {
	if m != nil {
		return m.StakingContractAddress
	}
	return ""
}

func (m *GenesisPoll) GetVoteThreshold() string {
	if m != nil {
		return m.VoteThreshold
	}
	return ""
}

func (m *GenesisPoll) GetScoreThreshold() string {
	if m != nil {
		return m.ScoreThreshold
	}
	return ""
}

func (m *GenesisPoll) GetSelfStakingThreshold() string {
	if m != nil {
		return m.SelfStakingThreshold
	}
	return ""
}

func (m *GenesisPoll) GetDelegates() []*GenesisDelegate {
	if m != nil {
		return m.Delegates
	}
	return nil
}

type GenesisDelegate struct {
	OperatorAddr         string   `protobuf:"bytes,1,opt,name=operatorAddr,proto3" json:"operatorAddr,omitempty"`
	RewardAddr           string   `protobuf:"bytes,2,opt,name=rewardAddr,proto3" json:"rewardAddr,omitempty"`
	Votes                string   `protobuf:"bytes,3,opt,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenesisDelegate) Reset()         { *m = GenesisDelegate{} }
func (m *GenesisDelegate) String() string { return proto.CompactTextString(m) }
func (*GenesisDelegate) ProtoMessage()    {}
func (*GenesisDelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{4}
}

func (m *GenesisDelegate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisDelegate.Unmarshal(m, b)
}
func (m *GenesisDelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisDelegate.Marshal(b, m, deterministic)
}
func (m *GenesisDelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDelegate.Merge(m, src)
}
func (m *GenesisDelegate) XXX_Size() int {
	return xxx_messageInfo_GenesisDelegate.Size(m)
}
func (m *GenesisDelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDelegate.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDelegate proto.InternalMessageInfo

func (m *GenesisDelegate) GetOperatorAddr() string {
	if m != nil {
		return m.OperatorAddr
	}
	return ""
}

func (m *GenesisDelegate) GetRewardAddr() string {
	if m != nil {
		return m.RewardAddr
	}
	return ""
}

func (m *GenesisDelegate) GetVotes() string {
	if m != nil {
		return m.Votes
	}
	return ""
}

type GenesisRewarding struct {
	InitAdminAddr                  string   `protobuf:"bytes,1,opt,name=initAdminAddr,proto3" json:"initAdminAddr,omitempty"`
	InitBalance                    string   `protobuf:"bytes,2,opt,name=initBalance,proto3" json:"initBalance,omitempty"`
	BlockReward                    string   `protobuf:"bytes,3,opt,name=blockReward,proto3" json:"blockReward,omitempty"`
	EpochReward                    string   `protobuf:"bytes,4,opt,name=epochReward,proto3" json:"epochReward,omitempty"`
	NumDelegatesForEpochReward     uint64   `protobuf:"varint,5,opt,name=numDelegatesForEpochReward,proto3" json:"numDelegatesForEpochReward,omitempty"`
	FoundationBonus                string   `protobuf:"bytes,6,opt,name=foundationBonus,proto3" json:"foundationBonus,omitempty"`
	NumDelegatesForFoundationBonus uint64   `protobuf:"varint,7,opt,name=numDelegatesForFoundationBonus,proto3" json:"numDelegatesForFoundationBonus,omitempty"`
	FoundationBonusLastEpoch       uint64   `protobuf:"varint,8,opt,name=foundationBonusLastEpoch,proto3" json:"foundationBonusLastEpoch,omitempty"`
	ProductivityThreshold          uint64   `protobuf:"varint,9,opt,name=productivityThreshold,proto3" json:"productivityThreshold,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *GenesisRewarding) Reset()         { *m = GenesisRewarding{} }
func (m *GenesisRewarding) String() string { return proto.CompactTextString(m) }
func (*GenesisRewarding) ProtoMessage()    {}
func (*GenesisRewarding) Descriptor() ([]byte, []int) {
	return fileDescriptor_8090b9f9a91af920, []int{5}
}

func (m *GenesisRewarding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisRewarding.Unmarshal(m, b)
}
func (m *GenesisRewarding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisRewarding.Marshal(b, m, deterministic)
}
func (m *GenesisRewarding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisRewarding.Merge(m, src)
}
func (m *GenesisRewarding) XXX_Size() int {
	return xxx_messageInfo_GenesisRewarding.Size(m)
}
func (m *GenesisRewarding) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisRewarding.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisRewarding proto.InternalMessageInfo

func (m *GenesisRewarding) GetInitAdminAddr() string {
	if m != nil {
		return m.InitAdminAddr
	}
	return ""
}

func (m *GenesisRewarding) GetInitBalance() string {
	if m != nil {
		return m.InitBalance
	}
	return ""
}

func (m *GenesisRewarding) GetBlockReward() string {
	if m != nil {
		return m.BlockReward
	}
	return ""
}

func (m *GenesisRewarding) GetEpochReward() string {
	if m != nil {
		return m.EpochReward
	}
	return ""
}

func (m *GenesisRewarding) GetNumDelegatesForEpochReward() uint64 {
	if m != nil {
		return m.NumDelegatesForEpochReward
	}
	return 0
}

func (m *GenesisRewarding) GetFoundationBonus() string {
	if m != nil {
		return m.FoundationBonus
	}
	return ""
}

func (m *GenesisRewarding) GetNumDelegatesForFoundationBonus() uint64 {
	if m != nil {
		return m.NumDelegatesForFoundationBonus
	}
	return 0
}

func (m *GenesisRewarding) GetFoundationBonusLastEpoch() uint64 {
	if m != nil {
		return m.FoundationBonusLastEpoch
	}
	return 0
}

func (m *GenesisRewarding) GetProductivityThreshold() uint64 {
	if m != nil {
		return m.ProductivityThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*Genesis)(nil), "iotextypes.Genesis")
	proto.RegisterType((*GenesisBlockchain)(nil), "iotextypes.GenesisBlockchain")
	proto.RegisterType((*GenesisAccount)(nil), "iotextypes.GenesisAccount")
	proto.RegisterType((*GenesisPoll)(nil), "iotextypes.GenesisPoll")
	proto.RegisterType((*GenesisDelegate)(nil), "iotextypes.GenesisDelegate")
	proto.RegisterType((*GenesisRewarding)(nil), "iotextypes.GenesisRewarding")
}

func init() { proto.RegisterFile("proto/types/genesis.proto", fileDescriptor_8090b9f9a91af920) }

var fileDescriptor_8090b9f9a91af920 = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xed, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0xd5, 0xa6, 0x5b, 0x97, 0xd3, 0xff, 0x7f, 0x2f, 0xd6, 0x60, 0x61, 0x8c, 0xa9, 0xaa,
	0x10, 0xaa, 0x78, 0x69, 0xa5, 0x31, 0x4d, 0x63, 0x12, 0x48, 0xeb, 0xd8, 0x06, 0xd2, 0x3e, 0x20,
	0x0f, 0xf1, 0x81, 0x4f, 0xb8, 0x89, 0x97, 0x9a, 0x25, 0x76, 0x64, 0x3b, 0x83, 0xdd, 0x16, 0x97,
	0xc0, 0x1d, 0x70, 0x15, 0xdc, 0x06, 0xb2, 0xd3, 0x36, 0x2f, 0x4b, 0xe1, 0x63, 0x9f, 0xf3, 0x7b,
	0xec, 0x9c, 0x9c, 0xe7, 0xa4, 0xf0, 0x20, 0x91, 0x42, 0x8b, 0xa1, 0xbe, 0x4d, 0xa8, 0x1a, 0x86,
	0x94, 0x53, 0xc5, 0xd4, 0xc0, 0x6a, 0x08, 0x98, 0xd0, 0xf4, 0xbb, 0xad, 0xf4, 0x7e, 0x37, 0xa0,
	0x7d, 0x9e, 0x55, 0xd1, 0x6b, 0x80, 0x71, 0x24, 0xfc, 0x6b, 0x7f, 0x42, 0x18, 0xf7, 0x1a, 0xdd,
	0x46, 0xbf, 0xb3, 0xf7, 0x68, 0x90, 0xc3, 0x83, 0x29, 0x38, 0x9a, 0x43, 0xb8, 0x60, 0x40, 0xfb,
	0xd0, 0x26, 0xbe, 0x2f, 0x52, 0xae, 0xbd, 0xa6, 0xf5, 0x6e, 0xd7, 0x78, 0x8f, 0x33, 0x02, 0xcf,
	0x50, 0xf4, 0x0c, 0x5a, 0x89, 0x88, 0x22, 0xcf, 0xb1, 0x96, 0xad, 0x1a, 0xcb, 0x07, 0x11, 0x45,
	0xd8, 0x42, 0xe8, 0x08, 0x5c, 0x49, 0xbf, 0x11, 0x19, 0x30, 0x1e, 0x7a, 0x2d, 0xeb, 0xd8, 0xa9,
	0x71, 0xe0, 0x19, 0x83, 0x73, 0xbc, 0xf7, 0xab, 0x09, 0x1b, 0x77, 0x1a, 0x40, 0x3b, 0xe0, 0x6a,
	0x16, 0x53, 0xa5, 0x49, 0x9c, 0xd8, 0x96, 0x1d, 0x9c, 0x0b, 0xe8, 0x31, 0xfc, 0x6f, 0x1b, 0x3c,
	0x27, 0xea, 0x82, 0xc5, 0x2c, 0x6b, 0xac, 0x85, 0xcb, 0x22, 0x7a, 0x02, 0xab, 0xc4, 0xd7, 0x4c,
	0xf0, 0x39, 0xe6, 0x58, 0xac, 0xa2, 0xce, 0x4f, 0x7b, 0xcf, 0x35, 0x95, 0x37, 0x24, 0xb2, 0x1d,
	0x38, 0xb8, 0x2c, 0xa2, 0x1e, 0xfc, 0xc7, 0xd3, 0xf8, 0x32, 0x1d, 0x9f, 0x26, 0xc2, 0x9f, 0x28,
	0x6f, 0xc9, 0x9e, 0x55, 0xd2, 0xa6, 0xcc, 0x5b, 0x1a, 0xd1, 0x90, 0x68, 0xaa, 0xbc, 0xe5, 0x39,
	0x33, 0xd7, 0xd0, 0x3e, 0xdc, 0xe3, 0x69, 0x7c, 0x42, 0x78, 0xc0, 0x02, 0xa2, 0x69, 0x0e, 0xb7,
	0x2d, 0x5c, 0x5f, 0x44, 0xcf, 0x61, 0xc3, 0xb4, 0x3f, 0x22, 0x8a, 0x06, 0x58, 0x68, 0x62, 0x1a,
	0xf0, 0x56, 0xba, 0x8d, 0xfe, 0x0a, 0xbe, 0x5b, 0xe8, 0x7d, 0x81, 0xd5, 0xf2, 0x5c, 0xd1, 0x53,
	0x58, 0x67, 0x9c, 0xe9, 0x11, 0x89, 0x08, 0xf7, 0xe9, 0x71, 0x10, 0x48, 0xe5, 0x35, 0xba, 0x4e,
	0xdf, 0xc5, 0x77, 0x74, 0xd3, 0x45, 0x41, 0x53, 0x5e, 0xd3, 0x72, 0x25, 0xad, 0xf7, 0xc3, 0x81,
	0x4e, 0x21, 0x07, 0xe8, 0x08, 0x3c, 0xca, 0xc9, 0x38, 0xa2, 0xe7, 0x92, 0xdc, 0x30, 0x7d, 0x7b,
	0x62, 0xa6, 0xf8, 0x49, 0x68, 0x13, 0x88, 0x86, 0x7d, 0xcc, 0x85, 0x75, 0x74, 0x08, 0x5b, 0x61,
	0x41, 0xbd, 0xd4, 0x44, 0xea, 0x77, 0x94, 0x85, 0x93, 0xd9, 0x5c, 0x17, 0x95, 0x8d, 0x53, 0xd2,
	0x90, 0x29, 0x4d, 0xe5, 0x89, 0xe0, 0x5a, 0x12, 0x5f, 0x9b, 0x16, 0xa8, 0x52, 0x76, 0xd4, 0x2e,
	0x5e, 0x54, 0x46, 0x07, 0x70, 0x5f, 0x69, 0x72, 0xcd, 0x78, 0x58, 0x35, 0xb6, 0xac, 0x71, 0x41,
	0xd5, 0x64, 0xe5, 0x46, 0x68, 0xfa, 0x71, 0x22, 0xa9, 0x9a, 0x88, 0x28, 0xb0, 0x31, 0x70, 0x71,
	0x59, 0x34, 0xc9, 0x53, 0xbe, 0x90, 0x05, 0x6c, 0xd9, 0x62, 0x15, 0x15, 0xed, 0xc1, 0xa6, 0xa2,
	0xd1, 0xd5, 0x65, 0x76, 0x57, 0x4e, 0xb7, 0x2d, 0x5d, 0x5b, 0x43, 0xaf, 0xc0, 0x0d, 0xe6, 0x99,
	0x59, 0xe9, 0x3a, 0xfd, 0xce, 0xde, 0xc3, 0x9a, 0x5d, 0x9b, 0x45, 0x07, 0xe7, 0x74, 0xef, 0x1a,
	0xd6, 0x2a, 0x55, 0x33, 0x6b, 0x91, 0x50, 0x49, 0xb4, 0x90, 0xa6, 0x45, 0x3b, 0x2b, 0x17, 0x97,
	0x34, 0xb4, 0x0b, 0x90, 0xad, 0xab, 0x25, 0x9a, 0x96, 0x28, 0x28, 0x68, 0x13, 0x96, 0x4c, 0xfb,
	0xb3, 0x77, 0x9e, 0xfd, 0xe8, 0xfd, 0x74, 0x60, 0xbd, 0xba, 0xf7, 0xe6, 0xf5, 0x99, 0x18, 0x1d,
	0x07, 0x31, 0xe3, 0x85, 0xfb, 0xca, 0x22, 0xea, 0x42, 0xa7, 0x10, 0xb6, 0xe9, 0x8d, 0x45, 0xc9,
	0x10, 0x76, 0x3b, 0xb3, 0x93, 0xa7, 0x17, 0x17, 0x25, 0x43, 0x50, 0xb3, 0x94, 0x53, 0x22, 0x9b,
	0x6a, 0x51, 0x42, 0x6f, 0x60, 0xbb, 0xb8, 0x98, 0x67, 0x42, 0x9e, 0x16, 0x0c, 0xd9, 0x7a, 0xff,
	0x85, 0x40, 0x7d, 0x58, 0xbb, 0x12, 0x29, 0x0f, 0xec, 0xca, 0x8d, 0x04, 0x4f, 0xd5, 0x74, 0xca,
	0x55, 0x19, 0x9d, 0xc1, 0x6e, 0xe5, 0x9c, 0xb3, 0x8a, 0x31, 0xdb, 0xfd, 0x7f, 0x50, 0x66, 0xc9,
	0x2a, 0x47, 0x5f, 0x10, 0xa5, 0xed, 0x33, 0xd9, 0x6f, 0x41, 0x0b, 0x2f, 0xac, 0x9b, 0xcf, 0x4e,
	0x22, 0x45, 0x90, 0xfa, 0x9a, 0x99, 0x55, 0xca, 0xb3, 0xe6, 0x66, 0x9f, 0x9d, 0xda, 0xe2, 0xe8,
	0xf0, 0xf3, 0x41, 0xc8, 0xf4, 0x24, 0x1d, 0x0f, 0x7c, 0x11, 0x0f, 0x6d, 0xca, 0x12, 0x29, 0xbe,
	0x52, 0x5f, 0x67, 0x3f, 0x5e, 0x98, 0x3c, 0x0f, 0xed, 0xdf, 0x57, 0x48, 0xf9, 0x30, 0x8f, 0xe1,
	0x78, 0xd9, 0x8a, 0x2f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x69, 0x94, 0x56, 0xf0, 0x06,
	0x00, 0x00,
}
