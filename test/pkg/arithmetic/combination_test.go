package arithmetic_test

import (
	"testing"

	"github.com/liampulles/countdown/pkg/arithmetic"
	"github.com/stretchr/testify/assert"
)

func TestCombination_String_ShouldReturnVisualSum(t *testing.T) {
	// Setup fixture
	sut := arithmetic.NewCombination(3, 4, &arithmetic.MultiplicationOperation{})

	// Setup expectations
	expected := "3 × 4 = 12"

	// Exercise SUT
	actual := sut.String()

	// Verify results
	assert.Equal(t, expected, actual)
}
