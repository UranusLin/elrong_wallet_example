package walletUtils

import (
	utils "elrondWallet/internal"
	"elrondWallet/internal/model"
	"fmt"
	"github.com/ElrondNetwork/elrond-sdk/erdgo"
	"github.com/ElrondNetwork/elrond-sdk/erdgo/data"
	"github.com/pelletier/go-toml"
	"math"
	"strconv"
)

var (
	conf, _ = toml.LoadFile("./configs/config.toml")
	url     = conf.Get("server.url").(string)
)

func GenerateNewMnemonic() (string, error) {
	return erdgo.GenerateNewMnemonic()
}

func GetPrivateKeyFromMnemonic(mnemonic string, account, index uint8) []byte {
	return erdgo.GetPrivateKeyFromMnemonic(mnemonic, account, index)
}

func GetAddressFromPrivateKey(privateKey []byte) (string, error) {
	return erdgo.GetAddressFromPrivateKey(privateKey)
}

func SavePrivateKeyToPemFile(privateKey []byte, filePath string) error {
	return erdgo.SavePrivateKeyToPemFile(privateKey, filePath+".pem")
}

func SavePrivateKeyToJsonFile(privateKey []byte, jsonPass, filePath string) error {
	return erdgo.SavePrivateKeyToJsonFile(privateKey, jsonPass, filePath)
}

func LoadPrivateKeyFromPemFile(filePath string) ([]byte, error) {
	return erdgo.LoadPrivateKeyFromPemFile(filePath)
}

func LoadPrivateKeyFromJsonFile(filePath, filePass string) ([]byte, error) {
	return erdgo.LoadPrivateKeyFromJsonFile(filePath, filePass)
}

func GetAddressBalance(address string) float64 {
	var balanceRes model.BalanceRes
	utils.HttpGet(url+"/address/"+address+"/balance", &balanceRes)
	fmt.Println(balanceRes)
	//fmt.Printf("balance: %s\n", balanceRes.Data.Balance)
	balance, _ := strconv.ParseFloat(balanceRes.Data.Balance, 64)
	fmt.Printf("balance: %f\n", balance/math.Pow(10, 18))
	return balance / math.Pow(10, 18)
}

func GetAddressNonce(address string) int {
	var addressNonceRes model.AddressNonceRes
	utils.HttpGet(url+"/address/"+address+"/nonce", &addressNonceRes)
	fmt.Println(addressNonceRes)
	fmt.Printf("nonce: %d\n", addressNonceRes.Data.Nonce)
	return addressNonceRes.Data.Nonce
}

func SentTransaction(sender, receiver, value string, nonce, gasPrice, gasLimit int, pemPrivateKey []byte) string {
	transaction := data.Transaction{
		Nonce:    uint64(nonce),
		Value:    value,
		SndAddr:  sender,
		RcvAddr:  receiver,
		GasPrice: uint64(gasPrice),
		GasLimit: uint64(gasLimit),
		ChainID:  "T",
		Version:  1,
	}
	err := erdgo.SignTransaction(&transaction, pemPrivateKey)
	if err != nil {
		return ""
	}
	fmt.Printf("Signature: %s", transaction.Signature)

	var transactionRes model.TransactionRes
	utils.HttpPost(url+"/transaction/send", transaction, &transactionRes)
	fmt.Println(transactionRes)
	return transactionRes.Code
}
