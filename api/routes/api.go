package routes

import "elrondWallet/api/controllers"

func init() {

	// unauthorized
	unauthorizedAPI := engine.Group("api/v1/wallet")
	{
		unauthorizedAPI.GET("/mnemonic", controllers.Mnemonic)
		unauthorizedAPI.POST("/pem", controllers.DownloadPem)
		unauthorizedAPI.POST("/:address/login", controllers.Login)
		unauthorizedAPI.GET("/:address/balance", controllers.GatBalance)
		unauthorizedAPI.POST("/:address/withdraw", controllers.Withdraw)
	}
}
