package main

import (
	"elrondWallet/api/routes"
	"log"
)

func main() {
	engine := routes.New()
	log.Fatal(engine.Run(":5001"))
}
