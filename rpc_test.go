package rpc_test

import (
	
	"Github.com/LSP/rpc"
	"testing"
)


type EncodingExample struct{
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing:true})
	if expected != actual {
		t.Fatalf("Expected %s, got: %s", expected, actual)
	}
}

func testDecode(t *testing.T){
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodeMessage([]byte(expected))
	if err != nil{
		t.Fatal(err)
	}

	if contentLength != 16{
		t.Fatalf("Expected: 16 Got :%d", contentLength)
	}
}
