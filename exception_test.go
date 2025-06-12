package redfox

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestException(t *testing.T) {
	exc := New("100", "Error 100")

	assert.Equal(t, "100", exc.ErrCode(), "Expected equal value")
	assert.Equal(t, "Error 100", exc.Message(), "Expected equal value")

	// Get display message
	assert.Equal(t, "", exc.DisplayMessage(), "Expected display message still empty")
	exc.WithDisplayMessage("Default error 100")
	assert.Equal(t, "Default error 100", exc.DisplayMessage(), "Expected display message updated")

	// Get metadata
	assert.Equal(t, map[string]interface{}{}, exc.Metadata(), "Expected metadata empty")
	exc.WithMetadata(map[string]interface{}{"key": "value"})
	assert.Equal(t, map[string]interface{}{"key": "value"}, exc.Metadata(), "Expected metadata updated")

	// Simulate with reason
	assert.Equal(t, nil, exc.Reason(), "Expected reason nil")
	exc.WithReason("Reason 100")
	assert.Equal(t, "Reason 100", exc.Reason(), "Expected reason updated")
}

func BenchmarkException(t *testing.B) {
	t.Run("Test Exception", func(t *testing.B) {
		metadata := map[string]interface{}{
			"key": "value",
		}
		_ = New("999", "Internal Server Error").
			WithBase(errors.New("Base internal server error")).
			WithDisplayMessage("Current server error").
			WithErrType("SERVER_ERROR").
			WithMetadata(metadata).
			WithReason("database not connected").
			WithStatusCode(500)
	})
}

func BenchmarkExceptionWithView(t *testing.B) {
	t.Run("Test Exception", func(t *testing.B) {
		metadata := map[string]interface{}{
			"key": "value",
		}
		err := New(999, "Internal Server Error").
			WithBase(errors.New("Base internal server error")).
			WithDisplayMessage("Current server error").
			WithErrType("SERVER_ERROR").
			WithMetadata(metadata).
			WithReason("database not connected").
			WithStatusCode(500)

		v, _ := json.MarshalIndent(err.View(), "", "  ")
		fmt.Println(string(v))
	})
}
