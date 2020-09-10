package command

import (
	"fmt"
)

// Conundrum answers Countdown Conundrums
func Conundrum(args []string) error {
	fmt.Println("And your countdown conundrum for this evening is...")
	return nil
}

// Check we match the type
var _ Command = Conundrum
