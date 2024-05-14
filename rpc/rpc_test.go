package rpc_test

import (
	
	"Github.com/LSP/rpc"
	"testing"
)


type EncodingExample struct{
	testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{testing:true})
	if expected != actual {
		t.Fatalf("Expected %s, got: %s", expected, actual)
	}
}
