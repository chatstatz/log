# logger

A simple logging interface for Go.

## Usage

```go
package main

import (
  "os"

  "github.com/chatstatz/logger"
)

func main() {
  logger := logger.New(logger.Debug, os.Stderr)

  logger.Debug("debug message")
  logger.Info("info message")
  logger.Warn("warning message")
  logger.Error("error message")
  logger.Fatal("fatal error message")
}
```

Check it out it the [Go Playground](https://play.golang.org/p/bH9G8TJ-_6U).
