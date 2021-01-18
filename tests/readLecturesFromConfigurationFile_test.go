package tests_test

import (
	"testing"

	"github.com/Stavrospanakakis/lfl/internal/services"
)

func TestReadLecturesFromConfigurationFile(t *testing.T) {
	s := services.MakeService()

	t.Run("Should return error", func(t *testing.T) {
		wrong_path := "this/is/a/wrong/path"
		err := s.ReadLecturesFromConfigurationFile(wrong_path)

		if err == nil {
			t.Fail()
		}

	})
}
