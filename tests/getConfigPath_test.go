package tests_test

import (
	"testing"

	"github.com/Stavrospanakakis/lfl/internal/services"
)

func TestGetConfigPath(t *testing.T) {
	s := services.MakeService()
	t.Run("Should return error", func(t *testing.T) {
		_, err := s.GetConfigPath()
		if err != nil {
			t.Fail()
		}
	})
}
