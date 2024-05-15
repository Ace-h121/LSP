package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"Github.com/LSP/lsp"
	"Github.com/LSP/rpc"
)



func main(){
	logger := getLogger("/home/ace/Documents/Programming/Networking/LSP/log.txt")
	logger.Println("Logger Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan(){
		msg := scanner.Bytes()
		method, content, err :=rpc.DecodeMessage(msg)
		if err != nil{
			logger.Printf("Got error: %s", err)
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte){
	logger.Printf("Received msg with method: %s", method)
	switch method{
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil{
			logger.Printf("Could not Parse: %s", err)
		}
		
		logger.Printf("Connected to %s %s", 
			request.Params.ClientInfo.Name, 
			request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil{
			logger.Printf("Couldnt Parse: %s", err)
		}
		logger.Printf("Opened: %s %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	} 
}

func getLogger(filename string) *log.Logger{
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil{
		panic("better file pls")
	}
	return log.New(logfile, "[LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
