package arithmetic_test

import (
	"testing"

	"github.com/liampulles/countdown/pkg/arithmetic"
	"github.com/stretchr/testify/assert"
)

func TestMultiplicationOperation_Sign_ShouldReturnTimesSign(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.MultiplicationOperation{}

	// Setup expectations
	expected := "×"

	// Exercise SUT
	actual := sut.Sign()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestMultiplicationOperation_Commutative_ShouldReturnTrue(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.MultiplicationOperation{}

	// Setup expectations
	expected := true

	// Exercise SUT
	actual := sut.Commutative()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestMultiplicationOperation_UsableOn_ShouldReturnTrue(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.MultiplicationOperation{}

	// Setup expectations
	expected := true

	// Exercise SUT
	actual := sut.UsableOn(9, 7)

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestMultiplicationOperation_Apply_ShouldReturnProduct(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.MultiplicationOperation{}

	// Setup expectations
	expected := 63

	// Exercise SUT
	actual := sut.Apply(9, 7)

	// Verify results
	assert.Equal(t, expected, actual)
}
