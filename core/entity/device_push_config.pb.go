// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/device_push_config.proto

package entity

import (
	fmt "fmt"
	io "io"
	math "math"
	time "time"

	push "berty.tech/core/push"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Local unencrypted push configuration
type DevicePushConfig struct {
	ID                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt            time.Time           `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	UpdatedAt            time.Time           `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at"`
	DeviceID             string              `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	PushType             push.DevicePushType `protobuf:"varint,5,opt,name=push_type,json=pushType,proto3,enum=berty.push.DevicePushType" json:"push_type,omitempty"`
	PushID               []byte              `protobuf:"bytes,6,opt,name=push_id,json=pushId,proto3" json:"push_id,omitempty"`
	RelayPubkey          string              `protobuf:"bytes,7,opt,name=relay_pubkey,json=relayPubkey,proto3" json:"relay_pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DevicePushConfig) Reset()         { *m = DevicePushConfig{} }
func (m *DevicePushConfig) String() string { return proto.CompactTextString(m) }
func (*DevicePushConfig) ProtoMessage()    {}
func (*DevicePushConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_469c6c4d4fab077f, []int{0}
}
func (m *DevicePushConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DevicePushConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DevicePushConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DevicePushConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevicePushConfig.Merge(m, src)
}
func (m *DevicePushConfig) XXX_Size() int {
	return m.Size()
}
func (m *DevicePushConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DevicePushConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DevicePushConfig proto.InternalMessageInfo

func (m *DevicePushConfig) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *DevicePushConfig) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *DevicePushConfig) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func (m *DevicePushConfig) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *DevicePushConfig) GetPushType() push.DevicePushType {
	if m != nil {
		return m.PushType
	}
	return push.DevicePushType_UnknownDevicePushType
}

func (m *DevicePushConfig) GetPushID() []byte {
	if m != nil {
		return m.PushID
	}
	return nil
}

func (m *DevicePushConfig) GetRelayPubkey() string {
	if m != nil {
		return m.RelayPubkey
	}
	return ""
}

func init() {
	proto.RegisterType((*DevicePushConfig)(nil), "berty.entity.DevicePushConfig")
}

func init() { proto.RegisterFile("entity/device_push_config.proto", fileDescriptor_469c6c4d4fab077f) }

