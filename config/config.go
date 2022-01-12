package config

var AppConfig = &Config{}

type Config struct {
	UploadFolder string `json:"upload_folder"`
}
