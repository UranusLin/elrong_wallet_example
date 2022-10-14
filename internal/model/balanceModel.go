package model

type BalanceRes struct {
	Code string         `json:"code"`
	Data BalanceResData `json:"data"`
}
type BalanceResData struct {
	Balance   string             `json:"balance"`
	BlockInfo BalanceResDataInfo `json:"blockInfo"`
}
type BalanceResDataInfo struct {
	Nonce    int    `json:"nonce"`
	Hash     string `json:"hash"`
	RootHash string `json:"rootHash"`
}
