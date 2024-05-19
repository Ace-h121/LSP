package analysis

import (
	"Github.com/LSP/lsp"
	"fmt"
)

type State struct {
	//map of filename + contents with it
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(document, text string) {
	s.Documents[document] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	//TODO this should look up type in code

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "1.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File : %s, Characters: %d", uri, len(document)),
		},
	}

}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	//TODO this should look up type in code

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "1.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}

}
