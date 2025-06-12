package redfox

import (
	"errors"
	"testing"
)

func TestManager(t *testing.T) {
	manager := NewManagerStr()

	t.Run("Add exceptions", func(t *testing.T) {
		manager.Add(New("100", "Error 100"))
		manager.Add(New("101", "Error 101"))
		manager.Add(New("102", "Error 102"))
	})

	t.Run("Get exceptions", func(t *testing.T) {
		expectedErrCode := map[string]string{
			"100": "Error 100",
			"101": "Error 101",
			"102": "Error 102",
		}

		for code, expected := range expectedErrCode {
			e := manager.Get(code)
			if e.Message() != expected {
				t.Errorf("Expected message '%s' for code '%s', but got '%s'", expected, code, e.Message())
				t.Fail()
				return
			}
		}
	})

	t.Run("Mutate exception", func(t *testing.T) {
		e100 := manager.Get("100")

		if e100.ErrCode() != "100" {
			t.Errorf("Expected ErrCode '100', but got '%s'", e100.ErrCode())
			t.Fail()
			return
		}
		// Update exception
		e100.WithDisplayMessage("Default error 100")
		if e100.DisplayMessage() != "Default error 100" {
			t.Errorf("Expected DisplayMessage 'Default error 100', but got '%s'", e100.DisplayMessage())
			t.Fail()
		}

		// Simulation exceptions cloned
		e100def := manager.Get("100")
		if e100def.ErrCode() != "100" {
			t.Errorf("Expected ErrCode '100', but got '%s'", e100def.ErrCode())
			t.Fail()
		} else if e100def.DisplayMessage() != "" {
			t.Errorf("Expected DisplayMessage still empty cause mutable, but got '%s'", e100def.DisplayMessage())
			t.Fail()
		}
	})
}

func BenchmarkManager(t *testing.B) {
	t.Run("Benchmark Manager", func(t *testing.B) {
		manager := NewManagerStr()
		manager.Add(New("100", "Error 100"))
		manager.Add(New("101", "Error 101"))
		manager.Add(New("102", "Error 102"))

		// Get error
		e100 := manager.Get("100")

		// Update exception
		e100.WithDisplayMessage("Default error 100")
		e100.WithBase(errors.New("Base Error 100"))
		e100.WithReason("Reason 100")
		e100.WithMetadata(map[string]interface{}{
			"key": "value",
		})

		_ = e100.View()
	})
}
