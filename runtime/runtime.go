package runtime

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
	"github.com/grpc-custom/graphql-gateway/runtime/codec"
	jsoniter "github.com/json-iterator/go"
)

const (
	acceptHeader              = "Accept"
	contentTypeHeader         = "Content-Type"
	authorizationHeader       = "Authorization"
	xForwardedForHeader       = "X-Forwarded-For"
	xForwardedHostHeader      = "X-Forwarded-Host"
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
	mutex       sync.Mutex
	pretty      bool
	graphiQL    bool
	schema      *graphql.Schema
	middlewares []func(http.Handler) http.Handler
	objects     map[string]*graphql.Object
	loaders     map[string]dataloader.BatchFunc
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
	ctx, err = AnnotateContext(ctx, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx = s.contextWithLoaders(ctx)
	params := graphql.Params{
		Context:        ctx,
		Schema:         *s.schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
		OperationName:  req.OperationName,
	}
	accept := r.Header.Get(acceptHeader)
	contentType := r.Header.Get(contentTypeHeader)
	ret := graphql.Do(params)
	w.Header().Set(acceptHeader, accept)
	w.Header().Set(contentTypeHeader, contentType)
	w.WriteHeader(http.StatusOK)
	if err := codec.NewEncoder(w).Encode(ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	defer s.closeBody(r)

	typ := r.Header.Get(contentTypeHeader)
	switch {
	case strings.HasPrefix(typ, applicationGraphQL):
		buf := &strings.Builder{}
		if _, err := io.Copy(buf, r.Body); err != nil {
			return nil, err
		}
		return &graphQLRequest{Query: buf.String()}, nil
	case strings.HasPrefix(typ, applicationFormURLEncoded):
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		return s.getQueryRequest(r.PostForm)
	default:
		req := &graphQLRequest{}
		if err := codec.NewDecoder(r).Decode(req); err != nil {
			return nil, err
		}
		return req, nil
	}
}

func (s *ServeMux) closeBody(r *http.Request) error {
	io.Copy(ioutil.Discard, r.Body)
	return r.Body.Close()
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

func (s *ServeMux) AddObjectType(obj *graphql.Object) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	object, ok := s.objects[obj.Name()]
	if !ok {
		s.objects[obj.Name()] = obj
		return
	}
	for name, field := range obj.Fields() {
		object.AddFieldConfig(name, &graphql.Field{
			Name:              field.Name,
			Type:              field.Type,
			Args:              s.fieldArgs(field.Args),
			Description:       field.Description,
			Resolve:           field.Resolve,
			DeprecationReason: field.DeprecationReason,
		})
	}
}

func (s *ServeMux) fieldArgs(args []*graphql.Argument) graphql.FieldConfigArgument {
	ret := make(graphql.FieldConfigArgument, len(args))
	for _, arg := range args {
		ret[arg.PrivateName] = &graphql.ArgumentConfig{
			Type:         arg.Type,
			Description:  arg.PrivateDescription,
			DefaultValue: arg.DefaultValue,
		}
	}
	return ret
}

func (s *ServeMux) AddField(typename, name string, field *graphql.Field) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	obj, ok := s.objects[typename]
	if !ok {
		s.newObject(typename, name, field)
		return
	}
	obj.AddFieldConfig(name, field)
}

func (s *ServeMux) newObject(typename, name string, field *graphql.Field) {
	s.objects[typename] = graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			name: field,
		},
	})
}

// AddLoader TODO: data loader
func (s *ServeMux) AddLoader(key string, loader dataloader.BatchFunc) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.loaders[key] = loader
}

func (s *ServeMux) contextWithLoaders(ctx context.Context) context.Context {
	for key := range s.loaders {
		loader := dataloader.NewBatchedLoader(s.loaders[key])
		ctx = context.WithValue(ctx, key, loader)
	}
	return ctx
}

func NewServeMux() (*ServeMux, error) {
	const (
		query        = "Query"
		mutation     = "Mutation"
		subscription = "Subscription"
	)
	dateField := dateResolver()
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: query,
			Fields: graphql.Fields{
				dateField.Name: dateField,
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: mutation,
			Fields: graphql.Fields{
				dateField.Name: dateField,
			},
		}),
		// TODO
		Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name: subscription,
			Fields: graphql.Fields{
				dateField.Name: dateField,
			},
		}),
	})
	if err != nil {
		return nil, err
	}
	serveMux := &ServeMux{
		mutex:    sync.Mutex{},
		graphiQL: false,
		schema:   &schema,
		objects:  map[string]*graphql.Object{},
		loaders:  map[string]dataloader.BatchFunc{},
	}
	return serveMux, nil
}

func dateResolver() *graphql.Field {
	return &graphql.Field{
		Name: "__date",
		Type: graphql.DateTime,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return time.Now(), nil
		},
	}
}
