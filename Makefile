protofmt:
	find . -name "*.proto" | xargs clang-format -i

protoc:
	protoc -I. \
		--go_out=plugins=grpc:${GOPATH}/src \
		graphql.proto

build-graphql-gateway-debug:
	go build -o protoc-gen-graphql-gateway-debug ./cmd/protoc-gen-graphql-gateway/main.go

build-graphql-schema-debug:
	go build -o protoc-gen-graphql-schema-debug ./cmd/protoc-gen-graphql-schema/main.go

gen-example-sample: build-graphql-gateway-debug
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/sample/proto/green/*.proto

	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/sample/proto/red/*.proto

gen-example-photo_share: build-graphql-gateway-debug
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/photo_share/proto/photo/*.proto

	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/photo_share/proto/user/*.proto

gen-example-photo_share-schema: build-graphql-schema-debug
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-schema-debug \
		--go_out=plugins=grpc:. \
		--graphql-schema-debug_out=logtostderr=true:. \
		example/photo_share/proto/photo/*.proto

gen-example-federation: build-graphql-gateway-debug
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/federation/proto/account/*.proto

	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/federation/proto/product/*.proto

	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/federation/proto/review/*.proto

.PHONY: gazelle protofmt protoc
