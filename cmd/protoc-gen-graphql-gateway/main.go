package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-custom/graphql-gateway/pkg/protoc"
	"github.com/grpc-custom/graphql-gateway/pkg/registry"
	"github.com/grpc-custom/graphql-gateway/pkg/template"
)

var (
	compilerType       = flag.String("compiler_type", "go", "")
	registerFuncSuffix = flag.String("register_func_suffix", "", "used to construct names of generated Register*<Suffix> methods.")
	traceType          = flag.String("trace_type", "", "")
	versionFlag        = flag.Bool("version", false, "print the current version")
)

// set version by goreleaser
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if *versionFlag {
		fmt.Printf("Version %s, commit %s, build at %s\n", version, commit, date)
		os.Exit(0)
	}

	req, err := protoc.ParseRequest(os.Stdin)
	if err != nil {
		glog.Fatal(err)
	}

	protoc.ParseParameter(req.GetParameter())

	reg := registry.New()
	if err := reg.Apply(req); err != nil {
		protoc.EmitError(err)
		return
	}

	files := make([]*plugin.CodeGeneratorResponse_File, 0, len(req.FileToGenerate))

	for _, filePath := range req.FileToGenerate {
		file, err := reg.LookupFile(filePath)
		if err != nil {
			protoc.EmitError(err)
			return
		}

		var imports []*registry.GoPackage
		for _, svc := range file.Services {
			for _, method := range svc.Methods {
				dep, err := reg.LookupMsg("", method.GetInputType())
				if err != nil {
					protoc.EmitError(err)
					return
				}
				if dep != nil && file.GoPkg != dep.File.GoPkg {
					imports = append(imports, dep.File.GoPkg)
				}
				for _, field := range method.Request.Fields {
					if !field.IsMessageType() {
						continue
					}
					dependence, err := reg.LookupMsg(file.GetPackage(), field.GetTypeName())
					if err != nil {
						protoc.EmitError(err)
						return
					}
					if file.GoPkg.String() == dependence.File.GoPkg.String() {
						continue
					}
					imports = append(imports, dependence.File.GoPkg)
				}
			}
		}

		for _, msg := range file.Messages {
			for _, field := range msg.Fields {
				switch {
				case field.IsEnumType():
					enum, err := reg.LookupEnum(file.GetPackage(), field.GetTypeName())
					if err != nil {
						protoc.EmitError(err)
						return
					}
					field.Enum = enum
				case field.IsMessageType():
					dependence, err := reg.LookupMsg(file.GetPackage(), field.GetTypeName())
					if err != nil {
						protoc.EmitError(err)
						return
					}
					field.Dependence = dependence
					field.Dependence.Type = field.GetTypeName()
				}
			}
		}

		// generate source code
		out, err := generate(file, imports)
		if err != nil {
			protoc.EmitError(err)
			return
		}
		files = append(files, out)
	}

	protoc.EmitFile(files)
}

func generate(
	file *registry.File,
	imports []*registry.GoPackage,
) (*plugin.CodeGeneratorResponse_File, error) {
	buf := new(bytes.Buffer)
	importCode, err := template.GenerateImports(file, imports...)
	if err != nil {
		return nil, err
	}
	buf.Write(importCode)
	schemaCode, err := template.GenerateSchemas(file)
	if err != nil {
		return nil, err
	}
	buf.Write(schemaCode)
	params := &template.HandlerParams{
		File:               file,
		RegisterFuncSuffix: *registerFuncSuffix,
	}
	handlerCode, err := template.GenerateHandler(params)
	if err != nil {
		return nil, err
	}
	buf.Write(handlerCode)
	code, err := protoc.SourceCode(buf)
	if err != nil {
		return nil, err
	}

	out := &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(protoc.Filename(file)),
		Content: proto.String(code),
	}
	return out, nil
}
