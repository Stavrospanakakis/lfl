package services

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

// ShowLecturesList shows the list of the lectures, scans the id
// the user gives and add the link of the lecture to the clipboard
func ShowLecturesList() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	numberOfLectures, err := strconv.Atoi(viper.GetString("totalSum"))

	if err != nil {
		return err
	}

	for i := 1; i <= numberOfLectures; i++ {
		numberOfLecture := strconv.Itoa(i)
		nameJSONQuery := numberOfLecture + ".name"
		name := viper.GetString(nameJSONQuery)
		fmt.Println(numberOfLecture + ") " + name)
	}

	fmt.Println("\nPlease enter the id of the lecture of your choice:")
	var lectureID string
	fmt.Scanf("%s", &lectureID)

	err = GetLinkByID(lectureID)

	if err != nil {
		return err
	}
	return nil
}
