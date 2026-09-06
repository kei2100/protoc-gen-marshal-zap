package main

import (
	"fmt"

	pbzap "github.com/kei2100/protoc-gen-marshal-zap"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	zapcorePkg = protogen.GoImportPath("go.uber.org/zap/zapcore")
	fmtPkg     = protogen.GoImportPath("fmt")
)

// wellKnownType describes how to encode a google.protobuf well-known type.
//
// The principle is to produce the same output that zap would produce for the
// equivalent Go value: e.g. Timestamp is encoded via AddTime so that it follows
// the EncoderConfig.EncodeTime setting, and wrapper types are unwrapped to
// their primitive values.
type wellKnownType struct {
	// addMethod is the zapcore.ObjectEncoder method name (e.g. "AddTime").
	addMethod string
	// appendMethod is the zapcore.ArrayEncoder method name (e.g. "AppendTime").
	appendMethod string
	// valueExpr returns the Go expression passed to the encoder method.
	// v is the Go expression of the (non-nil) well-known type message.
	valueExpr func(g *protogen.GeneratedFile, v string) string
}

func unwrapValue(_ *protogen.GeneratedFile, v string) string { return v + ".GetValue()" }

func asTime(_ *protogen.GeneratedFile, v string) string { return v + ".AsTime()" }

func asDuration(_ *protogen.GeneratedFile, v string) string { return v + ".AsDuration()" }

// fieldMaskArray returns an ArrayMarshalerFunc that appends each path of the FieldMask.
func fieldMaskArray(g *protogen.GeneratedFile, v string) string {
	return fmt.Sprintf("%s(func(enc %s) error { for _, p := range %s.GetPaths() { enc.AppendString(p) }; return nil })",
		g.QualifiedGoIdent(zapcorePkg.Ident("ArrayMarshalerFunc")),
		g.QualifiedGoIdent(zapcorePkg.Ident("ArrayEncoder")),
		v)
}

// emptyObject returns an ObjectMarshalerFunc that encodes nothing (`{}`).
func emptyObject(g *protogen.GeneratedFile, _ string) string {
	return fmt.Sprintf("%s(func(%s) error { return nil })",
		g.QualifiedGoIdent(zapcorePkg.Ident("ObjectMarshalerFunc")),
		g.QualifiedGoIdent(zapcorePkg.Ident("ObjectEncoder")))
}

// Any, Struct, Value and ListValue are intentionally not listed here and
// fall back to the generic message handling (ObjectMarshaler / AddReflected).
var wellKnownTypes = map[protoreflect.FullName]wellKnownType{
	"google.protobuf.Timestamp":   {addMethod: "AddTime", appendMethod: "AppendTime", valueExpr: asTime},
	"google.protobuf.Duration":    {addMethod: "AddDuration", appendMethod: "AppendDuration", valueExpr: asDuration},
	"google.protobuf.BoolValue":   {addMethod: "AddBool", appendMethod: "AppendBool", valueExpr: unwrapValue},
	"google.protobuf.StringValue": {addMethod: "AddString", appendMethod: "AppendString", valueExpr: unwrapValue},
	// zapcore.ArrayEncoder has no AppendBinary, so repeated elements are appended as a string
	// (same as a `repeated bytes` field).
	"google.protobuf.BytesValue":  {addMethod: "AddBinary", appendMethod: "AppendByteString", valueExpr: unwrapValue},
	"google.protobuf.Int32Value":  {addMethod: "AddInt32", appendMethod: "AppendInt32", valueExpr: unwrapValue},
	"google.protobuf.Int64Value":  {addMethod: "AddInt64", appendMethod: "AppendInt64", valueExpr: unwrapValue},
	"google.protobuf.UInt32Value": {addMethod: "AddUint32", appendMethod: "AppendUint32", valueExpr: unwrapValue},
	"google.protobuf.UInt64Value": {addMethod: "AddUint64", appendMethod: "AppendUint64", valueExpr: unwrapValue},
	"google.protobuf.FloatValue":  {addMethod: "AddFloat32", appendMethod: "AppendFloat32", valueExpr: unwrapValue},
	"google.protobuf.DoubleValue": {addMethod: "AddFloat64", appendMethod: "AppendFloat64", valueExpr: unwrapValue},
	// FieldMask -> array of path strings, same as a `repeated string` field.
	"google.protobuf.FieldMask": {addMethod: "AddArray", appendMethod: "AppendArray", valueExpr: fieldMaskArray},
	// Empty -> empty object `{}`.
	"google.protobuf.Empty": {addMethod: "AddObject", appendMethod: "AppendObject", valueExpr: emptyObject},
}

// lookupWellKnownType returns the well-known type definition for md, if any.
func lookupWellKnownType(md protoreflect.MessageDescriptor) (wellKnownType, bool) {
	wkt, ok := wellKnownTypes[md.FullName()]
	return wkt, ok
}

