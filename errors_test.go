package gostradamus

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatTokenIsNotMapped(t *testing.T) {
	actual := FormatTokenIsNotMapped("123")
	assert.Equal(
		t,
		errors.New("FormatToken: 123 is not mapped"),
		actual,
	)
}
