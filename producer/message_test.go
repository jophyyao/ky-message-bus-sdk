package producer

import (
	"fmt"
	"testing"

	"github.com/jophyyao/ky-message-bus-sdk/auth"
)

func TestProducer(t *testing.T) {
	auth.NewLogin("jophy", "asdf1234")
	code, msg := NewProducer("iaas", "resources", `{"name": "test_producer"}`)
	fmt.Println(code, msg)
}
