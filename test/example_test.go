package test

import (
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("test básico", func(t *testing.T) {
		if 1 != 1 {
			t.Error("Error en test básico")
		}
	})
}
