package gostradamus

import (
	"fmt"
)

func FormatTokenIsNotMapped(formatToken string) error {
	return fmt.Errorf("FormatToken: %s is not mapped", formatToken)
}
