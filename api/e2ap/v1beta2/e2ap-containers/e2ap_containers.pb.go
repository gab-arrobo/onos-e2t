// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/e2ap/v1beta2/e2ap_containers.proto

package e2ap_containers

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-constants"
	math "math"
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

// sequence from e2ap-v01.01.00.asn1:1710
// Param E2AP-PROTOCOL-IES:IEsSetParam
//{-}
type ProtocolIeContainer001 struct {
	Value                []*ProtocolIeField001 `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProtocolIeContainer001) Reset()         { *m = ProtocolIeContainer001{} }
func (m *ProtocolIeContainer001) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeContainer001) ProtoMessage()    {}
func (*ProtocolIeContainer001) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e4ab23fc25b4c0, []int{0}
}
func (m *ProtocolIeContainer001) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeContainer001.Unmarshal(m, b)
}
func (m *ProtocolIeContainer001) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeContainer001.Marshal(b, m, deterministic)
}
func (m *ProtocolIeContainer001) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeContainer001.Merge(m, src)
}
func (m *ProtocolIeContainer001) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeContainer001.Size(m)
}
func (m *ProtocolIeContainer001) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeContainer001.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeContainer001 proto.InternalMessageInfo

func (m *ProtocolIeContainer001) GetValue() []*ProtocolIeField001 {
	if m != nil {
		return m.Value
	}
	return nil
}

// reference from e2ap-v01.01.00.asn1:1713
// Param E2AP-PROTOCOL-IES:IEsSetParam
//{-}
type ProtocolIeSingleContainer001 struct {
	Value                *ProtocolIeField001 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProtocolIeSingleContainer001) Reset()         { *m = ProtocolIeSingleContainer001{} }
func (m *ProtocolIeSingleContainer001) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeSingleContainer001) ProtoMessage()    {}
func (*ProtocolIeSingleContainer001) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e4ab23fc25b4c0, []int{1}
}
func (m *ProtocolIeSingleContainer001) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeSingleContainer001.Unmarshal(m, b)
}
func (m *ProtocolIeSingleContainer001) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeSingleContainer001.Marshal(b, m, deterministic)
}
func (m *ProtocolIeSingleContainer001) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeSingleContainer001.Merge(m, src)
}
func (m *ProtocolIeSingleContainer001) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeSingleContainer001.Size(m)
}
func (m *ProtocolIeSingleContainer001) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeSingleContainer001.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeSingleContainer001 proto.InternalMessageInfo

func (m *ProtocolIeSingleContainer001) GetValue() *ProtocolIeField001 {
	if m != nil {
		return m.Value
	}
	return nil
}

// sequence from e2ap-v01.01.00.asn1:1716
// Param E2AP-PROTOCOL-IES:IEsSetParam
//{-}
type ProtocolIeField001 struct {
	Id                   *e2ap_constants.IdRicrequestId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Criticality          e2ap_commondatatypes.Criticality `protobuf:"varint,2,opt,name=criticality,proto3,enum=e2ap.v1beta2.Criticality" json:"criticality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ProtocolIeField001) Reset()         { *m = ProtocolIeField001{} }
func (m *ProtocolIeField001) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeField001) ProtoMessage()    {}
func (*ProtocolIeField001) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e4ab23fc25b4c0, []int{2}
}
func (m *ProtocolIeField001) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeField001.Unmarshal(m, b)
}
func (m *ProtocolIeField001) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeField001.Marshal(b, m, deterministic)
}
func (m *ProtocolIeField001) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeField001.Merge(m, src)
}
func (m *ProtocolIeField001) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeField001.Size(m)
}
func (m *ProtocolIeField001) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeField001.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeField001 proto.InternalMessageInfo

func (m *ProtocolIeField001) GetId() *e2ap_constants.IdRicrequestId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ProtocolIeField001) GetCriticality() e2ap_commondatatypes.Criticality {
	if m != nil {
		return m.Criticality
	}
	return e2ap_commondatatypes.Criticality_CRITICALITY_REJECT
}

