// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/types/types.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/kei2100/protoc-gen-marshal-zap"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enumerations
// https://developers.google.com/protocol-buffers/docs/proto#enum
type Types_Enum int32

const (
	Types_ENUM_0 Types_Enum = 0
	Types_ENUM_1 Types_Enum = 1
	Types_ENUM_2 Types_Enum = 2
)

var Types_Enum_name = map[int32]string{
	0: "ENUM_0",
	1: "ENUM_1",
	2: "ENUM_2",
}

var Types_Enum_value = map[string]int32{
	"ENUM_0": 0,
	"ENUM_1": 1,
	"ENUM_2": 2,
}

func (x Types_Enum) String() string {
	return proto.EnumName(Types_Enum_name, int32(x))
}

func (Types_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61f12dd0de6ed4ab, []int{0, 0}
}

type Types struct {
	SecretVal string `protobuf:"bytes,1,opt,name=secret_val,json=secretVal,proto3" json:"secret_val,omitempty"`
	// Scalar Value Types
	// https://developers.google.com/protocol-buffers/docs/proto#scalar
	DoubleVal   float64    `protobuf:"fixed64,2,opt,name=double_val,json=doubleVal,proto3" json:"double_val,omitempty"`
	FloatVal    float32    `protobuf:"fixed32,3,opt,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	Int32Val    int32      `protobuf:"varint,4,opt,name=int32_val,json=int32Val,proto3" json:"int32_val,omitempty"`
	Int64Val    int64      `protobuf:"varint,5,opt,name=int64_val,json=int64Val,proto3" json:"int64_val,omitempty"`
	Uint32Val   uint32     `protobuf:"varint,6,opt,name=uint32_val,json=uint32Val,proto3" json:"uint32_val,omitempty"`
	Uint64Val   uint64     `protobuf:"varint,7,opt,name=uint64_val,json=uint64Val,proto3" json:"uint64_val,omitempty"`
	Sint32Val   int32      `protobuf:"zigzag32,8,opt,name=sint32_val,json=sint32Val,proto3" json:"sint32_val,omitempty"`
	Sint64Val   int64      `protobuf:"zigzag64,9,opt,name=sint64_val,json=sint64Val,proto3" json:"sint64_val,omitempty"`
	Fixed32Val  uint32     `protobuf:"fixed32,10,opt,name=fixed32_val,json=fixed32Val,proto3" json:"fixed32_val,omitempty"`
	Fixed64Val  uint64     `protobuf:"fixed64,11,opt,name=fixed64_val,json=fixed64Val,proto3" json:"fixed64_val,omitempty"`
	Sfixed32Val int32      `protobuf:"fixed32,12,opt,name=sfixed32_val,json=sfixed32Val,proto3" json:"sfixed32_val,omitempty"`
	Sfixed64Val int64      `protobuf:"fixed64,13,opt,name=sfixed64_val,json=sfixed64Val,proto3" json:"sfixed64_val,omitempty"`
	BoolVal     bool       `protobuf:"varint,14,opt,name=bool_val,json=boolVal,proto3" json:"bool_val,omitempty"`
	StringVal   string     `protobuf:"bytes,15,opt,name=string_val,json=stringVal,proto3" json:"string_val,omitempty"`
	BytesVal    []byte     `protobuf:"bytes,16,opt,name=bytes_val,json=bytesVal,proto3" json:"bytes_val,omitempty"`
	EnumVal     Types_Enum `protobuf:"varint,17,opt,name=enum_val,json=enumVal,proto3,enum=types.Types_Enum" json:"enum_val,omitempty"`
	// Other Message Types
	// https://developers.google.com/protocol-buffers/docs/proto#other
	OtherTypeVal           *OtherType1            `protobuf:"bytes,18,opt,name=other_type_val,json=otherTypeVal,proto3" json:"other_type_val,omitempty"`
	NestedTypeVal          *Types_NestedType      `protobuf:"bytes,19,opt,name=nested_type_val,json=nestedTypeVal,proto3" json:"nested_type_val,omitempty"`
	OtherTypeNestedTypeVal *OtherType2_NestedType `protobuf:"bytes,20,opt,name=other_type_nested_type_val,json=otherTypeNestedTypeVal,proto3" json:"other_type_nested_type_val,omitempty"`
	// Oneof
	// https://developers.google.com/protocol-buffers/docs/proto#oneof
	//
	// Types that are valid to be assigned to OneofVal:
	//	*Types_OneofStringVal
	//	*Types_OneofInt64Val
	//	*Types_OneofBoolVal
	OneofVal isTypes_OneofVal `protobuf_oneof:"oneof_val"`
	// Maps
	// https://developers.google.com/protocol-buffers/docs/proto#updating
	MapVal1     map[string]string      `protobuf:"bytes,24,rep,name=map_val1,json=mapVal1,proto3" json:"map_val1,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapVal2     map[string]*OtherType1 `protobuf:"bytes,25,rep,name=map_val2,json=mapVal2,proto3" json:"map_val2,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapEmptyVal map[string]string      `protobuf:"bytes,26,rep,name=map_empty_val,json=mapEmptyVal,proto3" json:"map_empty_val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Repeated
	// * Repeated Scalar Values
	// * Repeated Enumerations
	// * Repeated MessageTypes
	// * (Repeated Maps -> NOT supported)
	// * (Repeated Repeated -> NOT supported)
	RepeatedVal1         []string      `protobuf:"bytes,27,rep,name=repeated_val1,json=repeatedVal1,proto3" json:"repeated_val1,omitempty"`
	RepeatedVal2         []Types_Enum  `protobuf:"varint,28,rep,packed,name=repeated_val2,json=repeatedVal2,proto3,enum=types.Types_Enum" json:"repeated_val2,omitempty"`
	RepeatedVal3         []*OtherType1 `protobuf:"bytes,29,rep,name=repeated_val3,json=repeatedVal3,proto3" json:"repeated_val3,omitempty"`
	RepeatedEmptyVal     []string      `protobuf:"bytes,30,rep,name=repeated_empty_val,json=repeatedEmptyVal,proto3" json:"repeated_empty_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Types) Reset()         { *m = Types{} }
func (m *Types) String() string { return proto.CompactTextString(m) }
func (*Types) ProtoMessage()    {}
func (*Types) Descriptor() ([]byte, []int) {
	return fileDescriptor_61f12dd0de6ed4ab, []int{0}
}

func (m *Types) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types.Unmarshal(m, b)
}
func (m *Types) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types.Marshal(b, m, deterministic)
}
func (m *Types) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types.Merge(m, src)
}
func (m *Types) XXX_Size() int {
	return xxx_messageInfo_Types.Size(m)
}
func (m *Types) XXX_DiscardUnknown() {
	xxx_messageInfo_Types.DiscardUnknown(m)
}

