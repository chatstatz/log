# logger

A simple logging interface for Go.

## Usage

```go
package main

import (
  "github.com/chatstatz/logger"
)

logger := logger.New(logger.Debug, os.Stderr)

logger.Debug("debug message")
logger.Info("info message")
logger.Warn("warning message")
logger.Error("error message")
logger.Fatal("fatal error message")
```
