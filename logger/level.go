package logger

type Level int8

const (
	LevelDebug Level = iota + 1
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string{
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

func ParseLevel(s string) Level {
	switch s {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelWarn
	case "FATAL":
		return LevelFatal
	}
	return LevelInfo
}