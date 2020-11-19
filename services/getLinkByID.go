package services

import (
	"github.com/spf13/viper"
)

// GetLinkByID gets the id the user gives and add the link of the lecture to the clipboard
func GetLinkByID(id string) error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	jsonQuery := id + ".link"
	link := viper.GetString(jsonQuery)

	err := OpenMeeting(link)
	if err != nil {
		return err
	}

	return nil
}
