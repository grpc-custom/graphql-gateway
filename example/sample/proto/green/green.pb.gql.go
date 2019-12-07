// Code generated by protoc-gen-graphql-gateway. DO NOT EDIT.
// source: example/sample/proto/green/green.proto

/*
Package green is a reverse proxy.

It translates gRPC into GraphQL.
*/
package green

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/grpc-custom/graphql-gateway/runtime"
	"github.com/grpc-custom/graphql-gateway/runtime/cache"
	"github.com/grpc-custom/graphql-gateway/runtime/errors"
	"github.com/grpc-custom/graphql-gateway/runtime/scalar"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	termType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Term",
		Fields: graphql.Fields{
			"startAt": &graphql.Field{
				Type: scalar.Int64,
			},
			"endAt": &graphql.Field{
				Type: scalar.Int64,
			},
		},
	})

	greenType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Green",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: scalar.String,
			},
			"title": &graphql.Field{
				Type: scalar.String,
			},
			"text": &graphql.Field{
				Type: scalar.String,
			},
			"url": &graphql.Field{
				Type: scalar.String,
			},
			"createdAt": &graphql.Field{
				Type: scalar.Timestamp,
			},
			"updatedAt": &graphql.Field{
				Type: scalar.Timestamp,
			},
			"term": &graphql.Field{
				Type: termType,
			},
			"published": &graphql.Field{
				Type: scalar.Bool,
			},
		},
	})

	listRequestType = graphql.NewObject(graphql.ObjectConfig{
		Name: "ListRequest",
		Fields: graphql.Fields{
			"nextToken": &graphql.Field{
				Type: scalar.String,
			},
			"size": &graphql.Field{
				Type: graphql.NewNonNull(scalar.Int32),
			},
		},
	})

	listResponseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "ListResponse",
		Fields: graphql.Fields{
			"nextToken": &graphql.Field{
				Type: scalar.String,
			},
			"greens": &graphql.Field{
				Type: graphql.NewList(greenType),
			},
		},
	})

	getRequestType = graphql.NewObject(graphql.ObjectConfig{
		Name: "GetRequest",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: scalar.String,
			},
		},
	})

	getResponseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "GetResponse",
		Fields: graphql.Fields{
			"green": &graphql.Field{
				Type: greenType,
			},
		},
	})
)

type GreenServiceResolver struct {
	client GreenServiceClient
	group  singleflight.Group
	c      cache.Cache
}

func newGreenServiceResolver(client GreenServiceClient) *GreenServiceResolver {
	return &GreenServiceResolver{
		client: client,
		group:  singleflight.Group{},
		c:      cache.New(100),
	}
}

func (r *GreenServiceResolver) FieldList() *graphql.Field {
	field := &graphql.Field{
		Name:        "/green.GreenService/List",
		Description: "",
		Type:        listResponseType,
		Args: graphql.FieldConfigArgument{
			"nextToken": &graphql.ArgumentConfig{
				Type: scalar.String,
			},
			"size": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(scalar.Int32),
			},
		},
		Resolve: r.resolveList,
	}
	return field
}

func (r *GreenServiceResolver) resolveList(p graphql.ResolveParams) (interface{}, error) {
	in := &ListRequest{}
	valueNextToken, ok := p.Args["nextToken"].(string)
	if !ok {
		valueNextToken = ""
	}
	in.NextToken = valueNextToken
	valueSize, ok := p.Args["size"].(int32)
	if !ok {
		return nil, errors.ErrInvalidArguments
	}
	in.Size = valueSize
	ctx := runtime.Context(p.Context)
	result, err := r.client.List(ctx, in)
	if err != nil {
		return nil, errors.ToGraphQLError(err)
	}
	return result, nil
}

func (r *GreenServiceResolver) FieldGet() *graphql.Field {
	field := &graphql.Field{
		Name:        "/green.GreenService/Get",
		Description: "",
		Type:        getResponseType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: scalar.String,
			},
		},
		Resolve: r.resolveGet,
	}
	return field
}

func (r *GreenServiceResolver) resolveGet(p graphql.ResolveParams) (interface{}, error) {
	in := &GetRequest{}
	valueId, ok := p.Args["id"].(string)
	if !ok {
		valueId = ""
	}
	in.Id = valueId
	ctx := runtime.Context(p.Context)
	result, err := r.client.Get(ctx, in)
	if err != nil {
		return nil, errors.ToGraphQLError(err)
	}
	return result, nil
}

func RegisterGreenServiceFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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
	return RegisterGreenServiceHandler(mux, conn)
}

func RegisterGreenServiceHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGreenServiceHandlerClient(mux, NewGreenServiceClient(conn))
}

func RegisterGreenServiceHandlerClient(mux *runtime.ServeMux, client GreenServiceClient) error {
	svc := newGreenServiceResolver(client)
	// gRPC /green.GreenService/List
	mux.AddQuery("listGreens", svc.FieldList())
	// gRPC /green.GreenService/Get
	mux.AddQuery("getGreen", svc.FieldGet())
	return nil
}
