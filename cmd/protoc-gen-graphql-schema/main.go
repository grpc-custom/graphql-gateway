package main

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-custom/graphql-gateway/pkg/protoc"
	"github.com/grpc-custom/graphql-gateway/pkg/registry"
)

func main() {
	flag.Parse()
	defer glog.Flush()

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

	// files := make([]*plugin.CodeGeneratorResponse_File, 0, len(req.FileToGenerate))

	for _, filePath := range req.FileToGenerate {
		file, err := reg.LookupFile(filePath)
		if err != nil {
			protoc.EmitError(err)
			return
		}

		for _, svc := range file.Services {
			for _, method := range svc.Methods {
				if method.Query {
					glog.Info(method)
				}
			}
		}
	}
}
