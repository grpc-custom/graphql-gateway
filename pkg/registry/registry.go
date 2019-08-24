package registry

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	options "github.com/grpc-custom/graphql-gateway/proto"
)

type Registry struct {
	Messages   map[string]*Message
	Enums      map[string]*Enum
	Files      map[string]*File
	Pkgs       map[string]string
	ImportPath string
	Prefix     string
}

func New() *Registry {
	return &Registry{
		Messages: make(map[string]*Message),
		Enums:    make(map[string]*Enum),
		Files:    make(map[string]*File),
		Pkgs:     make(map[string]string),
	}
}

func (r *Registry) Apply(req *plugin.CodeGeneratorRequest) error {
	for _, file := range req.GetProtoFile() {
		r.applyFile(file)
	}
	var targetPkg string
	for _, name := range req.FileToGenerate {
		target := r.Files[name]
		if target == nil {
			return fmt.Errorf("no such file: %s", name)
		}
		name := r.packageIdentityName(target.FileDescriptorProto)
		if targetPkg == "" {
			targetPkg = name
		} else {
			if targetPkg != name {
				return fmt.Errorf("inconsistent package names: %s %s", targetPkg, name)
			}
		}
		if err := r.applyServices(target); err != nil {
			return err
		}
	}

	return nil
}

func (r *Registry) applyFile(file *descriptor.FileDescriptorProto) {
	pkg := &GoPackage{
		Path: r.goPackagePath(file),
		Name: r.defaultGoPackageName(file),
	}
	f := &File{
		FileDescriptorProto: file,
		GoPkg:               pkg,
	}
	r.Files[file.GetName()] = f

	r.applyMsg(f, nil, file.GetMessageType())
	r.applyEnum(f, nil, file.GetEnumType())
}

func (r *Registry) applyMsg(file *File, outerPath []string, msgs []*descriptor.DescriptorProto) {
	for i, msg := range msgs {
		m := &Message{
			File:            file,
			Outers:          outerPath,
			DescriptorProto: msg,
			Index:           i,
		}
		for _, fd := range msg.GetField() {
			field := &Field{
				Message:              m,
				FieldDescriptorProto: fd,
			}
			if proto.HasExtension(fd.Options, options.E_Field) {
				ext, _ := proto.GetExtension(fd.Options, options.E_Field)
				if opts, ok := ext.(*options.Field); ok {
					field.Description = opts.Description
					field.Nullable = &opts.Nullable
				}
			}
			m.Fields = append(m.Fields, field)
		}
		file.Messages = append(file.Messages, m)
		r.Messages[m.MessageName()] = m
		glog.V(1).Infof("register name: %s", m.MessageName())

		var outers []string
		outers = append(outers, outerPath...)
		outers = append(outers, m.GetName())
		r.applyMsg(file, outers, m.GetNestedType())
		r.applyEnum(file, outers, m.GetEnumType())
	}
}

func (r *Registry) applyEnum(file *File, outerPath []string, enums []*descriptor.EnumDescriptorProto) {
	for i, enum := range enums {
		e := &Enum{
			File:                file,
			Outers:              outerPath,
			EnumDescriptorProto: enum,
			Index:               i,
		}
		file.Enums = append(file.Enums, e)
		r.Enums[e.EnumName()] = e
		glog.V(1).Infof("register enum name: %s", e.EnumName())
	}
}

func (r *Registry) applyServices(file *File) error {
	var svcs []*Service
	for _, sd := range file.GetService() {
		svc := &Service{
			File:                   file,
			ServiceDescriptorProto: sd,
		}
		for _, md := range sd.GetMethod() {
			m, err := r.newMethod(svc, md)
			if err != nil {
				return err
			}
			svc.Methods = append(svc.Methods, m)
		}
		if len(svc.Methods) == 0 {
			continue
		}
		glog.V(1).Infof("Registered %s with %d method(s)", svc.GetName(), len(svc.Methods))
		svcs = append(svcs, svc)
	}
	file.Services = svcs
	return nil
}

