package registry

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type GoPackage struct {
	Path  string
	Name  string
	Alias string
}

func (p *GoPackage) HasStdLib() bool {
	return !strings.Contains(p.Path, ".")
}

func (p *GoPackage) String() string {
	if p.Alias == "" {
		return fmt.Sprintf("%q", p.Path)
	}
	return fmt.Sprintf("%s %q", p.Alias, p.Path)
}

type File struct {
	*descriptor.FileDescriptorProto
	GoPkg    *GoPackage
	Messages []*Message
	Enums    []*Enum
	Services []*Service
}

type Message struct {
	*descriptor.DescriptorProto
	File   *File
	Outers []string
	Fields []*Field
	Index  int
	Type   string
}

func (m *Message) FieldName() string {
	return generator.CamelCase(m.GetName())
}

func (m *Message) GetSchemaTypeName() string {
	switch m.Type {
	case ".google.protobuf.Empty":
		return "scalar.Empty"
	default:
		name := m.FieldName()
		name = strings.ToLower(name[0:1]) + name[1:]
		return fmt.Sprintf("%sType", name)
	}
}

func (m *Message) MessageName() string {
	components := []string{""}
	if m.File.Package != nil {
		components = append(components, m.File.GetPackage())
	}
	components = append(components, m.Outers...)
	components = append(components, m.GetName())
	return strings.Join(components, ".")
}

type Enum struct {
	*descriptor.EnumDescriptorProto
	File   *File
	Outers []string
	Index  int
}

func (e *Enum) EnumName() string {
	components := []string{""}
	if e.File.Package != nil {
		components = append(components, e.File.GetPackage())
	}
	components = append(components, e.Outers...)
	components = append(components, e.GetName())
	return strings.Join(components, ".")
}

type Service struct {
	*descriptor.ServiceDescriptorProto
	File    *File
	Methods []*Method
}

type Method struct {
	*descriptor.MethodDescriptorProto
	Service  *Service
	Request  *Message
	Response *Message
	// GraphQL plugin options
	Description string
	FieldName   string
	Query       bool
	Mutation    bool
	Subscribe   bool
}

func (m *Method) FullMethod() string {
	return "/" + m.Service.File.GetPackage() + "." + m.Service.GetName() + "/" + m.GetName()
}

func (m *Method) HasGraphQLMethod() bool {
	return m.Query || m.Mutation || m.Subscribe
}

func (m *Method) Variable() string {
	name := generator.CamelCase(m.GetName()) + "Field"
	name = strings.ToLower(name[0:1]) + name[1:]
	return name
}

type Field struct {
	*descriptor.FieldDescriptorProto
	Message      *Message
	FieldMessage *Message
	Enum         *Enum
	// GraphQL plugin options
	Description string
	Nullable    *bool
	Dependence  *Message
}

func (f *Field) IsMessageType() bool {
	return f.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func (f *Field) IsEnumType() bool {
	return f.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM
}

func (f *Field) isRepeated() bool {
	return f.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED
}

func (f *Field) FieldName() string {
	name := generator.CamelCase(f.GetName())
	if isGoProtoMethod[name] {
		name += "_"
	}
	return name
}

func (f *Field) Variable() string {
	return "value" + f.FieldName()
}

func (f *Field) GoType() string {
	var typ string
	switch f.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		typ = "float64"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		typ = "float32"
	case
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		typ = "int64"
	case
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:
		typ = "uint64"
	case
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		typ = "int32"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		typ = "bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		typ = "string"
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		// unknown fields
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		if f.TypeName == nil {
			return "nil"
		}
		prefix := "*"
		if f.Nullable != nil && !*f.Nullable {
			prefix = ""
		}
		var depType string
		switch f.Dependence.Type {
		case ".google.protobuf.Timestamp":
			depType = "timestamp.Timestamp"
		default:
			idx := strings.LastIndex(f.GetTypeName(), ".")
			depType = f.GetTypeName()[idx+1:]
		}
		typ = prefix + depType
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		typ = "[]byte"
	case
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		typ = "uint32"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		name := f.GetTypeName()
		if strings.HasPrefix(name, ".") {
			name = name[1:]
		}
		vars := strings.SplitN(name, ".", 2)
		imp := vars[0]
		if imp == f.Message.File.GoPkg.Name {
			imp = ""
		} else {
			imp += "."
		}
		enumType := strings.ReplaceAll(vars[1], ".", "_")
		typ = imp + enumType
	default:
		typ = "nil"
	}
	if f.isRepeated() {
		return "[]" + typ
	}
	return typ
}

func (f *Field) ScalarType() string {
	nullable := f.Nullable == nil || *f.Nullable
	var scalar string
	switch f.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		scalar = "scalar.Float64"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		scalar = "scalar.Float32"
	case
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		scalar = "scalar.Int64"
	case
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:
		scalar = "scalar.Uint64"
	case
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		scalar = "scalar.Int32"
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		scalar = "scalar.Bool"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		scalar = "scalar.String"
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		// unknown fields
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		if f.Dependence != nil {
			switch f.Dependence.Type {
			case ".google.protobuf.Empty":
				scalar = "scalar.Empty"
			case ".google.protobuf.Timestamp":
				scalar = "scalar.Timestamp"
			case ".google.protobuf.DoubleValue":
				scalar = "scalar.Float64"
			case ".google.protobuf.FloatValue":
				scalar = "scalar.Float32"
			case ".google.protobuf.Int64Value":
				scalar = "scalar.Int64"
			case ".google.protobuf.UInt64Value":
				scalar = "scalar.Uint64"
			case ".google.protobuf.Int32Value":
				scalar = "scalar.Int32"
			case ".google.protobuf.UInt32Value":
				scalar = "scalar.Uint32"
			case ".google.protobuf.BoolValue":
				scalar = "scalar.Bool"
			case ".google.protobuf.StringValue":
				scalar = "scalar.String"
			default:
				scalar = f.Dependence.GetSchemaTypeName()
			}
		} else {
			scalar = f.Message.GetSchemaTypeName()
		}
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		scalar = "scalar.Bytes"
	case
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		scalar = "scalar.Uint32"
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		scalar = "scalar.Int32"
	default:
		scalar = "scalar.Nil"
	}
	if !nullable {
		scalar = fmt.Sprintf("graphql.NewNonNull(%s)", scalar)
	}
	if f.isRepeated() {
		scalar = fmt.Sprintf("graphql.NewList(%s)", scalar)
	}
	return scalar
}

var isGoProtoMethod = map[string]bool{
	"Reset":               true,
	"String":              true,
	"ProtoMessage":        true,
	"Marshal":             true,
	"Unmarshal":           true,
	"ExtensionRangeArray": true,
	"ExtensionMap":        true,
	"Descriptor":          true,
}
