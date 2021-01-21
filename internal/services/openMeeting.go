package services

import "os/exec"

// OpenMeeting opens the browser and pastes the url
func (s *Service) OpenMeeting(link string) error {
	err := exec.Command("xdg-open", link).Start()
	if err != nil {
		return err
	}

	return nil
}
