package run_test

import (
	"path"
	"testing"

	"github.com/liampulles/countdown/pkg/command"

	"github.com/stretchr/testify/assert"
)

func TestLetters_WhenGivenInvalidFlag_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := []string{"-invalidFlag"}

	// Setup expectations
	expectedErr := "could not parse args: flag provided but not defined: -invalidFlag"

	// Exercise SUT
	err := command.Letters(fixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestLetters_WhenGivenInvalidPositionalArgs_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := []string{"one.arg", "two.args"}

	// Setup expectations
	expectedErr := "expected 1 positional argument for the letters, but got 2"

	// Exercise SUT
	err := command.Letters(fixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestLetters_WhenGivenInvalidDictionaryLocation_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := []string{"-dictionary", "invalid.path", "conundrum"}

	// Setup expectations
	expectedErr := "could not open dictionary: open invalid.path: no such file or directory"

	// Exercise SUT
	err := command.Letters(fixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestLetters_WhenNoSolutionsForLetters_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := []string{"-dictionary", path.Join("testdata", "dictionary.txt"), "conundrum"}

	// Setup expectations
	expectedErr := "no solutions found for letters"

	// Exercise SUT
	err := command.Letters(fixture)

	// Verify results
	assert.EqualError(t, err, expectedErr)
}

func TestLetters_WhenThereIsASolutionForLetters_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := []string{"-dictionary", path.Join("testdata", "dictionary.txt"), "sotcidbuh"}

	// Exercise SUT
	err := command.Letters(fixture)

	// Verify results
	assert.NoError(t, err)
}
