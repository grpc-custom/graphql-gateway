package protoc

import (
	"bytes"
	"fmt"
	"go/format"
	"path/filepath"
	"strings"

	"github.com/grpc-custom/graphql-gateway/pkg/registry"
)

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
