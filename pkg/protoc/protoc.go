package protoc

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-custom/graphql-gateway/pkg/registry"
)

func ParseRequest(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read code generator request: %v", err)
	}
	req := new(plugin.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal code generator request: %v", err)
	}
	return req, nil
}

func ParseParameter(args string) {
	if args == "" {
		return
	}
	for _, arg := range strings.Split(args, ",") {
		spec := strings.SplitN(arg, "=", 2)
		if len(spec) == 1 {
			if err := flag.CommandLine.Set(spec[0], ""); err != nil {
				glog.Fatalf("Cannot set flag %s", args)
			}
			continue
		}
		key, value := spec[0], spec[1]
		if strings.HasPrefix(key, "M") {
			continue
		}
		if err := flag.CommandLine.Set(key, value); err != nil {
			glog.Fatal("Cannot set flag %s", arg)
		}
	}
}

func Filename(file *registry.File) string {
	name := file.GetName()
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	return fmt.Sprintf("%s.pb.gql.go", base)
}

func SourceCode(buf *bytes.Buffer) (string, error) {
	code, err := format.Source(buf.Bytes())
	return string(code), err
}

func EmitError(err error) {
	res := &plugin.CodeGeneratorResponse{
		Error: proto.String(err.Error()),
	}
	emitResp(res)
}

func EmitFile(out []*plugin.CodeGeneratorResponse_File) {
	res := &plugin.CodeGeneratorResponse{
		File: out,
	}
	emitResp(res)
}

func emitResp(resp *plugin.CodeGeneratorResponse) {
	buf, err := proto.Marshal(resp)
	if err != nil {
		glog.Fatal(err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		glog.Fatal(err)
	}
}
