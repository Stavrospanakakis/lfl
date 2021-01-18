package services

import (
	"fmt"

	"github.com/Stavrospanakakis/lfl/internal/models"
)

// GetLectureByID returns a lecture by ID
func (s *Service) GetLectureByID() models.Lecture {
	fmt.Printf("\nPlease enter the id of the lecture of your choice: ")
	var lectureID int

	fmt.Scanf("%d", &lectureID)

	lecturesLength := len(s.Lectures) - 1
	for lectureID < 0 || lectureID > lecturesLength {
		if lectureID < 0 || lectureID > lecturesLength {
			fmt.Printf("\033[31m"+"Invalid Option. Please try %d-%d: "+"\033[0m", 0, lecturesLength)
			fmt.Scanf("%d\n", &lectureID)
		}
	}

	lecture := s.Lectures[lectureID]

	return lecture

}
