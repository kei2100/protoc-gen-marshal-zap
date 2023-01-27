// Code generated by protoc-gen-marshal-zap. DO NOT EDIT.
//
// source: test/types/types.proto

package types

import (
	fmt "fmt"
	zapcore "go.uber.org/zap/zapcore"
)

func (x *Types) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

	enc.AddString("secret_val", "[MASKED]")

	enc.AddFloat64("double_val", x.DoubleVal)

	enc.AddFloat32("float_val", x.FloatVal)

	enc.AddInt32("int32_val", x.Int32Val)

	enc.AddInt64("int64_val", x.Int64Val)

	enc.AddUint32("uint32_val", x.Uint32Val)

	enc.AddUint64("uint64_val", x.Uint64Val)

	enc.AddInt32("sint32_val", x.Sint32Val)

	enc.AddInt64("sint64_val", x.Sint64Val)

	enc.AddUint32("fixed32_val", x.Fixed32Val)

	enc.AddUint64("fixed64_val", x.Fixed64Val)

	enc.AddInt32("sfixed32_val", x.Sfixed32Val)

	enc.AddInt64("sfixed64_val", x.Sfixed64Val)

	enc.AddBool("bool_val", x.BoolVal)

	enc.AddString("string_val", x.StringVal)

	enc.AddBinary("bytes_val", x.BytesVal)

	enc.AddString("enum_val", x.EnumVal.String())

	if obj, ok := interface{}(x.OtherTypeVal).(zapcore.ObjectMarshaler); ok {
		enc.AddObject("other_type_val", obj)
	} else {
		enc.AddReflected("other_type_val", x.OtherTypeVal)
	}

	if obj, ok := interface{}(x.NestedTypeVal).(zapcore.ObjectMarshaler); ok {
		enc.AddObject("nested_type_val", obj)
	} else {
		enc.AddReflected("nested_type_val", x.NestedTypeVal)
	}

	if obj, ok := interface{}(x.OtherTypeNestedTypeVal).(zapcore.ObjectMarshaler); ok {
		enc.AddObject("other_type_nested_type_val", obj)
	} else {
		enc.AddReflected("other_type_nested_type_val", x.OtherTypeNestedTypeVal)
	}

	if _, ok := x.GetOneofVal().(*Types_OneofStringVal); ok {
		enc.AddString("oneof_string_val", x.GetOneofStringVal())
	}

	if _, ok := x.GetOneofVal().(*Types_OneofInt64Val); ok {
		enc.AddInt64("oneof_int64_val", x.GetOneofInt64Val())
	}

	if _, ok := x.GetOneofVal().(*Types_OneofBoolVal); ok {
		enc.AddBool("oneof_bool_val", x.GetOneofBoolVal())
	}

	enc.AddObject("map_val1", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
		for k, v := range x.MapVal1 {
			enc.AddString(fmt.Sprintf("%v", k), v)
		}
		return nil
	}))

	enc.AddObject("map_val2", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
		for k, v := range x.MapVal2 {
			if obj, ok := interface{}(v).(zapcore.ObjectMarshaler); ok {
				enc.AddObject(fmt.Sprintf("%v", k), obj)
			} else {
				enc.AddReflected(fmt.Sprintf("%v", k), v)
			}
		}
		return nil
	}))

	enc.AddObject("map_empty_val", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
		for k, v := range x.MapEmptyVal {
			enc.AddString(fmt.Sprintf("%v", k), v)
		}
		return nil
	}))

	repeated_val1ArrMarshaller := func(enc zapcore.ArrayEncoder) error {
		for _, v := range x.RepeatedVal1 {
			enc.AppendString(v)
		}
		return nil
	}
	enc.AddArray("repeated_val1", zapcore.ArrayMarshalerFunc(repeated_val1ArrMarshaller))

	repeated_val2ArrMarshaller := func(enc zapcore.ArrayEncoder) error {
		for _, v := range x.RepeatedVal2 {
			enc.AppendString(v.String())
		}
		return nil
	}
	enc.AddArray("repeated_val2", zapcore.ArrayMarshalerFunc(repeated_val2ArrMarshaller))

	repeated_val3ArrMarshaller := func(enc zapcore.ArrayEncoder) error {
		for _, v := range x.RepeatedVal3 {
			if obj, ok := interface{}(v).(zapcore.ObjectMarshaler); ok {
				enc.AppendObject(obj)
			} else {
				enc.AppendReflected(v)
			}
		}
		return nil
	}
	enc.AddArray("repeated_val3", zapcore.ArrayMarshalerFunc(repeated_val3ArrMarshaller))

	repeated_empty_valArrMarshaller := func(enc zapcore.ArrayEncoder) error {
		for _, v := range x.RepeatedEmptyVal {
			enc.AppendString(v)
		}
		return nil
	}
	enc.AddArray("repeated_empty_val", zapcore.ArrayMarshalerFunc(repeated_empty_valArrMarshaller))

	if obj, ok := interface{}(x.StructVal).(zapcore.ObjectMarshaler); ok {
		enc.AddObject("struct_val", obj)
	} else {
		enc.AddReflected("struct_val", x.StructVal)
	}

	enc.AddString("_String", x.XString)

	if x.OptionalVal != nil {
		enc.AddString("optional_val", x.GetOptionalVal())
	}

	if x.OptionalNotPresentVal != nil {
		enc.AddString("optional_not_present_val", x.GetOptionalNotPresentVal())
	}

	if x.OptionalEnum != nil {
		enc.AddString("optional_enum", x.GetOptionalEnum().String())
	}

	if x.OptionalNotPresentEnum != nil {
		enc.AddString("optional_not_present_enum", x.GetOptionalNotPresentEnum().String())
	}

	if x.OptionalMessage != nil {
		if obj, ok := interface{}(x.GetOptionalMessage()).(zapcore.ObjectMarshaler); ok {
			enc.AddObject("optional_message", obj)
		} else {
			enc.AddReflected("optional_message", x.GetOptionalMessage())
		}
	}

	if x.OptionalNotPresentMessage != nil {
		if obj, ok := interface{}(x.GetOptionalNotPresentMessage()).(zapcore.ObjectMarshaler); ok {
			enc.AddObject("optional_not_present_message", obj)
		} else {
			enc.AddReflected("optional_not_present_message", x.GetOptionalNotPresentMessage())
		}
	}

	return nil
}

func (x *Types_NestedType) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

	enc.AddString("nested_string_val", x.NestedStringVal)

	enc.AddString("nested_secret_val", "[MASKED]")

	return nil
}

func (x *OtherType1) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

	enc.AddString("other_string_val", x.OtherStringVal)

	enc.AddString("other_secret_val", "[MASKED]")

	return nil
}

func (x *OtherType2) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

	return nil
}

func (x *OtherType2_NestedType) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

	enc.AddString("nested_string_val", x.NestedStringVal)

	enc.AddString("nested_secret_val", "[MASKED]")

	return nil
}

func (x *OtherType3) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

	enc.AddString("val", x.Val)

	return nil
}
