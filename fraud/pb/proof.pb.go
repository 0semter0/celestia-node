// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fraud/pb/proof.proto

package fraud_pb

import (
	fmt "fmt"
	pb "github.com/celestiaorg/celestia-node/ipld/pb"
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

type Axis int32

const (
	Axis_ROW Axis = 0
	Axis_COL Axis = 1
)

var Axis_name = map[int32]string{
	0: "ROW",
	1: "COL",
}

var Axis_value = map[string]int32{
	"ROW": 0,
	"COL": 1,
}

func (x Axis) String() string {
	return proto.EnumName(Axis_name, int32(x))
}

func (Axis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{0}
}

type ProofType int32

const (
	ProofType_BADENCODING ProofType = 0
)

var ProofType_name = map[int32]string{
	0: "BADENCODING",
}

var ProofType_value = map[string]int32{
	"BADENCODING": 0,
}

func (x ProofType) String() string {
	return proto.EnumName(ProofType_name, int32(x))
}

func (ProofType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{1}
}

type BadEncoding struct {
	HeaderHash []byte      `protobuf:"bytes,1,opt,name=HeaderHash,proto3" json:"HeaderHash,omitempty"`
	Height     uint64      `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	Shares     []*pb.Share `protobuf:"bytes,3,rep,name=Shares,proto3" json:"Shares,omitempty"`
	Index      uint32      `protobuf:"varint,4,opt,name=Index,proto3" json:"Index,omitempty"`
	Axis       Axis        `protobuf:"varint,5,opt,name=Axis,proto3,enum=fraud.pb.Axis" json:"Axis,omitempty"`
}

func (m *BadEncoding) Reset()         { *m = BadEncoding{} }
func (m *BadEncoding) String() string { return proto.CompactTextString(m) }
func (*BadEncoding) ProtoMessage()    {}
func (*BadEncoding) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{0}
}
func (m *BadEncoding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadEncoding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadEncoding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadEncoding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadEncoding.Merge(m, src)
}
func (m *BadEncoding) XXX_Size() int {
	return m.Size()
}
func (m *BadEncoding) XXX_DiscardUnknown() {
	xxx_messageInfo_BadEncoding.DiscardUnknown(m)
}

var xxx_messageInfo_BadEncoding proto.InternalMessageInfo

func (m *BadEncoding) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *BadEncoding) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BadEncoding) GetShares() []*pb.Share {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *BadEncoding) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BadEncoding) GetAxis() Axis {
	if m != nil {
		return m.Axis
	}
	return Axis_ROW
}

type FraudMessageRequest struct {
	RequestedProofType []ProofType `protobuf:"varint,1,rep,packed,name=RequestedProofType,proto3,enum=fraud.pb.ProofType" json:"RequestedProofType,omitempty"`
}

func (m *FraudMessageRequest) Reset()         { *m = FraudMessageRequest{} }
func (m *FraudMessageRequest) String() string { return proto.CompactTextString(m) }
func (*FraudMessageRequest) ProtoMessage()    {}
func (*FraudMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{1}
}
func (m *FraudMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FraudMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FraudMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FraudMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FraudMessageRequest.Merge(m, src)
}
func (m *FraudMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *FraudMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FraudMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FraudMessageRequest proto.InternalMessageInfo

func (m *FraudMessageRequest) GetRequestedProofType() []ProofType {
	if m != nil {
		return m.RequestedProofType
	}
	return nil
}

type ProofResponse struct {
	Type  ProofType `protobuf:"varint,1,opt,name=Type,proto3,enum=fraud.pb.ProofType" json:"Type,omitempty"`
	Value []byte    `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *ProofResponse) Reset()         { *m = ProofResponse{} }
func (m *ProofResponse) String() string { return proto.CompactTextString(m) }
func (*ProofResponse) ProtoMessage()    {}
func (*ProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{2}
}
func (m *ProofResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofResponse.Merge(m, src)
}
func (m *ProofResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProofResponse proto.InternalMessageInfo

func (m *ProofResponse) GetType() ProofType {
	if m != nil {
		return m.Type
	}
	return ProofType_BADENCODING
}

func (m *ProofResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type FraudMessageResponse struct {
	Proofs []*ProofResponse `protobuf:"bytes,1,rep,name=Proofs,proto3" json:"Proofs,omitempty"`
}

func (m *FraudMessageResponse) Reset()         { *m = FraudMessageResponse{} }
func (m *FraudMessageResponse) String() string { return proto.CompactTextString(m) }
func (*FraudMessageResponse) ProtoMessage()    {}
func (*FraudMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_318cb87a8bb2d394, []int{3}
}
func (m *FraudMessageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FraudMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FraudMessageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FraudMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FraudMessageResponse.Merge(m, src)
}
func (m *FraudMessageResponse) XXX_Size() int {
	return m.Size()
}
func (m *FraudMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FraudMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FraudMessageResponse proto.InternalMessageInfo

func (m *FraudMessageResponse) GetProofs() []*ProofResponse {
	if m != nil {
		return m.Proofs
	}
	return nil
}

func init() {
	proto.RegisterEnum("fraud.pb.Axis", Axis_name, Axis_value)
	proto.RegisterEnum("fraud.pb.ProofType", ProofType_name, ProofType_value)
	proto.RegisterType((*BadEncoding)(nil), "fraud.pb.BadEncoding")
	proto.RegisterType((*FraudMessageRequest)(nil), "fraud.pb.FraudMessageRequest")
	proto.RegisterType((*ProofResponse)(nil), "fraud.pb.ProofResponse")
	proto.RegisterType((*FraudMessageResponse)(nil), "fraud.pb.FraudMessageResponse")
}

func init() { proto.RegisterFile("fraud/pb/proof.proto", fileDescriptor_318cb87a8bb2d394) }

var fileDescriptor_318cb87a8bb2d394 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xcd, 0x6e, 0xda, 0x40,
	0x18, 0xf4, 0xd6, 0xc6, 0x6d, 0x3f, 0x83, 0x8b, 0x16, 0xd4, 0xae, 0xaa, 0xca, 0xb2, 0x7c, 0x68,
	0x2d, 0x0e, 0x46, 0x72, 0x9f, 0x80, 0xbf, 0x02, 0x52, 0x0b, 0xd5, 0xb6, 0x6a, 0xa5, 0xdc, 0xd6,
	0xf1, 0x02, 0x96, 0x90, 0xbd, 0xf1, 0x82, 0x44, 0xde, 0x22, 0x6f, 0x91, 0x57, 0xc9, 0x91, 0x63,
	0x8e, 0x11, 0xbc, 0x48, 0xe4, 0xc5, 0x84, 0x20, 0x25, 0xb7, 0x9d, 0x99, 0x4f, 0x33, 0xf3, 0x7d,
	0x0b, 0xcd, 0x59, 0xce, 0xd6, 0x71, 0x5b, 0x44, 0x6d, 0x91, 0x67, 0xd9, 0x2c, 0x10, 0x79, 0xb6,
	0xca, 0xf0, 0x3b, 0xc5, 0x06, 0x22, 0xfa, 0xdc, 0x48, 0xc4, 0x52, 0xc9, 0x72, 0xc1, 0x72, 0x7e,
	0x90, 0xbd, 0x5b, 0x04, 0x56, 0x97, 0xc5, 0x83, 0xf4, 0x32, 0x8b, 0x93, 0x74, 0x8e, 0x1d, 0x80,
	0x11, 0x67, 0x31, 0xcf, 0x47, 0x4c, 0x2e, 0x08, 0x72, 0x91, 0x5f, 0xa5, 0xcf, 0x18, 0xfc, 0x11,
	0xcc, 0x11, 0x4f, 0xe6, 0x8b, 0x15, 0x79, 0xe3, 0x22, 0xdf, 0xa0, 0x25, 0xc2, 0x5f, 0xc1, 0xfc,
	0x53, 0xd8, 0x4a, 0xa2, 0xbb, 0xba, 0x6f, 0x85, 0x76, 0x50, 0xa4, 0x05, 0x22, 0x0a, 0x14, 0x4d,
	0x4b, 0x15, 0x37, 0xa1, 0x32, 0x4e, 0x63, 0xbe, 0x21, 0x86, 0x8b, 0xfc, 0x1a, 0x3d, 0x00, 0xec,
	0x81, 0xd1, 0xd9, 0x24, 0x92, 0x54, 0x5c, 0xe4, 0xdb, 0xa1, 0x1d, 0x1c, 0x3b, 0x07, 0x6c, 0x93,
	0x48, 0xaa, 0x34, 0xef, 0x02, 0x1a, 0x3f, 0x0a, 0xfa, 0x17, 0x97, 0x92, 0xcd, 0x39, 0xe5, 0x57,
	0x6b, 0x2e, 0x57, 0xb8, 0x07, 0xb8, 0x7c, 0xf2, 0xf8, 0x77, 0xb1, 0xf7, 0xdf, 0x6b, 0xc1, 0x09,
	0x72, 0x75, 0xdf, 0x0e, 0x1b, 0x27, 0x23, 0x71, 0x94, 0xe8, 0x0b, 0xe3, 0xde, 0x04, 0x6a, 0x0a,
	0x50, 0x2e, 0x45, 0x96, 0x4a, 0x8e, 0xbf, 0x81, 0x51, 0xfa, 0xa0, 0xd7, 0x7c, 0xd4, 0x40, 0xb1,
	0xcf, 0x3f, 0xb6, 0x5c, 0x73, 0x75, 0x8e, 0x2a, 0x3d, 0x00, 0x6f, 0x08, 0xcd, 0xf3, 0xae, 0xa5,
	0x6d, 0x1b, 0x4c, 0x95, 0x23, 0x55, 0x41, 0x2b, 0xfc, 0x74, 0x32, 0x3e, 0xcb, 0xa7, 0xe5, 0x58,
	0x8b, 0x80, 0x51, 0x9c, 0x00, 0xbf, 0x05, 0x9d, 0x4e, 0xff, 0xd7, 0xb5, 0xe2, 0xd1, 0x9b, 0xfe,
	0xac, 0xa3, 0xd6, 0x17, 0x78, 0xff, 0xd4, 0x05, 0x7f, 0x00, 0xab, 0xdb, 0xe9, 0x0f, 0x26, 0xbd,
	0x69, 0x7f, 0x3c, 0x19, 0xd6, 0xb5, 0x2e, 0xb9, 0xdb, 0x39, 0x68, 0xbb, 0x73, 0xd0, 0xc3, 0xce,
	0x41, 0x37, 0x7b, 0x47, 0xdb, 0xee, 0x1d, 0xed, 0x7e, 0xef, 0x68, 0x91, 0xa9, 0xfe, 0xfd, 0xfb,
	0x63, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xd2, 0x3d, 0x5b, 0x2e, 0x02, 0x00, 0x00,
}

func (m *BadEncoding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadEncoding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadEncoding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Axis != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Axis))
		i--
		dAtA[i] = 0x28
	}
	if m.Index != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProof(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.HeaderHash) > 0 {
		i -= len(m.HeaderHash)
		copy(dAtA[i:], m.HeaderHash)
		i = encodeVarintProof(dAtA, i, uint64(len(m.HeaderHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FraudMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FraudMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FraudMessageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequestedProofType) > 0 {
		dAtA2 := make([]byte, len(m.RequestedProofType)*10)
		var j1 int
		for _, num := range m.RequestedProofType {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintProof(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProofResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintProof(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FraudMessageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FraudMessageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FraudMessageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for iNdEx := len(m.Proofs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proofs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProof(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BadEncoding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovProof(uint64(m.Height))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	if m.Index != 0 {
		n += 1 + sovProof(uint64(m.Index))
	}
	if m.Axis != 0 {
		n += 1 + sovProof(uint64(m.Axis))
	}
	return n
}

func (m *FraudMessageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RequestedProofType) > 0 {
		l = 0
		for _, e := range m.RequestedProofType {
			l += sovProof(uint64(e))
		}
		n += 1 + sovProof(uint64(l)) + l
	}
	return n
}

func (m *ProofResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovProof(uint64(m.Type))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	return n
}

func (m *FraudMessageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for _, e := range m.Proofs {
			l = e.Size()
			n += 1 + l + sovProof(uint64(l))
		}
	}
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BadEncoding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: BadEncoding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadEncoding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, &pb.Share{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Axis", wireType)
			}
			m.Axis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Axis |= Axis(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func (m *FraudMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: FraudMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FraudMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v ProofType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProof
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ProofType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RequestedProofType = append(m.RequestedProofType, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProof
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProof
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProof
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.RequestedProofType) == 0 {
					m.RequestedProofType = make([]ProofType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ProofType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProof
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ProofType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RequestedProofType = append(m.RequestedProofType, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedProofType", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func (m *ProofResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: ProofResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ProofType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func (m *FraudMessageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: FraudMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FraudMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proofs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proofs = append(m.Proofs, &ProofResponse{})
			if err := m.Proofs[len(m.Proofs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)
