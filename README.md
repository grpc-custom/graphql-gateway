# graphql-gateway

This project is currently in WIP.

## Installation

```bash
go get -u github.com/grpc-custom/graphql-gateway/cmd/protoc-gen-graphql-gateway
```

## Usage

1. Define your gRPC service using protocol buffers

    `example.proto`
    
    ```proto
    syntax = "proto3";
    package example;
    
    import "github.com/grpc-custom/graphql-gateway/graphql.proto";
    
    message User {
      int32 id    = 1;
      string name = 2;
    }
    
    message GetRequest {
      int32 id = 1;
    }
    
    message GetResponse {
      User user = 1;
    }
    
    service UserService {
      rpc Get(GetRequest) returns (GetResponse) {
        option (grpc_custom.graphql.schema) = {
          query: "getUser"
        };
      }
    }
    ```

1. Generate GraphQL server using `protoc-gen-graphql-gateway`

    ```bash
    protoc -I=${GOPATH}/src:. \
        --go_out=plugins=grpc:. \
        --graphql-gateway_out=. \
        path/to/example.proto
    ```
    
    It will generate a GraphQL server `path/to/example.pb.gql.go`

1. Write an entrypoint for the GraphQL server

    ```go
    package main
    
    import (
        "context"
        "log"
        "net/http"
    
        "github.com/grpc-custom/graphql-gateway/runtime"
        "google.golang.org/grpc"
    
        gql "path/to/example_service_package"
    )
    
    func main() {
        ctx := context.Background()
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()
    
        mux, err := runtime.NewServeMux()
        if err != nil {
            log.Fatal(err)
        }
        // Register gRPC server endpoint
        opts := []grpc.DialOption{grpc.WithInsecure()}
        host := "localhost:9090" // gRPC server endpoint
        err = gql.RegisterUserServiceFromEndpoint(ctx, mux, host, opts)
        if err != nil {
            log.Fatal(err)
        }
        // Start GraphQL server
        err = http.ListenAndServe(":8080", mux)
        if err != nil {
            log.Fatal(err)
        }
    }
    ```

## Examples

[examples](https://github.com/grpc-custom/graphql-gateway/tree/master/example)

## TODO

- Apollo Persisted Queries
- complexity
- data loader
- custom error handling
- subscribe
  - websocket
- enum (string)
- message pack
- cache control
  - redis
  - memcache
- graphql schema
- open trace
