// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/bootstrap/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BootstrapProposal struct {
	// title specifies the title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description specifies the description of the proposal
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// proposer_address specifies the proposer of the proposal
	ProposerAddress string                                   `protobuf:"bytes,3,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
	OfferCoins      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=offer_coins,json=offerCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"offer_coins"`
	QuoteCoinDenom  string                                   `protobuf:"bytes,5,opt,name=quote_coin_denom,json=quoteCoinDenom,proto3" json:"quote_coin_denom,omitempty"`
	MinPrice        github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,6,opt,name=min_price,json=minPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_price"`
	MaxPrice        github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,7,opt,name=max_price,json=maxPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_price"`
	PairId          uint64                                   `protobuf:"varint,8,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
	PoolId          uint64                                   `protobuf:"varint,9,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	InitialOrders   []InitialOrder                           `protobuf:"bytes,10,rep,name=initial_orders,json=initialOrders,proto3" json:"initial_orders"`
	StartTime       time.Time                                `protobuf:"bytes,11,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"end_time"`
	// number of stages of the bootstrap pool
	NumOfStages uint32 `protobuf:"varint,12,opt,name=num_of_stages,json=numOfStages,proto3" json:"num_of_stages,omitempty"`
	// duration of the stage
	StageDuration time.Duration `protobuf:"bytes,13,opt,name=stage_duration,json=stageDuration,proto3,stdduration" json:"stage_duration"`
}

func (m *BootstrapProposal) Reset()      { *m = BootstrapProposal{} }
func (*BootstrapProposal) ProtoMessage() {}
func (*BootstrapProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91229c2c9a1c550, []int{0}
}
func (m *BootstrapProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BootstrapProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BootstrapProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapProposal.Merge(m, src)
}
func (m *BootstrapProposal) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapProposal.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapProposal proto.InternalMessageInfo

type InitialOrder struct {
	OfferCoin types.Coin                             `protobuf:"bytes,1,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin"`
	Price     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	// direction specifies the order direction; either buy or sell
	Direction OrderDirection `protobuf:"varint,3,opt,name=direction,proto3,enum=crescent.bootstrap.v1beta1.OrderDirection" json:"direction,omitempty"`
}

func (m *InitialOrder) Reset()         { *m = InitialOrder{} }
func (m *InitialOrder) String() string { return proto.CompactTextString(m) }
func (*InitialOrder) ProtoMessage()    {}
func (*InitialOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a91229c2c9a1c550, []int{1}
}
func (m *InitialOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InitialOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InitialOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InitialOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitialOrder.Merge(m, src)
}
func (m *InitialOrder) XXX_Size() int {
	return m.Size()
}
func (m *InitialOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_InitialOrder.DiscardUnknown(m)
}

var xxx_messageInfo_InitialOrder proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BootstrapProposal)(nil), "crescent.bootstrap.v1beta1.BootstrapProposal")
	proto.RegisterType((*InitialOrder)(nil), "crescent.bootstrap.v1beta1.InitialOrder")
}

func init() {
	proto.RegisterFile("crescent/bootstrap/v1beta1/proposal.proto", fileDescriptor_a91229c2c9a1c550)
}

