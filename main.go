package main

import (
	"fmt"
	"github.com/ran-seagate/files-app/api"
	"github.com/ran-seagate/files-app/config"
)

func init() {
	err := config.ReadConf("config/local.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("[init]: successfully read config file: %+v\n", *config.AppConfig)

	err = api.CreateUploadFolder()
	if err != nil {
		panic(err)
	}
}

func main() {
	router := api.InitRouter()
	err := router.Run("localhost:8081")
	if err != nil {
		panic("Failed to run http server")
	}
}
