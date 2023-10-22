// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: share/ipldv2/pb/ipldv2pb.proto

package ipldv2pb

import (
	fmt "fmt"
	pb "github.com/celestiaorg/nmt/pb"
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

type AxisType int32

const (
	AxisType_Row AxisType = 0
	AxisType_Col AxisType = 1
)

var AxisType_name = map[int32]string{
	0: "Row",
	1: "Col",
}

var AxisType_value = map[string]int32{
	"Row": 0,
	"Col": 1,
}

func (x AxisType) String() string {
	return proto.EnumName(AxisType_name, int32(x))
}

func (AxisType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{0}
}

type SampleType int32

const (
	SampleType_Data   SampleType = 0
	SampleType_Parity SampleType = 1
)

var SampleType_name = map[int32]string{
	0: "Data",
	1: "Parity",
}

var SampleType_value = map[string]int32{
	"Data":   0,
	"Parity": 1,
}

func (x SampleType) String() string {
	return proto.EnumName(SampleType_name, int32(x))
}

func (SampleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{1}
}

type Axis struct {
	AxisId   []byte   `protobuf:"bytes,1,opt,name=axis_id,json=axisId,proto3" json:"axis_id,omitempty"`
	AxisHalf [][]byte `protobuf:"bytes,2,rep,name=axis_half,json=axisHalf,proto3" json:"axis_half,omitempty"`
}

func (m *Axis) Reset()         { *m = Axis{} }
func (m *Axis) String() string { return proto.CompactTextString(m) }
func (*Axis) ProtoMessage()    {}
func (*Axis) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{0}
}
func (m *Axis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Axis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Axis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Axis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Axis.Merge(m, src)
}
func (m *Axis) XXX_Size() int {
	return m.Size()
}
func (m *Axis) XXX_DiscardUnknown() {
	xxx_messageInfo_Axis.DiscardUnknown(m)
}

var xxx_messageInfo_Axis proto.InternalMessageInfo

func (m *Axis) GetAxisId() []byte {
	if m != nil {
		return m.AxisId
	}
	return nil
}

func (m *Axis) GetAxisHalf() [][]byte {
	if m != nil {
		return m.AxisHalf
	}
	return nil
}

