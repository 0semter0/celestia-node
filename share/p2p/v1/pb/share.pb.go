// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: share/p2p/v1/pb/share.proto

package share_p2p_v1

import (
	fmt "fmt"
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

type StatusCode int32

const (
	StatusCode_INVALID          StatusCode = 0
	StatusCode_OK               StatusCode = 1
	StatusCode_NOT_FOUND        StatusCode = 2
	StatusCode_INVALID_ARGUMENT StatusCode = 3
	StatusCode_INTERNAL         StatusCode = 4
)

var StatusCode_name = map[int32]string{
	0: "INVALID",
	1: "OK",
	2: "NOT_FOUND",
	3: "INVALID_ARGUMENT",
	4: "INTERNAL",
}

var StatusCode_value = map[string]int32{
	"INVALID":          0,
	"OK":               1,
	"NOT_FOUND":        2,
	"INVALID_ARGUMENT": 3,
	"INTERNAL":         4,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12ce7bddbfd98489, []int{0}
}

type GetSharesByNamespaceRequest struct {
	// Request network head by passing 0 in height field.
	Height      int64  `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	NamespaceId []byte `protobuf:"bytes,2,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	WithProofs  bool   `protobuf:"varint,3,opt,name=with_proofs,json=withProofs,proto3" json:"with_proofs,omitempty"`
}

func (m *GetSharesByNamespaceRequest) Reset()         { *m = GetSharesByNamespaceRequest{} }
func (m *GetSharesByNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*GetSharesByNamespaceRequest) ProtoMessage()    {}
func (*GetSharesByNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_12ce7bddbfd98489, []int{0}
}
func (m *GetSharesByNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSharesByNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSharesByNamespaceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSharesByNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSharesByNamespaceRequest.Merge(m, src)
}
func (m *GetSharesByNamespaceRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetSharesByNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSharesByNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSharesByNamespaceRequest proto.InternalMessageInfo

func (m *GetSharesByNamespaceRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetSharesByNamespaceRequest) GetNamespaceId() []byte {
	if m != nil {
		return m.NamespaceId
	}
	return nil
}

func (m *GetSharesByNamespaceRequest) GetWithProofs() bool {
	if m != nil {
		return m.WithProofs
	}
	return false
}

type GetSharesByNamespaceResponse struct {
	Status StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=share.p2p.v1.StatusCode" json:"status,omitempty"`
	Rows   []*Row     `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (m *GetSharesByNamespaceResponse) Reset()         { *m = GetSharesByNamespaceResponse{} }
func (m *GetSharesByNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*GetSharesByNamespaceResponse) ProtoMessage()    {}
func (*GetSharesByNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_12ce7bddbfd98489, []int{1}
}
func (m *GetSharesByNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSharesByNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSharesByNamespaceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSharesByNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSharesByNamespaceResponse.Merge(m, src)
}
func (m *GetSharesByNamespaceResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetSharesByNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSharesByNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSharesByNamespaceResponse proto.InternalMessageInfo

func (m *GetSharesByNamespaceResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_INVALID
}

func (m *GetSharesByNamespaceResponse) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Row struct {
	Shares [][]byte `protobuf:"bytes,1,rep,name=shares,proto3" json:"shares,omitempty"`
	Proof  *Proof   `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_12ce7bddbfd98489, []int{2}
}
func (m *Row) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Row.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return m.Size()
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetShares() [][]byte {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *Row) GetProof() *Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type Proof struct {
	Start int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Nodes [][]byte `protobuf:"bytes,3,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_12ce7bddbfd98489, []int{3}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Proof) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Proof) GetNodes() [][]byte {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterEnum("share.p2p.v1.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*GetSharesByNamespaceRequest)(nil), "share.p2p.v1.GetSharesByNamespaceRequest")
	proto.RegisterType((*GetSharesByNamespaceResponse)(nil), "share.p2p.v1.GetSharesByNamespaceResponse")
	proto.RegisterType((*Row)(nil), "share.p2p.v1.Row")
	proto.RegisterType((*Proof)(nil), "share.p2p.v1.Proof")
}

func init() { proto.RegisterFile("share/p2p/v1/pb/share.proto", fileDescriptor_12ce7bddbfd98489) }

var fileDescriptor_12ce7bddbfd98489 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0x94, 0x40,
	0x1c, 0xc7, 0x77, 0x76, 0xba, 0x58, 0x7f, 0xa0, 0xc1, 0xb1, 0x31, 0x24, 0x35, 0x88, 0x24, 0x26,
	0xe8, 0x01, 0x2c, 0x3e, 0xc1, 0xd6, 0xae, 0x95, 0x58, 0x67, 0xcd, 0x94, 0x7a, 0x25, 0x54, 0x46,
	0xd9, 0x83, 0x3b, 0x23, 0x33, 0x2d, 0xe9, 0x5b, 0xf8, 0x58, 0x1e, 0x7b, 0xf4, 0x68, 0x76, 0x5f,
	0xc4, 0xcc, 0x40, 0xd5, 0x26, 0xde, 0xf8, 0x7c, 0xe7, 0x93, 0xdf, 0xbf, 0x00, 0xfb, 0xaa, 0xad,
	0x3b, 0x9e, 0xc9, 0x5c, 0x66, 0x97, 0x07, 0x99, 0x3c, 0xcf, 0x2c, 0xa7, 0xb2, 0x13, 0x5a, 0x10,
	0x6f, 0x84, 0x5c, 0xa6, 0x97, 0x07, 0xf1, 0x15, 0xec, 0x1f, 0x73, 0x7d, 0x6a, 0x22, 0x75, 0x78,
	0x45, 0xeb, 0xaf, 0x5c, 0xc9, 0xfa, 0x13, 0x67, 0xfc, 0xdb, 0x05, 0x57, 0x9a, 0x3c, 0x02, 0xa7,
	0xe5, 0xab, 0x2f, 0xad, 0x0e, 0x50, 0x84, 0x12, 0xcc, 0x46, 0x22, 0x4f, 0xc1, 0x5b, 0xdf, 0xb8,
	0xd5, 0xaa, 0x09, 0xa6, 0x11, 0x4a, 0x3c, 0xe6, 0xfe, 0xc9, 0x8a, 0x86, 0x3c, 0x01, 0xb7, 0x5f,
	0xe9, 0xb6, 0x92, 0x9d, 0x10, 0x9f, 0x55, 0x80, 0x23, 0x94, 0xec, 0x32, 0x30, 0xd1, 0x07, 0x9b,
	0xc4, 0x3d, 0x3c, 0xfe, 0x7f, 0x6b, 0x25, 0xc5, 0x5a, 0x71, 0xf2, 0x12, 0x1c, 0xa5, 0x6b, 0x7d,
	0xa1, 0x6c, 0xef, 0xfb, 0x79, 0x90, 0xfe, 0x3b, 0x79, 0x7a, 0x6a, 0xdf, 0x5e, 0x8b, 0x86, 0xb3,
	0xd1, 0x23, 0xcf, 0x60, 0xa7, 0x13, 0xbd, 0x0a, 0xa6, 0x11, 0x4e, 0xdc, 0xfc, 0xc1, 0x6d, 0x9f,
	0x89, 0x9e, 0xd9, 0xe7, 0xf8, 0x2d, 0x60, 0x26, 0x7a, 0xb3, 0x9b, 0x15, 0x4c, 0x7d, 0x9c, 0x78,
	0x6c, 0x24, 0xf2, 0x1c, 0x66, 0x76, 0x66, 0xbb, 0x94, 0x9b, 0x3f, 0xbc, 0x5d, 0xc6, 0x0e, 0xcf,
	0x06, 0x23, 0x5e, 0xc0, 0xcc, 0x32, 0xd9, 0x83, 0x99, 0xd2, 0x75, 0x77, 0x73, 0xa6, 0x01, 0x88,
	0x0f, 0x98, 0xaf, 0x87, 0xe3, 0x60, 0x66, 0x3e, 0x8d, 0x47, 0x45, 0xc3, 0xcd, 0x39, 0x4c, 0xcb,
	0x01, 0x5e, 0x94, 0x00, 0x7f, 0xb7, 0x21, 0x2e, 0xdc, 0x29, 0xe8, 0xc7, 0xf9, 0x49, 0x71, 0xe4,
	0x4f, 0x88, 0x03, 0xd3, 0xe5, 0x3b, 0x1f, 0x91, 0x7b, 0x70, 0x97, 0x2e, 0xcb, 0xea, 0xcd, 0xf2,
	0x8c, 0x1e, 0xf9, 0x53, 0xb2, 0x07, 0xfe, 0xe8, 0x54, 0x73, 0x76, 0x7c, 0xf6, 0x7e, 0x41, 0x4b,
	0x1f, 0x13, 0x0f, 0x76, 0x0b, 0x5a, 0x2e, 0x18, 0x9d, 0x9f, 0xf8, 0x3b, 0x87, 0xc1, 0x8f, 0x4d,
	0x88, 0xae, 0x37, 0x21, 0xfa, 0xb5, 0x09, 0xd1, 0xf7, 0x6d, 0x38, 0xb9, 0xde, 0x86, 0x93, 0x9f,
	0xdb, 0x70, 0x72, 0xee, 0xd8, 0x3f, 0xe1, 0xd5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x7e,
	0x37, 0x27, 0x28, 0x02, 0x00, 0x00,
}

func (m *GetSharesByNamespaceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSharesByNamespaceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSharesByNamespaceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithProofs {
		i--
		if m.WithProofs {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.NamespaceId) > 0 {
		i -= len(m.NamespaceId)
		copy(dAtA[i:], m.NamespaceId)
		i = encodeVarintShare(dAtA, i, uint64(len(m.NamespaceId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetSharesByNamespaceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSharesByNamespaceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSharesByNamespaceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rows) > 0 {
		for iNdEx := len(m.Rows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShare(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Status != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Row) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Row) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Row) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintShare(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Shares[iNdEx])
			copy(dAtA[i:], m.Shares[iNdEx])
			i = encodeVarintShare(dAtA, i, uint64(len(m.Shares[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Nodes[iNdEx])
			copy(dAtA[i:], m.Nodes[iNdEx])
			i = encodeVarintShare(dAtA, i, uint64(len(m.Nodes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.End != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x10
	}
	if m.Start != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetSharesByNamespaceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovShare(uint64(m.Height))
	}
	l = len(m.NamespaceId)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	if m.WithProofs {
		n += 2
	}
	return n
}

func (m *GetSharesByNamespaceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovShare(uint64(m.Status))
	}
	if len(m.Rows) > 0 {
		for _, e := range m.Rows {
			l = e.Size()
			n += 1 + l + sovShare(uint64(l))
		}
	}
	return n
}

func (m *Row) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Shares) > 0 {
		for _, b := range m.Shares {
			l = len(b)
			n += 1 + l + sovShare(uint64(l))
		}
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovShare(uint64(l))
	}
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovShare(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovShare(uint64(m.End))
	}
	if len(m.Nodes) > 0 {
		for _, b := range m.Nodes {
			l = len(b)
			n += 1 + l + sovShare(uint64(l))
		}
	}
	return n
}

func sovShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShare(x uint64) (n int) {
	return sovShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetSharesByNamespaceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: GetSharesByNamespaceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSharesByNamespaceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceId = append(m.NamespaceId[:0], dAtA[iNdEx:postIndex]...)
			if m.NamespaceId == nil {
				m.NamespaceId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithProofs", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithProofs = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *GetSharesByNamespaceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: GetSharesByNamespaceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSharesByNamespaceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= StatusCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rows = append(m.Rows, &Row{})
			if err := m.Rows[len(m.Rows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *Row) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: Row: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Row: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, make([]byte, postIndex-iNdEx))
			copy(m.Shares[len(m.Shares)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, make([]byte, postIndex-iNdEx))
			copy(m.Nodes[len(m.Nodes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func skipShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
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
				return 0, ErrInvalidLengthShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShare = fmt.Errorf("proto: unexpected end of group")
)
