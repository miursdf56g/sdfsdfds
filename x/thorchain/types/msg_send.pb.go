// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/msg_send.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSend struct {
	FromAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from_address,omitempty"`
	ToAddress   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"to_address,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *MsgSend) Reset()         { *m = MsgSend{} }
func (m *MsgSend) String() string { return proto.CompactTextString(m) }
func (*MsgSend) ProtoMessage()    {}
func (*MsgSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d732b92ac4191cd, []int{0}
}
func (m *MsgSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSend.Merge(m, src)
}
func (m *MsgSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSend proto.InternalMessageInfo

func (m *MsgSend) GetFromAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.FromAddress
	}
	return nil
}

func (m *MsgSend) GetToAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *MsgSend) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSend)(nil), "types.MsgSend")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/msg_send.proto", fileDescriptor_1d732b92ac4191cd)
}

var fileDescriptor_1d732b92ac4191cd = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0x93, 0x56, 0x14, 0xe1, 0x76, 0xaa, 0x18, 0x4a, 0x07, 0xb7, 0x62, 0xea, 0x12, 0x9b,
	0xc0, 0x09, 0x1a, 0x26, 0x06, 0x24, 0x54, 0x98, 0x58, 0x2a, 0xc7, 0x36, 0x4e, 0x04, 0xf1, 0xab,
	0x62, 0xb7, 0x82, 0x5b, 0x30, 0x71, 0x08, 0x4e, 0xd2, 0xb1, 0x23, 0x53, 0x41, 0xc9, 0x2d, 0x98,
	0x50, 0xe2, 0xf0, 0x47, 0x62, 0x41, 0x4c, 0x7e, 0xef, 0xc9, 0xdf, 0xef, 0xf9, 0xf3, 0x87, 0x02,
	0x9b, 0x40, 0xce, 0x13, 0x96, 0x6a, 0xba, 0x0a, 0xe9, 0x3d, 0xfd, 0x6e, 0xed, 0xc3, 0x42, 0x1a,
	0x9a, 0x19, 0x35, 0x37, 0x52, 0x0b, 0xb2, 0xc8, 0xc1, 0x42, 0x7f, 0xa7, 0x9e, 0x0e, 0x31, 0x07,
	0x93, 0x81, 0xa1, 0x31, 0x33, 0x92, 0xae, 0xc2, 0x58, 0x5a, 0x16, 0x52, 0x0e, 0xa9, 0x76, 0xd7,
	0x86, 0xfb, 0x0a, 0x14, 0xd4, 0x25, 0xad, 0x2a, 0x37, 0x3d, 0x7c, 0x6a, 0xa1, 0xdd, 0x73, 0xa3,
	0x2e, 0xa5, 0x16, 0xfd, 0x2b, 0xd4, 0xbb, 0xc9, 0x21, 0x9b, 0x33, 0x21, 0x72, 0x69, 0xcc, 0xc0,
	0x1f, 0xfb, 0x93, 0x5e, 0x14, 0xbe, 0x6f, 0x47, 0x81, 0x4a, 0x6d, 0xb2, 0x8c, 0x09, 0x87, 0x8c,
	0x36, 0x6b, 0xdc, 0x11, 0x18, 0x71, 0xeb, 0x5e, 0x45, 0xa6, 0x9c, 0x4f, 0x9d, 0x70, 0xd6, 0xad,
	0x30, 0x4d, 0xd3, 0xbf, 0x40, 0xc8, 0xc2, 0x17, 0xb3, 0xf5, 0x5f, 0xe6, 0x9e, 0x85, 0x4f, 0x22,
	0x47, 0x1d, 0x96, 0xc1, 0x52, 0xdb, 0x41, 0x7b, 0xdc, 0x9e, 0x74, 0x8f, 0x0f, 0x88, 0x13, 0x92,
	0xca, 0x3a, 0x69, 0xac, 0x93, 0x53, 0x48, 0x75, 0x74, 0xb4, 0xde, 0x8e, 0xbc, 0xe7, 0xd7, 0xd1,
	0xe4, 0x0f, 0xcb, 0x2a, 0x81, 0x99, 0x35, 0xe8, 0xe8, 0x6c, 0x5d, 0x60, 0x7f, 0x53, 0x60, 0xff,
	0xad, 0xc0, 0xfe, 0x63, 0x89, 0xbd, 0x4d, 0x89, 0xbd, 0x97, 0x12, 0x7b, 0xd7, 0x54, 0xa5, 0xf6,
	0x8e, 0x39, 0xd6, 0x8f, 0x68, 0x12, 0xc8, 0x35, 0x08, 0xf9, 0x3b, 0xaf, 0xb8, 0x53, 0x7f, 0xf5,
	0xc9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x15, 0xca, 0x3f, 0xd8, 0x01, 0x00, 0x00,
}

func (m *MsgSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgSend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintMsgSend(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMsgSend(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgSend(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgSend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMsgSend(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovMsgSend(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovMsgSend(uint64(l))
		}
	}
	return n
}

func sovMsgSend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgSend(x uint64) (n int) {
	return sovMsgSend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgSend
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
			return fmt.Errorf("proto: MsgSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSend
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
				return ErrInvalidLengthMsgSend
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = append(m.FromAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.FromAddress == nil {
				m.FromAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSend
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
				return ErrInvalidLengthMsgSend
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = append(m.ToAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ToAddress == nil {
				m.ToAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgSend
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
				return ErrInvalidLengthMsgSend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgSend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgSend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgSend
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
func skipMsgSend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgSend
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
					return 0, ErrIntOverflowMsgSend
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
					return 0, ErrIntOverflowMsgSend
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
				return 0, ErrInvalidLengthMsgSend
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgSend
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgSend
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgSend        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgSend          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgSend = fmt.Errorf("proto: unexpected end of group")
)
