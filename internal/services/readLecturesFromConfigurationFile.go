package services

import (
	"encoding/json"
	"os"

	"github.com/Stavrospanakakis/lfl/internal/models"
)

// ReadLecturesFromConfigurationFile reads lectures from configuration file
func (s *Service) ReadLecturesFromConfigurationFile() error {
	file, err := os.Open(s.ConfigPath)
	defer file.Close()

	if err != nil {
		return err
	}

	var lectures []models.Lecture

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&lectures)
	if err != nil {
		return err
	}
	s.Lectures = lectures
	return nil
}
