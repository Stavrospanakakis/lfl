package services

import "os"

// ConfigurationFileExists checks if configuration file exists
func (s *Service) ConfigurationFileExists(configPath string) bool {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
