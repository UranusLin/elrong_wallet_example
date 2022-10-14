package controllers

import (
	walletUtils "elrondWallet/internal/wallet"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"path/filepath"
)

func Mnemonic(c *gin.Context) {
	mnemonic, _ := walletUtils.GenerateNewMnemonic()
	c.JSON(200, gin.H{
		"code":     "00",
		"mnemonic": mnemonic,
	})
}

type MnemonicReq struct {
	Mnemonic string `json:"mnemonic"`
}

type WithdrawReq struct {
	Receiver string  `json:"receiver"`
	Sender   string  `json:"sender"`
	Amount   float64 `json:"amount"`
}

func DownloadPem(c *gin.Context) {
	var ap MnemonicReq
	if err := c.ShouldBindJSON(&ap); err != nil {
		c.JSON(400, gin.H{"error": "Calculator error"})
		return
	}
	privateKey := walletUtils.GetPrivateKeyFromMnemonic(ap.Mnemonic, 0, 0)
	address, _ := walletUtils.GetAddressFromPrivateKey(privateKey)
	err := walletUtils.SavePrivateKeyToPemFile(privateKey, address)
	if err != nil {
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+address+".pem")
	c.Header("Content-Type", "application/octet-stream")
	c.File(address + ".pem")

}

func Login(c *gin.Context) {
	file, _ := c.FormFile("file")
	filename := filepath.Base(file.Filename)
	fmt.Println("filename", filename)
	filePath := "saved/" + filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		// 存檔失敗時的錯誤處理
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	privateKey, _ := walletUtils.LoadPrivateKeyFromPemFile(filePath)
	address, _ := walletUtils.GetAddressFromPrivateKey(privateKey)
	if address != c.Param("address") {
		c.String(http.StatusBadRequest, fmt.Sprintf("File %s address and param address not match.", file.Filename))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "00",
		"message": "login success!",
	})
}

func GatBalance(c *gin.Context) {
	address, _ := getAddressFromReq(c)
	balance := walletUtils.GetAddressBalance(address)
	c.JSON(http.StatusOK, gin.H{
		"code":    "00",
		"balance": balance,
	})
}

func Withdraw(c *gin.Context) {
	address, privateKey := getAddressFromReq(c)
	var withdrawReq WithdrawReq
	if err := c.ShouldBindJSON(&withdrawReq); err != nil {
		c.JSON(400, gin.H{"error": "Calculator error"})
		return
	}
	if address != c.Param("address") {
		c.String(http.StatusBadRequest, fmt.Sprintf("Json address and param address not match."))
		return
	}
	addressNonce := walletUtils.GetAddressNonce(address)

	code := walletUtils.SentTransaction(withdrawReq.Sender, withdrawReq.Receiver, fmt.Sprint(
		int(withdrawReq.Amount*(math.Pow(10, 18)))), addressNonce, 1000000000, 50000, privateKey)
	c.JSON(http.StatusOK, gin.H{
		"code":    "00",
		"message": code,
	})
}

func getAddressFromReq(c *gin.Context) (string, []byte) {
	address := c.Param("address")
	filePath := "saved/" + address + ".pem"
	privateKey, _ := walletUtils.LoadPrivateKeyFromPemFile(filePath)
	address, _ = walletUtils.GetAddressFromPrivateKey(privateKey)
	return address, privateKey
}
