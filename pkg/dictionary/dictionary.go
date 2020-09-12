package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var alphaReg, _ = regexp.Compile("[^a-zA-Z]+")

// Dictionary contains words
type Dictionary []string

// Load reads a dictionary from a file
func Load(path string) (Dictionary, error) {
	var result Dictionary

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open dictionary: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanned := scanner.Text()
		regScanned := regularize(scanned)
		if len(regScanned) == 0 {
			continue
		}
		result = append(result, regScanned)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not scan dictionary: %w", err)
	}

	sort.Strings(result)
	return result, nil
}

// Contains returns true if the dictionary contains the given
// word, otherwise false
func (d Dictionary) Contains(word string) bool {
	regWord := regularize(word)
	for _, dictWord := range d {
		if regWord == dictWord {
			return true
		}
	}
	return false
}

// Anagrams returns words made up from the same letters
// as input in the dictionary
func (d Dictionary) Anagrams(input string) []string {
	inKey := convertToKey(input)
	return d.findKeyMatching(input, func(dictWord string) bool {
		dictKey := convertToKey(dictWord)
		return inKey == dictKey
	})
}

// ContainedWords returns words made up from some of the letters
// comprising input from the dictionary
func (d Dictionary) ContainedWords(input string) []string {
	return d.findKeyMatching(input, func(dictWord string) bool {
		return containsSubstrChars(input, dictWord)
	})
}

func (d Dictionary) findKeyMatching(input string, matcher func(dictWord string) bool) []string {
	regInput := regularize(input)
	var results []string
	for _, dictWord := range d {
		if dictWord == regInput {
			continue
		}
		if matcher(dictWord) {
			results = append(results, dictWord)
		}
	}
	return results
}

func regularize(input string) string {
	alphaNumStr := alphaReg.ReplaceAllString(input, "")
	return strings.ToLower(alphaNumStr)
}

func convertToKey(input string) string {
	var chars []rune = []rune(input)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func containsSubstrChars(base string, substr string) bool {
	used := make([]bool, len(base))
	matched := 0
	for _, subChar := range []rune(substr) {
		for i, baseChar := range base {
			if subChar == baseChar && !used[i] {
				used[i] = true
				matched++
			}
		}
	}
	return matched == len(substr)
}
