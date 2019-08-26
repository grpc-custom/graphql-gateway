package template

import (
	"bytes"
	"path"

	"github.com/grpc-custom/graphql-gateway/pkg/registry"
)

type params struct {
	*registry.File
	Imports []*registry.GoPackage
}

//go:generate go run ./gen.go imports.tmpl imports
func GenerateImports(file *registry.File, imports ...*registry.GoPackage) ([]byte, error) {
	pkgs := make([]*registry.GoPackage, 0)
	for _, pkgPath := range []string{
		"context",
		"github.com/graphql-go/graphql",
		"github.com/grpc-custom/graphql-gateway/runtime",
		"github.com/grpc-custom/graphql-gateway/runtime/scalar",
		"google.golang.org/grpc",
		"google.golang.org/grpc/grpclog",
	} {
		pkg := &registry.GoPackage{
			Path: pkgPath,
			Name: path.Base(pkgPath),
		}
		pkgs = append(pkgs, pkg)
	}
	pkgs = append(pkgs, imports...)
	params := &params{
		File:    file,
		Imports: pkgs,
	}
	buf := new(bytes.Buffer)
	if err := importsTemplate.Execute(buf, &params); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//go:generate go run ./gen.go schemas.tmpl schemas
func GenerateSchemas(file *registry.File) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := schemasTemplate.Execute(buf, file); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type HandlerParams struct {
	*registry.File
	RegisterFuncSuffix string
}

//go:generate go run ./gen.go handler.tmpl handler
func GenerateHandler(params *HandlerParams) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := handlerTemplate.Execute(buf, params); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