// sequence from e2ap-v01.01.00.asn1:1729
// Param E2AP-PROTOCOL-IES-PAIR:IEsSetParam
//{-}
type ProtocolIeContainerPair struct {
	Value                []*ProtocolIeFieldPair `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProtocolIeContainerPair) Reset()         { *m = ProtocolIeContainerPair{} }
func (m *ProtocolIeContainerPair) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeContainerPair) ProtoMessage()    {}
func (*ProtocolIeContainerPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e4ab23fc25b4c0, []int{3}
}
func (m *ProtocolIeContainerPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeContainerPair.Unmarshal(m, b)
}
func (m *ProtocolIeContainerPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeContainerPair.Marshal(b, m, deterministic)
}
func (m *ProtocolIeContainerPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeContainerPair.Merge(m, src)
}
func (m *ProtocolIeContainerPair) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeContainerPair.Size(m)
}
func (m *ProtocolIeContainerPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeContainerPair.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeContainerPair proto.InternalMessageInfo

func (m *ProtocolIeContainerPair) GetValue() []*ProtocolIeFieldPair {
	if m != nil {
		return m.Value
	}
	return nil
}

// sequence from e2ap-v01.01.00.asn1:1732
// Param E2AP-PROTOCOL-IES-PAIR:IEsSetParam
//{-}
type ProtocolIeFieldPair struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtocolIeFieldPair) Reset()         { *m = ProtocolIeFieldPair{} }
func (m *ProtocolIeFieldPair) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeFieldPair) ProtoMessage()    {}
func (*ProtocolIeFieldPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e4ab23fc25b4c0, []int{4}
}
func (m *ProtocolIeFieldPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeFieldPair.Unmarshal(m, b)
}
func (m *ProtocolIeFieldPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeFieldPair.Marshal(b, m, deterministic)
}
func (m *ProtocolIeFieldPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeFieldPair.Merge(m, src)
}
func (m *ProtocolIeFieldPair) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeFieldPair.Size(m)
}
func (m *ProtocolIeFieldPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeFieldPair.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeFieldPair proto.InternalMessageInfo

// sequence from e2ap-v01.01.00.asn1:1747
// Param INTEGER:lowerBound
// Param INTEGER:upperBound
// Param E2AP-PROTOCOL-IES:IEsSetParam
//{-}
type ProtocolIeContainerList struct {
	Value                []*ProtocolIeSingleContainer001 `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ProtocolIeContainerList) Reset()         { *m = ProtocolIeContainerList{} }
