package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config holds login data
type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppKey   string `json:"api_key"`
	CertFile string `json:"ssl_cert"`
	KeyFile  string `json:"ssl_key"`
	Locale   string
}

// LoadConfig loads a specified config.json file.
func LoadConfig(configPath string) (*Config, error) {
	jsonFile, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	return &config, nil

}
