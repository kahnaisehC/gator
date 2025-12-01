package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigLocation() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configLocation := homeDir + "/.gatorconfig.json"
	return configLocation, nil
}

func Read() (Config, error) {
	cfg := Config{}

	configLocation, err := getConfigLocation()
	if err != nil {
		return cfg, err
	}

	file, err := os.Open(configLocation)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data[:n], &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	configLocation, err := getConfigLocation()
	if err != nil {
		return err
	}

	f, err := os.Create(configLocation)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(data))

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
