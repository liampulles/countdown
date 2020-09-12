package command

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/liampulles/countdown/pkg/dictionary"
)

// Letters provides answers for the Countdown "Letters round"
func Letters(args []string) error {
	flagSet := flag.NewFlagSet("letters", flag.ContinueOnError)
	dictionaryPathPtr := flagSet.String("dictionary", "dictionary.txt", "Path to the dictionary file to use")
	if err := parseLettersFlags(flagSet, args); err != nil {
		return err
	}

	letters, err := parseLetters(flagSet)
	if err != nil {
		return err
	}

	solution, err := solveLetters(letters, *dictionaryPathPtr)
	if err != nil {
		return err
	}

	fmt.Println(solution)
	return nil
}

// Check we match the type
var _ Command = Letters

func parseLettersFlags(flagSet *flag.FlagSet, args []string) error {
	if err := flagSet.Parse(args); err != nil {
		return fmt.Errorf("could not parse args: %w", err)
	}
	return nil
}

func parseLetters(flagSet *flag.FlagSet) (string, error) {
	if flagSet.NArg() != 1 {
		return "", fmt.Errorf("expected 1 positional argument for the letters, but got %d", flagSet.NArg())
	}
	return flagSet.Arg(0), nil
}

func solveLetters(letters string, dictionaryPath string) (string, error) {
	dict, err := dictionary.Load(dictionaryPath)
	if err != nil {
		return "", err
	}

	contained := dict.ContainedWords(letters)
	if len(contained) == 0 {
		return "", fmt.Errorf("no solutions found for letters")
	}

	sort.Slice(contained, func(i, j int) bool {
		return len(contained[j]) < len(contained[i])
	})

	slice := 20
	if len(contained) < slice {
		slice = len(contained)
	}
	return strings.Join(contained[:slice], ", "), nil
}
