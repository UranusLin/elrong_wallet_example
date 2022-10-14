package node

import (
	utils "elrondWallet/internal"
	"elrondWallet/internal/model"
	"fmt"
	"github.com/pelletier/go-toml"
)

var (
	conf, _ = toml.LoadFile("./configs/config.toml")
	url     = conf.Get("server.url").(string)
)

func GetNodeInfo() {
	var networkStatusRes model.NetworkStatusRes
	utils.HttpGet(url+"/network/status/4294967295", &networkStatusRes)
	fmt.Println(networkStatusRes)
	fmt.Printf("erd_highest_final_nonce: %d\n", networkStatusRes.Data.Status.ErdHighestFinalNonce)
}

func GetHyperBlock(nonce string) {
	var hyperBlockRes model.HyperBlockRes
	utils.HttpGet(url+"/hyperblock/by-nonce/"+nonce+"?withTxs=true", &hyperBlockRes)
	//fmt.Println(hyperBlockRes)
	fmt.Printf("transactions: %d\n", hyperBlockRes)
	utils.CheckTransactions(hyperBlockRes.Data.HyperBLock.Transactions)
}
