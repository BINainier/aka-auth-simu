package main

import (
	"core/hss"
	"core/mme"
	"core/ue"
)

func main() {
	go ue.Run("8082")
	go mme.Run("8083")
	go hss.Run("8084")

	select {}
}


