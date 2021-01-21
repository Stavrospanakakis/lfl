package tests_test

import (
	"testing"

	"github.com/Stavrospanakakis/lfl/internal/services"
)

func TestOpenMeeting(t *testing.T) {
	s := services.MakeService()
	link := "https://google.com"
	t.Run("Should not return error", func(t *testing.T) {
		err := s.OpenMeeting(link)
		if err != nil {
			t.Fail()
		}
	})
}