var fileDescriptor_469c6c4d4fab077f = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x6e, 0x9c, 0x30,
	0x14, 0xc6, 0xe3, 0x49, 0x3b, 0x01, 0x0f, 0x6a, 0x2a, 0xab, 0xaa, 0x10, 0xaa, 0x30, 0x9a, 0x6e,
	0xa8, 0x54, 0x81, 0x34, 0x5d, 0x54, 0xea, 0x2e, 0x84, 0x0d, 0xbb, 0x08, 0x65, 0xd5, 0x0d, 0x02,
	0xec, 0x30, 0x56, 0x42, 0x6c, 0x79, 0x4c, 0x25, 0x9f, 0xa0, 0xdb, 0x2e, 0x7b, 0xa4, 0x2c, 0x7b,
	0x02, 0x5a, 0xd1, 0x1b, 0xf4, 0x04, 0x15, 0x36, 0xa3, 0x59, 0x67, 0x83, 0xde, 0x9f, 0xef, 0xbd,
	0xef, 0xf1, 0x33, 0xc4, 0xf4, 0x51, 0x31, 0xa5, 0x53, 0x42, 0xbf, 0xb1, 0x96, 0x56, 0x62, 0x38,
	0xec, 0xab, 0x96, 0x3f, 0xde, 0xb1, 0x2e, 0x11, 0x92, 0x2b, 0x8e, 0xbc, 0x86, 0x4a, 0xa5, 0x13,
	0x2b, 0x0b, 0x2e, 0x67, 0x41, 0x3a, 0x7f, 0x6c, 0x3b, 0x78, 0xd3, 0xf1, 0x8e, 0x9b, 0x30, 0x9d,
	0xa3, 0xa5, 0x8a, 0x3b, 0xce, 0xbb, 0x07, 0x9a, 0x9a, 0xac, 0x19, 0xee, 0x52, 0xc5, 0x7a, 0x7a,
	0x50, 0x75, 0x2f, 0xac, 0x60, 0xfb, 0xfd, 0x1c, 0xbe, 0xce, 0x8d, 0xe5, 0xcd, 0x70, 0xd8, 0x5f,
	0x1b, 0x43, 0xf4, 0x11, 0xae, 0x18, 0xf1, 0x41, 0x04, 0x62, 0x37, 0x7b, 0x37, 0x8d, 0x78, 0x55,
	0xe4, 0xff, 0x46, 0x8c, 0x3a, 0x2e, 0xfb, 0x2f, 0x5b, 0x21, 0x59, 0x5f, 0x4b, 0x5d, 0xdd, 0x53,
	0xbd, 0x2d, 0x57, 0x8c, 0xa0, 0x6b, 0x08, 0x5b, 0x49, 0x6b, 0x45, 0x49, 0x55, 0x2b, 0x7f, 0x15,
	0x81, 0x78, 0xb3, 0x0b, 0x12, 0x6b, 0x9c, 0x1c, 0x8d, 0x93, 0xdb, 0xa3, 0x71, 0xe6, 0x3c, 0x8d,
	0xf8, 0xec, 0xc7, 0x6f, 0x0c, 0x4a, 0x77, 0x99, 0xbb, 0x52, 0xf3, 0x92, 0x41, 0x90, 0xe3, 0x92,
	0xf3, 0xe7, 0x2c, 0x59, 0xe6, 0xae, 0x14, 0xfa, 0x00, 0xdd, 0x05, 0x1f, 0x23, 0xfe, 0x0b, 0x73,
	0xbe, 0x37, 0x8d, 0xd8, 0xb1, 0x3f, 0x58, 0xe4, 0xa5, 0x63, 0xdb, 0x05, 0x41, 0x9f, 0xa1, 0x6b,
	0x10, 0x2b, 0x2d, 0xa8, 0xff, 0x32, 0x02, 0xf1, 0xab, 0x5d, 0x90, 0x58, 0xc2, 0x06, 0xea, 0x89,
	0xc9, 0xad, 0x16, 0xb4, 0x74, 0xc4, 0x12, 0xa1, 0xf7, 0xf0, 0xc2, 0x0c, 0x32, 0xe2, 0xaf, 0x23,
	0x10, 0x7b, 0x19, 0x9c, 0x46, 0xbc, 0x9e, 0x85, 0x45, 0x5e, 0xae, 0xe7, 0x56, 0x41, 0xd0, 0x0e,
	0x7a, 0x92, 0x3e, 0xd4, 0xba, 0x12, 0x43, 0x73, 0x4f, 0xb5, 0x7f, 0x61, 0x6e, 0xb9, 0x9c, 0x46,
	0xbc, 0x29, 0xe7, 0xfa, 0x8d, 0x29, 0x97, 0x1b, 0x79, 0x4a, 0xb2, 0xf8, 0x69, 0x0a, 0xc1, 0xaf,
	0x29, 0x04, 0x7f, 0xa6, 0x10, 0xfc, 0xfc, 0x1b, 0x9e, 0x7d, 0x7d, 0x6b, 0xef, 0x51, 0xb4, 0xdd,
	0xa7, 0x2d, 0x97, 0x34, 0xb5, 0x6f, 0xdf, 0xac, 0x0d, 0x8f, 0x4f, 0xff, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x65, 0x47, 0x9a, 0xe8, 0x33, 0x02, 0x00, 0x00,
}

func (m *DevicePushConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DevicePushConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintDevicePushConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err1 != nil {
		return 0, err1
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintDevicePushConfig(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)))
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err2 != nil {
		return 0, err2
	}
	i += n2
	if len(m.DeviceID) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.DeviceID)))
		i += copy(dAtA[i:], m.DeviceID)
	}
	if m.PushType != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(m.PushType))
	}
	if len(m.PushID) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.PushID)))
		i += copy(dAtA[i:], m.PushID)
	}
	if len(m.RelayPubkey) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintDevicePushConfig(dAtA, i, uint64(len(m.RelayPubkey)))
		i += copy(dAtA[i:], m.RelayPubkey)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDevicePushConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DevicePushConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovDevicePushConfig(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovDevicePushConfig(uint64(l))
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	if m.PushType != 0 {
		n += 1 + sovDevicePushConfig(uint64(m.PushType))
	}
	l = len(m.PushID)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	l = len(m.RelayPubkey)
	if l > 0 {
		n += 1 + l + sovDevicePushConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDevicePushConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDevicePushConfig(x uint64) (n int) {
	return sovDevicePushConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DevicePushConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevicePushConfig
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
			return fmt.Errorf("proto: DevicePushConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DevicePushConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
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
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
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
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
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
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
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
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushType", wireType)
			}
			m.PushType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PushType |= push.DevicePushType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
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
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PushID = append(m.PushID[:0], dAtA[iNdEx:postIndex]...)
			if m.PushID == nil {
				m.PushID = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayPubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevicePushConfig
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
				return ErrInvalidLengthDevicePushConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayPubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevicePushConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDevicePushConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDevicePushConfig
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
func skipDevicePushConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDevicePushConfig
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
					return 0, ErrIntOverflowDevicePushConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDevicePushConfig
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
				return 0, ErrInvalidLengthDevicePushConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDevicePushConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDevicePushConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDevicePushConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDevicePushConfig
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDevicePushConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDevicePushConfig   = fmt.Errorf("proto: integer overflow")
)