func generateListField(g *protogen.GeneratedFile, f *protogen.Field) {
	fname := f.Desc.Name()
	g.P(fname, "ArrMarshaller := func(enc ", g.QualifiedGoIdent(zapcorePkg.Ident("ArrayEncoder")), ") error {")
	g.P("for _, v := range x.", f.GoName, " {")
	switch f.Desc.Kind() {
	case protoreflect.BoolKind:
		g.P("enc.AppendBool(v)")
	case protoreflect.BytesKind:
		g.P("enc.AppendByteString(v)")
	case protoreflect.DoubleKind:
		g.P("enc.AppendFloat64(v)")
	case protoreflect.EnumKind:
		g.P("enc.AppendString(v.String())")
	case protoreflect.Fixed32Kind, protoreflect.Uint32Kind:
		g.P("enc.AppendUint32(v)")
	case protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
		g.P("enc.AppendUint64(v)")
	case protoreflect.FloatKind:
		g.P("enc.AppendFloat32(v)")
	case protoreflect.Int32Kind, protoreflect.Sfixed32Kind, protoreflect.Sint32Kind:
		g.P("enc.AppendInt32(v)")
	case protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Sint64Kind:
		g.P("enc.AppendInt64(v)")
	case protoreflect.GroupKind:
		g.P("enc.AppendReflected(v)")
	case protoreflect.MessageKind:
		if wkt, ok := lookupWellKnownType(f.Desc.Message()); ok {
			g.P("if v != nil {")
			g.P("enc.", wkt.appendMethod, "(", wkt.valueExpr(g, "v"), ")")
			g.P("} else {")
			g.P("enc.AppendReflected(nil)")
			g.P("}")
			break
		}
		g.P("if obj, ok := interface{}(v).(", g.QualifiedGoIdent(zapcorePkg.Ident("ObjectMarshaler")), "); ok {")
		g.P("enc.AppendObject(obj)")
		g.P("} else {")
		g.P("enc.AppendReflected(v)")
		g.P("}")
	case protoreflect.StringKind:
		g.P("enc.AppendString(v)")
	default:
		g.P("enc.AppendReflected(v)")
	}
	g.P("}")
	g.P("return nil")
	g.P("}")
	g.P("enc.AddArray(\"", fname, "\",", g.QualifiedGoIdent(zapcorePkg.Ident("ArrayMarshalerFunc")), "(", fname, "ArrMarshaller))")
}

