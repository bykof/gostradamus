package gostradamus

import (
	"errors"
	"fmt"
)

func FormatTokenIsNotMapped(formatToken string) error {
	return errors.New(fmt.Sprintf("FormatToken: %s is not mapped", formatToken))
}
