package arithmetic_test

import (
	"testing"

	"github.com/liampulles/countdown/pkg/arithmetic"
	"github.com/stretchr/testify/assert"
)

func TestAdditionOperation_Sign_ShouldReturnPlusSign(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.AdditionOperation{}

	// Setup expectations
	expected := "+"

	// Exercise SUT
	actual := sut.Sign()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestAdditionOperation_Commutative_ShouldReturnTrue(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.AdditionOperation{}

	// Setup expectations
	expected := true

	// Exercise SUT
	actual := sut.Commutative()

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestAdditionOperation_UsableOn_ShouldReturnTrue(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.AdditionOperation{}

	// Setup expectations
	expected := true

	// Exercise SUT
	actual := sut.UsableOn(9, 7)

	// Verify results
	assert.Equal(t, expected, actual)
}

func TestAdditionOperation_Apply_ShouldReturnSum(t *testing.T) {
	// Setup fixture
	sut := &arithmetic.AdditionOperation{}

	// Setup expectations
	expected := 16

	// Exercise SUT
	actual := sut.Apply(9, 7)

	// Verify results
	assert.Equal(t, expected, actual)
}
