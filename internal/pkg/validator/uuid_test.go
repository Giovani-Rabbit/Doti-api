package val_test

import (
	"testing"

	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
	"github.com/google/uuid"
)

func TestUUidValidator(t *testing.T) {
	t.Run("The uuid must be valid", func(t *testing.T) {
		isValid := val.IsValidUUID(uuid.New().String())

		if !isValid {
			t.Fatal("uuid validation should return true")
		}
	})

	t.Run("The uuid must be invalid", func(t *testing.T) {
		isValid := val.IsValidUUID("12345")

		if isValid {
			t.Fatal("uuid validation should return false")
		}
	})
}
