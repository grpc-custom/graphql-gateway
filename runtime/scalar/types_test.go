package scalar

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	src, err := base64.StdEncoding.DecodeString("AQEB")
	fmt.Println(src)
	fmt.Println(err)
}
