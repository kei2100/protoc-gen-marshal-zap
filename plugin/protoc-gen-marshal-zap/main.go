package main

import (
	"fmt"

	pbzap "github.com/kei2100/protoc-gen-marshal-zap"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	zapcorePkg = protogen.GoImportPath("go.uber.org/zap/zapcore")
	fmtPkg     = protogen.GoImportPath("fmt")
)

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
	g.P()
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
	g.P()
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
	g.P()
}

func isMasked(opts *descriptorpb.FieldOptions) bool {
	return proto.GetExtension(opts, pbzap.E_Mask).(bool)
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
			g.P()
		} else if f.Desc.IsList() {
			generateListField(g, f)
		} else if f.Desc.IsMap() {
			generateMapField(g, f)
		} else {
			generatePrimitiveField(g, f)
		}
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

	filename := fmt.Sprintf("%s_marshall_zap.pb.go", file.GeneratedFilenamePrefix)
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
		for _, file := range plugin.FilesByPath {
			if !file.Generate {
				continue
			}

			generateFile(plugin, file)
		}
		return nil
	})
}
