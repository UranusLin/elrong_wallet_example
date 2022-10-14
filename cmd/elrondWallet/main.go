package main

import (
	"fmt"
	"math"
)

func main() {
	//address := "erd1aym5zlt5dhmjr4gqkcyuswpejvr4aam06mns0gmfc4unehxw337sykx33c"
	//walletUtils.GetAddressBalance(address)
	//walletUtils.GetAddressNonce(address)
	//nodeUtils.GetNodeInfo()
	//nodeUtils.GetHyperBlock("1131724")
	////1. get New Mnemonic
	//mnemonic, _ := walletUtils.GenerateNewMnemonic()
	//fmt.Printf("Generated mnemonic: %s \n", mnemonic)
	//// 2. get Private Key from Mnemonic
	//privateKey := walletUtils.GetPrivateKeyFromMnemonic(mnemonic, 0, 0)
	//fmt.Printf("Generated privateKey: %s \n", privateKey)
	//// 3. get Address from private key
	//address, _ = walletUtils.GetAddressFromPrivateKey(privateKey)
	//fmt.Printf("Generated wallet: %s \n", address)
	//// 4. export pem file from private key
	//err := walletUtils.SavePrivateKeyToPemFile(privateKey, address+".pem")
	//if err != nil {
	//	return
	//}
	//// generate password from json
	//jsonPass, _ := password.Generate(64, 10, 10, false, false)
	//fmt.Printf("Generated json password: %s \n", jsonPass)
	//// 5. export json file from private key
	//err = walletUtils.SavePrivateKeyToJsonFile(privateKey, jsonPass, address+".json")
	//if err != nil {
	//	return
	//}
	//// 6. import from pem
	//pemPrivateKey, err := walletUtils.LoadPrivateKeyFromPemFile("/Users/morris.lin/GolandProjects/elrondWallet/" + address + ".pem")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("PrivateKey from pem: %s \n", pemPrivateKey)
	//pemAddress, _ := walletUtils.GetAddressFromPrivateKey(pemPrivateKey)
	//fmt.Printf("Address from pem: %s \n", pemAddress)
	//// 7. import from json
	//jsonPrivateKey, _ := walletUtils.LoadPrivateKeyFromJsonFile(address+".json", jsonPass)
	//fmt.Printf("PrivateKey from json: %s \n", jsonPrivateKey)
	//jsonAddress, _ := walletUtils.GetAddressFromPrivateKey(jsonPrivateKey)
	//fmt.Printf("Address from json: %s \n", jsonAddress)
	//walletUtils.SentTransaction(address,
	//	"erd1dtfqa76nzjgy0wcg859ls4rh5fet9pnqnmh0337sy3usjmuplmksjhs5l2",
	//	"1000000000000000000",
	//	1,
	//	1000000000,
	//	50000,
	//	pemPrivateKey,
	//)
	var a float64
	a = 1
	a = a * math.Pow(10, 18)
	fmt.Printf(fmt.Sprint(int(a)))

}
