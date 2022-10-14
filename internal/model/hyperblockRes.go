package model

type HyperBlockRes struct {
	Code string         `json:"code"`
	Data HyperBlockInfo `json:"data"`
}

type HyperBlockInfo struct {
	HyperBLock HyperBLock `json:"hyperblock"`
}

type HyperBLock struct {
	Hash                   string        `json:"hash"`
	PrevBlockHash          string        `json:"prevBlockHash"`
	StateRootHash          string        `json:"stateRootHash"`
	Nonce                  int           `json:"nonce"`
	Round                  int           `json:"round"`
	Epoch                  int           `json:"epoch"`
	NumTxs                 int           `json:"numTxs"`
	AccumulatedFees        string        `json:"accumulatedFees"`
	DeveloperFees          string        `json:"developerFees"`
	AccumulatedFeesInEpoch string        `json:"accumulatedFeesInEpoch"`
	DeveloperFeesInEpoch   string        `json:"developerFeesInEpoch"`
	Timestamp              int           `json:"timestamp"`
	ShardBlocks            []ShardBlock  `json:"shardBlocks"`
	Transactions           []Transaction `json:"transactions"`
	Status                 string        `json:"status"`
}

type ShardBlock struct {
	Hash  string `json:"hash"`
	Nonce int    `json:"nonce"`
	Round int    `json:"round"`
	Shard int    `json:"shard"`
}

type Transaction struct {
	Type                        string   `json:"type"`
	ProcessingTypeOnSource      string   `json:"processingTypeOnSource"`
	ProcessingTypeOnDestination string   `json:"processingTypeOnDestination"`
	Hash                        string   `json:"hash"`
	Nonce                       int      `json:"nonce"`
	Round                       int      `json:"round"`
	Epoch                       int      `json:"epoch"`
	Value                       string   `json:"value"`
	Receiver                    string   `json:"receiver"`
	Sender                      string   `json:"sender"`
	GasPrice                    int      `json:"gasPrice"`
	Data                        string   `json:"data"`
	PreviousTransactionHash     string   `json:"previousTransactionHash"`
	OriginalTransactionHash     string   `json:"originalTransactionHash"`
	SourceShard                 int      `json:"sourceShard"`
	DestinationShard            int      `json:"destinationShard"`
	MiniblockType               string   `json:"miniblockType"`
	MiniblockHash               string   `json:"miniblockHash"`
	Status                      string   `json:"status"`
	Tokens                      []string `json:"tokens"`
	EsdtValues                  []string `json:"esdtValues"`
	Operation                   string   `json:"operation"`
	InitiallyPaidFee            string   `json:"initiallyPaidFee"`
}
