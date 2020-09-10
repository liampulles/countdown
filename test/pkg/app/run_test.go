package run_test

import (
	"testing"

	"github.com/liampulles/countdown/pkg/app"
	"github.com/stretchr/testify/assert"
)

func TestRun_WhenMissingCommand_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := []string{"some.exe.name"}

	// Exercise SUT
	actual := app.Run(fixture)

	// Verify results
	assert.Equal(t, 1, actual)
}

func TestRun_WhenGivenInvalidCommand_ShouldFail(t *testing.T) {
	// Setup fixture
	fixture := []string{"some.exe.name", "invalid.command"}

	// Exercise SUT
	actual := app.Run(fixture)

	// Verify results
	assert.Equal(t, 2, actual)
}

func TestRun_WhenGivenValidPassingCommand_ShouldPass(t *testing.T) {
	// Setup fixture
	fixture := []string{"some.exe.name", "conundrum"}

	// Exercise SUT
	actual := app.Run(fixture)

	// Verify results
	assert.Equal(t, 0, actual)
}
