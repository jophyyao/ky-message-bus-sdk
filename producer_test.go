package kymessage

import (
	"fmt"
	"testing"
)

func TestProducer(t *testing.T) {
	NewLogin("jophy", "asdf1234")
	code, msg := NewProducer("iaas", "resources", `{"name": "test_producer2", "type": "sdk"}`)
	fmt.Println(code, msg)
}
