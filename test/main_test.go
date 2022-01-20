package test

import (
	"newEmpTask/config"
	"os"
	"path/filepath"
)

func init() {
	err := config.ReadConf("../config/tests.json")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(".", config.AppConfig.UploadFolder), os.ModePerm)
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
