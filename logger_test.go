package logger

import (
	"os"
	"testing"

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
