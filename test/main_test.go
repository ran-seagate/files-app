package test

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"newEmpTask/api"
	"newEmpTask/config"
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
	err := os.MkdirAll(filepath.Join(config.AppConfig.UploadFolder), os.ModePerm)
	if err != nil {
		panic(err)
	}
	return router, w
}