var fileDescriptor_a91229c2c9a1c550 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xbf, 0x4f, 0xdc, 0x4a,
	0x10, 0xb6, 0xf9, 0x71, 0xdc, 0xed, 0x71, 0x07, 0x6f, 0x85, 0xf4, 0xcc, 0x3d, 0xc9, 0x3e, 0x5d,
	0xf1, 0x64, 0x90, 0xb0, 0x1f, 0xbc, 0x34, 0xa1, 0x88, 0x14, 0xe7, 0x8a, 0x90, 0x14, 0x20, 0x93,
	0xa4, 0x48, 0x63, 0xed, 0xd9, 0x7b, 0x97, 0x15, 0xb6, 0xd7, 0xd9, 0xdd, 0x23, 0x50, 0xa7, 0x49,
	0x49, 0x49, 0x49, 0x9d, 0xbf, 0x84, 0x92, 0x32, 0x4a, 0x01, 0x11, 0x34, 0xd4, 0xf9, 0x0b, 0xa2,
	0xdd, 0xb5, 0x8f, 0x53, 0xa2, 0xa0, 0x88, 0xea, 0x3c, 0x33, 0xdf, 0x37, 0xdf, 0xcc, 0xf8, 0xf3,
	0x81, 0xb5, 0x98, 0x61, 0x1e, 0xe3, 0x5c, 0xf8, 0x03, 0x4a, 0x05, 0x17, 0x0c, 0x15, 0xfe, 0xe1,
	0xe6, 0x00, 0x0b, 0xb4, 0xe9, 0x17, 0x8c, 0x16, 0x94, 0xa3, 0xd4, 0x2b, 0x18, 0x15, 0x14, 0x76,
	0x2a, 0xa8, 0x37, 0x81, 0x7a, 0x25, 0xb4, 0xb3, 0x32, 0xa2, 0x23, 0xaa, 0x60, 0xbe, 0x7c, 0xd2,
	0x8c, 0xce, 0x6a, 0x4c, 0x79, 0x46, 0x79, 0xa4, 0x0b, 0x3a, 0x28, 0x4b, 0xb6, 0x8e, 0xfc, 0x01,
	0xe2, 0x78, 0x22, 0x18, 0x53, 0x92, 0x97, 0xf5, 0xf5, 0x7b, 0xe6, 0xba, 0x93, 0xd7, 0x58, 0x67,
	0x44, 0xe9, 0x28, 0xc5, 0xbe, 0x8a, 0x06, 0xe3, 0xa1, 0x2f, 0x48, 0x86, 0xb9, 0x40, 0x59, 0x05,
	0xb0, 0x7f, 0x06, 0x24, 0x63, 0x86, 0x04, 0xa1, 0xa5, 0x58, 0xef, 0x63, 0x0d, 0xfc, 0x15, 0x54,
	0x4d, 0xf7, 0xca, 0xad, 0xe1, 0x0a, 0x98, 0x17, 0x44, 0xa4, 0xd8, 0x32, 0xbb, 0xa6, 0xdb, 0x08,
	0x75, 0x00, 0xbb, 0xa0, 0x99, 0x60, 0x1e, 0x33, 0x52, 0xc8, 0x06, 0xd6, 0x8c, 0xaa, 0x4d, 0xa7,
	0xe0, 0x1a, 0x58, 0xd6, 0x97, 0xc3, 0x2c, 0x42, 0x49, 0xc2, 0x30, 0xe7, 0xd6, 0xac, 0x82, 0x2d,
	0x55, 0xf9, 0xa7, 0x3a, 0x0d, 0x53, 0xd0, 0xa4, 0xc3, 0x21, 0x66, 0x91, 0xdc, 0x9c, 0x5b, 0x73,
	0xdd, 0x59, 0xb7, 0xb9, 0xb5, 0xea, 0x95, 0x97, 0x92, 0xb7, 0xa9, 0x2e, 0xec, 0x3d, 0xa3, 0x24,
	0x0f, 0xfe, 0x3b, 0xbf, 0x74, 0x8c, 0xcf, 0x57, 0x8e, 0x3b, 0x22, 0xe2, 0xdd, 0x78, 0xe0, 0xc5,
	0x34, 0x2b, 0xcf, 0x5a, 0xfe, 0x6c, 0xf0, 0xe4, 0xc0, 0x17, 0xc7, 0x05, 0xe6, 0x8a, 0xc0, 0x43,
	0xa0, 0xfa, 0xab, 0x67, 0xe8, 0x82, 0xe5, 0xf7, 0x63, 0x2a, 0xb0, 0x52, 0x8b, 0x12, 0x9c, 0xd3,
	0xcc, 0x9a, 0x57, 0x83, 0xb5, 0x55, 0x5e, 0xa2, 0xfa, 0x32, 0x0b, 0x5f, 0x82, 0x46, 0x46, 0xf2,
	0xa8, 0x60, 0x24, 0xc6, 0x56, 0x4d, 0x42, 0x02, 0x4f, 0x4a, 0x7f, 0xbd, 0x74, 0xfe, 0xfd, 0x03,
	0xe9, 0x3e, 0x8e, 0xc3, 0x7a, 0x46, 0xf2, 0x3d, 0xc9, 0x57, 0xcd, 0xd0, 0x51, 0xd9, 0x6c, 0xe1,
	0x81, 0xcd, 0xd0, 0x91, 0x6e, 0xf6, 0x37, 0x58, 0x28, 0x10, 0x61, 0x11, 0x49, 0xac, 0x7a, 0xd7,
	0x74, 0xe7, 0xc2, 0x9a, 0x0c, 0x77, 0x12, 0x55, 0xa0, 0x34, 0x95, 0x85, 0x46, 0x59, 0xa0, 0x34,
	0xdd, 0x49, 0xe0, 0x6b, 0xd0, 0x26, 0x39, 0x11, 0x04, 0xa5, 0x11, 0x65, 0x09, 0x66, 0xdc, 0x02,
	0xea, 0xcc, 0xae, 0xf7, 0x7b, 0x3f, 0x7b, 0x3b, 0x9a, 0xb1, 0x2b, 0x09, 0xc1, 0x9c, 0x9c, 0x36,
	0x6c, 0x91, 0xa9, 0x1c, 0x87, 0x6f, 0x00, 0xe0, 0x02, 0x31, 0x11, 0x49, 0xb3, 0x59, 0xcd, 0xae,
	0xe9, 0x36, 0xb7, 0x3a, 0x9e, 0x36, 0x9a, 0x57, 0x19, 0xcd, 0x7b, 0x55, 0x39, 0x31, 0xf8, 0x47,
	0x36, 0xf9, 0x7e, 0xe9, 0x2c, 0x1d, 0xa3, 0x2c, 0xdd, 0xee, 0xe1, 0x3c, 0x51, 0xcc, 0xde, 0xc9,
	0x95, 0x63, 0x86, 0x0d, 0xd5, 0x4a, 0x82, 0x61, 0x0f, 0xb4, 0xf2, 0x71, 0x16, 0xd1, 0x61, 0xc4,
	0x05, 0x1a, 0x61, 0x6e, 0x2d, 0x76, 0x4d, 0xb7, 0x15, 0x36, 0xf3, 0x71, 0xb6, 0x3b, 0xdc, 0x57,
	0x29, 0xf8, 0x02, 0xb4, 0x55, 0x31, 0xaa, 0x7c, 0x6c, 0xb5, 0x94, 0xfe, 0xea, 0x2f, 0xfa, 0xfd,
	0x12, 0x10, 0xd4, 0xa5, 0xfc, 0xa9, 0xd4, 0x6a, 0x29, 0x6a, 0x55, 0xd8, 0xae, 0x7f, 0x3a, 0x73,
	0x8c, 0xd3, 0x33, 0xc7, 0xe8, 0xdd, 0x9a, 0x60, 0x71, 0x7a, 0x6f, 0xf8, 0x04, 0x80, 0x3b, 0x77,
	0xaa, 0xaf, 0xe0, 0x5e, 0x73, 0xea, 0x33, 0x35, 0x26, 0x86, 0x83, 0x7d, 0x30, 0xaf, 0x5f, 0xfa,
	0xcc, 0x83, 0x5e, 0xba, 0x26, 0xc3, 0xe7, 0xa0, 0x91, 0x10, 0x86, 0x63, 0xb5, 0xa7, 0xfc, 0x8e,
	0xda, 0x5b, 0xeb, 0xf7, 0xbd, 0x3a, 0x35, 0x7b, 0xbf, 0x62, 0x84, 0x77, 0x64, 0xbd, 0xea, 0xed,
	0x99, 0x63, 0x04, 0xfb, 0xe7, 0xd7, 0xb6, 0x79, 0x71, 0x6d, 0x9b, 0xdf, 0xae, 0x6d, 0xf3, 0xe4,
	0xc6, 0x36, 0x2e, 0x6e, 0x6c, 0xe3, 0xcb, 0x8d, 0x6d, 0xbc, 0x7d, 0x3c, 0x3d, 0x5c, 0x29, 0xb2,
	0x91, 0x63, 0xf1, 0x81, 0xb2, 0x83, 0x49, 0xc2, 0x3f, 0x7c, 0xe4, 0x1f, 0x4d, 0xfd, 0x31, 0xa9,
	0x99, 0x07, 0x35, 0x75, 0xf5, 0xff, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x18, 0x3e, 0x5f,
	0x53, 0x05, 0x00, 0x00,
}

