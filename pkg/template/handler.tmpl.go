// Code generated by pkg/template/gen.go DO NOT EDIT.
package template

import (
	"text/template"
)

var handlerTemplate = template.Must(template.New("handler").
	Parse(`
{{ range $svc := .Services }}

    {{ range $method := $svc.Methods }}
        {{ if $method.Extend -}}
            type extend{{ $method.FieldName }} interface {
            {{ range $field := $method.Request.Fields -}}
                {{ if $field.IsExternal -}}
                    {{ $field.GetterExternalName }}() {{ $field.GoType }}
                {{ end -}}
            {{ end -}}
            }
        {{ end -}}
    {{ end }}

    type {{ $svc.PrivateServiceName }}Resolver struct {
        client {{ $svc.GetName }}Client
        group singleflight.Group
        c cache.Cache
    }

    func new{{ $svc.GetName }}Resolver(client {{ $svc.GetName }}Client) *{{ $svc.PrivateServiceName }}Resolver {
        return &{{ $svc.PrivateServiceName }}Resolver{
            client: client,
            group: singleflight.Group{},
            c: cache.New(100),
        }
    }

    {{ range $method := $svc.Methods -}}
        {{ if $method.HasGraphQLMethod -}}
            func (r *{{ $svc.PrivateServiceName }}Resolver) Field{{ $method.GetName }}() *graphql.Field {
                field := &graphql.Field{
                    Name: "{{ $method.FullMethod }}",
                    Description: "{{ $method.Description }}",
                    {{ if $method.Response.Inline -}}
                        Type: {{ $method.Response.GetInlineSchemaTypeName }},
                    {{ else -}}
                        Type: {{ $method.Response.GetSchemaTypeName }},
                    {{ end -}}
                    Args: graphql.FieldConfigArgument{
                    {{ range $field := $method.Request.Fields -}}
                        "{{ $field.GetJsonName }}": &graphql.ArgumentConfig{
                            Type: {{ $field.ScalarType }},
                        },
                    {{ end }}
                    },
                    Resolve: r.resolve{{ $method.GetName }},
                }
                return field
            }

            func (r *{{ $svc.PrivateServiceName }}Resolver) resolve{{ $method.GetName }}(p graphql.ResolveParams) (interface{}, error) {
                in := &{{ $method.Request.GetGoTypeName }}{}
                {{ range $field := $method.Request.Fields -}}
                    {{ $field.Variable }}, ok := p.Args["{{ $field.GetJsonName }}"].({{ $field.GoType }})
                    if !ok {
                    {{ if $field.IsNullable -}}
                        {{ $field.Variable }} = {{ $field.GoDefaultValue }}
                    {{ else -}}
                        return nil, errors.ErrInvalidArguments
                    {{ end -}}
                    }
                    in.{{ $field.FieldName }} = {{ $field.Variable }}
                {{ end -}}
                ctx := runtime.Context(p.Context)
                {{ if $method.CacheControl -}}
                    // cache control max age: {{ $method.CacheControl.MaxAge.Seconds }} second
                    key := cache.GenerateKey("{{ $method.FullMethod }}", in)
                    value, ok := r.c.Get(key)
                    if ok {
                        {{ if $method.Response.Inline -}}
                            ret := value.(*{{ $method.Response.Name }})
                            return ret.{{ $method.Response.InlineFieldName }}, nil
                        {{ else -}}
                            return value, nil
                        {{ end -}}
                    }
                    result, err, _ := r.group.Do(key, func() (interface{}, error) {
                        if timeout := runtime.GrpcTimeout(ctx); timeout > 0 {
                            var cancel context.CancelFunc
                            ctx, cancel = context.WithTimeout(ctx, timeout)
                            defer cancel()
                        }
                        return r.client.{{ $method.GetName }}(ctx, in)
                    })
                    if err != nil {
                        return nil, errors.ToGraphQLError(err)
                    }
                    r.c.Set(key, result, {{ $method.CacheControl.MaxAge.Seconds }}*time.Second)
                    {{ if $method.Response.Inline -}}
                        ret := result.(*{{ $method.Response.Name }})
                        return ret.{{ $method.Response.InlineFieldName }}, nil
                    {{ else -}}
                        return result, nil
                    {{ end -}}
                {{ else -}}
                    if timeout := runtime.GrpcTimeout(ctx); timeout > 0 {
                        var cancel context.CancelFunc
                        ctx, cancel = context.WithTimeout(ctx, timeout)
                        defer cancel()
                    }
                    out, err := r.client.{{ $method.GetName }}(ctx, in)
                    if err != nil {
                        return nil, errors.ToGraphQLError(err)
                    }
                    {{ if $method.Response.Inline -}}
                        return out.{{ $method.Response.InlineFieldName }}, nil
                    {{ else -}}
                        return out, nil
                    {{ end -}}
                {{ end -}}
            }

        {{ end -}}

        {{ if $method.Extend -}}
            func (r *{{ $svc.PrivateServiceName }}Resolver) extend{{ $method.FieldName }}() *graphql.Field {
                return &graphql.Field{
                    {{ if $method.Response.Inline -}}
                        Type: {{ $method.Response.GetInlineSchemaTypeName }},
                    {{ else -}}
                        Type: {{ $method.Response.GetSchemaTypeName }},
                    {{ end -}}
                    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                        ctx := runtime.Context(p.Context)
                        src, ok := p.Source.(extend{{ $method.FieldName }})
                        if !ok {
                            return nil, errors.ErrWrongType
                        }
                        args := p.Args
                        {{ range $field := $method.Request.Fields -}}
                            {{ if $field.IsExternal -}}
                                args["{{ $field.GetJsonName }}"] = src.{{ $field.GetterExternalName }}()
                            {{ end -}}
                        {{ end -}}
                        return r.resolve{{ $method.GetName }}(ctx, args)
                    },
                }
            }

            func (r *{{ $svc.PrivateServiceName }}Resolver) resolve{{ $method.GetName }}(ctx context.Context, args map[string]interface{}) (interface{}, error) {
                in := new({{ $method.Request.GetGoTypeName }})
                {{ range $field := $method.Request.Fields -}}
                    {{ $field.Variable }}, ok := args["{{ $field.GetJsonName }}"].({{ $field.GoType }})
                    if !ok {
                        {{ if $field.GetJsonName -}}
                            {{ $field.Variable }} = {{ $field.GoDefaultValue }}
                        {{ else -}}
                            return nil, errors.ErrInvalidArguments
                        {{ end -}}
                    }
                    in.{{ $field.FieldName }} = {{ $field.Variable }}
                {{ end -}}
                {{ if $method.CacheControl -}}
                    // cache control max age: {{ $method.CacheControl.MaxAge.Seconds }} second
                    key := cache.GenerateKey("{{ $method.FullMethod }}", in)
                    value, ok := r.c.Get(key)
                    if ok {
                        {{ if $method.Response.Inline -}}
                            ret := value.(*{{ $method.Response.Name }})
                            return ret.{{ $method.Response.InlineFieldName }}, nil
                        {{ else -}}
                            return value, nil
                        {{ end -}}
                    }
                    result, err, _ := r.group.Do(key, func() (interface{}, error) {
                        if timeout := runtime.GrpcTimeout(ctx); timeout > 0 {
                            var cancel context.CancelFunc
                            ctx, cancel = context.WithTimeout(ctx, timeout)
                            defer cancel()
                        }
                        return r.client.{{ $method.GetName }}(ctx, in)
                    })
                    if err != nil {
                        return nil, errors.ToGraphQLError(err)
                    }
                    r.c.Set(key, result, {{ $method.CacheControl.MaxAge.Seconds }}*time.Second)
                    {{ if $method.Response.Inline -}}
                        ret := result.(*{{ $method.Response.Name }})
                        return ret.{{ $method.Response.InlineFieldName }}, nil
                    {{ else -}}
                        return result, nil
                    {{ end -}}
                {{ else -}}
                    if timeout := runtime.GrpcTimeout(ctx); timeout > 0 {
                        var cancel context.CancelFunc
                        ctx, cancel = context.WithTimeout(ctx, timeout)
                        defer cancel()
                    }
                    out, err := r.client.{{ $method.GetName }}(ctx, in)
                    if err != nil {
                        return nil, errors.ToGraphQLError(err)
                    }
                    {{ if $method.Response.Inline -}}
                        return out.{{ $method.Response.InlineFieldName }}, nil
                    {{ else -}}
                        return out, nil
                    {{ end -}}
                {{ end -}}
            }

        {{ end -}}
    {{ end -}}

{{ end -}}

{{ range $svc := .Services }}
func Register{{ $svc.GetName }}{{ $.RegisterFuncSuffix }}FromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
    conn, err := grpc.DialContext(ctx, endpoint, opts...)
    if err != nil {
        return
    }
    defer func() {
        if err != nil {
            if e := conn.Close(); e != nil {
                grpclog.Infof("Failed to close conn to %s: %v", endpoint, e)
            }
            return
        }
        go func() {
            <-ctx.Done()
            if e := conn.Close(); e != nil {
                grpclog.Infof("Failed to close conn to %s: %v", endpoint, e)
            }
        }()
    }()
    return Register{{ $svc.GetName }}{{ $.RegisterFuncSuffix }}Handler(mux, conn)
}

func Register{{ $svc.GetName }}{{ $.RegisterFuncSuffix }}Handler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
    return Register{{ $svc.GetName }}{{ $.RegisterFuncSuffix }}HandlerClient(mux, New{{ $svc.GetName }}Client(conn))
}

func Register{{ $svc.GetName }}{{ $.RegisterFuncSuffix }}HandlerClient(mux *runtime.ServeMux, client {{ $svc.GetName }}Client) error {
    svc := new{{ $svc.GetName }}Resolver(client)
{{ range $method := $svc.Methods -}}
    {{ if $method.HasGraphQLMethod -}}
        // gRPC {{ $method.FullMethod }}
        {{ if $method.Query -}}
            mux.AddQuery("{{ $method.Name }}", svc.Field{{ $method.GetName }}())
        {{ else if $method.Mutation -}}
            mux.AddMutation("{{ $method.Name }}", svc.Field{{ $method.GetName }}())
        {{ else if $method.Subscribe -}}
            mux.AddSubscribe("{{ $method.Name }}", svc.Field{{ $method.GetName }}())
        {{ end -}}
    {{ end -}}
    {{ if $method.Extend -}}
        // extend {{ $method.Name }}.{{ $method.Field }}
        mux.AddField("{{ $method.Name }}", "{{ $method.Field }}", svc.extend{{ $method.FieldName }}())
    {{ end -}}
{{ end -}}
    return nil
}
{{ end -}}

`))
