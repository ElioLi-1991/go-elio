package logger

import (
	"fmt"
	"testing"
)

func TestLevel_String(t *testing.T) {
	var l Level = LevelInfo
	fmt.Println(l.String())
}

func TestParseLevel(t *testing.T) {
	levelStr := "DEBUG"
	l := ParseLevel(levelStr)
	fmt.Println(l)
}
