// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_tx_out.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/thorchain/thornode/common"
	gitlab_com_thorchain_thornode_common "gitlab.com/thorchain/thornode/common"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type TxOutItem struct {
	Chain       gitlab_com_thorchain_thornode_common.Chain   `protobuf:"bytes,1,opt,name=chain,proto3,casttype=gitlab.com/thorchain/thornode/common.Chain" json:"chain,omitempty"`
	ToAddress   gitlab_com_thorchain_thornode_common.Address `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3,casttype=gitlab.com/thorchain/thornode/common.Address" json:"to_address,omitempty"`
	VaultPubKey gitlab_com_thorchain_thornode_common.PubKey  `protobuf:"bytes,3,opt,name=vault_pub_key,json=vaultPubKey,proto3,casttype=gitlab.com/thorchain/thornode/common.PubKey" json:"vault_pub_key,omitempty"`
	Coin        common.Coin                                  `protobuf:"bytes,4,opt,name=coin,proto3" json:"coin"`
	Memo        string                                       `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	MaxGas      gitlab_com_thorchain_thornode_common.Gas     `protobuf:"bytes,6,rep,name=max_gas,json=maxGas,proto3,castrepeated=gitlab.com/thorchain/thornode/common.Gas" json:"max_gas"`
	GasRate     int64                                        `protobuf:"varint,7,opt,name=gas_rate,json=gasRate,proto3" json:"gas_rate,omitempty"`
	InHash      gitlab_com_thorchain_thornode_common.TxID    `protobuf:"bytes,8,opt,name=in_hash,json=inHash,proto3,casttype=gitlab.com/thorchain/thornode/common.TxID" json:"in_hash,omitempty"`
	OutHash     gitlab_com_thorchain_thornode_common.TxID    `protobuf:"bytes,9,opt,name=out_hash,json=outHash,proto3,casttype=gitlab.com/thorchain/thornode/common.TxID" json:"out_hash,omitempty"`
	ModuleName  string                                       `protobuf:"bytes,10,opt,name=module_name,json=-,proto3" json:"-"`
}

func (m *TxOutItem) Reset()      { *m = TxOutItem{} }
func (*TxOutItem) ProtoMessage() {}
func (*TxOutItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a671465d0e1e96b8, []int{0}
}
func (m *TxOutItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxOutItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxOutItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxOutItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOutItem.Merge(m, src)
}
func (m *TxOutItem) XXX_Size() int {
	return m.Size()
}
func (m *TxOutItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOutItem.DiscardUnknown(m)
}

var xxx_messageInfo_TxOutItem proto.InternalMessageInfo

type TxOut struct {
	Height  int64       `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	TxArray []TxOutItem `protobuf:"bytes,2,rep,name=tx_array,json=txArray,proto3" json:"tx_array"`
}

func (m *TxOut) Reset()      { *m = TxOut{} }
func (*TxOut) ProtoMessage() {}
func (*TxOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_a671465d0e1e96b8, []int{1}
}
func (m *TxOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOut.Merge(m, src)
}
func (m *TxOut) XXX_Size() int {
	return m.Size()
}
func (m *TxOut) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOut.DiscardUnknown(m)
}

var xxx_messageInfo_TxOut proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TxOutItem)(nil), "types.TxOutItem")
	proto.RegisterType((*TxOut)(nil), "types.TxOut")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_tx_out.proto", fileDescriptor_a671465d0e1e96b8)
}

