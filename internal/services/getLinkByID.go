package services

import "fmt"

// GetLinkByID gets the id the user gives and add the link of the lecture to the clipboard
func (s *Service) GetLinkByID() string {
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
	link := lecture.Link

	return link
}
