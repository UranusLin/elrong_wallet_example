package model

type NetworkStatusRes struct {
	Code string     `json:"code"`
	Data StatusData `json:"data"`
}

type StatusData struct {
	Status Status `json:"status"`
}

type Status struct {
	ErdCrossCheckBlockHeight      string `json:"erd_cross_check_block_height"`
	ErdCurrentRound               int    `json:"erd_current_round"`
	ErdEpochNumber                int    `json:"erd_epoch_number"`
	ErdHighestFinalNonce          int    `json:"erd_highest_final_nonce"`
	ErdNonce                      int    `json:"erd_nonce"`
	ErdNonceAtEpochStart          int    `json:"erd_nonce_at_epoch_start"`
	ErdNoncesPassedInCurrentEpoch int    `json:"erd_nonces_passed_in_current_epoch"`
	ErdRoundAtEpochStart          int    `json:"erd_round_at_epoch_start"`
	ErdRoundsPassedInCurrentEpoch int    `json:"erd_rounds_passed_in_current_epoch"`
	ErdRoundsPerEpoch             int    `json:"erd_rounds_per_epoch"`
}
