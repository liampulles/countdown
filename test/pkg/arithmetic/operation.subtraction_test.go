package arithmetic_test

import (
	"testing"

	"github.com/liampulles/countdown/pkg/arithmetic"
	"github.com/stretchr/testify/assert"
)

func TestSubtractionOperation_Sign_ShouldReturnMinusSign(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.SubtractionOperation{}

	// Setup expectations
	expected := "-"

	// Exercise SUT
	actual := sut.Sign()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestSubtractionOperation_Commutative_ShouldReturnFalse(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.SubtractionOperation{}

	// Setup expectations
	expected := false

	// Exercise SUT
	actual := sut.Commutative()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestSubtractionOperation_UsableOn_WhenSubtractionGreaterThanZero_ShouldReturnTrue(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.SubtractionOperation{}

	// Setup expectations
	expected := true

	// Exercise SUT
	actual := sut.UsableOn(9, 7)

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestSubtractionOperation_UsableOn_WhenSubtractionSmallerThanZero_ShouldReturnFalse(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.SubtractionOperation{}

	// Setup expectations
	expected := false

	// Exercise SUT
	actual := sut.UsableOn(7, 9)

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestSubtractionOperation_Apply_ShouldReturnDifference(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.SubtractionOperation{}

	// Setup expectations
	expected := 2

	// Exercise SUT
	actual := sut.Apply(9, 7)

	// Verify results
	assert.Equal(t, expected, actual)
}
