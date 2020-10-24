// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: test/types/types.proto

package types

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/kei2100/protoc-gen-marshal-zap"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Enumerations
// https://developers.google.com/protocol-buffers/docs/proto#enum
type Types_Enum int32

const (
	Types_ENUM_0 Types_Enum = 0
	Types_ENUM_1 Types_Enum = 1
	Types_ENUM_2 Types_Enum = 2
)

// Enum value maps for Types_Enum.
var (
	Types_Enum_name = map[int32]string{
		0: "ENUM_0",
		1: "ENUM_1",
		2: "ENUM_2",
	}
	Types_Enum_value = map[string]int32{
		"ENUM_0": 0,
		"ENUM_1": 1,
		"ENUM_2": 2,
	}
)

func (x Types_Enum) Enum() *Types_Enum {
	p := new(Types_Enum)
	*p = x
	return p
}

func (x Types_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Types_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_test_types_types_proto_enumTypes[0].Descriptor()
}

func (Types_Enum) Type() protoreflect.EnumType {
	return &file_test_types_types_proto_enumTypes[0]
}

func (x Types_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Types_Enum.Descriptor instead.
func (Types_Enum) EnumDescriptor() ([]byte, []int) {
	return file_test_types_types_proto_rawDescGZIP(), []int{0, 0}
}

type Types struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	// Types that are assignable to OneofVal:
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
	RepeatedVal1     []string      `protobuf:"bytes,27,rep,name=repeated_val1,json=repeatedVal1,proto3" json:"repeated_val1,omitempty"`
	RepeatedVal2     []Types_Enum  `protobuf:"varint,28,rep,packed,name=repeated_val2,json=repeatedVal2,proto3,enum=types.Types_Enum" json:"repeated_val2,omitempty"`
	RepeatedVal3     []*OtherType1 `protobuf:"bytes,29,rep,name=repeated_val3,json=repeatedVal3,proto3" json:"repeated_val3,omitempty"`
	RepeatedEmptyVal []string      `protobuf:"bytes,30,rep,name=repeated_empty_val,json=repeatedEmptyVal,proto3" json:"repeated_empty_val,omitempty"`
	// TODO wellknown types
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf
	StructVal *_struct.Struct `protobuf:"bytes,31,opt,name=struct_val,json=structVal,proto3" json:"struct_val,omitempty"`
}

func (x *Types) Reset() {
	*x = Types{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_types_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Types) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Types) ProtoMessage() {}

