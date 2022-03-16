// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/msg_solvency.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/thorchain/thornode/common"
	gitlab_com_thorchain_thornode_common "gitlab.com/thorchain/thornode/common"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSolvency struct {
	Id     gitlab_com_thorchain_thornode_common.TxID     `protobuf:"bytes,1,opt,name=id,proto3,casttype=gitlab.com/thorchain/thornode/common.TxID" json:"id,omitempty"`
	Chain  gitlab_com_thorchain_thornode_common.Chain    `protobuf:"bytes,2,opt,name=chain,proto3,casttype=gitlab.com/thorchain/thornode/common.Chain" json:"chain,omitempty"`
	PubKey gitlab_com_thorchain_thornode_common.PubKey   `protobuf:"bytes,3,opt,name=pub_key,json=pubKey,proto3,casttype=gitlab.com/thorchain/thornode/common.PubKey" json:"pub_key,omitempty"`
	Coins  gitlab_com_thorchain_thornode_common.Coins    `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=gitlab.com/thorchain/thornode/common.Coins" json:"coins"`
	Height int64                                         `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Signer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgSolvency) Reset()         { *m = MsgSolvency{} }
func (m *MsgSolvency) String() string { return proto.CompactTextString(m) }
func (*MsgSolvency) ProtoMessage()    {}
func (*MsgSolvency) Descriptor() ([]byte, []int) {
	return fileDescriptor_324cd31d37ed2e28, []int{0}
}
func (m *MsgSolvency) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSolvency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSolvency.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSolvency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSolvency.Merge(m, src)
}
func (m *MsgSolvency) XXX_Size() int {
	return m.Size()
}
func (m *MsgSolvency) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSolvency.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSolvency proto.InternalMessageInfo

func (m *MsgSolvency) GetId() gitlab_com_thorchain_thornode_common.TxID {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgSolvency) GetChain() gitlab_com_thorchain_thornode_common.Chain {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *MsgSolvency) GetPubKey() gitlab_com_thorchain_thornode_common.PubKey {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *MsgSolvency) GetCoins() gitlab_com_thorchain_thornode_common.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *MsgSolvency) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MsgSolvency) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSolvency)(nil), "types.MsgSolvency")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/msg_solvency.proto", fileDescriptor_324cd31d37ed2e28)
}

var fileDescriptor_324cd31d37ed2e28 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x6e, 0xa3, 0x40,
	0x10, 0x86, 0xc1, 0x1c, 0x9c, 0x6e, 0xed, 0x0a, 0x9d, 0x4e, 0xc8, 0x05, 0xa0, 0xab, 0xb8, 0x8b,
	0x0c, 0xb2, 0x53, 0xa7, 0x30, 0x71, 0x11, 0x2b, 0x8a, 0x14, 0x91, 0x28, 0x45, 0x1a, 0xcb, 0x2c,
	0xab, 0x65, 0x65, 0xc3, 0x20, 0x16, 0x5b, 0xe6, 0x2d, 0xd2, 0xe5, 0x1d, 0xf2, 0x24, 0x2e, 0x5d,
	0xa6, 0x22, 0x91, 0xfd, 0x16, 0x54, 0x91, 0x01, 0x2b, 0x89, 0x22, 0x45, 0xae, 0xfe, 0x99, 0xd9,
	0x99, 0x6f, 0x7e, 0x69, 0x16, 0xf5, 0xb3, 0x10, 0x52, 0x1c, 0x4e, 0x59, 0xec, 0x2c, 0xfb, 0xce,
	0xca, 0x79, 0x4f, 0xb3, 0x3c, 0x21, 0xdc, 0x89, 0x38, 0x9d, 0x70, 0x98, 0x2f, 0x49, 0x8c, 0x73,
	0x3b, 0x49, 0x21, 0x03, 0x55, 0xae, 0x5e, 0xba, 0xe6, 0xa7, 0x49, 0x0c, 0x51, 0x04, 0x71, 0x23,
	0x75, 0x63, 0xf7, 0x37, 0x05, 0x0a, 0x55, 0xe8, 0xec, 0xa3, 0xba, 0xfa, 0xf7, 0x51, 0x42, 0xed,
	0x2b, 0x4e, 0x6f, 0x1a, 0xa8, 0x7a, 0x86, 0x5a, 0x2c, 0xd0, 0x44, 0x53, 0xb4, 0x7e, 0xb9, 0xbd,
	0xb2, 0x30, 0xfe, 0x51, 0x96, 0xcd, 0xa7, 0xbe, 0x8d, 0x21, 0xfa, 0xe8, 0x26, 0x84, 0x34, 0x86,
	0x80, 0x1c, 0x36, 0xdc, 0xae, 0xc6, 0x23, 0xaf, 0xc5, 0x02, 0x75, 0x84, 0xe4, 0xaa, 0x43, 0x6b,
	0x55, 0x04, 0xbb, 0x2c, 0x8c, 0xff, 0x47, 0x11, 0xce, 0xf7, 0x55, 0xaf, 0x1e, 0x56, 0x2f, 0xd0,
	0xcf, 0x64, 0xe1, 0x4f, 0x66, 0x24, 0xd7, 0xa4, 0x8a, 0xe3, 0x94, 0x85, 0x71, 0x72, 0x14, 0xe7,
	0x7a, 0xe1, 0x5f, 0x92, 0xdc, 0x53, 0x92, 0x4a, 0xd5, 0x3b, 0x24, 0x63, 0x60, 0x31, 0xd7, 0x7e,
	0x98, 0x92, 0xd5, 0x1e, 0x74, 0xec, 0xc3, 0x3a, 0x60, 0xb1, 0x3b, 0x58, 0x17, 0x86, 0xf0, 0xf4,
	0x72, 0xac, 0xc3, 0x3d, 0xc7, 0xab, 0x71, 0xea, 0x1f, 0xa4, 0x84, 0x84, 0xd1, 0x30, 0xd3, 0x64,
	0x53, 0xb4, 0x24, 0xaf, 0xc9, 0xd4, 0x31, 0x52, 0x38, 0xa3, 0x31, 0x49, 0x35, 0xc5, 0x14, 0xad,
	0x8e, 0xdb, 0x2f, 0x0b, 0xa3, 0x47, 0x59, 0x16, 0x2e, 0x6a, 0x3c, 0x06, 0x1e, 0x01, 0x6f, 0xa4,
	0xc7, 0x83, 0x59, 0x7d, 0x58, 0x7b, 0x88, 0xf1, 0x30, 0x08, 0x52, 0xc2, 0xb9, 0xd7, 0x00, 0xdc,
	0xf1, 0x7a, 0xab, 0x8b, 0x9b, 0xad, 0x2e, 0xbe, 0x6e, 0x75, 0xf1, 0x61, 0xa7, 0x0b, 0x9b, 0x9d,
	0x2e, 0x3c, 0xef, 0x74, 0xe1, 0xde, 0xf9, 0xde, 0xef, 0x97, 0x6f, 0xe3, 0x2b, 0xd5, 0xad, 0x4f,
	0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x88, 0xbd, 0xed, 0x5f, 0x02, 0x00, 0x00,
}

func (m *MsgSolvency) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSolvency) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSolvency) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgSolvency(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x32
	}
	if m.Height != 0 {
		i = encodeVarintMsgSolvency(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgSolvency(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintMsgSolvency(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintMsgSolvency(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintMsgSolvency(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgSolvency(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSolvency(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSolvency) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMsgSolvency(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovMsgSolvency(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovMsgSolvency(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovMsgSolvency(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovMsgSolvency(uint64(m.Height))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgSolvency(uint64(l))
	}
	return n
}

func sovMsgSolvency(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSolvency(x uint64) (n int) {
	return sovMsgSolvency(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSolvency) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSolvency
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
			return fmt.Errorf("proto: MsgSolvency: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSolvency: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSolvency
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
				return ErrInvalidLengthMsgSolvency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSolvency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = gitlab_com_thorchain_thornode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSolvency
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
				return ErrInvalidLengthMsgSolvency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSolvency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = gitlab_com_thorchain_thornode_common.Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSolvency
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
				return ErrInvalidLengthMsgSolvency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSolvency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = gitlab_com_thorchain_thornode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSolvency
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
				return ErrInvalidLengthMsgSolvency
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSolvency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, common.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSolvency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSolvency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgSolvency
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSolvency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgSolvency(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgSolvency
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgSolvency
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
func skipMsgSolvency(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSolvency
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
					return 0, ErrIntOverflowMsgSolvency
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
					return 0, ErrIntOverflowMsgSolvency
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
				return 0, ErrInvalidLengthMsgSolvency
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSolvency
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSolvency
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSolvency        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSolvency          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSolvency = fmt.Errorf("proto: unexpected end of group")
)