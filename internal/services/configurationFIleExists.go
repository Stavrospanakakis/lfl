package services

import "os"

// ConfigurationFileExists checks if configuration file exists
func (s *Service) ConfigurationFileExists() bool {
	_, err := os.Stat(s.ConfigPath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
