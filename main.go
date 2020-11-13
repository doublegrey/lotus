package main

import (
	"github.com/doublegrey/lotus/api"
	"github.com/doublegrey/lotus/utils"
)

func main() {
	utils.ParseConfig()
	utils.InitializeDB()

	router := api.Setup()
	router.Run(utils.Config.Server.Addr)
}
