package logger

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		level    string
		logLevel LogLevel
	}{
		{"debug", Debug},
		{"info", Info},
		{"warning", Warning},
		{"error", Error},
		{"fatal", Fatal},
	}

	for _, test := range tests {
		logger := New(test.level, os.Stderr)
		assert.Equal(t, test.logLevel, logger.level)
	}

}

func TestNewPanicsWithIncorrectLogLevel(t *testing.T) {
	tests := []struct {
		level  string
		errMsg string
	}{
		{"foo", "logger: \"foo\" is not a valid log level"},
		{"bar", "logger: \"bar\" is not a valid log level"},
	}

	for _, test := range tests {
		assert.PanicsWithValue(t, test.errMsg, func() {
			New(test.level, os.Stderr)
		})
	}
}

// DEBUG TESTS

func TestDebugIsCalledWhenDebugLogLevel(t *testing.T) {
	logger := New("debug", os.Stderr)
	logger.Debug("debug message should be called 1")
}

func TestDebugIsNotCalledWhenInfoLogLevel(t *testing.T) {
	logger := New("info", os.Stderr)
	logger.Debug("debug message should not be called")
}

// // DEBUGF TESTS

func TestDebugfIsCalledWhenDebugLogLevel(t *testing.T) {
	logger := New("debug", os.Stderr)
	logger.Debugf("debugf message should be called %d", 1)
}

func TestDebugfIsNotCalledWhenInfoLogLevel(t *testing.T) {
	logger := New("info", os.Stderr)
	logger.Debugf("debugf message should not be called %d", 1)
}

// INFO TESTS

func TestInfoIsCalledWhenInfoLogLevel(t *testing.T) {
	logger := New("info", os.Stderr)
	logger.Info("info message should be called 1")
}

func TestInfoIsCalledWhenDebugLogLevel(t *testing.T) {
	logger := New("debug", os.Stderr)
	logger.Info("info message should be called 2")
}

func TestInfoIsNotCalledWhenWarningLogLevel(t *testing.T) {
	logger := New("warning", os.Stderr)
	logger.Info("info message should not be called")
}

// INFOF TESTS

func TestInfofIsCalledWhenInfoLogLevel(t *testing.T) {
	logger := New("info", os.Stderr)
	logger.Infof("infof message should be called %d", 1)
}

func TestInfofIsCalledWhenDebugLogLevel(t *testing.T) {
	logger := New("debug", os.Stderr)
	logger.Infof("infof message should be called %d", 2)
}

func TestInfofIsNotCalledWhenWarningLogLevel(t *testing.T) {
	logger := New("warning", os.Stderr)
	logger.Infof("infof message should not be called %d", 1)
}

// WARN TESTS

func TestWarnIsCalledWhenWarningLogLevel(t *testing.T) {
	logger := New("warning", os.Stderr)
	logger.Warn("warn message should be called 1")
}

func TestWarnIsCalledWhenInfoLogLevel(t *testing.T) {
	logger := New("info", os.Stderr)
	logger.Warn("warn message should be called 2")
}

func TestWarnIsNotCalledWhenErrorLogLevel(t *testing.T) {
	logger := New("error", os.Stderr)
	logger.Warn("warn message should not be called")
}

// WARNF TESTS

func TestWarnfIsCalledWhenWarningLogLevel(t *testing.T) {
	logger := New("warning", os.Stderr)
	logger.Warnf("warnf message should be called %d", 1)
}

func TestWarnfIsCalledWhenInfoLogLevel(t *testing.T) {
	logger := New("info", os.Stderr)
	logger.Warnf("warnf message should be called %d", 2)
}

func TestWarnfIsNotCalledWhenErrorLogLevel(t *testing.T) {
	logger := New("error", os.Stderr)
	logger.Warnf("warnf message should not be called %d", 1)
}

// ERROR TESTS

func TestErrorIsCalledWhenErrorLogLevel(t *testing.T) {
	logger := New("error", os.Stderr)
	logger.Error("error message should be called 1")
}

func TestErrorIsCalledWhenWarningLogLevel(t *testing.T) {
	logger := New("warning", os.Stderr)
	logger.Error("error message should be called 2")
}

func TestErrorIsNotCalledWhenFatalLogLevel(t *testing.T) {
	logger := New("fatal", os.Stderr)
	logger.Error("error message should not be called")
}

// ERRORF TESTS

func TestErrorfIsCalledWhenErrorLogLevel(t *testing.T) {
	logger := New("error", os.Stderr)
	logger.Errorf("errorf message should be called %d", 1)
}

func TestErrorfIsCalledWhenWarningLogLevel(t *testing.T) {
	logger := New("warning", os.Stderr)
	logger.Errorf("errorf message should be called %d", 2)
}

func TestErrorfIsNotCalledWhenFatalLogLevel(t *testing.T) {
	logger := New("fatal", os.Stderr)
	logger.Errorf("errorf message should not be called %d", 1)
}

// FATAL TESTS

func TestFatalIsCalledWhenFatalLogLevel(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.Panics(t, func() {
		logger := New("fatal", os.Stderr)
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
		logger := New("error", os.Stderr)
		logger.Fatal("fatal message should be called 2")
	})
}

// FATALF TESTS

func TestFatalfIsCalledWhenFatalLogLevel(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.Panics(t, func() {
		logger := New("fatal", os.Stderr)
		logger.Fatalf("fatalf error message should be called %d", 1)
	})
}

func TestFatalfIsCalledWhenErrorLogLevel(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}

	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.Panics(t, func() {
		logger := New("error", os.Stderr)
		logger.Fatalf("fatalf error message should be called %d", 2)
	})
}
