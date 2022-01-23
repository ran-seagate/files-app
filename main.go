package main

import (
	"fmt"
	"newEmpTask/api"
	"newEmpTask/config"
	"os"
	"path/filepath"
)

func init() {
	err := config.ReadConf(filepath.Join("config", "local.json"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("[init]: successfully read config file: %+v\n", *config.AppConfig)
	err = os.MkdirAll(filepath.Join(".", config.AppConfig.UploadFolder), os.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Println("[init]: successfully created upload folder")
}

func main() {
	router := api.InitRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		panic("Failed to run http server")
	}
}
