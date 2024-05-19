package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"Github.com/LSP/analysis"
	"Github.com/LSP/lsp"
	"Github.com/LSP/rpc"
)

func main() {
	logger := getLogger("/home/ace/Documents/Programming/Networking/LSP/log.txt")
	logger.Println("Logger Started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout

	state := analysis.NewState()
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got error: %s", err)
		}
		handleMessage(logger, writer, state, method, content)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Could not Parse: %s", err)
		}

		logger.Printf("Connected to %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
		msg := lsp.NewInitializeResponse(request.ID)

		writeResponse(writer, msg)

		logger.Print("Sent the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Couldnt Parse: %s", err)
			return
		}
		logger.Printf("Opened: %s", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChanceNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("TextDocument/DidChange error: %s", err)
			return
		}
		logger.Printf("Changed %s", request.Params.TextDocument.URI)

		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover error: %s, err", err)
			return
		}
		//create a response
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		writeResponse(writer, response)

	case "textDocument/definition":
		var request lsp.DefinitionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/definition error: %s, err", err)
			return
		}
		//create a response
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("better file pls")
	}
	return log.New(logfile, "[LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
