package logger

import (
	"testing"
)

func TestCriticalF(t *testing.T) {
	logger := NewLogger("logger")

	logger.CriticalF("This is critical error %d", 15)
}

func TestDebugF(t *testing.T) {
	logger := NewLogger("logger")

	logger.DebugF("This is debug error %d", 15)
}

func TestWarningF(t *testing.T) {
	logger := NewLogger("logger")

	logger.WarningF("This is warning error %d", 15)
}

func TestErrorF(t *testing.T) {
	logger := NewLogger("logger")

	logger.ErrorF("This is error error %d", 15)
}

func TestInfoF(t *testing.T) {
	logger := NewLogger("logger")

	logger.InfoF("This is info error %d", 15)
}
