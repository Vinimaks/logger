package logger

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println()
	logger := DefaultLogger
	logger.Notice("notice message:)")
	logger.Warning("warning message:)")
	logger.Error("error message:)")
	logger.Fatal("fatal message:)")
}
