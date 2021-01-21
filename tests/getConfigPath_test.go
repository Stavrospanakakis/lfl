package tests_test

import (
	"testing"

	"github.com/Stavrospanakakis/lfl/internal/services"
)

func TestGetConfigPath(t *testing.T) {
	s := services.MakeService()
	t.Run("Should return error", func(t *testing.T) {
		err := s.SetConfigPath()
		if err != nil {
			t.Fail()
		}
	})
}
