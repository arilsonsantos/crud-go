package validation

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateUserError(t *testing.T) {
	// Test case 1: Invalid field type error
	validationErr := errors.New("Invalid field type error")
	errorDto := ValidateUserError(validationErr)
	assert.Equal(t, "Error trying to convert fields.", errorDto.Message)
	assert.Equal(t, "Bad Request", errorDto.Err)

}
