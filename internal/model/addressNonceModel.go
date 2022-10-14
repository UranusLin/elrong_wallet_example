package model

type AddressNonceRes struct {
	Code string           `json:"code"`
	Data AddressNonceData `json:"data"`
}

type AddressNonceData struct {
	Nonce     int              `json:"nonce"`
	BlockInfo AddressNonceInfo `json:"blockInfo"`
}

type AddressNonceInfo struct {
	Nonce    int    `json:"nonce"`
	Hash     string `json:"hash"`
	RootHash string `json:"rootHash"`
}
