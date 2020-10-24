package main

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type protoFile struct {
	Path      string
	GoPackage string
	Messages  []*protoMessage
}

type protoMessage struct {
	Name   string
	Fields []*protoField
}

type protoField struct {
	// Name is protobuf-style field name (snake_case)
	Name string
	// Accessor name when marshaling this field
	Accessor   string
	Type       protoType
	IsRepeated bool
	IsMap      bool
	MapType    *mapType
	IsMask     bool
}

type mapType struct {
	KeyType   keyType
	ValueType protoType
}

type keyType struct {
	protoType
}

func (kt keyType) KeyToString(variableName string) string {
	// https://developers.google.com/protocol-buffers/docs/proto3#maps
	// where the key_type can be any integral or string type (so, any scalar type except for floating point types and bytes).
	// Note that enum is not a valid key_type
	switch descriptor.FieldDescriptorProto_Type(kt.protoType) {
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return variableName
	case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_SINT64, descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return fmt.Sprintf("strconv.FormatInt(%s, 10)", variableName)
	case descriptor.FieldDescriptorProto_TYPE_UINT64, descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return fmt.Sprintf("strconv.FormatUint(%s, 10)", variableName)
	default:
		return fmt.Sprintf("strconv.FormatInt(int64(%s), 10)", variableName)
	}
}

type protoType descriptor.FieldDescriptorProto_Type

func (t protoType) IsEnum() bool {
	return t == protoType(descriptor.FieldDescriptorProto_TYPE_ENUM)
}

func (t protoType) IsMessage() bool {
	return t == protoType(descriptor.FieldDescriptorProto_TYPE_MESSAGE)
}

func (t protoType) IsScalar() bool {
	return !t.IsEnum() && !t.IsMessage()
}

func (t protoType) ScalarName() string {
	switch descriptor.FieldDescriptorProto_Type(t) {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "Float64"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return "Float32"
	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_SINT32, descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		return "Int32"
	case descriptor.FieldDescriptorProto_TYPE_UINT32, descriptor.FieldDescriptorProto_TYPE_FIXED32:
		return "Uint32"
	case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_SINT64, descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return "Int64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64, descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return "Uint64"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		return "Bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "String"
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "Binary"
	}
	panic(fmt.Sprintf("unknown scalar type %s", t))
}

func (t protoType) String() string {
	return descriptor.FieldDescriptorProto_Type(t).String()
}

// # NOTE
//
// ### About `repeated`
//
// repeated scalar  -> supported
// repeated message -> supported
// repeated enum    -> supported
// repeated map<...>      -> NOT supported
// repeated repeated ...  -> NOT supported
//
const tmpl = `// Code generated by protoc-gen-marshal-zap. DO NOT EDIT.
// source: {{ .Path }}

package {{ .GoPackage }}

import  (
	"strconv"
	"go.uber.org/zap/zapcore"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = zapcore.NewNopCore
var _ = strconv.FormatInt

{{ range .Messages }}
func (x *{{ .Name }}) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}

{{ range .Fields }}
	{{ if .IsMask }}
		enc.AddString("{{ .Name }}", "[MASKED]")
	{{ else if .IsRepeated }}
		{{- /* NOTE */ -}}
		{{- /* repeated scalar  -> supported */ -}}
		{{- /* repeated message -> supported */ -}}
		{{- /* repeated enum    -> supported */ -}}
		{{- /* repeated map      -> NOT supported */ -}}
		{{- /* repeated repeated -> NOT supported */ -}}

		{{ .Name }}ArrMarshaller := func (enc zapcore.ArrayEncoder) error {
	        for _, v := range x.{{ .Accessor }} {
				{{ if .Type.IsScalar }}
	            	enc.Append{{ .Type.ScalarName }}(v)
				{{ else if .Type.IsEnum }}
	            	enc.AppendString(v.String())
				{{ else if and .Type.IsMessage }}
			    	if obj, ok := interface{}(v).(zapcore.ObjectMarshaler); ok {
			    	    enc.AppendObject(obj)
			    	} else {
			    	    enc.AppendReflected(v)
			    	}
				{{ else }}	
			    	enc.AppendReflected(v)
				{{ end }}	
	        }
	        return nil
	    }
	    enc.AddArray("{{ .Name }}", zapcore.ArrayMarshalerFunc({{ .Name }}ArrMarshaller))

	{{ else if .IsMap }}
		enc.AddObject("{{ .Name }}", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
			for k, v := range x.{{ .Accessor }} {
				{{ if .MapType.ValueType.IsScalar }}
					enc.Add{{ .MapType.ValueType.ScalarName }}({{ .MapType.KeyType.KeyToString "k" }}, v)
				{{ else if .MapType.ValueType.IsEnum }}
					enc.AddString({{ .MapType.KeyType.KeyToString "k" }}, v.String())
				{{ else if .MapType.ValueType.IsMessage }}
					if obj, ok := interface{}(v).(zapcore.ObjectMarshaler); ok {
						enc.AddObject({{ .MapType.KeyType.KeyToString "k" }}, obj)
					} else {
						enc.AddReflected({{ .MapType.KeyType.KeyToString "k" }}, v)
					}
				{{ else }}	
					enc.AddReflected({{ .MapType.KeyType.KeyToString "k" }}, v)
				{{ end }}	
			}
			return nil
		}))

	{{ else }}	
		{{ if .Type.IsScalar }}
			enc.Add{{ .Type.ScalarName }}("{{ .Name }}", x.{{ .Accessor }})
		{{ else if .Type.IsEnum }}
			enc.AddString("{{ .Name }}", x.{{ .Accessor }}.String())
		{{ else if and .Type.IsMessage }}
			if obj, ok := interface{}(x.{{ .Accessor }}).(zapcore.ObjectMarshaler); ok {
				enc.AddObject("{{ .Name }}", obj)
			} else {
				enc.AddReflected("{{ .Name }}", x.{{ .Accessor }})
			}
		{{ else }}	
			enc.AddReflected("{{ .Name }}", x.{{ .Accessor }})
		{{ end }}	
	{{ end }}	
{{ end }}	

	return nil
}
{{ end }}
`
