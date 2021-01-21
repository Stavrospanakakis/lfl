package services

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Stavrospanakakis/lfl/internal/models"
)

// AddLectures adds the lectures
func (s *Service) AddLectures() []models.Lecture {
	var lectures []models.Lecture
	var lecture models.Lecture
	var ID int

	if s.Lectures == nil {
		ID = 0
	} else {
		lectures = s.Lectures
		lecturesLength := len(lectures)
		ID = lecturesLength
	}

	stop := "y"
	for stop != "n" {
		var name string
		var link string

		fmt.Printf("%d) Name: ", ID)
		inputReader := bufio.NewReader(os.Stdin)
		name, _ = inputReader.ReadString('\n')

		fmt.Printf("   Link: ")
		fmt.Scanf("%s\n", &link)

		lecture.ID = ID
		lecture.Name = name
		lecture.Link = link

		lectures = append(lectures, lecture)
		ID++

		fmt.Print("Would you like to add another lecture? Type y/n: ")

		fmt.Scanf("%s\n", &stop)
		for stop != "y" && stop != "n" {
			if stop != "y" && stop != "n" {
				fmt.Printf("\033[31m" + "Invalid Option. Please type y/n: " + "\033[0m")
				fmt.Scanf("%s\n", &stop)
			}
		}
	}
	return lectures
}
