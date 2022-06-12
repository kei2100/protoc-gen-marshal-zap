package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestTypes_MarshalLogObject(t *testing.T) {
	m := &Types{
		SecretVal:   "secret",
		DoubleVal:   0.1,
		FloatVal:    0.1,
		Int32Val:    1,
		Int64Val:    1,
		Uint32Val:   1,
		Uint64Val:   1,
		Sint32Val:   2,
		Sint64Val:   2,
		Fixed32Val:  2,
		Fixed64Val:  2,
		Sfixed32Val: 3,
		Sfixed64Val: 3,
		BoolVal:     true,
		StringVal:   "string",
		BytesVal:    []byte{1, 2, 3},
		EnumVal:     Types_ENUM_1,
		OtherTypeVal: &OtherType1{
			OtherStringVal: "other_string",
			OtherSecretVal: "other_secret",
		},
		NestedTypeVal: &Types_NestedType{
			NestedStringVal: "nested_string",
			NestedSecretVal: "nested_secret",
		},
		OtherTypeNestedTypeVal: &OtherType2_NestedType{
			NestedStringVal: "other_nested_string",
			NestedSecretVal: "other_nested_secret",
		},
		OneofVal: &Types_OneofStringVal{
			OneofStringVal: "oneof_string",
		},
		MapVal1: map[string]string{
			"foo": "bar",
		},
		MapVal2: map[string]*OtherType1{
			"foo": {
				OtherStringVal: "other_string",
				OtherSecretVal: "other_secret",
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
				OtherStringVal: "other_string",
				OtherSecretVal: "other_secret",
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
		XString: "foo",
	}
	enc := zapcore.NewMapObjectEncoder()
	err := m.MarshalLogObject(enc)
	if !assert.NoError(t, err) {
		return
	}
	assert.EqualValues(t, enc.Fields, map[string]interface{}{
		"secret_val":   "[MASKED]",
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
		"oneof_string_val": "oneof_string",
		"oneof_int64_val":  int64(0),
		"oneof_bool_val":   false,
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
		"_String": "foo",
	})
}
