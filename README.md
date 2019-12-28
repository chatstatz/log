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
  logger := logger.New("debug", os.Stderr)

  logger.Debug("debug message")
  logger.Info("info message")
  logger.Warn("warning message")
  logger.Error("error message")
  logger.Fatal("fatal error message")
}
```

Check it out it the [Go Playground](https://play.golang.org/p/tFHRLgkCdnv).

## License

This package is distributed under the terms of the [MIT](LICENSE) License.
