package command

import (
	"flag"
	"fmt"

	"github.com/liampulles/countdown/pkg/dictionary"
)

// Conundrum answers Countdown Conundrums
func Conundrum(args []string) error {
	flagSet := flag.NewFlagSet("conundrum", flag.ContinueOnError)
	dictionaryPathPtr := flagSet.String("dictionary", "dictionary.txt", "Path to the dictionary file to use")
	if err := parseFlags(flagSet, args); err != nil {
		return err
	}

	conundrum, err := parseConundrum(flagSet)
	if err != nil {
		return err
	}

	solution, err := solveConundrum(conundrum, *dictionaryPathPtr)
	if err != nil {
		return err
	}

	fmt.Println(solution)
	return nil
}

// Check we match the type
var _ Command = Conundrum

func parseFlags(flagSet *flag.FlagSet, args []string) error {
	if err := flagSet.Parse(args); err != nil {
		return fmt.Errorf("could not parse args: %w", err)
	}
	return nil
}

func parseConundrum(flagSet *flag.FlagSet) (string, error) {
	if flagSet.NArg() != 1 {
		return "", fmt.Errorf("expected 1 positional argument for the conundrum, but got %d", flagSet.NArg())
	}
	return flagSet.Arg(0), nil
}

func solveConundrum(conundrum string, dictionaryPath string) (string, error) {
	dict, err := dictionary.Load(dictionaryPath)
	if err != nil {
		return "", err
	}

	anagrams := dict.Anagrams(conundrum)
	if len(anagrams) == 0 {
		return "", fmt.Errorf("no solutions found for conundrum")
	}

	return anagrams[0], nil
}
