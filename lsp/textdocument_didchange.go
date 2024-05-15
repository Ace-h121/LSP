package lsp

type TextDocumentDidChanceNotification struct {
	Notification
	Params DidChangeTextDocumentsParams`json:"params"`
	
}

type DidChangeTextDocumentsParams struct {
	TextDocument VersionTextDocumentIdentifer `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct{
	Text string `json:"text"`
}
