package runtime

import (
	"fmt"
	"testing"

	user "github.com/grpc-custom/graphql-gateway/test/basic/proto"
)

func TestNewServeMux(t *testing.T) {
	args := map[string]interface{}{
		"test": 100,
	}
	n, ok := args["aaa"].(*user.User)
	fmt.Println(ok)
	fmt.Println(n)
}
