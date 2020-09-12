package arithmetic_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/countdown/pkg/arithmetic"
	"github.com/stretchr/testify/assert"
)

func TestDivisionOperation_Sign_ShouldReturnTimesSign(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.DivisionOperation{}

	// Setup expectations
	expected := "÷"

	// Exercise SUT
	actual := sut.Sign()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestDivisionOperation_Commutative_ShouldReturnFalse(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.DivisionOperation{}

	// Setup expectations
	expected := false

	// Exercise SUT
	actual := sut.Commutative()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestDivisionOperation_UsableOn_GivenValidInput_ShouldReturnAsExpected(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.DivisionOperation{}
	var tests = []struct {
		a        int
		b        int
		expected bool
	}{
		{1, 0, false},
		{0, 1, false},
		{2, 3, false},
		{9, 2, false},
		{9, 3, true},
		{9, 9, true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := sut.UsableOn(test.a, test.b)

			// Verify result
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestDivisionOperation_Apply_ShouldReturnFraction(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.DivisionOperation{}

	// Setup expectations
	expected := 3

	// Exercise SUT
	actual := sut.Apply(9, 3)

	// Verify results
	assert.Equal(t, expected, actual)
}
