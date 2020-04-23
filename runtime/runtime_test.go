package runtime

import (
	"fmt"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/grpc-custom/graphql-gateway/runtime/scalar"
	"github.com/stretchr/testify/require"
)

type Review struct {
	ID string `json:"id"`
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: scalar.String,
		},
		"name": &graphql.Field{
			Type: scalar.String,
		},
		"username": &graphql.Field{
			Type: scalar.String,
		},
	},
})

func TestServeMux_AddField(t *testing.T) {
	serve, err := NewServeMux()
	require.NoError(t, err)

	reviewType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Review",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	field := &graphql.Field{
		Type: graphql.NewList(reviewType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			list := []*Review{
				{
					ID: "aaa",
				},
			}
			return list, nil
		},
	}
	serve.AddObjectType(userType)
	serve.AddField("User", "reviews", field)
	fmt.Println(serve.objects)
}
