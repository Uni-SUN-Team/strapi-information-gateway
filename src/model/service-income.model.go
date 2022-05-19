package model

type Payload struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Body   []byte `json:"body"`
}
