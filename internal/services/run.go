package services

// Run runs the main functionallity of the program
func (s *Service) Run(configPath string) error {
	err := s.ReadLecturesFromConfigurationFile(configPath)

	if err != nil {
		return err
	}

	s.ShowLecturesList()
	link := s.GetLinkByID()
	s.OpenMeeting(link)

	return nil
}