func (x *Types) ProtoReflect() protoreflect.Message {
	mi := &file_test_types_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Types.ProtoReflect.Descriptor instead.
func (*Types) Descriptor() ([]byte, []int) {
	return file_test_types_types_proto_rawDescGZIP(), []int{0}
}

func (x *Types) GetSecretVal() string {
	if x != nil {
		return x.SecretVal
	}
	return ""
}

func (x *Types) GetDoubleVal() float64 {
	if x != nil {
		return x.DoubleVal
	}
	return 0
}

func (x *Types) GetFloatVal() float32 {
	if x != nil {
		return x.FloatVal
	}
	return 0
}

func (x *Types) GetInt32Val() int32 {
	if x != nil {
		return x.Int32Val
	}
	return 0
}

func (x *Types) GetInt64Val() int64 {
	if x != nil {
		return x.Int64Val
	}
	return 0
}

func (x *Types) GetUint32Val() uint32 {
	if x != nil {
		return x.Uint32Val
	}
	return 0
}

func (x *Types) GetUint64Val() uint64 {
	if x != nil {
		return x.Uint64Val
	}
	return 0
}

func (x *Types) GetSint32Val() int32 {
	if x != nil {
		return x.Sint32Val
	}
	return 0
}

func (x *Types) GetSint64Val() int64 {
	if x != nil {
		return x.Sint64Val
	}
	return 0
}

func (x *Types) GetFixed32Val() uint32 {
	if x != nil {
		return x.Fixed32Val
	}
	return 0
}

func (x *Types) GetFixed64Val() uint64 {
	if x != nil {
		return x.Fixed64Val
	}
	return 0
}

func (x *Types) GetSfixed32Val() int32 {
	if x != nil {
		return x.Sfixed32Val
	}
	return 0
}

func (x *Types) GetSfixed64Val() int64 {
	if x != nil {
		return x.Sfixed64Val
	}
	return 0
}

func (x *Types) GetBoolVal() bool {
	if x != nil {
		return x.BoolVal
	}
	return false
}

func (x *Types) GetStringVal() string {
	if x != nil {
		return x.StringVal
	}
	return ""
}

func (x *Types) GetBytesVal() []byte {
	if x != nil {
		return x.BytesVal
	}
	return nil
}

func (x *Types) GetEnumVal() Types_Enum {
	if x != nil {
		return x.EnumVal
	}
	return Types_ENUM_0
}

func (x *Types) GetOtherTypeVal() *OtherType1 {
	if x != nil {
		return x.OtherTypeVal
	}
	return nil
}

func (x *Types) GetNestedTypeVal() *Types_NestedType {
	if x != nil {
		return x.NestedTypeVal
	}
	return nil
}

func (x *Types) GetOtherTypeNestedTypeVal() *OtherType2_NestedType {
	if x != nil {
		return x.OtherTypeNestedTypeVal
	}
	return nil
}

func (m *Types) GetOneofVal() isTypes_OneofVal {
	if m != nil {
		return m.OneofVal
	}
	return nil
}

func (x *Types) GetOneofStringVal() string {
	if x, ok := x.GetOneofVal().(*Types_OneofStringVal); ok {
		return x.OneofStringVal
	}
	return ""
}

func (x *Types) GetOneofInt64Val() int64 {
	if x, ok := x.GetOneofVal().(*Types_OneofInt64Val); ok {
		return x.OneofInt64Val
	}
	return 0
}

func (x *Types) GetOneofBoolVal() bool {
	if x, ok := x.GetOneofVal().(*Types_OneofBoolVal); ok {
		return x.OneofBoolVal
	}
	return false
}

func (x *Types) GetMapVal1() map[string]string {
	if x != nil {
		return x.MapVal1
	}
	return nil
}

func (x *Types) GetMapVal2() map[string]*OtherType1 {
	if x != nil {
		return x.MapVal2
	}
	return nil
}

func (x *Types) GetMapEmptyVal() map[string]string {
	if x != nil {
		return x.MapEmptyVal
	}
	return nil
}

func (x *Types) GetRepeatedVal1() []string {
	if x != nil {
		return x.RepeatedVal1
	}
	return nil
}

func (x *Types) GetRepeatedVal2() []Types_Enum {
	if x != nil {
		return x.RepeatedVal2
	}
	return nil
}

func (x *Types) GetRepeatedVal3() []*OtherType1 {
	if x != nil {
		return x.RepeatedVal3
	}
	return nil
}

func (x *Types) GetRepeatedEmptyVal() []string {
	if x != nil {
		return x.RepeatedEmptyVal
	}
	return nil
}

func (x *Types) GetStructVal() *_struct.Struct {
	if x != nil {
		return x.StructVal
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

type OtherType1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherStringVal string `protobuf:"bytes,1,opt,name=other_string_val,json=otherStringVal,proto3" json:"other_string_val,omitempty"`
	OtherSecretVal string `protobuf:"bytes,2,opt,name=other_secret_val,json=otherSecretVal,proto3" json:"other_secret_val,omitempty"`
}

func (x *OtherType1) Reset() {
	*x = OtherType1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_types_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherType1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherType1) ProtoMessage() {}

func (x *OtherType1) ProtoReflect() protoreflect.Message {
	mi := &file_test_types_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherType1.ProtoReflect.Descriptor instead.
func (*OtherType1) Descriptor() ([]byte, []int) {
	return file_test_types_types_proto_rawDescGZIP(), []int{1}
}

func (x *OtherType1) GetOtherStringVal() string {
	if x != nil {
		return x.OtherStringVal
	}
	return ""
}

func (x *OtherType1) GetOtherSecretVal() string {
	if x != nil {
		return x.OtherSecretVal
	}
	return ""
}

type OtherType2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OtherType2) Reset() {
	*x = OtherType2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_types_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherType2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherType2) ProtoMessage() {}

func (x *OtherType2) ProtoReflect() protoreflect.Message {
	mi := &file_test_types_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherType2.ProtoReflect.Descriptor instead.
func (*OtherType2) Descriptor() ([]byte, []int) {
	return file_test_types_types_proto_rawDescGZIP(), []int{2}
}

// Nested Types
// https://developers.google.com/protocol-buffers/docs/proto#nested
type Types_NestedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedStringVal string `protobuf:"bytes,1,opt,name=nested_string_val,json=nestedStringVal,proto3" json:"nested_string_val,omitempty"`
	NestedSecretVal string `protobuf:"bytes,2,opt,name=nested_secret_val,json=nestedSecretVal,proto3" json:"nested_secret_val,omitempty"`
}

func (x *Types_NestedType) Reset() {
	*x = Types_NestedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_types_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Types_NestedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Types_NestedType) ProtoMessage() {}

func (x *Types_NestedType) ProtoReflect() protoreflect.Message {
	mi := &file_test_types_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Types_NestedType.ProtoReflect.Descriptor instead.
func (*Types_NestedType) Descriptor() ([]byte, []int) {
	return file_test_types_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Types_NestedType) GetNestedStringVal() string {
	if x != nil {
		return x.NestedStringVal
	}
	return ""
}

func (x *Types_NestedType) GetNestedSecretVal() string {
	if x != nil {
		return x.NestedSecretVal
	}
	return ""
}

type OtherType2_NestedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedStringVal string `protobuf:"bytes,1,opt,name=nested_string_val,json=nestedStringVal,proto3" json:"nested_string_val,omitempty"`
	NestedSecretVal string `protobuf:"bytes,2,opt,name=nested_secret_val,json=nestedSecretVal,proto3" json:"nested_secret_val,omitempty"`
}

func (x *OtherType2_NestedType) Reset() {
	*x = OtherType2_NestedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_types_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherType2_NestedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherType2_NestedType) ProtoMessage() {}

func (x *OtherType2_NestedType) ProtoReflect() protoreflect.Message {
	mi := &file_test_types_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherType2_NestedType.ProtoReflect.Descriptor instead.
func (*OtherType2_NestedType) Descriptor() ([]byte, []int) {
	return file_test_types_types_proto_rawDescGZIP(), []int{2, 0}
}

func (x *OtherType2_NestedType) GetNestedStringVal() string {
	if x != nil {
		return x.NestedStringVal
	}
	return ""
}

func (x *OtherType2_NestedType) GetNestedSecretVal() string {
	if x != nil {
		return x.NestedSecretVal
	}
	return ""
}

var File_test_types_types_proto protoreflect.FileDescriptor

var file_test_types_types_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a,
	0x11, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x2d, 0x7a, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x0d, 0x0a, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0x88, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x11, 0x52, 0x09, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0a, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0a, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x0b, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0b, 0x73, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c,
	0x12, 0x2c, 0x0a, 0x08, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x07, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x12, 0x37,
	0x0a, 0x0e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x31, 0x52, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x58, 0x0a, 0x1a, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x32, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x16, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x56,
	0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x28,
	0x0a, 0x0f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61,
	0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x12, 0x34, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x31, 0x18, 0x18, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d,
	0x61, 0x70, 0x56, 0x61, 0x6c, 0x31, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x76, 0x61,
	0x6c, 0x32, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x32, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x32, 0x12, 0x41, 0x0a, 0x0d,
	0x6d, 0x61, 0x70, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x1a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x31,
	0x18, 0x1b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x31, 0x12, 0x36, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x32, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x32, 0x12, 0x36, 0x0a, 0x0d,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x33, 0x18, 0x1d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x31, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x33, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56,
	0x61, 0x6c, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x09, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x1a, 0x6a, 0x0a, 0x0a, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x11, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x1a, 0x3a, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c,
	0x31, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x32, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x31, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x61, 0x6c,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x2a, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x30, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x31, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x32, 0x10, 0x02, 0x42, 0x0b, 0x0a,
	0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x22, 0x66, 0x0a, 0x0a, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x31, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x10, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x88, 0xb5,
	0x18, 0x01, 0x52, 0x0e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x22, 0x78, 0x0a, 0x0a, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x32,
	0x1a, 0x6a, 0x0a, 0x0a, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x11, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x0f, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x69, 0x32, 0x31,
	0x30, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x2d, 0x7a, 0x61, 0x70, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_test_types_types_proto_rawDescOnce sync.Once
	file_test_types_types_proto_rawDescData = file_test_types_types_proto_rawDesc
)

func file_test_types_types_proto_rawDescGZIP() []byte {
	file_test_types_types_proto_rawDescOnce.Do(func() {
		file_test_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_types_types_proto_rawDescData)
	})
	return file_test_types_types_proto_rawDescData
}

var file_test_types_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_types_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_test_types_types_proto_goTypes = []interface{}{
	(Types_Enum)(0),               // 0: types.Types.Enum
	(*Types)(nil),                 // 1: types.Types
	(*OtherType1)(nil),            // 2: types.OtherType1
	(*OtherType2)(nil),            // 3: types.OtherType2
	(*Types_NestedType)(nil),      // 4: types.Types.NestedType
	nil,                           // 5: types.Types.MapVal1Entry
	nil,                           // 6: types.Types.MapVal2Entry
	nil,                           // 7: types.Types.MapEmptyValEntry
	(*OtherType2_NestedType)(nil), // 8: types.OtherType2.NestedType
	(*_struct.Struct)(nil),        // 9: google.protobuf.Struct
}
var file_test_types_types_proto_depIdxs = []int32{
	0,  // 0: types.Types.enum_val:type_name -> types.Types.Enum
	2,  // 1: types.Types.other_type_val:type_name -> types.OtherType1
	4,  // 2: types.Types.nested_type_val:type_name -> types.Types.NestedType
	8,  // 3: types.Types.other_type_nested_type_val:type_name -> types.OtherType2.NestedType
	5,  // 4: types.Types.map_val1:type_name -> types.Types.MapVal1Entry
	6,  // 5: types.Types.map_val2:type_name -> types.Types.MapVal2Entry
	7,  // 6: types.Types.map_empty_val:type_name -> types.Types.MapEmptyValEntry
	0,  // 7: types.Types.repeated_val2:type_name -> types.Types.Enum
	2,  // 8: types.Types.repeated_val3:type_name -> types.OtherType1
	9,  // 9: types.Types.struct_val:type_name -> google.protobuf.Struct
	2,  // 10: types.Types.MapVal2Entry.value:type_name -> types.OtherType1
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_test_types_types_proto_init() }
func file_test_types_types_proto_init() {
	if File_test_types_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_types_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Types); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_types_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherType1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_types_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherType2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_types_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Types_NestedType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_types_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherType2_NestedType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_test_types_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Types_OneofStringVal)(nil),
		(*Types_OneofInt64Val)(nil),
		(*Types_OneofBoolVal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_types_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_types_types_proto_goTypes,
		DependencyIndexes: file_test_types_types_proto_depIdxs,
		EnumInfos:         file_test_types_types_proto_enumTypes,
		MessageInfos:      file_test_types_types_proto_msgTypes,
	}.Build()
	File_test_types_types_proto = out.File
	file_test_types_types_proto_rawDesc = nil
	file_test_types_types_proto_goTypes = nil
	file_test_types_types_proto_depIdxs = nil
}
