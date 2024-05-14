package rpc

import (
	"encoding/json"
	"fmt"

)


func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil{
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg any) string {
	return "woah" 
}