var xxx_messageInfo_Types proto.InternalMessageInfo

func (m *Types) GetSecretVal() string {
	if m != nil {
		return m.SecretVal
	}
	return ""
}

func (m *Types) GetDoubleVal() float64 {
	if m != nil {
		return m.DoubleVal
	}
	return 0
}

func (m *Types) GetFloatVal() float32 {
	if m != nil {
		return m.FloatVal
	}
	return 0
}

func (m *Types) GetInt32Val() int32 {
	if m != nil {
		return m.Int32Val
	}
	return 0
}

func (m *Types) GetInt64Val() int64 {
	if m != nil {
		return m.Int64Val
	}
	return 0
}

func (m *Types) GetUint32Val() uint32 {
	if m != nil {
		return m.Uint32Val
	}
	return 0
}

func (m *Types) GetUint64Val() uint64 {
	if m != nil {
		return m.Uint64Val
	}
	return 0
}

func (m *Types) GetSint32Val() int32 {
	if m != nil {
		return m.Sint32Val
	}
	return 0
}

func (m *Types) GetSint64Val() int64 {
	if m != nil {
		return m.Sint64Val
	}
	return 0
}

func (m *Types) GetFixed32Val() uint32 {
	if m != nil {
		return m.Fixed32Val
	}
	return 0
}

func (m *Types) GetFixed64Val() uint64 {
	if m != nil {
		return m.Fixed64Val
	}
	return 0
}

func (m *Types) GetSfixed32Val() int32 {
	if m != nil {
		return m.Sfixed32Val
	}
	return 0
}

func (m *Types) GetSfixed64Val() int64 {
	if m != nil {
		return m.Sfixed64Val
	}
	return 0
}

func (m *Types) GetBoolVal() bool {
	if m != nil {
		return m.BoolVal
	}
	return false
}

