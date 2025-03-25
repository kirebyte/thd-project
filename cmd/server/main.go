package main

import (
	"fmt"

	"github.com/kirebyte/thd-project/internal/logger"
)

func main() {
	fmt.Println("Hello, World!")
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
}
