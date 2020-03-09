package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
)

type ReviewInterface interface {
	GetAuthorID() string
}

type Review struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	AuthorID  string `json:"author_id"`
	ProductID string `json:"product_id"`
}

func (r *Review) GetAuthorID() string {
	return r.AuthorID
}

type Product struct {
	UPC   string `json:"upc"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

var key = "key"

var (
	reviewType = graphql.NewObject(graphql.ObjectConfig{
		Name: "review",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"body": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Println("---> 10")
					fmt.Println(p.Info.RootValue)
					val, ok := p.Source.(ReviewInterface)
					if !ok {
						return nil, nil
					}
					loader, ok := p.Context.Value("loader3").(*dataloader.Loader)
					fmt.Println(ok)
					thunk := loader.Load(p.Context, dataloader.StringKey(val.GetAuthorID()))
					return func() (interface{}, error) {
						ret, err := thunk()
						fmt.Println("->10 ret", ret)
						fmt.Println("->10 err", err)
						return ret, nil
					}, nil
				},
			},
			"product": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Println("---> 20")
					val, ok := p.Source.(*Review)
					if !ok {
						return nil, nil
					}
					loader, ok := p.Context.Value("loader4").(*dataloader.Loader)
					fmt.Println(ok)
					thunk := loader.Load(p.Context, dataloader.StringKey(val.ProductID))
					return func() (interface{}, error) {
						ret, err := thunk()
						fmt.Println("->20 ret", ret)
						fmt.Println("->20 err", err)
						return val.ProductID, nil
					}, nil
				},
			},
		},
	})
	productType = graphql.NewObject(graphql.ObjectConfig{
		Name: "product",
		Fields: graphql.Fields{
			"upc": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"reviews": &graphql.Field{
				Type: graphql.NewList(reviewType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Println("---> 50")

					loader, ok := p.Context.Value("loader2").(*dataloader.Loader)
					fmt.Println(ok)
					product := p.Source.(*Product)
					thunk := loader.Load(p.Context, dataloader.StringKey(product.UPC))

					return func() (interface{}, error) {
						ret, err := thunk()
						fmt.Println("->50 ret", ret)
						fmt.Println("->50 err", err)
						return []*Review{
							{
								ID:        "10",
								Body:      "body-10",
								AuthorID:  "author-id-10",
								ProductID: "product-id-10",
							},
							{
								ID:        "20",
								Body:      "body-20",
								AuthorID:  "author-id-20",
								ProductID: "product-id-20",
							},
						}, nil
					}, nil
				},
			},
		},
	})
	userType = graphql.NewObject(graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"reviews": &graphql.Field{
				Type: graphql.NewList(reviewType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fmt.Println("---> 5")

					loader, ok := p.Context.Value("loader1").(*dataloader.Loader)
					fmt.Println(ok)
					user := p.Source.(*User)
					thunk := loader.Load(p.Context, dataloader.StringKey(user.ID))

					return func() (interface{}, error) {
						ret, err := thunk()
						fmt.Println("->5 ret", ret)
						fmt.Println("->5 err", err)
						return []*Review{
							{
								ID:        "1",
								Body:      "body-1",
								AuthorID:  "author-id-1",
								ProductID: "product-id-1",
							},
							{
								ID:        "2",
								Body:      "body-2",
								AuthorID:  "author-id-2",
								ProductID: "product-id-2",
							},
						}, nil
					}, nil
				},
			},
		},
	})
)

func batch1(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	fmt.Println("---> batch1", keys)
	ret := make([]*dataloader.Result, 0, len(keys))
	for _, key := range keys {
		ret = append(ret, &dataloader.Result{
			Data:  key,
			Error: nil,
		})
	}
	return ret
}

func batch2(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	fmt.Println("---> batch2", keys)
	ret := make([]*dataloader.Result, 0, len(keys))
	for _, key := range keys {
		ret = append(ret, &dataloader.Result{
			Data:  key,
			Error: nil,
		})
	}
	return ret
}

func batch3(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	fmt.Println("---> batch3", keys)
	ret := make([]*dataloader.Result, 0, len(keys))
	for _, key := range keys {
		ret = append(ret, &dataloader.Result{
			Data:  key,
			Error: nil,
		})
	}
	return ret
}

func batch4(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	fmt.Println("---> batch4", keys)
	ret := make([]*dataloader.Result, 0, len(keys))
	for _, key := range keys {
		ret = append(ret, &dataloader.Result{
			Data:  key,
			Error: nil,
		})
	}
	return ret
}

func main() {
	// Schema
	fields := graphql.Fields{
		"me": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println("---> 1")
				rootValue := p.Info.RootValue.(map[string]interface{})
				rootValue["hello-resolve"] = "ok"
				u := &User{
					ID:       "user-id",
					Name:     "bbb",
					Username: "ccc",
				}
				return u, nil
			},
		},
		"topProducts": &graphql.Field{
			Type: graphql.NewList(productType),
			Args: graphql.FieldConfigArgument{
				"first": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 5,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println("---> 2")
				return []*Product{
					{
						UPC:   "product-1",
						Name:  "product-name-1",
						Price: 100,
					},
					{
						UPC:   "product-2",
						Name:  "product-name-2",
						Price: 200,
					},
				}, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			me {
				id
				name
				username
				reviews {
					id
					author
					product
				}
			}
			topProducts {
				upc
				reviews {
					id
					product
				}
			}
		}
	`
	ctx := context.WithValue(context.Background(), key, "root")

	loader1 := dataloader.NewBatchedLoader(batch1)
	ctx = context.WithValue(ctx, "loader1", loader1)
	loader2 := dataloader.NewBatchedLoader(batch2)
	ctx = context.WithValue(ctx, "loader2", loader2)
	loader3 := dataloader.NewBatchedLoader(batch3)
	ctx = context.WithValue(ctx, "loader3", loader3)
	loader4 := dataloader.NewBatchedLoader(batch4)
	ctx = context.WithValue(ctx, "loader4", loader4)

	params := graphql.Params{
		Context:       ctx,
		Schema:        schema,
		RequestString: query,
		RootObject:    map[string]interface{}{},
	}
	start := time.Now()
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
	fmt.Println(time.Now().Sub(start))
}
