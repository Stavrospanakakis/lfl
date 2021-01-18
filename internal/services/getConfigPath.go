package services

import (
	homedir "github.com/mitchellh/go-homedir"
)

// GetConfigPath returns the configuration file path
func (s *Service) GetConfigPath() (string, error) {
	home, err := homedir.Dir()

	if err != nil {
		return "", err
	}

	configPath := home + "/.lfl.json"
	return configPath, err
}
