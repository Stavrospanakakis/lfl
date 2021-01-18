package services

// AddNewLectures adds new lectures
func (s *Service) AddNewLectures(configPath string) error {
	err := s.ReadLecturesFromConfigurationFile(configPath)
	if err != nil {
		return err
	}

	lectures := s.AddLectures()

	err = s.WriteLecturesToConfigurationFile(lectures, configPath)

	if err != nil {
		return err
	}

	return nil
}
