package proto2types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestTypes_MarshalLogObject(t *testing.T) {
	m := &Types{
		SecretVal:   refs("secret"),
		DoubleVal:   refs(0.1),
		FloatVal:    refs(float32(0.1)),
		Int32Val:    refs(int32(1)),
		Int64Val:    refs(int64(1)),
		Uint32Val:   refs(uint32(1)),
		Uint64Val:   refs(uint64(1)),
		Sint32Val:   refs(int32(2)),
		Sint64Val:   refs(int64(2)),
		Fixed32Val:  refs(uint32(2)),
		Fixed64Val:  refs(uint64(2)),
		Sfixed32Val: refs(int32(3)),
		Sfixed64Val: refs(int64(3)),
		BoolVal:     refs(true),
		StringVal:   refs("string"),
		BytesVal:    []byte{1, 2, 3},
		EnumVal:     Types_ENUM_1.Enum(),
		OtherTypeVal: &OtherType1{
			OtherStringVal: refs("other_string"),
			OtherSecretVal: refs("other_secret"),
		},
		NestedTypeVal: &Types_NestedType{
			NestedStringVal: refs("nested_string"),
			NestedSecretVal: refs("nested_secret"),
		},
		OtherTypeNestedTypeVal: &OtherType2_NestedType{
			NestedStringVal: refs("other_nested_string"),
			NestedSecretVal: refs("other_nested_secret"),
		},
		OneofVal: &Types_OneofStringVal{
			OneofStringVal: "", // set zero value explicitly
		},
		MapVal1: map[string]string{
			"foo": "bar",
		},
		MapVal2: map[string]*OtherType1{
			"foo": {
				OtherStringVal: refs("other_string"),
				OtherSecretVal: refs("other_secret"),
			},
		},
		RepeatedVal1: []string{
			"foo", "bar",
		},
		RepeatedVal2: []Types_Enum{
			Types_ENUM_1, Types_ENUM_2,
		},
		RepeatedVal3: []*OtherType1{
			{
				OtherStringVal: refs("other_string"),
				OtherSecretVal: refs("other_secret"),
			},
		},
		RepeatedEmptyVal: []string{},
		StructVal: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"struct_string": {
					Kind: &structpb.Value_StringValue{StringValue: "foo"},
				},
				"struct_number": {
					Kind: &structpb.Value_NumberValue{NumberValue: 0.1},
				},
				"struct_bool": {
					Kind: &structpb.Value_BoolValue{BoolValue: true},
				},
			},
		},
		XString:                   refs("foo"),
		OptionalVal:               refs("foo"),
		OptionalNotPresentVal:     nil,
		OptionalEnum:              refs(Types_ENUM_1),
		OptionalNotPresentEnum:    nil,
		OptionalMessage:           &OtherType3{Val: refs("foo")},
		OptionalNotPresentMessage: nil,
		PresentMessage:            &OtherType3{Val: refs("foo")},
		NotPresentMessage:         nil,
	}
	enc := zapcore.NewMapObjectEncoder()
	err := m.MarshalLogObject(enc)
	if !assert.NoError(t, err) {
		return
	}
	assert.EqualValues(t, map[string]interface{}{
		"secret_val":   "[MASKED]",
		"secret_val2":  "[MASKED]",
		"double_val":   float64(0.1),
		"float_val":    float32(0.1),
		"int32_val":    int32(1),
		"int64_val":    int64(1),
		"uint32_val":   uint32(1),
		"uint64_val":   uint64(1),
		"sint32_val":   int32(2),
		"sint64_val":   int64(2),
		"fixed32_val":  uint32(2),
		"fixed64_val":  uint64(2),
		"sfixed32_val": int32(3),
		"sfixed64_val": int64(3),
		"bool_val":     true,
		"string_val":   "string",
		"bytes_val":    []byte{1, 2, 3},
		"enum_val":     Types_ENUM_1.String(),
		"other_type_val": map[string]interface{}{
			"other_string_val": "other_string",
			"other_secret_val": "[MASKED]",
		},
		"nested_type_val": map[string]interface{}{
			"nested_string_val": "nested_string",
			"nested_secret_val": "[MASKED]",
		},
		"other_type_nested_type_val": map[string]interface{}{
			"nested_string_val": "other_nested_string",
			"nested_secret_val": "[MASKED]",
		},
		"oneof_string_val": "",
		"map_val1": map[string]interface{}{
			"foo": "bar",
		},
		"map_val2": map[string]interface{}{
			"foo": map[string]interface{}{
				"other_string_val": "other_string",
				"other_secret_val": "[MASKED]",
			},
		},
		"map_empty_val": map[string]interface{}{},
		"repeated_val1": []interface{}{"foo", "bar"},
		"repeated_val2": []interface{}{Types_ENUM_1.String(), Types_ENUM_2.String()},
		"repeated_val3": []interface{}{
			map[string]interface{}{
				"other_string_val": "other_string",
				"other_secret_val": "[MASKED]",
			},
		},
		"repeated_empty_val": []interface{}{},
		"struct_val": &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"struct_string": {
					Kind: &structpb.Value_StringValue{StringValue: "foo"},
				},
				"struct_number": {
					Kind: &structpb.Value_NumberValue{NumberValue: 0.1},
				},
				"struct_bool": {
					Kind: &structpb.Value_BoolValue{BoolValue: true},
				},
			},
		},
		"_String":       "foo",
		"optional_val":  "foo",
		"optional_enum": Types_ENUM_1.String(),
		"optional_message": map[string]interface{}{
			"val": "foo",
		},
		"present_message": map[string]interface{}{
			"val": "foo",
		},
	}, enc.Fields)

	assert.NotContains(t, enc.Fields, "oneof_int64_val")
	assert.NotContains(t, enc.Fields, "oneof_bool_val")
	assert.NotContains(t, enc.Fields, "optional_not_present_val")
	assert.NotContains(t, enc.Fields, "optional_not_present_enum")
	assert.NotContains(t, enc.Fields, "optional_not_present_message")
	assert.NotContains(t, enc.Fields, "not_present_message")
}

func refs[T any](v T) *T {
	return &v
}
