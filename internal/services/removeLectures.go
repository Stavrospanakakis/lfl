package services

import "github.com/Stavrospanakakis/lfl/internal/models"

// RemoveLectures removes lectures from list
func (s *Service) RemoveLectures(configPath string) error {
	err := s.ReadLecturesFromConfigurationFile(configPath)
	if err != nil {
		return err
	}

	s.ShowLecturesList()

	lectureToByRemoved := s.GetLectureByID()
	lectures := s.Lectures

	var updatedLectures []models.Lecture
	counter := 0
	for i := 0; i <= len(lectures)-1; i++ {
		if lectures[i] != lectureToByRemoved {
			updatedLectures = append(updatedLectures, lectures[i])
			updatedLectures[counter].ID = counter
			counter++
		}
	}

	err = s.WriteLecturesToConfigurationFile(updatedLectures, configPath)
	if err != nil {
		return err
	}

	return nil
}
