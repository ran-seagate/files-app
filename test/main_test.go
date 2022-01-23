package test

import (
	"github.com/gin-gonic/gin"
	"github.com/ran-seagate/files-app/api"
	"github.com/ran-seagate/files-app/config"
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
