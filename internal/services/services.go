package services

import (
	"github.com/Stavrospanakakis/lfl/internal/models"
)

// Service contains the lectures
type Service struct {
	Lectures   []models.Lecture
	ConfigPath string
}

// MakeService makes service
func MakeService() *Service {
	s := &Service{}
	return s
}
