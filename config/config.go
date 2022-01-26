package config

import (
	"encoding/json"
	"io/ioutil"
)

var AppConfig = &Config{}

type Config struct {
	UploadFolder string `json:"upload_folder"`
}

// ReadConf reads conf from config dir
func ReadConf(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, AppConfig)
	return err
}