func (m *ProtocolIeContainerList) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeContainerList) ProtoMessage()    {}
func (*ProtocolIeContainerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e4ab23fc25b4c0, []int{5}
}
func (m *ProtocolIeContainerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeContainerList.Unmarshal(m, b)
}
func (m *ProtocolIeContainerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeContainerList.Marshal(b, m, deterministic)
}
func (m *ProtocolIeContainerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeContainerList.Merge(m, src)
}
func (m *ProtocolIeContainerList) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeContainerList.Size(m)
}
func (m *ProtocolIeContainerList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeContainerList.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeContainerList proto.InternalMessageInfo

func (m *ProtocolIeContainerList) GetValue() []*ProtocolIeSingleContainer001 {
	if m != nil {
		return m.Value
	}
	return nil
}

// sequence from e2ap-v01.01.00.asn1:1751
// Param INTEGER:lowerBound
// Param INTEGER:upperBound
// Param E2AP-PROTOCOL-IES-PAIR:IEsSetParam
//{-}
type ProtocolIeContainerPairList struct {
	Value                []*ProtocolIeContainerPair `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ProtocolIeContainerPairList) Reset()         { *m = ProtocolIeContainerPairList{} }
func (m *ProtocolIeContainerPairList) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeContainerPairList) ProtoMessage()    {}
func (*ProtocolIeContainerPairList) Descriptor() ([]byte, []int) {
	return fileDescriptor_69e4ab23fc25b4c0, []int{6}
}
func (m *ProtocolIeContainerPairList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeContainerPairList.Unmarshal(m, b)
}
func (m *ProtocolIeContainerPairList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeContainerPairList.Marshal(b, m, deterministic)
}
func (m *ProtocolIeContainerPairList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeContainerPairList.Merge(m, src)
}
func (m *ProtocolIeContainerPairList) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeContainerPairList.Size(m)
}
func (m *ProtocolIeContainerPairList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeContainerPairList.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeContainerPairList proto.InternalMessageInfo

func (m *ProtocolIeContainerPairList) GetValue() []*ProtocolIeContainerPair {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtocolIeContainer001)(nil), "e2ap.v1beta2.ProtocolIeContainer001")
	proto.RegisterType((*ProtocolIeSingleContainer001)(nil), "e2ap.v1beta2.ProtocolIeSingleContainer001")
	proto.RegisterType((*ProtocolIeField001)(nil), "e2ap.v1beta2.ProtocolIeField001")
	proto.RegisterType((*ProtocolIeContainerPair)(nil), "e2ap.v1beta2.ProtocolIeContainerPair")
	proto.RegisterType((*ProtocolIeFieldPair)(nil), "e2ap.v1beta2.ProtocolIeFieldPair")
	proto.RegisterType((*ProtocolIeContainerList)(nil), "e2ap.v1beta2.ProtocolIeContainerList")
	proto.RegisterType((*ProtocolIeContainerPairList)(nil), "e2ap.v1beta2.ProtocolIeContainerPairList")
}

func init() {
	proto.RegisterFile("api/e2ap/v1beta2/e2ap_containers.proto", fileDescriptor_69e4ab23fc25b4c0)
}

var fileDescriptor_69e4ab23fc25b4c0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4b, 0x4b, 0xc3, 0x40,
	0x10, 0xa6, 0x15, 0x3d, 0x6c, 0xc5, 0xc3, 0x8a, 0x5a, 0x6b, 0x0f, 0x6d, 0x40, 0x2d, 0x62, 0xf3,
	0x12, 0xf4, 0xd0, 0x4b, 0x69, 0x41, 0x28, 0x78, 0x28, 0x11, 0x3c, 0xd4, 0x83, 0x6c, 0x77, 0x97,
	0x3a, 0x92, 0x66, 0x63, 0x32, 0x0d, 0xf4, 0xe4, 0x5f, 0x97, 0xa6, 0x09, 0x69, 0xe8, 0x46, 0xf0,
	0x36, 0xc9, 0xf7, 0x9a, 0x19, 0x76, 0xc8, 0x0d, 0x0b, 0xc1, 0x92, 0x2e, 0x0b, 0xad, 0xc4, 0x99,
	0x4b, 0x64, 0x6e, 0xfa, 0xf1, 0xc1, 0x55, 0x80, 0x0c, 0x02, 0x19, 0xc5, 0x66, 0x18, 0x29, 0x54,
	0xf4, 0x78, 0xf3, 0xdb, 0xcc, 0x38, 0xad, 0x5b, 0x9d, 0x62, 0xb9, 0x54, 0x81, 0x60, 0xc8, 0x70,
	0x1d, 0xca, 0x4c, 0xd6, 0xea, 0x6a, 0xad, 0x63, 0x64, 0x01, 0xe6, 0x94, 0x56, 0xc2, 0x7c, 0x10,
	0x0c, 0xa5, 0x95, 0x38, 0x56, 0x5e, 0x6f, 0x31, 0x63, 0x4a, 0xce, 0xa7, 0x9b, 0x82, 0x2b, 0x7f,
	0x22, 0xc7, 0x79, 0x4f, 0xb6, 0xed, 0xd0, 0x47, 0x72, 0x98, 0x30, 0x7f, 0x25, 0x9b, 0xb5, 0xce,
	0x41, 0xaf, 0xe1, 0x76, 0xcc, 0xdd, 0xfe, 0xcc, 0x42, 0xf4, 0x0c, 0xd2, 0x17, 0xb6, 0xed, 0x78,
	0x5b, 0xba, 0xf1, 0x46, 0xda, 0x05, 0xf8, 0x0a, 0xc1, 0xc2, 0xaf, 0xf4, 0xad, 0xfd, 0xc7, 0xf7,
	0x87, 0xd0, 0x7d, 0x90, 0xde, 0x93, 0x3a, 0x88, 0xcc, 0xaa, 0x5d, 0xb6, 0x9a, 0x08, 0x0f, 0x78,
	0x24, 0xbf, 0x57, 0x32, 0xc6, 0x89, 0xf0, 0xea, 0x20, 0xe8, 0x80, 0x34, 0x78, 0x04, 0x08, 0x9c,
	0xf9, 0x80, 0xeb, 0x66, 0xbd, 0x53, 0xeb, 0x9d, 0xb8, 0x97, 0x65, 0xd9, 0xb8, 0x20, 0x78, 0xbb,
	0x6c, 0xc3, 0x23, 0x17, 0x9a, 0x55, 0x4d, 0x19, 0x44, 0xf4, 0xa9, 0xbc, 0xab, 0xee, 0x9f, 0x33,
	0x6d, 0x14, 0xf9, 0x50, 0x67, 0xe4, 0x54, 0x83, 0x1a, 0xef, 0xda, 0xa8, 0x17, 0x88, 0x91, 0x0e,
	0xcb, 0x51, 0x77, 0x55, 0x51, 0xfb, 0x9b, 0xcf, 0x33, 0x67, 0xe4, 0xaa, 0x62, 0x8e, 0x34, 0x60,
	0x50, 0x0e, 0xb8, 0xae, 0x0a, 0x28, 0x29, 0x33, 0xef, 0xd1, 0x68, 0x36, 0x5c, 0x00, 0x7e, 0xae,
	0xe6, 0x26, 0x57, 0x4b, 0x4b, 0x05, 0x2a, 0x0e, 0x23, 0xf5, 0x25, 0x39, 0xa6, 0x75, 0x5f, 0xba,
	0x68, 0x69, 0xcf, 0xa1, 0x5f, 0x9c, 0xc3, 0xfc, 0x28, 0x7d, 0x99, 0x0f, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0xce, 0xe7, 0x57, 0x39, 0x03, 0x00, 0x00,
}