func (m *Types) GetStringVal() string {
	if m != nil {
		return m.StringVal
	}
	return ""
}

func (m *Types) GetBytesVal() []byte {
	if m != nil {
		return m.BytesVal
	}
	return nil
}

func (m *Types) GetEnumVal() Types_Enum {
	if m != nil {
		return m.EnumVal
	}
	return Types_ENUM_0
}

func (m *Types) GetOtherTypeVal() *OtherType1 {
	if m != nil {
		return m.OtherTypeVal
	}
	return nil
}

func (m *Types) GetNestedTypeVal() *Types_NestedType {
	if m != nil {
		return m.NestedTypeVal
	}
	return nil
}

func (m *Types) GetOtherTypeNestedTypeVal() *OtherType2_NestedType {
	if m != nil {
		return m.OtherTypeNestedTypeVal
	}
	return nil
}

type isTypes_OneofVal interface {
	isTypes_OneofVal()
}

type Types_OneofStringVal struct {
	OneofStringVal string `protobuf:"bytes,21,opt,name=oneof_string_val,json=oneofStringVal,proto3,oneof"`
}

type Types_OneofInt64Val struct {
	OneofInt64Val int64 `protobuf:"varint,22,opt,name=oneof_int64_val,json=oneofInt64Val,proto3,oneof"`
}

type Types_OneofBoolVal struct {
	OneofBoolVal bool `protobuf:"varint,23,opt,name=oneof_bool_val,json=oneofBoolVal,proto3,oneof"`
}

func (*Types_OneofStringVal) isTypes_OneofVal() {}

func (*Types_OneofInt64Val) isTypes_OneofVal() {}

func (*Types_OneofBoolVal) isTypes_OneofVal() {}

func (m *Types) GetOneofVal() isTypes_OneofVal {
	if m != nil {
		return m.OneofVal
	}
	return nil
}

func (m *Types) GetOneofStringVal() string {
	if x, ok := m.GetOneofVal().(*Types_OneofStringVal); ok {
		return x.OneofStringVal
	}
	return ""
}

func (m *Types) GetOneofInt64Val() int64 {
	if x, ok := m.GetOneofVal().(*Types_OneofInt64Val); ok {
		return x.OneofInt64Val
	}
	return 0
}

func (m *Types) GetOneofBoolVal() bool {
	if x, ok := m.GetOneofVal().(*Types_OneofBoolVal); ok {
		return x.OneofBoolVal
	}
	return false
}

func (m *Types) GetMapVal1() map[string]string {
	if m != nil {
		return m.MapVal1
	}
	return nil
}

func (m *Types) GetMapVal2() map[string]*OtherType1 {
	if m != nil {
		return m.MapVal2
	}
	return nil
}

func (m *Types) GetMapEmptyVal() map[string]string {
	if m != nil {
		return m.MapEmptyVal
	}
	return nil
}

func (m *Types) GetRepeatedVal1() []string {
	if m != nil {
		return m.RepeatedVal1
	}
	return nil
}

func (m *Types) GetRepeatedVal2() []Types_Enum {
	if m != nil {
		return m.RepeatedVal2
	}
	return nil
}

func (m *Types) GetRepeatedVal3() []*OtherType1 {
	if m != nil {
		return m.RepeatedVal3
	}
	return nil
}

func (m *Types) GetRepeatedEmptyVal() []string {
	if m != nil {
		return m.RepeatedEmptyVal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Types) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Types_OneofStringVal)(nil),
		(*Types_OneofInt64Val)(nil),
		(*Types_OneofBoolVal)(nil),
	}
}