func generateMapField(g *protogen.GeneratedFile, f *protogen.Field) {
	fname := f.Desc.Name()
	g.P("enc.AddObject(\"", fname, "\", ", g.QualifiedGoIdent(zapcorePkg.Ident("ObjectMarshalerFunc")), "(func(enc ", g.QualifiedGoIdent(zapcorePkg.Ident("ObjectEncoder")), ") error {")
	g.P("for k, v := range x.", f.GoName, " {")
	switch f.Desc.MapValue().Kind() {
	case protoreflect.BoolKind:
		g.P("enc.AddBool(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.BytesKind:
		g.P("enc.AddBinary(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.DoubleKind:
		g.P("enc.AddFloat64(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.EnumKind:
		g.P("enc.AddString(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v.String())")
	case protoreflect.Fixed32Kind, protoreflect.Uint32Kind:
		g.P("enc.AddUint32(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
		g.P("enc.AddUint64(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.FloatKind:
		g.P("enc.AddFloat32(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.Int32Kind, protoreflect.Sfixed32Kind, protoreflect.Sint32Kind:
		g.P("enc.AddInt32(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Sint64Kind:
		g.P("enc.AddInt64(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.GroupKind:
		g.P("enc.AddReflected(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	case protoreflect.MessageKind:
		if wkt, ok := lookupWellKnownType(f.Desc.MapValue().Message()); ok {
			key := g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")) + "(\"%v\", k)"
			g.P("if v != nil {")
			g.P("enc.", wkt.addMethod, "(", key, ", ", wkt.valueExpr(g, "v"), ")")
			g.P("} else {")
			g.P("enc.AddReflected(", key, ", nil)")
			g.P("}")
			break
		}
		g.P("if obj, ok := interface{}(v).(", g.QualifiedGoIdent(zapcorePkg.Ident("ObjectMarshaler")), "); ok {")
		g.P("enc.AddObject(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), obj)")
		g.P("} else {")
		g.P("enc.AddReflected(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
		g.P("}")
	case protoreflect.StringKind:
		g.P("enc.AddString(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	default:
		g.P("enc.AddReflected(", g.QualifiedGoIdent(fmtPkg.Ident("Sprintf")), "(\"%v\", k), v)")
	}
	g.P("}")
	g.P("return nil")
	g.P("}))")
}

func generatePrimitiveField(g *protogen.GeneratedFile, f *protogen.Field) {
	fname := f.Desc.Name()
	var gname string
	if f.Oneof != nil {
		gname = fmt.Sprintf("Get%s()", f.GoName)
	} else {
		gname = f.GoName
	}
	switch f.Desc.Kind() {
	case protoreflect.BoolKind:
		g.P("enc.AddBool(\"", fname, "\", x.", gname, ")")
	case protoreflect.BytesKind:
		g.P("enc.AddBinary(\"", fname, "\", x.", gname, ")")
	case protoreflect.DoubleKind:
		g.P("enc.AddFloat64(\"", fname, "\", x.", gname, ")")
	case protoreflect.EnumKind:
		g.P("enc.AddString(\"", fname, "\", x.", gname, ".String())")
	case protoreflect.Fixed32Kind, protoreflect.Uint32Kind:
		g.P("enc.AddUint32(\"", fname, "\", x.", gname, ")")
	case protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
		g.P("enc.AddUint64(\"", fname, "\", x.", gname, ")")
	case protoreflect.FloatKind:
		g.P("enc.AddFloat32(\"", fname, "\", x.", gname, ")")
	case protoreflect.Int32Kind, protoreflect.Sfixed32Kind, protoreflect.Sint32Kind:
		g.P("enc.AddInt32(\"", fname, "\", x.", gname, ")")
	case protoreflect.Int64Kind, protoreflect.Sfixed64Kind, protoreflect.Sint64Kind:
		g.P("enc.AddInt64(\"", fname, "\", x.", gname, ")")
	case protoreflect.GroupKind:
		g.P("enc.AddReflected(\"", fname, "\", x.", gname, ")")
	case protoreflect.MessageKind:
		if wkt, ok := lookupWellKnownType(f.Desc.Message()); ok {
			// nil is guarded by handleExplicitPresence for singular and optional fields.
			// For a oneof member, only the wrapper type is checked; a nil inner pointer is
			// treated as the zero message (AsTime, AsDuration, GetValue and GetPaths are nil-safe).
			g.P("enc.", wkt.addMethod, "(\"", fname, "\", ", wkt.valueExpr(g, "x."+gname), ")")
			break
		}
		g.P("if obj, ok := interface{}(x.", gname, ").(", g.QualifiedGoIdent(zapcorePkg.Ident("ObjectMarshaler")), "); ok {")
		g.P("enc.AddObject(\"", fname, "\", obj)")
		g.P("} else {")
		g.P("enc.AddReflected(\"", fname, "\", x.", gname, ")")
		g.P("}")
	case protoreflect.StringKind:
		g.P("enc.AddString(\"", fname, "\", x.", gname, ")")
	default:
		g.P("enc.AddReflected(\"", fname, "\", x.", gname, ")")
	}
}

func isMasked(opts *descriptorpb.FieldOptions) bool {
	return proto.GetExtension(opts, pbzap.E_Mask).(bool) || opts.GetDebugRedact()
}

func handleExplicitPresence(g *protogen.GeneratedFile, f *protogen.Field, generateFunc func(*protogen.GeneratedFile, *protogen.Field)) {
	// Omit the fields that are defined as `Explicit Presence` and the value is not present.
	// https://protobuf.dev/programming-guides/field_presence/#presence-in-proto3-apis
	switch {
	case f.Oneof != nil && f.Desc.HasOptionalKeyword():
		// handle optional fields
		g.P("if x.", f.GoName, " != nil {")
		defer g.P("}")
	case f.Oneof != nil && !f.Desc.HasOptionalKeyword():
		// handle oneof fields
		g.P("if _, ok := x.Get", f.Oneof.GoName, "().(*", f.GoIdent, "); ok {")
		defer g.P("}")
	case f.Desc.Kind() == protoreflect.MessageKind || f.Desc.Kind() == protoreflect.GroupKind:
		// handle message fields
		g.P("if x.", f.GoName, " != nil {")
		defer g.P("}")
	}
	generateFunc(g, f)
}

func generateMessage(g *protogen.GeneratedFile, m *protogen.Message) {
	ident := g.QualifiedGoIdent(m.GoIdent)
	g.P("func (x *", ident, ") MarshalLogObject(enc ", g.QualifiedGoIdent(zapcorePkg.Ident("ObjectEncoder")), ") error {")
	g.P("if x == nil {")
	g.P("return nil")
	g.P("}")
	g.P()
	for _, f := range m.Fields {
		if isMasked(f.Desc.Options().(*descriptorpb.FieldOptions)) {
			g.P("enc.AddString(\"", f.Desc.Name(), "\", \"[MASKED]\")")
		} else if f.Desc.IsList() {
			generateListField(g, f)
		} else if f.Desc.IsMap() {
			generateMapField(g, f)
		} else {
			handleExplicitPresence(g, f, generatePrimitiveField)
		}
		g.P()
	}
	g.P("return nil")
	g.P("}")
	g.P()
	for _, submsg := range m.Messages {
		if submsg.Desc.IsMapEntry() {
			continue
		}
		generateMessage(g, submsg)
	}
}

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Messages) == 0 {
		return nil
	}

	filename := fmt.Sprintf("%s.pb.marshal-zap.go", file.GeneratedFilenamePrefix)
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-marshal-zap. DO NOT EDIT.")
	g.P("//")
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, m := range file.Messages {
		generateMessage(g, m)
	}

	return g
}

func main() {
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, file := range plugin.FilesByPath {
			if !file.Generate {
				continue
			}
			generateFile(plugin, file)
		}
		return nil
	})
}
