package services

import (
	"fmt"
)

// GenerateConfig generates the configuration file
func (s *Service) GenerateConfig(configPath string) error {

	fmt.Println("\033[31m" + "It seems you don't have any lectures saved. Let's add them!" + "\033[0m")
	lectures := s.AddLectures()

	err := s.WriteLecturesToConfigurationFile(lectures, configPath)

	if err != nil {
		return err
	}

	return nil
}