func (m *BootstrapProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BootstrapProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.StageDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.StageDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProposal(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x6a
	if m.NumOfStages != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.NumOfStages))
		i--
		dAtA[i] = 0x60
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintProposal(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x5a
	if len(m.InitialOrders) > 0 {
		for iNdEx := len(m.InitialOrders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialOrders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.PoolId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x48
	}
	if m.PairId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.MaxPrice.Size()
		i -= size
		if _, err := m.MaxPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.MinPrice.Size()
		i -= size
		if _, err := m.MinPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.QuoteCoinDenom) > 0 {
		i -= len(m.QuoteCoinDenom)
		copy(dAtA[i:], m.QuoteCoinDenom)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.QuoteCoinDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OfferCoins) > 0 {
		for iNdEx := len(m.OfferCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OfferCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ProposerAddress) > 0 {
		i -= len(m.ProposerAddress)
		copy(dAtA[i:], m.ProposerAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ProposerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InitialOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InitialOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InitialOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Direction != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.OfferCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BootstrapProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ProposerAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.OfferCoins) > 0 {
		for _, e := range m.OfferCoins {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	l = len(m.QuoteCoinDenom)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = m.MinPrice.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.MaxPrice.Size()
	n += 1 + l + sovProposal(uint64(l))
	if m.PairId != 0 {
		n += 1 + sovProposal(uint64(m.PairId))
	}
	if m.PoolId != 0 {
		n += 1 + sovProposal(uint64(m.PoolId))
	}
	if len(m.InitialOrders) > 0 {
		for _, e := range m.InitialOrders {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovProposal(uint64(l))
	if m.NumOfStages != 0 {
		n += 1 + sovProposal(uint64(m.NumOfStages))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.StageDuration)
	n += 1 + l + sovProposal(uint64(l))
	return n
}

func (m *InitialOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OfferCoin.Size()
	n += 1 + l + sovProposal(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovProposal(uint64(l))
	if m.Direction != 0 {
		n += 1 + sovProposal(uint64(m.Direction))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BootstrapProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BootstrapProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OfferCoins = append(m.OfferCoins, types.Coin{})
			if err := m.OfferCoins[len(m.OfferCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialOrders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialOrders = append(m.InitialOrders, InitialOrder{})
			if err := m.InitialOrders[len(m.InitialOrders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumOfStages", wireType)
			}
			m.NumOfStages = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumOfStages |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.StageDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InitialOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InitialOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InitialOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OfferCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= OrderDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
