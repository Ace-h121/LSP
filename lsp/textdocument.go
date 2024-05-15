package lsp

type TextDocumentItem struct {
	URI string `json:"uri"`

	LanguageId string `json:"languageId"`

	Version int `json:"version"`

	Text string `json:"text"`
}

type TextDocumentIdentifer struct {
	URI string `json:"uri"`
}

type VersionTextDocumentIdentifer struct {
	TextDocumentIdentifer
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifer `json:"textDocument"`
	Position     Position              `json:"position"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
