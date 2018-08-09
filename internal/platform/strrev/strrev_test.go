package strrev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseShouldReturnReversedInputString(t *testing.T) {
	actualResult, err := Reverse("hello")
	expectedResult := "olleh"
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedResult, actualResult)
}

func TestReverseShouldReturnErrorWhenInputEmptyString(t *testing.T) {
	actualResult, err := Reverse("")
	expectedResult := ""
	assert.Equal(t, ErrInputEmptyStr, err)
	assert.Equal(t, expectedResult, actualResult)
}
