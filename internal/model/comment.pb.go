// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/model/comment.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CommentCountData struct {
	Like                 int32    `protobuf:"varint,1,opt,name=like,proto3" json:"like,omitempty"`
	Hate                 int32    `protobuf:"varint,2,opt,name=hate,proto3" json:"hate,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentCountData) Reset()         { *m = CommentCountData{} }
func (m *CommentCountData) String() string { return proto.CompactTextString(m) }
func (*CommentCountData) ProtoMessage()    {}
func (*CommentCountData) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e7792f7d8ba284, []int{0}
}
func (m *CommentCountData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommentCountData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommentCountData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommentCountData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentCountData.Merge(m, src)
}
func (m *CommentCountData) XXX_Size() int {
	return m.Size()
}
func (m *CommentCountData) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentCountData.DiscardUnknown(m)
}

var xxx_messageInfo_CommentCountData proto.InternalMessageInfo

func (m *CommentCountData) GetLike() int32 {
	if m != nil {
		return m.Like
	}
	return 0
}

func (m *CommentCountData) GetHate() int32 {
	if m != nil {
		return m.Hate
	}
	return 0
}

func (m *CommentCountData) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CommentMetaData struct {
	ObjId                int64    `protobuf:"varint,1,opt,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
	ObjType              int32    `protobuf:"varint,2,opt,name=obj_type,json=objType,proto3" json:"obj_type,omitempty"`
	Uid                  int64    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	CommentId            int64    `protobuf:"varint,4,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	Root                 int64    `protobuf:"varint,5,opt,name=root,proto3" json:"root,omitempty"`
	Parent               int64    `protobuf:"varint,6,opt,name=parent,proto3" json:"parent,omitempty"`
	Floor                int32    `protobuf:"varint,7,opt,name=floor,proto3" json:"floor,omitempty"`
	Ip                   int32    `protobuf:"varint,8,opt,name=ip,proto3" json:"ip,omitempty"`
	CreateTime           int64    `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Message              string   `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentMetaData) Reset()         { *m = CommentMetaData{} }
func (m *CommentMetaData) String() string { return proto.CompactTextString(m) }
func (*CommentMetaData) ProtoMessage()    {}
func (*CommentMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e7792f7d8ba284, []int{1}
}
func (m *CommentMetaData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommentMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommentMetaData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommentMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentMetaData.Merge(m, src)
}
func (m *CommentMetaData) XXX_Size() int {
	return m.Size()
}
func (m *CommentMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_CommentMetaData proto.InternalMessageInfo

func (m *CommentMetaData) GetObjId() int64 {
	if m != nil {
		return m.ObjId
	}
	return 0
}

func (m *CommentMetaData) GetObjType() int32 {
	if m != nil {
		return m.ObjType
	}
	return 0
}

func (m *CommentMetaData) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CommentMetaData) GetCommentId() int64 {
	if m != nil {
		return m.CommentId
	}
	return 0
}

func (m *CommentMetaData) GetRoot() int64 {
	if m != nil {
		return m.Root
	}
	return 0
}

func (m *CommentMetaData) GetParent() int64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *CommentMetaData) GetFloor() int32 {
	if m != nil {
		return m.Floor
	}
	return 0
}

func (m *CommentMetaData) GetIp() int32 {
	if m != nil {
		return m.Ip
	}
	return 0
}

func (m *CommentMetaData) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CommentMetaData) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CommentCountData)(nil), "model.CommentCountData")
	proto.RegisterType((*CommentMetaData)(nil), "model.CommentMetaData")
}

func init() { proto.RegisterFile("internal/model/comment.proto", fileDescriptor_92e7792f7d8ba284) }

var fileDescriptor_92e7792f7d8ba284 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x6d, 0x4b, 0x0b, 0x8c, 0x89, 0x92, 0x8d, 0x9a, 0x35, 0xd1, 0x4a, 0x38, 0x71, 0x92,
	0x83, 0x6f, 0x20, 0x5e, 0x38, 0x98, 0x98, 0x86, 0x3b, 0xd9, 0xd2, 0x51, 0x17, 0xdb, 0x4e, 0xb3,
	0x0c, 0x07, 0x8e, 0xbe, 0x85, 0x8f, 0xe4, 0xd1, 0x47, 0x30, 0xf8, 0x22, 0x66, 0xa7, 0xe5, 0xf6,
	0xff, 0xdf, 0xa4, 0x5f, 0xfe, 0x2e, 0xdc, 0xd8, 0x9a, 0xd1, 0xd5, 0xa6, 0x9c, 0x55, 0x54, 0x60,
	0x39, 0x5b, 0x53, 0x55, 0x61, 0xcd, 0xf7, 0x8d, 0x23, 0x26, 0x15, 0x0b, 0x9c, 0xbc, 0xc0, 0x68,
	0xde, 0xf2, 0x39, 0xed, 0x6a, 0x7e, 0x32, 0x6c, 0x94, 0x82, 0x5e, 0x69, 0x3f, 0x50, 0x07, 0xe3,
	0x60, 0x1a, 0x67, 0x92, 0x3d, 0x7b, 0x37, 0x8c, 0x3a, 0x6c, 0x99, 0xcf, 0xea, 0x02, 0xe2, 0xb5,
	0xff, 0x48, 0x47, 0x02, 0xdb, 0x32, 0xf9, 0x0c, 0xe1, 0xbc, 0x53, 0x3e, 0x23, 0x1b, 0x31, 0x5e,
	0x42, 0x42, 0xf9, 0x66, 0x65, 0x0b, 0x71, 0x46, 0x59, 0x4c, 0xf9, 0x66, 0x51, 0xa8, 0x6b, 0x18,
	0x78, 0xcc, 0xfb, 0xe6, 0x28, 0xee, 0x53, 0xbe, 0x59, 0xee, 0x1b, 0x54, 0x23, 0x88, 0x76, 0xb6,
	0x10, 0x73, 0x94, 0xf9, 0xa8, 0x6e, 0x01, 0xba, 0x3f, 0xf0, 0x9e, 0x9e, 0x1c, 0x86, 0x1d, 0x59,
	0x14, 0x7e, 0xa0, 0x23, 0x62, 0x1d, 0xcb, 0x41, 0xb2, 0xba, 0x82, 0xa4, 0x31, 0x0e, 0x6b, 0xd6,
	0x89, 0xd0, 0xae, 0xf9, 0xe1, 0xaf, 0x25, 0x91, 0xd3, 0xfd, 0x76, 0xb8, 0x14, 0x75, 0x06, 0xa1,
	0x6d, 0xf4, 0x40, 0x50, 0x68, 0x1b, 0x75, 0x07, 0xa7, 0x6b, 0x87, 0x86, 0x71, 0xc5, 0xb6, 0x42,
	0x3d, 0x14, 0x05, 0xb4, 0x68, 0x69, 0x2b, 0x54, 0x1a, 0xfa, 0x15, 0x6e, 0xb7, 0xe6, 0x0d, 0x35,
	0x8c, 0x83, 0xe9, 0x30, 0x3b, 0xd6, 0xc7, 0xd1, 0xf7, 0x21, 0x0d, 0x7e, 0x0e, 0x69, 0xf0, 0x7b,
	0x48, 0x83, 0xaf, 0xbf, 0xf4, 0x24, 0x4f, 0xe4, 0xd5, 0x1f, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x93, 0x71, 0x4b, 0x33, 0x95, 0x01, 0x00, 0x00,
}

func (m *CommentCountData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommentCountData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommentCountData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Count != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x18
	}
	if m.Hate != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Hate))
		i--
		dAtA[i] = 0x10
	}
	if m.Like != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Like))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CommentMetaData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommentMetaData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommentMetaData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintComment(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x52
	}
	if m.CreateTime != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.CreateTime))
		i--
		dAtA[i] = 0x48
	}
	if m.Ip != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Ip))
		i--
		dAtA[i] = 0x40
	}
	if m.Floor != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Floor))
		i--
		dAtA[i] = 0x38
	}
	if m.Parent != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Parent))
		i--
		dAtA[i] = 0x30
	}
	if m.Root != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Root))
		i--
		dAtA[i] = 0x28
	}
	if m.CommentId != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.CommentId))
		i--
		dAtA[i] = 0x20
	}
	if m.Uid != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x18
	}
	if m.ObjType != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.ObjType))
		i--
		dAtA[i] = 0x10
	}
	if m.ObjId != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.ObjId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintComment(dAtA []byte, offset int, v uint64) int {
	offset -= sovComment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommentCountData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Like != 0 {
		n += 1 + sovComment(uint64(m.Like))
	}
	if m.Hate != 0 {
		n += 1 + sovComment(uint64(m.Hate))
	}
	if m.Count != 0 {
		n += 1 + sovComment(uint64(m.Count))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommentMetaData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ObjId != 0 {
		n += 1 + sovComment(uint64(m.ObjId))
	}
	if m.ObjType != 0 {
		n += 1 + sovComment(uint64(m.ObjType))
	}
	if m.Uid != 0 {
		n += 1 + sovComment(uint64(m.Uid))
	}
	if m.CommentId != 0 {
		n += 1 + sovComment(uint64(m.CommentId))
	}
	if m.Root != 0 {
		n += 1 + sovComment(uint64(m.Root))
	}
	if m.Parent != 0 {
		n += 1 + sovComment(uint64(m.Parent))
	}
	if m.Floor != 0 {
		n += 1 + sovComment(uint64(m.Floor))
	}
	if m.Ip != 0 {
		n += 1 + sovComment(uint64(m.Ip))
	}
	if m.CreateTime != 0 {
		n += 1 + sovComment(uint64(m.CreateTime))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComment(x uint64) (n int) {
	return sovComment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommentCountData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
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
			return fmt.Errorf("proto: CommentCountData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommentCountData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Like", wireType)
			}
			m.Like = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Like |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hate", wireType)
			}
			m.Hate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hate |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommentMetaData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
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
			return fmt.Errorf("proto: CommentMetaData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommentMetaData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjId", wireType)
			}
			m.ObjId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObjId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjType", wireType)
			}
			m.ObjType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObjType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommentId", wireType)
			}
			m.CommentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommentId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			m.Root = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Root |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			m.Parent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Parent |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Floor", wireType)
			}
			m.Floor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Floor |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			m.Ip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ip |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTime", wireType)
			}
			m.CreateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipComment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComment
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
					return 0, ErrIntOverflowComment
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
					return 0, ErrIntOverflowComment
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
				return 0, ErrInvalidLengthComment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComment = fmt.Errorf("proto: unexpected end of group")
)
