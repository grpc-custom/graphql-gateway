protoc-go:
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--go_out=plugins=grpc:. \
		./proto/account/*.proto
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--go_out=plugins=grpc:. \
		./proto/product/*.proto
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--go_out=plugins=grpc:. \
		./proto/review/*.proto

run-account-server:
	go run ./account/main.go

run-review-server:
	go run ./review/main.go

run-product-server:
	go run ./product/main.go
