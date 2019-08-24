package protoc

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
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
