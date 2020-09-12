package dictionary_test

import (
	"fmt"
	"path"
	"testing"

	"github.com/liampulles/countdown/pkg/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestLoad_WhenGivenInvalidPath_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := "not.a.file"

	// Setup expectations
	expectedErr := "could not open dictionary: open not.a.file: no such file or directory"

	// Exercise SUT
	actual, err := dictionary.Load(fixture)

	// Verify results
	assert.Error(t, err, expectedErr)
	assert.Nil(t, actual)
}

func TestLoad_WhenGivenFileCannotBeRead_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := path.Join("testdata", "not.readable.txt")

	// Setup expectations
	expectedErr := "could not open dictionary: open not.a.file: no such file or directory"

	// Exercise SUT
	actual, err := dictionary.Load(fixture)

	// Verify results
	assert.Error(t, err, expectedErr)
	assert.Nil(t, actual)
}

func TestLoad_WhenGivenFileCanBeRead_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := path.Join("testdata", "valid.txt")

	// Setup expectations
	expected := dictionary.Dictionary{
		"and",
		"another",
		"word",
		"word",
	}

	// Exercise SUT
	actual, err := dictionary.Load(fixture)

	// Verify results
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestDictionary_Contains_WhenGivenValidInput_ShouldReturnAsExpected(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		sut      dictionary.Dictionary
		fixture  string
		expected bool
	}{
		// -> Not contained, return false
		{
			dictionary.Dictionary{"duck", "rabbit", "goose"},
			"cat",
			false,
		},
		// -> Contained, return true
		{
			dictionary.Dictionary{"duck", "rabbit", "goose"},
			"duck",
			true,
		},
		// -> Irregular, but still contained, return true
		{
			dictionary.Dictionary{"duck", "rabbit", "goose"},
			"Du ck",
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.sut.Contains(test.fixture)

			// Verify result
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestDictionary_Anagrams_WhenGivenValidInput_ShouldReturnAsExpected(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		sut      dictionary.Dictionary
		fixture  string
		expected []string
	}{
		// -> No anagrams, return empty
		{
			dictionary.Dictionary{"cat", "duck", "rabbit", "goose"},
			"cat",
			nil,
		},
		// -> One anagram - return it
		{
			dictionary.Dictionary{"cat", "duck", "rabbit", "goose", "dog"},
			"god",
			[]string{"dog"},
		},
		// -> Multiple anagrams - return them
		{
			dictionary.Dictionary{"cat", "duck", "rabbit", "goose", "dog", "ogd"},
			"god",
			[]string{"dog", "ogd"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.sut.Anagrams(test.fixture)

			// Verify result
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestDictionary_ContainedWords_WhenGivenValidInput_ShouldReturnAsExpected(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		sut      dictionary.Dictionary
		fixture  string
		expected []string
	}{
		// -> No contained words, return empty
		{
			dictionary.Dictionary{"cat", "duck", "rabbit", "goose"},
			"cat",
			nil,
		},
		// -> One contained word - return it
		{
			dictionary.Dictionary{"cat", "duck", "rabbit", "goose", "dog"},
			"godman",
			[]string{"dog"},
		},
		// -> Multiple contained words - return them
		{
			dictionary.Dictionary{"cat", "duck", "rabbit", "goose", "dog", "ogd"},
			"godman",
			[]string{"dog", "ogd"},
		},
		// Variation of above
		{
			dictionary.Dictionary{"cat", "duck", "rabbit", "goose", "dog", "ogd", "man"},
			"mangod",
			[]string{"dog", "ogd", "man"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.sut.ContainedWords(test.fixture)

			// Verify result
			assert.Equal(t, test.expected, actual)
		})
	}
}
