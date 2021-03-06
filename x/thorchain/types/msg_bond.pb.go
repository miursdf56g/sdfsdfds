// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/msg_bond.proto

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

type MsgBond struct {
	TxIn                common.Tx                                     `protobuf:"bytes,1,opt,name=tx_in,json=txIn,proto3" json:"tx_in"`
	NodeAddress         github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"node_address,omitempty"`
	Bond                github_com_cosmos_cosmos_sdk_types.Uint       `protobuf:"bytes,3,opt,name=bond,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"bond"`
	BondAddress         gitlab_com_thorchain_thornode_common.Address  `protobuf:"bytes,4,opt,name=bond_address,json=bondAddress,proto3,casttype=gitlab.com/thorchain/thornode/common.Address" json:"bond_address,omitempty"`
	Signer              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
	BondProviderAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=bond_provider_address,json=bondProviderAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"bond_provider_address,omitempty"`
}

func (m *MsgBond) Reset()         { *m = MsgBond{} }
func (m *MsgBond) String() string { return proto.CompactTextString(m) }
func (*MsgBond) ProtoMessage()    {}
func (*MsgBond) Descriptor() ([]byte, []int) {
	return fileDescriptor_02a49a5ca01154e8, []int{0}
}
func (m *MsgBond) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBond.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBond.Merge(m, src)
}
func (m *MsgBond) XXX_Size() int {
	return m.Size()
}
func (m *MsgBond) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBond.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBond proto.InternalMessageInfo

func (m *MsgBond) GetTxIn() common.Tx {
	if m != nil {
		return m.TxIn
	}
	return common.Tx{}
}

func (m *MsgBond) GetNodeAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.NodeAddress
	}
	return nil
}

func (m *MsgBond) GetBondAddress() gitlab_com_thorchain_thornode_common.Address {
	if m != nil {
		return m.BondAddress
	}
	return ""
}

func (m *MsgBond) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *MsgBond) GetBondProviderAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.BondProviderAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgBond)(nil), "types.MsgBond")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/msg_bond.proto", fileDescriptor_02a49a5ca01154e8)
}

var fileDescriptor_02a49a5ca01154e8 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x93, 0xdb, 0xb4, 0x17, 0xa7, 0x5d, 0x45, 0x85, 0xd0, 0x45, 0x12, 0x04, 0xb1, 0x0b,
	0xdb, 0xb1, 0xfa, 0x04, 0x8d, 0xab, 0x2e, 0x04, 0xa9, 0x75, 0xe3, 0xa6, 0xb4, 0x49, 0x98, 0x0c,
	0x9a, 0x39, 0x65, 0x66, 0x2c, 0xf1, 0x2d, 0x7c, 0xac, 0x2e, 0x8b, 0x2b, 0x71, 0x11, 0xa4, 0x7d,
	0x8b, 0xae, 0x64, 0x92, 0x69, 0x55, 0x04, 0x91, 0xae, 0xce, 0x1f, 0xe6, 0xfb, 0xcd, 0x77, 0x0e,
	0x07, 0xb5, 0x65, 0x02, 0x3c, 0x4c, 0xc6, 0x94, 0xe1, 0x59, 0x17, 0x67, 0xf8, 0xb3, 0x94, 0x4f,
	0xd3, 0x58, 0xe0, 0x54, 0x90, 0xd1, 0x04, 0x58, 0xd4, 0x99, 0x72, 0x90, 0x60, 0x57, 0x8b, 0x6e,
	0xd3, 0xff, 0xa6, 0x0a, 0x21, 0x4d, 0x81, 0xe9, 0x50, 0x3e, 0x6c, 0x1e, 0x10, 0x20, 0x50, 0xa4,
	0x58, 0x65, 0x65, 0xf7, 0xe8, 0xa5, 0x82, 0xfe, 0x5f, 0x09, 0x12, 0x00, 0x8b, 0xec, 0x63, 0x54,
	0x95, 0xd9, 0x88, 0x32, 0xc7, 0xf4, 0xcd, 0x56, 0xfd, 0x1c, 0x75, 0xb4, 0x7e, 0x98, 0x05, 0xd6,
	0x3c, 0xf7, 0x8c, 0x81, 0x25, 0xb3, 0x3e, 0xb3, 0x87, 0xa8, 0xc1, 0x20, 0x8a, 0x47, 0xe3, 0x28,
	0xe2, 0xb1, 0x10, 0xce, 0x3f, 0xdf, 0x6c, 0x35, 0x82, 0xee, 0x3a, 0xf7, 0xda, 0x84, 0xca, 0xe4,
	0x71, 0xa2, 0x74, 0x38, 0x04, 0x91, 0x82, 0xd0, 0xa1, 0x2d, 0xa2, 0xfb, 0xd2, 0x7e, 0xa7, 0x17,
	0x86, 0xbd, 0x52, 0x38, 0xa8, 0x2b, 0x8c, 0x2e, 0xec, 0x4b, 0x64, 0xa9, 0xa9, 0x9c, 0x8a, 0x6f,
	0xb6, 0xf6, 0x02, 0xac, 0xfe, 0x7b, 0xcb, 0xbd, 0x93, 0x3f, 0x10, 0x6f, 0x29, 0x93, 0x83, 0x42,
	0x6c, 0xdf, 0xa0, 0x86, 0x8a, 0x5b, 0x6b, 0x56, 0x01, 0x3b, 0x5b, 0xe7, 0xde, 0x29, 0xa1, 0xf2,
	0x61, 0x5c, 0x82, 0xbe, 0x6c, 0x34, 0x01, 0xae, 0x4c, 0x6c, 0x36, 0xb5, 0x75, 0xa6, 0x28, 0x1b,
	0x67, 0x7d, 0x54, 0x13, 0x94, 0xb0, 0x98, 0x3b, 0xd5, 0x5d, 0x27, 0xd5, 0x00, 0x3b, 0x46, 0x87,
	0x85, 0xbf, 0x29, 0x87, 0x19, 0x8d, 0x62, 0xbe, 0x35, 0x5a, 0xdb, 0x95, 0xbc, 0xaf, 0x78, 0xd7,
	0x1a, 0xa7, 0x9b, 0x41, 0x7f, 0xbe, 0x74, 0xcd, 0xc5, 0xd2, 0x35, 0xdf, 0x97, 0xae, 0xf9, 0xbc,
	0x72, 0x8d, 0xc5, 0xca, 0x35, 0x5e, 0x57, 0xae, 0x71, 0x87, 0x7f, 0x5f, 0xc3, 0x8f, 0x6b, 0x9b,
	0xd4, 0x8a, 0x33, 0xb9, 0xf8, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x34, 0xfa, 0xa5, 0x96, 0x02,
	0x00, 0x00,
}

func (m *MsgBond) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBond) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBond) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BondProviderAddress) > 0 {
		i -= len(m.BondProviderAddress)
		copy(dAtA[i:], m.BondProviderAddress)
		i = encodeVarintMsgBond(dAtA, i, uint64(len(m.BondProviderAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgBond(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BondAddress) > 0 {
		i -= len(m.BondAddress)
		copy(dAtA[i:], m.BondAddress)
		i = encodeVarintMsgBond(dAtA, i, uint64(len(m.BondAddress)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Bond.Size()
		i -= size
		if _, err := m.Bond.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsgBond(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintMsgBond(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.TxIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgBond(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMsgBond(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgBond(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgBond) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TxIn.Size()
	n += 1 + l + sovMsgBond(uint64(l))
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovMsgBond(uint64(l))
	}
	l = m.Bond.Size()
	n += 1 + l + sovMsgBond(uint64(l))
	l = len(m.BondAddress)
	if l > 0 {
		n += 1 + l + sovMsgBond(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgBond(uint64(l))
	}
	l = len(m.BondProviderAddress)
	if l > 0 {
		n += 1 + l + sovMsgBond(uint64(l))
	}
	return n
}

func sovMsgBond(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgBond(x uint64) (n int) {
	return sovMsgBond(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgBond) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgBond
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
			return fmt.Errorf("proto: MsgBond: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBond: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBond
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
				return ErrInvalidLengthMsgBond
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBond
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
				return ErrInvalidLengthMsgBond
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = append(m.NodeAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeAddress == nil {
				m.NodeAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bond", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBond
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
				return ErrInvalidLengthMsgBond
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bond.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBond
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
				return ErrInvalidLengthMsgBond
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BondAddress = gitlab_com_thorchain_thornode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBond
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
				return ErrInvalidLengthMsgBond
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondProviderAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgBond
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
				return ErrInvalidLengthMsgBond
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgBond
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BondProviderAddress = append(m.BondProviderAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.BondProviderAddress == nil {
				m.BondProviderAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgBond(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgBond
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgBond
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
func skipMsgBond(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgBond
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
					return 0, ErrIntOverflowMsgBond
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
					return 0, ErrIntOverflowMsgBond
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
				return 0, ErrInvalidLengthMsgBond
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgBond
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgBond
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgBond        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgBond          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgBond = fmt.Errorf("proto: unexpected end of group")
)
