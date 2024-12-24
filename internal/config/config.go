package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFilename = ".gatorconfig.json"

type Config struct {
	Url      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	new_config := Config{}
	if err := json.Unmarshal(file, &new_config); err != nil {
		return Config{}, err
	}
	return new_config, nil
}
func (c *Config) SetUser(username string) error {
	c.Username = username
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		return err
	}
	return nil

}
func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home_dir, configFilename)
	return path, nil

}
