package test

import (
	"files-app/api"
	"files-app/config"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func init() {
	err := config.ReadConf(filepath.Join("..", "config", "tests.json"))
	if err != nil {
		panic(err)
	}
}

func deleteTestUploadFolder() {
	err := os.RemoveAll(config.AppConfig.UploadFolder)
	if err != nil {
		panic(err)
	}
}

func setup() (*gin.Engine, *httptest.ResponseRecorder) {
	router := api.InitRouter()
	w := httptest.NewRecorder()
	err := api.CreateUploadFolder()
	if err != nil {
		panic(err)
	}
	return router, w
}
