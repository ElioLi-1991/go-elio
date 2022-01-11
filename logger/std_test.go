package logger

import (
	"log"
	"testing"
)

func TestNewStdLogger(t *testing.T) {
	std := NewStdLogger(log.Writer())
	std.Log(LevelInfo,"this is info!")
	std.Log(LevelWarn,"this is Warn!")
}
