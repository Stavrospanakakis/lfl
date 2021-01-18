package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Stavrospanakakis/lfl/internal/models"
)

// WriteLecturesToConfigurationFile writes lectures to configuration file
func (s *Service) WriteLecturesToConfigurationFile(lectures []models.Lecture, configPath string) error {
	configFile, err := json.MarshalIndent(lectures, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, configFile, 0644)

	if err != nil {
		return err
	}

	fmt.Println("\033[32m" + "Proccess completed successfully!" + "\033[0m")

	return nil
}
