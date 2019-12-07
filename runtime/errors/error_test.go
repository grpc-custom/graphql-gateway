package errors

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestError(t *testing.T) {
	st := status.Newf(codes.Internal, "hoge %s", "test")
	e := NewError(st)
	fmt.Println(e)
}
