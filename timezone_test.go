package gostradamus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimezone_Location(t *testing.T) {
	actual := Timezone("notexist")
	assert.PanicsWithError(
		t,
		"unknown time zone notexist",
		func() {
			actual.Location()
		},
	)
}
