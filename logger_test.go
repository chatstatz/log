package logger

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		level LogLevel
	}{
		{Debug},
		{Info},
		{Warning},
		{Error},
		{Fatal},
	}

	for _, test := range tests {
		logger := New(test.level, os.Stderr)

		assert.Equal(t, test.level, logger.level)
	}

}

func TestNewPanicsWithIncorrectLogLevel(t *testing.T) {
	tests := []struct {
		level LogLevel
	}{
		{-1},
		{5},
	}

	for _, test := range tests {
		assert.Panics(t, func() { New(test.level, os.Stderr) })
	}
}

// DEBUG TESTS

func TestDebugIsCalledWhenDebugLogLevel(t *testing.T) {
	logger := New(Debug, os.Stderr)
	logger.Debug("debug message should be called 1")
}

func TestDebugIsNotCalledWhenInfoLogLevel(t *testing.T) {
	logger := New(Info, os.Stderr)
	logger.Debug("debug message should not be called")
}

// INFO TESTS

func TestInfoIsCalledWhenInfoLogLevel(t *testing.T) {
	logger := New(Info, os.Stderr)
	logger.Info("info message should be called 1")
}

func TestInfoIsCalledWhenDebugLogLevel(t *testing.T) {
	logger := New(Debug, os.Stderr)
	logger.Info("info message should be called 2")
}

func TestInfoIsNotCalledWhenWarningLogLevel(t *testing.T) {
	logger := New(Warning, os.Stderr)
	logger.Info("info message should not be called")
}

// WARNING TESTS

func TestWarnIsCalledWhenWarningLogLevel(t *testing.T) {
	logger := New(Warning, os.Stderr)
	logger.Warn("warning message should be called 1")
}

func TestWarnIsCalledWhenInfoLogLevel(t *testing.T) {
	logger := New(Info, os.Stderr)
	logger.Warn("warning message should be called 2")
}

func TestWarnIsNotCalledWhenErrorLogLevel(t *testing.T) {
	logger := New(Error, os.Stderr)
	logger.Warn("warning message should not be called")
}

// ERROR TESTS

func TestErrorIsCalledWhenErrorLogLevel(t *testing.T) {
	logger := New(Error, os.Stderr)
	logger.Error("error message should be called 1")
}

func TestErrorIsCalledWhenWarningLogLevel(t *testing.T) {
	logger := New(Warning, os.Stderr)
	logger.Error("error message should be called 2")
}

func TestErrorIsNotCalledWhenFatalLogLevel(t *testing.T) {
	logger := New(Fatal, os.Stderr)
	logger.Error("error message should not be called")
}

// FATAL TESTS

func TestFatalIsCalledWhenFatalLogLevel(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.Panics(t, func() {
		logger := New(Fatal, os.Stderr)
		logger.Fatal("fatal message should be called 1")
	})
}

func TestFatalIsCalledWhenErrorLogLevel(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.Panics(t, func() {
		logger := New(Error, os.Stderr)
		logger.Fatal("fatal message should be called 2")
	})
}
