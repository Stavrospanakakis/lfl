package services

// AddNewLectures adds new lectures
func (s *Service) AddNewLectures() error {
	err := s.ReadLecturesFromConfigurationFile()
	if err != nil {
		return err
	}

	lectures := s.AddLectures()

	err = s.WriteLecturesToConfigurationFile(lectures)

	if err != nil {
		return err
	}

	return nil
}
