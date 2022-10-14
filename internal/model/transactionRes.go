package model

type TransactionRes struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}
