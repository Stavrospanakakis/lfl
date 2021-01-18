package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GenerateConfig generates the configuration file
func (s *Service) GenerateConfig(configPath string) error {

	lectures := s.AddLectures()

	configFile, err := json.MarshalIndent(lectures, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, configFile, 0644)

	if err != nil {
		return err
	}

	fmt.Println("\033[32m" + "Lectures added successfully!" + "\033[0m")
	return nil
}