var fileDescriptor_a671465d0e1e96b8 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0x35, 0x6d, 0x5a, 0x17, 0x24, 0x64, 0x21, 0x64, 0x26, 0x91, 0x44, 0x3b, 0xa0,
	0xf2, 0xd2, 0xa4, 0x83, 0x1b, 0xb7, 0x95, 0x89, 0x6d, 0xe2, 0x30, 0x64, 0x86, 0x40, 0x5c, 0x2c,
	0xb7, 0xb5, 0x92, 0x88, 0x3a, 0xae, 0x62, 0x67, 0x4a, 0x6f, 0xfb, 0x08, 0x7c, 0x0e, 0x3e, 0x49,
	0x8f, 0x3b, 0xf6, 0x80, 0x02, 0x6b, 0x6f, 0x7c, 0x84, 0x9e, 0x50, 0x9d, 0x8e, 0x17, 0x4d, 0x42,
	0x15, 0x97, 0xe4, 0x79, 0x6c, 0xff, 0x7f, 0x8f, 0x9f, 0xe4, 0xff, 0xc0, 0x9e, 0x8e, 0x65, 0x36,
	0x8c, 0x59, 0x92, 0x86, 0xe7, 0xfb, 0x61, 0x11, 0xfe, 0x4e, 0xf5, 0x74, 0xc2, 0x95, 0x79, 0x52,
	0x5d, 0x50, 0x99, 0xeb, 0x60, 0x92, 0x49, 0x2d, 0x51, 0xdd, 0x6c, 0xec, 0xfa, 0x7f, 0x09, 0x87,
	0x52, 0x08, 0x99, 0x6e, 0x5e, 0xd5, 0xc1, 0xdd, 0xbb, 0x91, 0x8c, 0xa4, 0x09, 0xc3, 0x75, 0x54,
	0xad, 0xee, 0xcd, 0x6d, 0xd8, 0x3a, 0x2b, 0x4e, 0x73, 0x7d, 0xa2, 0xb9, 0x40, 0x87, 0xb0, 0x6e,
	0x18, 0x18, 0xf8, 0xa0, 0xd3, 0xea, 0x07, 0xab, 0xd2, 0x7b, 0x1c, 0x25, 0x7a, 0xcc, 0x06, 0xc1,
	0x50, 0x8a, 0x3f, 0x6f, 0x13, 0xcb, 0x2c, 0x95, 0x23, 0x7e, 0x5d, 0xe2, 0xe5, 0x7a, 0x95, 0x54,
	0x62, 0x74, 0x0a, 0xa1, 0x96, 0x94, 0x8d, 0x46, 0x19, 0x57, 0x0a, 0xef, 0x18, 0x54, 0x6f, 0x55,
	0x7a, 0x4f, 0xb7, 0x42, 0x1d, 0x54, 0x3a, 0xd2, 0xd2, 0x72, 0x13, 0xa2, 0xb7, 0xf0, 0xf6, 0x39,
	0xcb, 0xc7, 0x9a, 0x4e, 0xf2, 0x01, 0xfd, 0xc4, 0xa7, 0xb8, 0x66, 0x98, 0xe1, 0xaa, 0xf4, 0x9e,
	0x6c, 0xc5, 0x7c, 0x93, 0x0f, 0x5e, 0xf3, 0x29, 0x69, 0x1b, 0x4a, 0x95, 0xa0, 0x87, 0xd0, 0x1e,
	0xca, 0x24, 0xc5, 0xb6, 0x0f, 0x3a, 0xed, 0x67, 0xb7, 0x82, 0xeb, 0x4e, 0x64, 0x92, 0xf6, 0xed,
	0x59, 0xe9, 0x59, 0xc4, 0xec, 0x23, 0x04, 0x6d, 0xc1, 0x85, 0xc4, 0xf5, 0x75, 0x4d, 0x62, 0x62,
	0xf4, 0x1e, 0x3a, 0x82, 0x15, 0x34, 0x62, 0x0a, 0x37, 0xfc, 0xda, 0x0d, 0x79, 0x6f, 0x2d, 0xff,
	0xf2, 0xcd, 0xeb, 0x6c, 0x75, 0xb9, 0x23, 0xa6, 0x48, 0x43, 0xb0, 0xe2, 0x88, 0x29, 0x74, 0x1f,
	0x36, 0x23, 0xa6, 0x68, 0xc6, 0x34, 0xc7, 0x8e, 0x0f, 0x3a, 0x35, 0xe2, 0x44, 0x4c, 0x11, 0xa6,
	0x39, 0x7a, 0x05, 0x9d, 0x24, 0xa5, 0x31, 0x53, 0x31, 0x6e, 0x9a, 0xf6, 0xbb, 0xab, 0xd2, 0x7b,
	0xb4, 0x55, 0x85, 0xb3, 0xe2, 0xe4, 0x90, 0x34, 0x92, 0xf4, 0x98, 0xa9, 0x18, 0x1d, 0xc3, 0xa6,
	0xcc, 0x75, 0x05, 0x6a, 0xfd, 0x0f, 0xc8, 0x91, 0xb9, 0x36, 0xa4, 0x07, 0xb0, 0x2d, 0xe4, 0x28,
	0x1f, 0x73, 0x9a, 0x32, 0xc1, 0x31, 0x34, 0xb0, 0xfa, 0x8f, 0xd2, 0x03, 0x5d, 0x02, 0xba, 0x7b,
	0x1f, 0x60, 0xdd, 0x38, 0x0b, 0xdd, 0x83, 0x8d, 0x98, 0x27, 0x51, 0xac, 0x8d, 0xad, 0x6a, 0x64,
	0x93, 0xa1, 0x7d, 0xd8, 0xd4, 0x05, 0x65, 0x59, 0xc6, 0xa6, 0x78, 0xc7, 0x7c, 0xc6, 0x3b, 0x81,
	0x71, 0x73, 0xf0, 0xcb, 0x91, 0x9b, 0x3f, 0xe1, 0xe8, 0xe2, 0x60, 0x7d, 0xec, 0x85, 0x7d, 0xf1,
	0xd5, 0x07, 0xfd, 0x77, 0xb3, 0x2b, 0xd7, 0x9a, 0x5f, 0xb9, 0xd6, 0xc5, 0xc2, 0xb5, 0x66, 0x0b,
	0x17, 0x5c, 0x2e, 0x5c, 0xf0, 0x7d, 0xe1, 0x82, 0xcf, 0x4b, 0xd7, 0xba, 0x5c, 0xba, 0xd6, 0x7c,
	0xe9, 0x5a, 0x1f, 0xc3, 0x7f, 0xb7, 0x74, 0x63, 0xb8, 0x06, 0x0d, 0x33, 0x12, 0xcf, 0x7f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xe5, 0x35, 0x75, 0x15, 0x85, 0x03, 0x00, 0x00,
}

