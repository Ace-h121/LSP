package rpc_test

import (
	"testing"

	"Github.com/LSP/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected %s, got: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, contentLength, err := rpc.DecodeMessage([]byte(expected))
	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 15 {
		t.Fatalf("Expected: 15 Got :%d", contentLength)
	}

	if method != "hi" {
		t.Fatalf("Expected hi, got %s", method)
	}
}