// Nested Types
// https://developers.google.com/protocol-buffers/docs/proto#nested
type Types_NestedType struct {
	NestedStringVal      string   `protobuf:"bytes,1,opt,name=nested_string_val,json=nestedStringVal,proto3" json:"nested_string_val,omitempty"`
	NestedSecretVal      string   `protobuf:"bytes,2,opt,name=nested_secret_val,json=nestedSecretVal,proto3" json:"nested_secret_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Types_NestedType) Reset()         { *m = Types_NestedType{} }
func (m *Types_NestedType) String() string { return proto.CompactTextString(m) }
func (*Types_NestedType) ProtoMessage()    {}
func (*Types_NestedType) Descriptor() ([]byte, []int) {
	return fileDescriptor_61f12dd0de6ed4ab, []int{0, 0}
}

func (m *Types_NestedType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types_NestedType.Unmarshal(m, b)
}
func (m *Types_NestedType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types_NestedType.Marshal(b, m, deterministic)
}
func (m *Types_NestedType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types_NestedType.Merge(m, src)
}
func (m *Types_NestedType) XXX_Size() int {
	return xxx_messageInfo_Types_NestedType.Size(m)
}
func (m *Types_NestedType) XXX_DiscardUnknown() {
	xxx_messageInfo_Types_NestedType.DiscardUnknown(m)
}

var xxx_messageInfo_Types_NestedType proto.InternalMessageInfo

func (m *Types_NestedType) GetNestedStringVal() string {
	if m != nil {
		return m.NestedStringVal
	}
	return ""
}

func (m *Types_NestedType) GetNestedSecretVal() string {
	if m != nil {
		return m.NestedSecretVal
	}
	return ""
}

type OtherType1 struct {
	OtherStringVal       string   `protobuf:"bytes,1,opt,name=other_string_val,json=otherStringVal,proto3" json:"other_string_val,omitempty"`
	OtherSecretVal       string   `protobuf:"bytes,2,opt,name=other_secret_val,json=otherSecretVal,proto3" json:"other_secret_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtherType1) Reset()         { *m = OtherType1{} }
func (m *OtherType1) String() string { return proto.CompactTextString(m) }
func (*OtherType1) ProtoMessage()    {}
func (*OtherType1) Descriptor() ([]byte, []int) {
	return fileDescriptor_61f12dd0de6ed4ab, []int{1}
}

func (m *OtherType1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherType1.Unmarshal(m, b)
}
func (m *OtherType1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherType1.Marshal(b, m, deterministic)
}
func (m *OtherType1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherType1.Merge(m, src)
}
func (m *OtherType1) XXX_Size() int {
	return xxx_messageInfo_OtherType1.Size(m)
}
func (m *OtherType1) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherType1.DiscardUnknown(m)
}

var xxx_messageInfo_OtherType1 proto.InternalMessageInfo

func (m *OtherType1) GetOtherStringVal() string {
	if m != nil {
		return m.OtherStringVal
	}
	return ""
}

func (m *OtherType1) GetOtherSecretVal() string {
	if m != nil {
		return m.OtherSecretVal
	}
	return ""
}

type OtherType2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtherType2) Reset()         { *m = OtherType2{} }
func (m *OtherType2) String() string { return proto.CompactTextString(m) }
func (*OtherType2) ProtoMessage()    {}
func (*OtherType2) Descriptor() ([]byte, []int) {
	return fileDescriptor_61f12dd0de6ed4ab, []int{2}
}

func (m *OtherType2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherType2.Unmarshal(m, b)
}
func (m *OtherType2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherType2.Marshal(b, m, deterministic)
}
func (m *OtherType2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherType2.Merge(m, src)
}
func (m *OtherType2) XXX_Size() int {
	return xxx_messageInfo_OtherType2.Size(m)
}
func (m *OtherType2) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherType2.DiscardUnknown(m)
}

var xxx_messageInfo_OtherType2 proto.InternalMessageInfo