func (m *TxOutItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxOutItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxOutItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintTypeTxOut(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.OutHash) > 0 {
		i -= len(m.OutHash)
		copy(dAtA[i:], m.OutHash)
		i = encodeVarintTypeTxOut(dAtA, i, uint64(len(m.OutHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.InHash) > 0 {
		i -= len(m.InHash)
		copy(dAtA[i:], m.InHash)
		i = encodeVarintTypeTxOut(dAtA, i, uint64(len(m.InHash)))
		i--
		dAtA[i] = 0x42
	}
	if m.GasRate != 0 {
		i = encodeVarintTypeTxOut(dAtA, i, uint64(m.GasRate))
		i--
		dAtA[i] = 0x38
	}
	if len(m.MaxGas) > 0 {
		for iNdEx := len(m.MaxGas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxGas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeTxOut(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintTypeTxOut(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypeTxOut(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.VaultPubKey) > 0 {
		i -= len(m.VaultPubKey)
		copy(dAtA[i:], m.VaultPubKey)
		i = encodeVarintTypeTxOut(dAtA, i, uint64(len(m.VaultPubKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTypeTxOut(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypeTxOut(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxArray) > 0 {
		for iNdEx := len(m.TxArray) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxArray[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeTxOut(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintTypeTxOut(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeTxOut(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeTxOut(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxOutItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypeTxOut(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTypeTxOut(uint64(l))
	}
	l = len(m.VaultPubKey)
	if l > 0 {
		n += 1 + l + sovTypeTxOut(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovTypeTxOut(uint64(l))
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovTypeTxOut(uint64(l))
	}
	if len(m.MaxGas) > 0 {
		for _, e := range m.MaxGas {
			l = e.Size()
			n += 1 + l + sovTypeTxOut(uint64(l))
		}
	}
	if m.GasRate != 0 {
		n += 1 + sovTypeTxOut(uint64(m.GasRate))
	}
	l = len(m.InHash)
	if l > 0 {
		n += 1 + l + sovTypeTxOut(uint64(l))
	}
	l = len(m.OutHash)
	if l > 0 {
		n += 1 + l + sovTypeTxOut(uint64(l))
	}
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovTypeTxOut(uint64(l))
	}
	return n
}

func (m *TxOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypeTxOut(uint64(m.Height))
	}
	if len(m.TxArray) > 0 {
		for _, e := range m.TxArray {
			l = e.Size()
			n += 1 + l + sovTypeTxOut(uint64(l))
		}
	}
	return n
}

func sovTypeTxOut(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeTxOut(x uint64) (n int) {
	return sovTypeTxOut(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TxOut) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForTxArray := "[]TxOutItem{"
	for _, f := range this.TxArray {
		repeatedStringForTxArray += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForTxArray += "}"
	s := strings.Join([]string{`&TxOut{`,
		`Height:` + fmt.Sprintf("%v", this.Height) + `,`,
		`TxArray:` + repeatedStringForTxArray + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypeTxOut(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TxOutItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTxOut
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
			return fmt.Errorf("proto: TxOutItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxOutItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = gitlab_com_thorchain_thornode_common.Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = gitlab_com_thorchain_thornode_common.Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VaultPubKey = gitlab_com_thorchain_thornode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxGas = append(m.MaxGas, common.Coin{})
			if err := m.MaxGas[len(m.MaxGas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasRate", wireType)
			}
			m.GasRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasRate |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InHash = gitlab_com_thorchain_thornode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutHash = gitlab_com_thorchain_thornode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTxOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeTxOut
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
func (m *TxOut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeTxOut
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
			return fmt.Errorf("proto: TxOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxArray", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeTxOut
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
				return ErrInvalidLengthTypeTxOut
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxArray = append(m.TxArray, TxOutItem{})
			if err := m.TxArray[len(m.TxArray)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypeTxOut(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeTxOut
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeTxOut
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
func skipTypeTxOut(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeTxOut
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
					return 0, ErrIntOverflowTypeTxOut
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
					return 0, ErrIntOverflowTypeTxOut
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
				return 0, ErrInvalidLengthTypeTxOut
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeTxOut
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeTxOut
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeTxOut        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeTxOut          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeTxOut = fmt.Errorf("proto: unexpected end of group")
)