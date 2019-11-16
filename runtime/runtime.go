package runtime

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/graphql-go/graphql"
	jsoniter "github.com/json-iterator/go"
)

const (
	acceptHeader              = "Accept"
	contentTypeHeader         = "Context-Type"
	applicationJSON           = "application/json"
	applicationGraphQL        = "application/graphql"
	applicationFormURLEncoded = "application/x-www-form-urlencoded"
	graphqlPath               = "/graphql"
)

var _ http.Handler = (*ServeMux)(nil)

type graphQLRequest struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

type ServeMux struct {
	pretty      bool
	graphiQL    bool
	schema      *graphql.Schema
	middlewares []func(http.Handler) http.Handler
}

func (s *ServeMux) Use(middlewares ...func(http.Handler) http.Handler) {
	s.middlewares = append(s.middlewares, middlewares...)
}

func (s *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	path := r.URL.Path
	if path != graphqlPath {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	req, err := s.newGraphQLRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := graphql.Params{
		Context:        ctx,
		Schema:         *s.schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
		OperationName:  req.OperationName,
	}
	ret := graphql.Do(params)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.NewEncoder(w).Encode(ret)
}

func (s *ServeMux) newGraphQLRequest(r *http.Request) (*graphQLRequest, error) {
	req, err := s.getQueryRequest(r.URL.Query())
	if err != nil {
		return nil, err
	}
	if req != nil {
		return req, nil
	}

	if r.Method != http.MethodPost || r.Body == nil {
		return &graphQLRequest{}, nil
	}

	defer r.Body.Close()

	typ := r.Header.Get(contentTypeHeader)
	switch {
	case strings.HasPrefix(typ, applicationGraphQL):
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		return &graphQLRequest{Query: string(body)}, nil
	case strings.HasPrefix(typ, applicationFormURLEncoded):
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		return s.getQueryRequest(r.PostForm)
	default:
		req := &graphQLRequest{}
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, err
		}
		return req, nil
	}
}

func (s *ServeMux) getQueryRequest(values url.Values) (*graphQLRequest, error) {
	const (
		queryKey     = "query"
		variablesKey = "variables"
		operationKey = "operationName"
	)
	query := values.Get(queryKey)
	if query == "" {
		return nil, nil
	}
	variables := make(map[string]interface{})
	variablesStr := values.Get(variablesKey)
	if variablesStr != "" {
		if err := jsoniter.UnmarshalFromString(variablesStr, &variables); err != nil {
			return nil, err
		}
	}
	req := &graphQLRequest{
		Query:         query,
		Variables:     variables,
		OperationName: values.Get(operationKey),
	}
	return req, nil
}

func (s *ServeMux) AddQuery(name string, field *graphql.Field) {
	s.schema.QueryType().AddFieldConfig(name, field)
}

func (s *ServeMux) AddMutation(name string, field *graphql.Field) {
	s.schema.MutationType().AddFieldConfig(name, field)
}

func (s *ServeMux) AddSubscribe(name string, field *graphql.Field) {
	s.schema.SubscriptionType().AddFieldConfig(name, field)
}

func NewServeMux() (*ServeMux, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"__date": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"__date": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		}),
		Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name: "Subscription",
			Fields: graphql.Fields{
				"__date": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		}),
	})
	if err != nil {
		return nil, err
	}
	serveMux := &ServeMux{
		graphiQL: false,
		schema:   &schema,
	}
	return serveMux, nil
}