type Sample struct {
	Id    []byte     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type  SampleType `protobuf:"varint,2,opt,name=type,proto3,enum=SampleType" json:"type,omitempty"`
	Share []byte     `protobuf:"bytes,3,opt,name=share,proto3" json:"share,omitempty"`
	Proof *pb.Proof  `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{1}
}
func (m *Sample) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
}
func (m *Sample) XXX_Size() int {
	return m.Size()
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Sample) GetType() SampleType {
	if m != nil {
		return m.Type
	}
	return SampleType_Data
}

func (m *Sample) GetShare() []byte {
	if m != nil {
		return m.Share
	}
	return nil
}

func (m *Sample) GetProof() *pb.Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterEnum("AxisType", AxisType_name, AxisType_value)
	proto.RegisterEnum("SampleType", SampleType_name, SampleType_value)
	proto.RegisterType((*Axis)(nil), "Axis")
	proto.RegisterType((*Sample)(nil), "Sample")
}

func init() { proto.RegisterFile("share/ipldv2/pb/ipldv2pb.proto", fileDescriptor_cb41c3a4f982a271) }

var fileDescriptor_cb41c3a4f982a271 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4f, 0x4b, 0x84, 0x40,
	0x18, 0xc6, 0x1d, 0x75, 0x5d, 0x7b, 0x77, 0x31, 0x19, 0x82, 0x86, 0x8a, 0x49, 0x84, 0x40, 0xf6,
	0xa0, 0x60, 0xd7, 0x2e, 0xfd, 0x39, 0xd4, 0x6d, 0xb1, 0xee, 0x31, 0xa2, 0xb2, 0x03, 0xc6, 0x0c,
	0x2a, 0xdb, 0xfa, 0x2d, 0xfa, 0x58, 0x1d, 0xf7, 0xd8, 0x31, 0xf4, 0x8b, 0xc4, 0x8c, 0x4b, 0x7b,
	0x7b, 0x9e, 0x79, 0xe6, 0x79, 0xdf, 0x1f, 0x2f, 0xd0, 0x76, 0xc3, 0x9a, 0x32, 0xe1, 0xb2, 0x2e,
	0xb6, 0x69, 0x22, 0xf3, 0x83, 0x92, 0x79, 0x2c, 0x1b, 0xd1, 0x89, 0x0b, 0x4f, 0xe6, 0x89, 0x6c,
	0x84, 0xa8, 0x26, 0x1f, 0xde, 0x81, 0x7d, 0xbf, 0xe3, 0x2d, 0x3e, 0x87, 0x39, 0xdb, 0xf1, 0xf6,
	0x9d, 0x17, 0x04, 0x05, 0x28, 0x5a, 0x66, 0x8e, 0xb2, 0x2f, 0x05, 0xbe, 0x84, 0x13, 0x1d, 0x6c,
	0x58, 0x5d, 0x11, 0x33, 0xb0, 0xa2, 0x65, 0xe6, 0xaa, 0x87, 0x67, 0x56, 0x57, 0xe1, 0x16, 0x9c,
	0x57, 0xf6, 0x21, 0xeb, 0x12, 0x7b, 0x60, 0xfe, 0x57, 0x4d, 0x5e, 0xe0, 0x6b, 0xb0, 0xbb, 0x5e,
	0x96, 0xc4, 0x0c, 0x50, 0xe4, 0xa5, 0x8b, 0x78, 0xfa, 0xf6, 0xd6, 0xcb, 0x32, 0xd3, 0x01, 0x3e,
	0x83, 0x99, 0x46, 0x25, 0x96, 0xee, 0x4c, 0x06, 0xdf, 0xc0, 0x4c, 0xd3, 0x11, 0x3b, 0x40, 0xd1,
	0x22, 0x3d, 0x8d, 0x0f, 0xac, 0x79, 0xbc, 0x56, 0x22, 0x9b, 0xd2, 0xd5, 0x15, 0xb8, 0x8a, 0x5a,
	0x8d, 0xc3, 0x73, 0xb0, 0x32, 0xf1, 0xe9, 0x1b, 0x4a, 0x3c, 0x8a, 0xda, 0x47, 0xab, 0x10, 0xe0,
	0xb8, 0x0e, 0xbb, 0x60, 0x3f, 0xb1, 0x8e, 0xf9, 0x06, 0x06, 0x70, 0xd6, 0xac, 0xe1, 0x5d, 0xef,
	0xa3, 0x07, 0xf2, 0x3d, 0x50, 0xb4, 0x1f, 0x28, 0xfa, 0x1d, 0x28, 0xfa, 0x1a, 0xa9, 0xb1, 0x1f,
	0xa9, 0xf1, 0x33, 0x52, 0x23, 0x77, 0xf4, 0x61, 0x6e, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2e,
	0x20, 0xe3, 0x8f, 0x4a, 0x01, 0x00, 0x00,
}

func (m *Axis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Axis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Axis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AxisHalf) > 0 {
		for iNdEx := len(m.AxisHalf) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AxisHalf[iNdEx])
			copy(dAtA[i:], m.AxisHalf[iNdEx])
			i = encodeVarintIpldv2Pb(dAtA, i, uint64(len(m.AxisHalf[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AxisId) > 0 {
		i -= len(m.AxisId)
		copy(dAtA[i:], m.AxisId)
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(len(m.AxisId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Sample) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sample) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sample) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintIpldv2Pb(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Share) > 0 {
		i -= len(m.Share)
		copy(dAtA[i:], m.Share)
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(len(m.Share)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIpldv2Pb(dAtA []byte, offset int, v uint64) int {
	offset -= sovIpldv2Pb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Axis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AxisId)
	if l > 0 {
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	if len(m.AxisHalf) > 0 {
		for _, b := range m.AxisHalf {
			l = len(b)
			n += 1 + l + sovIpldv2Pb(uint64(l))
		}
	}
	return n
}

func (m *Sample) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovIpldv2Pb(uint64(m.Type))
	}
	l = len(m.Share)
	if l > 0 {
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	return n
}

func sovIpldv2Pb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIpldv2Pb(x uint64) (n int) {
	return sovIpldv2Pb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Axis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpldv2Pb
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
			return fmt.Errorf("proto: Axis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Axis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AxisId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AxisId = append(m.AxisId[:0], dAtA[iNdEx:postIndex]...)
			if m.AxisId == nil {
				m.AxisId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AxisHalf", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AxisHalf = append(m.AxisHalf, make([]byte, postIndex-iNdEx))
			copy(m.AxisHalf[len(m.AxisHalf)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpldv2Pb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpldv2Pb
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
func (m *Sample) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpldv2Pb
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
			return fmt.Errorf("proto: Sample: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sample: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SampleType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Share", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Share = append(m.Share[:0], dAtA[iNdEx:postIndex]...)
			if m.Share == nil {
				m.Share = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &pb.Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpldv2Pb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpldv2Pb
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
func skipIpldv2Pb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIpldv2Pb
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
					return 0, ErrIntOverflowIpldv2Pb
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
					return 0, ErrIntOverflowIpldv2Pb
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
				return 0, ErrInvalidLengthIpldv2Pb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIpldv2Pb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIpldv2Pb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIpldv2Pb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIpldv2Pb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIpldv2Pb = fmt.Errorf("proto: unexpected end of group")
)
