package runtime

import (
	"fmt"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/grpc-custom/graphql-gateway/runtime/scalar"
	"github.com/stretchr/testify/require"
)

func TestServeMux_AddFields(t *testing.T) {
	userFields := graphql.Fields{
		"id": &graphql.Field{
			Type: scalar.String,
		},
	}
	userReviewsFields := graphql.Fields{
		"reviews": &graphql.Field{
			Type: graphql.NewList(scalar.String),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: scalar.String,
				},
				"first": &graphql.ArgumentConfig{
					Type:         scalar.Int32,
					DefaultValue: 5,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ret := []string{"AAA", "BBB", "CCC"}
				return ret, nil
			},
		},
	}

	mux, err := NewServeMux()
	require.NoError(t, err)

	mux.AddFields("User", userFields)
	mux.AddFields("User", userReviewsFields)

	for key, value := range mux.entities["User"] {
		fmt.Println(key)
		fmt.Println(value)
	}
}
