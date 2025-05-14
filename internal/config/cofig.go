package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_url string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	fileName, err := getConfigDir();
	if err != nil {
		return Config{}, fmt.Errorf("%w", err);
	}

	data, err := os.ReadFile(fileName);
	if err != nil {
		return Config{}, fmt.Errorf("Error while reading file: %w", err)
	}

	config := Config{}
	
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Error during unmarshal: %w", err)
	}

	return config, nil
}

func (c *Config) SetUser (user_name string) error {
	c.Current_user_name = user_name;
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error during marshal: %w", err)
	}

	fileName, err := getConfigDir();
	if err != nil {
		return fmt.Errorf("%w", err);
	}

	err = os.WriteFile(fileName, data, 0666);
	if err != nil {
		return fmt.Errorf("Error during writing to file: %w", err)	
	}
	
	return nil
}

func getConfigDir() (string, error) {
	home_dir, err := os.UserHomeDir();
	if err != nil {
		return "", fmt.Errorf("Can't get home dir: %w", err);
	}

	return home_dir + "/.gatorconfig.json", nil
}
