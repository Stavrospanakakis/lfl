package services

// Run runs the main functionallity of the program
func (s *Service) Run(configPath string) {
	s.ReadLecturesFromConfigurationFile(configPath)
	s.ShowLecturesList()
	link := s.GetLinkByID()
	s.OpenMeeting(link)
}
