package lsp

type Request struct {
	RPC string `json:"jsonrpc"`
	ID int `json:"id"`
	Method string `json:"method"`
	
	//We will specify at a later point
	//Params
}

//will always have a result or an error
type Response struct {
	RPC string `json:"jsonrpc"`
	ID *int `json:"id,omitempty"`
	
}

type Notification struct {
	RPC string `json:"jsonrpc"`
	Method string `json:"method"`
}