func (r *Registry) newMethod(svc *Service, md *descriptor.MethodDescriptorProto) (*Method, error) {
	request, err := r.LookupMsg(svc.File.GetPackage(), md.GetInputType())
	if err != nil {
		return nil, err
	}
	request.Type = md.GetInputType()

	response, err := r.LookupMsg(svc.File.GetPackage(), md.GetOutputType())
	if err != nil {
		return nil, err
	}
	response.Type = md.GetOutputType()

	m := &Method{
		Service:               svc,
		MethodDescriptorProto: md,
		Request:               request,
		Response:              response,
	}

	if proto.HasExtension(md.Options, options.E_Schema) {
		ext, _ := proto.GetExtension(md.Options, options.E_Schema)
		if opts, ok := ext.(*options.Schema); ok {
			m.Description = opts.Description
			switch t := opts.Type.(type) {
			case *options.Schema_Query:
				m.Query = true
				m.FieldName = t.Query
			case *options.Schema_Mutation:
				m.Mutation = true
				m.FieldName = t.Mutation
			case *options.Schema_Subscribe:
				m.Subscribe = true
				m.FieldName = t.Subscribe
			}
		}
	}

	return m, nil
}

func (r *Registry) goPackagePath(file *descriptor.FileDescriptorProto) string {
	name := file.GetName()
	if pkg, ok := r.Pkgs[name]; ok {
		return path.Join(r.Prefix, pkg)
	}
	gopkg := file.Options.GetGoPackage()
	idx := strings.LastIndex(gopkg, "/")
	if idx >= 0 {
		if sc := strings.LastIndex(gopkg, ";"); sc > 0 {
			gopkg = gopkg[:sc+1-1]
		}
		return gopkg
	}
	return path.Join(r.Prefix, path.Dir(name))
}

func (r *Registry) defaultGoPackageName(file *descriptor.FileDescriptorProto) string {
	name := r.packageIdentityName(file)
	return sanitizePackageName(name)
}

func (r *Registry) packageIdentityName(f *descriptor.FileDescriptorProto) string {
	if f.Options != nil && f.Options.GoPackage != nil {
		gopkg := f.Options.GetGoPackage()
		idx := strings.LastIndex(gopkg, "/")
		if idx < 0 {
			gopkg = gopkg[idx+1:]
		}

		gopkg = gopkg[idx+1:]
		// package name is overrided with the string after the
		// ';' character
		sc := strings.IndexByte(gopkg, ';')
		if sc < 0 {
			return sanitizePackageName(gopkg)

		}
		return sanitizePackageName(gopkg[sc+1:])
	}
	if p := r.ImportPath; len(p) != 0 {
		if i := strings.LastIndex(p, "/"); i >= 0 {
			p = p[i+1:]
		}
		return p
	}

	if f.Package == nil {
		base := filepath.Base(f.GetName())
		ext := filepath.Ext(base)
		return strings.TrimSuffix(base, ext)
	}
	return f.GetPackage()
}

func sanitizePackageName(pkgName string) string {
	pkgName = strings.Replace(pkgName, ".", "_", -1)
	pkgName = strings.Replace(pkgName, "-", "_", -1)
	return pkgName
}

func (r *Registry) LookupMsg(location, name string) (*Message, error) {
	glog.V(1).Infof("lookup %s from %s", name, location)
	if strings.HasPrefix(name, ".") {
		m, ok := r.Messages[name]
		if !ok {
			return nil, fmt.Errorf("no message found: %s", name)
		}
		return m, nil
	}

	if !strings.HasPrefix(location, ".") {
		location = fmt.Sprintf(".%s", location)
	}
	components := strings.Split(location, ".")
	for len(components) > 0 {
		n := strings.Join(append(components, name), ".")
		if m, ok := r.Messages[n]; ok {
			return m, nil
		}
		components = components[:len(components)-1]
	}
	return nil, fmt.Errorf("no message found: %s", name)
}

func (r *Registry) LookupEnum(location, name string) (*Enum, error) {
	glog.V(1).Infof("lookup enum %s from %s", name, location)
	if strings.HasPrefix(name, ".") {
		enum, ok := r.Enums[name]
		if !ok {
			return nil, fmt.Errorf("no enum found: %s", name)
		}
		return enum, nil
	}
	if !strings.HasPrefix(location, ".") {
		location = fmt.Sprintf(".%s", location)
	}
	components := strings.Split(location, ".")
	for len(components) > 0 {
		n := strings.Join(append(components, name), ".")
		if enum, ok := r.Enums[n]; ok {
			return enum, nil
		}
		components = components[:len(components)-1]
	}
	return nil, fmt.Errorf("no enum found: %s", name)
}

func (r *Registry) LookupFile(name string) (*File, error) {
	f, ok := r.Files[name]
	if !ok {
		return nil, fmt.Errorf("no such file given: %s", name)
	}
	return f, nil
}
