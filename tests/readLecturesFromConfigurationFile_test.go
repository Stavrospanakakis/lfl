package tests_test

import (
	"testing"

	"github.com/Stavrospanakakis/lfl/internal/services"
)

func TestReadLecturesFromConfigurationFile(t *testing.T) {
	s := services.MakeService()

	t.Run("Should return error", func(t *testing.T) {
		wrong_path := "this/is/a/wrong/path"
		s.ConfigPath = wrong_path
		err := s.ReadLecturesFromConfigurationFile()

		if err == nil {
			t.Fail()
		}

	})
}