type OtherType2_NestedType struct {
	NestedStringVal      string   `protobuf:"bytes,1,opt,name=nested_string_val,json=nestedStringVal,proto3" json:"nested_string_val,omitempty"`
	NestedSecretVal      string   `protobuf:"bytes,2,opt,name=nested_secret_val,json=nestedSecretVal,proto3" json:"nested_secret_val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtherType2_NestedType) Reset()         { *m = OtherType2_NestedType{} }
func (m *OtherType2_NestedType) String() string { return proto.CompactTextString(m) }
func (*OtherType2_NestedType) ProtoMessage()    {}
func (*OtherType2_NestedType) Descriptor() ([]byte, []int) {
	return fileDescriptor_61f12dd0de6ed4ab, []int{2, 0}
}

func (m *OtherType2_NestedType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherType2_NestedType.Unmarshal(m, b)
}
func (m *OtherType2_NestedType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherType2_NestedType.Marshal(b, m, deterministic)
}
func (m *OtherType2_NestedType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherType2_NestedType.Merge(m, src)
}
func (m *OtherType2_NestedType) XXX_Size() int {
	return xxx_messageInfo_OtherType2_NestedType.Size(m)
}
func (m *OtherType2_NestedType) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherType2_NestedType.DiscardUnknown(m)
}

var xxx_messageInfo_OtherType2_NestedType proto.InternalMessageInfo

func (m *OtherType2_NestedType) GetNestedStringVal() string {
	if m != nil {
		return m.NestedStringVal
	}
	return ""
}

func (m *OtherType2_NestedType) GetNestedSecretVal() string {
	if m != nil {
		return m.NestedSecretVal
	}
	return ""
}

func init() {
	proto.RegisterEnum("types.Types_Enum", Types_Enum_name, Types_Enum_value)
	proto.RegisterType((*Types)(nil), "types.Types")
	proto.RegisterMapType((map[string]string)(nil), "types.Types.MapEmptyValEntry")
	proto.RegisterMapType((map[string]string)(nil), "types.Types.MapVal1Entry")
	proto.RegisterMapType((map[string]*OtherType1)(nil), "types.Types.MapVal2Entry")
	proto.RegisterType((*Types_NestedType)(nil), "types.Types.NestedType")
	proto.RegisterType((*OtherType1)(nil), "types.OtherType1")
	proto.RegisterType((*OtherType2)(nil), "types.OtherType2")
	proto.RegisterType((*OtherType2_NestedType)(nil), "types.OtherType2.NestedType")
}

func init() {
	proto.RegisterFile("test/types/types.proto", fileDescriptor_61f12dd0de6ed4ab)
}

var fileDescriptor_61f12dd0de6ed4ab = []byte{
	// 811 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xdb, 0x6e, 0xf3, 0x44,
	0x10, 0xc7, 0xb3, 0x39, 0xda, 0x13, 0x27, 0x71, 0x4c, 0x69, 0xd3, 0xb4, 0x85, 0x25, 0x95, 0x60,
	0x55, 0xb5, 0x39, 0x38, 0x55, 0x41, 0x05, 0x81, 0x88, 0x14, 0xa9, 0x5c, 0xb4, 0x48, 0x0b, 0x54,
	0x88, 0x9b, 0xc8, 0x69, 0x37, 0x6d, 0xa8, 0x0f, 0x51, 0xec, 0x54, 0x0d, 0x4f, 0xc0, 0x0b, 0xf1,
	0x3c, 0xbc, 0x0a, 0xda, 0x5d, 0x1f, 0x36, 0x21, 0xbd, 0xf8, 0x6e, 0xbe, 0x9b, 0x68, 0x3d, 0xff,
	0xff, 0x6f, 0x3c, 0x9e, 0x99, 0xd8, 0xb0, 0x1f, 0xb1, 0x30, 0xea, 0x45, 0xeb, 0x05, 0x0b, 0xe5,
	0x6f, 0x77, 0xb1, 0x0c, 0xa2, 0xc0, 0x2a, 0x89, 0x8b, 0x76, 0xd3, 0x73, 0x96, 0xe1, 0xb3, 0xe3,
	0x5e, 0xfc, 0xe5, 0x2c, 0xa4, 0xd2, 0xf9, 0xd7, 0x80, 0xd2, 0xaf, 0x5c, 0xb4, 0x4e, 0x01, 0x42,
	0xf6, 0xb0, 0x64, 0xd1, 0xe4, 0xd5, 0x71, 0x5b, 0x08, 0x23, 0xa2, 0x8f, 0x8a, 0x7f, 0xff, 0xd3,
	0x42, 0x54, 0x97, 0xf1, 0x7b, 0xc7, 0xb5, 0x4e, 0x00, 0x1e, 0x83, 0xd5, 0xd4, 0x65, 0xc2, 0x94,
	0xc7, 0x88, 0x20, 0xaa, 0xcb, 0x08, 0x97, 0x8f, 0x40, 0x9f, 0xb9, 0x81, 0x23, 0x53, 0x14, 0x30,
	0x22, 0x79, 0xaa, 0x89, 0x40, 0x2c, 0xce, 0xfd, 0x68, 0x68, 0x0b, 0xb1, 0x88, 0x11, 0x29, 0x51,
	0x4d, 0x04, 0x32, 0xf1, 0xea, 0x52, 0x88, 0x25, 0x8c, 0x48, 0x41, 0x88, 0x57, 0x97, 0xf1, 0x5d,
	0x57, 0x19, 0x5a, 0xc6, 0x88, 0xd4, 0xa8, 0xbe, 0x4a, 0xd9, 0x58, 0x8e, 0xe1, 0x0a, 0x46, 0xa4,
	0x28, 0xe5, 0x94, 0x0e, 0x33, 0x5a, 0xc3, 0x88, 0x34, 0xa9, 0x1e, 0xaa, 0x74, 0x98, 0xd1, 0x3a,
	0x46, 0xc4, 0x92, 0xb2, 0xa4, 0x3f, 0x87, 0xea, 0x6c, 0xfe, 0xc6, 0x1e, 0x63, 0x1c, 0x30, 0x22,
	0x15, 0x0a, 0x71, 0x48, 0x35, 0xc4, 0x09, 0xaa, 0x18, 0x91, 0x72, 0x6c, 0x90, 0x19, 0xbe, 0x00,
	0x23, 0x54, 0x53, 0x18, 0x18, 0x91, 0x06, 0xad, 0x86, 0x4a, 0x8e, 0xd4, 0x12, 0x27, 0xa9, 0x61,
	0x44, 0xcc, 0xc4, 0x22, 0xb3, 0x1c, 0x82, 0x36, 0x0d, 0x02, 0x57, 0xc8, 0x75, 0x8c, 0x88, 0x46,
	0x2b, 0xfc, 0x3a, 0x79, 0x82, 0x68, 0x39, 0xf7, 0x9f, 0x84, 0xd8, 0xe0, 0x93, 0xa3, 0xba, 0x8c,
	0xc4, 0xad, 0x9d, 0xae, 0x23, 0x16, 0x0a, 0xd5, 0xc4, 0x88, 0x18, 0x54, 0x13, 0x01, 0x2e, 0x9e,
	0x83, 0xc6, 0xfc, 0x95, 0x27, 0xb4, 0x26, 0x46, 0xa4, 0x6e, 0x37, 0xbb, 0x72, 0x73, 0xc4, 0x56,
	0x74, 0xc7, 0xfe, 0xca, 0xa3, 0x15, 0x6e, 0xe1, 0xee, 0xaf, 0xa1, 0x1e, 0x44, 0xcf, 0x6c, 0x39,
	0xe1, 0x16, 0xc1, 0x58, 0x18, 0x91, 0x6a, 0xca, 0xfc, 0xcc, 0x45, 0x0e, 0x0e, 0xa8, 0x11, 0x24,
	0x67, 0x0e, 0xfe, 0x00, 0x0d, 0x9f, 0x85, 0x11, 0x7b, 0xcc, 0xc8, 0x4f, 0x04, 0x79, 0xb0, 0x71,
	0xb7, 0x3b, 0xe1, 0xe1, 0x67, 0x5a, 0xf3, 0xd3, 0x33, 0x4f, 0xf0, 0x3b, 0xb4, 0x95, 0x3b, 0x6f,
	0xe7, 0xda, 0x13, 0xb9, 0x8e, 0xb7, 0xab, 0xb0, 0xd5, 0x84, 0xfb, 0x69, 0x41, 0x77, 0x1b, 0x99,
	0xcf, 0xc0, 0x0c, 0x7c, 0x16, 0xcc, 0x26, 0x4a, 0x0f, 0x3f, 0xe5, 0x3d, 0xbc, 0xc9, 0xd1, 0xba,
	0x50, 0x7e, 0x49, 0x5b, 0x49, 0xa0, 0x21, 0xbd, 0xd9, 0xc2, 0xec, 0xf3, 0x5d, 0xbd, 0xc9, 0xd1,
	0x9a, 0x10, 0x7e, 0x4a, 0xd6, 0xe6, 0x4b, 0x90, 0xec, 0x24, 0x1d, 0xda, 0x01, 0x1f, 0xda, 0x4d,
	0x8e, 0x1a, 0x22, 0x3e, 0x8a, 0x67, 0x77, 0x09, 0x9a, 0xe7, 0x2c, 0xb8, 0x61, 0xd0, 0x6a, 0xe1,
	0x02, 0xa9, 0xda, 0x87, 0x1b, 0x1d, 0xb9, 0x75, 0x16, 0xf7, 0x8e, 0x3b, 0x18, 0xfb, 0xd1, 0x72,
	0x4d, 0x2b, 0x9e, 0xbc, 0x52, 0x28, 0xbb, 0x75, 0xf8, 0x2e, 0x65, 0x6f, 0x50, 0xb6, 0xf5, 0x23,
	0xd4, 0x38, 0xc5, 0xbc, 0x45, 0xb4, 0x16, 0x25, 0xb5, 0x05, 0x7a, 0xb2, 0x8d, 0x8e, 0xb9, 0xe1,
	0xde, 0x71, 0x25, 0x5e, 0xf5, 0xb2, 0x88, 0x75, 0x0a, 0xb5, 0x25, 0x5b, 0x30, 0x87, 0x77, 0x5f,
	0xd4, 0x7c, 0x84, 0x0b, 0x44, 0xa7, 0x46, 0x12, 0x14, 0xd5, 0x5d, 0x6d, 0x9a, 0xec, 0xd6, 0x31,
	0x2e, 0xec, 0x5e, 0x2c, 0x95, 0xb3, 0xb7, 0xb9, 0x61, 0xeb, 0x44, 0xd4, 0xb7, 0x6b, 0xb9, 0x14,
	0x6e, 0x68, 0x9d, 0x83, 0x95, 0x72, 0xd9, 0xc3, 0x7d, 0x26, 0x2a, 0x33, 0x13, 0x25, 0x79, 0x84,
	0xf6, 0x9f, 0x00, 0xd9, 0x02, 0x58, 0x67, 0xd0, 0x8c, 0x97, 0x49, 0x19, 0xbf, 0x78, 0xf9, 0xd1,
	0x78, 0x63, 0xb3, 0xe9, 0xf7, 0x33, 0x6f, 0xf6, 0xa2, 0xcc, 0x2b, 0x2f, 0xca, 0x84, 0x48, 0x5e,
	0x97, 0xed, 0x6b, 0x30, 0xd4, 0x01, 0x5a, 0x26, 0x14, 0x5e, 0xd8, 0x3a, 0xce, 0xcf, 0x8f, 0xd6,
	0x1e, 0x94, 0x5e, 0x1d, 0x77, 0xc5, 0x64, 0x1e, 0x2a, 0x2f, 0xae, 0xf3, 0xdf, 0xa0, 0xf6, 0x6d,
	0xc2, 0xda, 0xef, 0xb1, 0x5f, 0xa9, 0xec, 0xce, 0x3e, 0x29, 0xe9, 0xbe, 0x07, 0x73, 0x7b, 0xb4,
	0x1f, 0x52, 0x4e, 0xe7, 0x0c, 0x8a, 0x7c, 0x64, 0x16, 0x40, 0x79, 0x7c, 0xf7, 0xdb, 0xed, 0xa4,
	0x6f, 0xe6, 0xd2, 0xf3, 0xc0, 0x44, 0xe9, 0xd9, 0x36, 0xf3, 0xa3, 0x2a, 0xe8, 0x72, 0xf9, 0x5f,
	0x1d, 0xb7, 0x33, 0x03, 0xc8, 0x2a, 0xb2, 0x08, 0x98, 0xf2, 0x7f, 0xfc, 0xbf, 0x76, 0xcb, 0x37,
	0x4b, 0xd6, 0xed, 0x6e, 0xea, 0xdc, 0xdd, 0xec, 0xd8, 0x9f, 0xf4, 0xba, 0xf3, 0xa6, 0xdc, 0xc7,
	0xfe, 0x98, 0x53, 0x1e, 0x7d, 0xf7, 0xc7, 0xf5, 0xd3, 0x3c, 0x7a, 0x5e, 0x4d, 0xbb, 0x0f, 0x81,
	0xd7, 0x7b, 0x61, 0x73, 0x7b, 0xd0, 0xef, 0xf7, 0xc4, 0xf7, 0xf5, 0xe1, 0xe2, 0x89, 0xf9, 0x17,
	0xca, 0x67, 0xb7, 0x97, 0x7d, 0xa1, 0xbf, 0x15, 0xbf, 0xd3, 0xb2, 0x30, 0x0e, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x4c, 0xaf, 0xdc, 0x98, 0xbc, 0x07, 0x00, 0x00,
}
