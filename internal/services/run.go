package services

// Run runs the main functionallity of the program
func (s *Service) Run() error {
	err := s.ReadLecturesFromConfigurationFile()

	if err != nil {
		return err
	}

	s.ShowLecturesList()
	link := s.GetLinkByID()
	s.OpenMeeting(link)

	return nil
}
