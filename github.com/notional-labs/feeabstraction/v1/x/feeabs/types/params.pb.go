// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeabstraction/absfee/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Params defines the parameters for the feeabs module.
type Params struct {
	// allowed
	AllowedToken string `protobuf:"bytes,1,opt,name=allowed_token,json=allowedToken,proto3" json:"allowed_token,omitempty" yaml:"allowed_token"`
	// this is query period
	QueryPeriod time.Duration `protobuf:"bytes,2,opt,name=query_period,json=queryPeriod,proto3,stdduration" json:"query_period"`
	// this is swap period
	SwapPeriod time.Duration `protobuf:"bytes,3,opt,name=swap_period,json=swapPeriod,proto3,stdduration" json:"swap_period"`
	// contract address from fee provider side
	ContractAddress string `protobuf:"bytes,4,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e66a0978c84086, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAllowedToken() string {
	if m != nil {
		return m.AllowedToken
	}
	return ""
}

func (m *Params) GetQueryPeriod() time.Duration {
	if m != nil {
		return m.QueryPeriod
	}
	return 0
}

func (m *Params) GetSwapPeriod() time.Duration {
	if m != nil {
		return m.SwapPeriod
	}
	return 0
}

func (m *Params) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "feeabstraction.absfee.v1beta1.Params")
}

func init() {
	proto.RegisterFile("feeabstraction/absfee/v1beta1/params.proto", fileDescriptor_64e66a0978c84086)
}

var fileDescriptor_64e66a0978c84086 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x3b, 0xf7, 0xde, 0x90, 0x6b, 0xc1, 0x98, 0x34, 0x2c, 0x2a, 0xd1, 0x42, 0x58, 0x11,
	0x12, 0x3b, 0x41, 0x57, 0x9a, 0xb8, 0x90, 0x10, 0xd7, 0x04, 0x5d, 0xb9, 0x21, 0x67, 0xda, 0xa1,
	0x36, 0xb6, 0x3d, 0xb5, 0x33, 0x05, 0x79, 0x0b, 0x97, 0x3e, 0x12, 0x4b, 0x96, 0xae, 0xd0, 0xc0,
	0x1b, 0xf8, 0x02, 0x9a, 0x0e, 0x43, 0x02, 0xac, 0xdc, 0xcd, 0x39, 0xe7, 0xff, 0xff, 0xf3, 0xcd,
	0x8c, 0xd9, 0x1e, 0x71, 0x0e, 0x4c, 0xc8, 0x0c, 0x3c, 0x19, 0x62, 0x42, 0x81, 0x89, 0x11, 0xe7,
	0x74, 0xdc, 0x61, 0x5c, 0x42, 0x87, 0xa6, 0x90, 0x41, 0x2c, 0xdc, 0x34, 0x43, 0x89, 0xd6, 0xe9,
	0xae, 0xd6, 0x5d, 0x6b, 0x5d, 0xad, 0xad, 0x55, 0x03, 0x0c, 0x50, 0x29, 0x69, 0x71, 0x5a, 0x9b,
	0x6a, 0x27, 0x01, 0x62, 0x10, 0x71, 0x0a, 0x69, 0x48, 0x21, 0x49, 0x50, 0x42, 0xe1, 0xd5, 0x91,
	0xb5, 0xb6, 0x87, 0x22, 0x46, 0x41, 0x19, 0x08, 0x4e, 0x9f, 0x73, 0x9e, 0x4d, 0xb7, 0x56, 0x07,
	0x61, 0xa2, 0xc4, 0x5a, 0xeb, 0xe8, 0x24, 0x55, 0xb1, 0x7c, 0x44, 0xfd, 0x3c, 0xdb, 0x9a, 0x37,
	0xbf, 0x89, 0x59, 0xea, 0x2b, 0x5e, 0xeb, 0xda, 0x3c, 0x84, 0x28, 0xc2, 0x09, 0xf7, 0x87, 0x12,
	0x9f, 0x78, 0x62, 0x93, 0x06, 0x69, 0x1d, 0x74, 0xed, 0xaf, 0x45, 0xbd, 0x3a, 0x85, 0x38, 0xba,
	0x6a, 0xee, 0x8c, 0x9b, 0x83, 0x8a, 0xae, 0xef, 0x8b, 0xd2, 0xba, 0x35, 0x2b, 0x8a, 0x65, 0x98,
	0xf2, 0x2c, 0x44, 0xdf, 0xfe, 0xd3, 0x20, 0xad, 0xf2, 0xf9, 0xb1, 0xbb, 0x06, 0x70, 0x37, 0x00,
	0x6e, 0x4f, 0x03, 0x74, 0xff, 0xcf, 0x16, 0x75, 0xe3, 0xed, 0xa3, 0x4e, 0x06, 0x65, 0x65, 0xec,
	0x2b, 0x9f, 0xd5, 0x33, 0xcb, 0x62, 0x02, 0xe9, 0x26, 0xe6, 0xef, 0xef, 0x63, 0xcc, 0xc2, 0xa7,
	0x53, 0x5a, 0xe6, 0x91, 0x87, 0x89, 0x7a, 0xf5, 0x1b, 0xdf, 0xcf, 0xb8, 0x10, 0xf6, 0xbf, 0xe2,
	0x3a, 0x83, 0xfd, 0x76, 0xf7, 0x6e, 0xb6, 0x74, 0xc8, 0x7c, 0xe9, 0x90, 0xcf, 0xa5, 0x43, 0x5e,
	0x57, 0x8e, 0x31, 0x5f, 0x39, 0xc6, 0xfb, 0xca, 0x31, 0x1e, 0x2e, 0x83, 0x50, 0x3e, 0xe6, 0xcc,
	0xf5, 0x30, 0xa6, 0x09, 0x16, 0xcb, 0x20, 0x3a, 0x8b, 0x80, 0x09, 0xba, 0xf7, 0xff, 0xe3, 0x0e,
	0x7d, 0xd1, 0x3d, 0x2a, 0xa7, 0x29, 0x17, 0xac, 0xa4, 0x38, 0x2f, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xe1, 0x45, 0x79, 0x64, 0x2a, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x22
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.SwapPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.SwapPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.QueryPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.QueryPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.AllowedToken) > 0 {
		i -= len(m.AllowedToken)
		copy(dAtA[i:], m.AllowedToken)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AllowedToken)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AllowedToken)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.QueryPeriod)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.SwapPeriod)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.QueryPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.SwapPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
