package gostradamus

import (
	"fmt"
)

// FormatTokenIsNotMapped errors the given formatToken
func FormatTokenIsNotMapped(formatToken string) error {
	return fmt.Errorf("FormatToken: %s is not mapped", formatToken)
}
