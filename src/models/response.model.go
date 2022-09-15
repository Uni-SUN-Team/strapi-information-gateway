package models

type Response struct {
	Error   string `json:"error"`
	Status  bool   `json:"status"`
	Payload string `json:"payload"`
}
