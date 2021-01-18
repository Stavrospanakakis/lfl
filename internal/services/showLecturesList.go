package services

import (
	"fmt"
)

// ShowLecturesList shows the list of the lectures
func (s *Service) ShowLecturesList() {
	fmt.Println("Lectures List:")

	lectures := s.Lectures
	for numberOfLecture := 0; numberOfLecture <= len(lectures)-1; numberOfLecture++ {
		fmt.Printf("%d) %s", numberOfLecture, lectures[numberOfLecture].Name)
	}
}
