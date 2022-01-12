package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"newEmpTask/api"
	"newEmpTask/config"
	"os"
	"path/filepath"
)

func readConf(path string, conf *config.Config) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, conf)
	return err
}

func init() {
	err := readConf("config/local.json", config.AppConfig)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(".", config.AppConfig.UploadFolder), os.ModePerm)
	if err != nil {
		panic(err)
	}

}

func main() {
	router := gin.Default()
	files := router.Group("/files")
	{
		files.GET("/", api.GetFilesList)
		files.GET("/:file_name", api.GetFile)
		files.DELETE("/:file_name", api.DeleteFile)
	}

	err := router.Run("localhost:8080")
	if err != nil {
		panic("Failed to run http server")
	}
}
