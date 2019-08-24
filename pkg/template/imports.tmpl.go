// Code generated by pkg/template/gen.go DO NOT EDIT.
package template

import (
    "text/template"
)

var importsTemplate = template.Must(template.New("imports").
Parse(`
// Code generated by protoc-gen-alcyone. DO NOT EDIT.
// source: {{ .GetName }}

/*
Package {{ .GoPkg.Name }} is a reverse proxy.

It translates gRPC into GraphQL.
*/
package {{ .GoPkg.Name }}

import (
{{ range $i := .Imports }}{{ if $i.HasStdLib }}{{ $i | printf "%s\n" }}{{ end }}{{ end }}

{{ range $i := .Imports }}{{ if not $i.HasStdLib }}{{ $i | printf "%s\n" }}{{ end }}{{ end }}
)

`))
