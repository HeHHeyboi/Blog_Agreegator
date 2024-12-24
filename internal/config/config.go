package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

const configFilename = "/.gatorconfig.json"

type Config struct {
	Url      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func Read() Config {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	file, err := os.ReadFile(home_dir + configFilename)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	new_config := Config{}
	if err := json.Unmarshal(file, &new_config); err != nil {
		fmt.Println(err)
		return Config{}
	}
	return new_config
}
func (c *Config) SetUser(username string) {
	c.Username = username
	home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonData, err := json.Marshal(c)
	fmt.Println(string(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := os.WriteFile(home_dir+configFilename, jsonData, fs.ModePerm); err != nil {
		fmt.Println(err)
		return
	}
}
