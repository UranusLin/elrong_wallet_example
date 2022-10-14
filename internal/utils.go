package internal

import (
	"bytes"
	"elrondWallet/internal/model"
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml"
	"io"
	"net/http"
)

const (
	Withdraw string = "withdraw"
	Deposit         = "deposit"
	Self            = "self"
	Nothing         = "nothing"
	Transfer        = "transfer"
)

var (
	conf, _ = toml.LoadFile("./configs/config.toml")
	address = conf.Get("wallet.address").(string)
)

func HttpGet(url string, model any) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&model)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func HttpPost(url string, request any, model any) {
	body, _ := json.Marshal(request)
	fmt.Println(request)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(resp)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&model)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CheckTransactions(transactions []model.Transaction) {
	for _, transaction := range transactions {
		if transaction.Operation == Transfer {
			fmt.Printf("The derection is: %s", decideDirection(transaction.Sender, transaction.Receiver))
		}
	}
}

func decideDirection(sender, receiver string) string {
	if sender == address {
		return Withdraw
	} else if receiver == address {
		return Deposit
	} else if sender == address && receiver == address {
		return Self
	} else {
		return Nothing
	}
}
