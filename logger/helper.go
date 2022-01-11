package logger

import (
	"fmt"
	"os"
)

type Option func(helper *Helper)

type Helper struct {
	logger Logger
}

func NewHelper(logger Logger,options ...Option) *Helper{
	h := &Helper{
		logger:logger,
	}
	for _,o := range options {
		o(h)
	}
	return h
}

func (h *Helper) Log(level Level,msg ...interface{}) {
	_ = h.logger.Log(level,msg...)
}

func (h *Helper) Debug (msg ...interface{}) {
	h.Log(LevelDebug,fmt.Sprint(msg...))
}

func (h *Helper) Debugf (format string,msg ...interface{}) {
	h.Log(LevelDebug,fmt.Sprintf(format,msg...))
}

func (h *Helper) Debugw (msg ...interface{}) {
	h.Log(LevelDebug,msg...)
}

func (h *Helper) Info (msg ...interface{}) {
	h.Log(LevelInfo,fmt.Sprint(msg...))
}

func (h *Helper) Infof (format string,msg ...interface{}) {
	h.Log(LevelInfo,fmt.Sprintf(format,msg...))
}

func (h *Helper) Infow (msg ...interface{}) {
	h.Log(LevelInfo,msg...)
}

func (h *Helper) Warn (msg ...interface{}) {
	h.Log(LevelWarn,fmt.Sprint(msg...))
}

func (h *Helper) Warnf (format string,msg ...interface{}) {
	h.Log(LevelWarn,fmt.Sprintf(format,msg...))
}

func (h *Helper) Warnw (msg ...interface{}) {
	h.Log(LevelWarn,msg...)
}

func (h *Helper) Error (msg ...interface{}) {
	h.Log(LevelError,fmt.Sprint(msg...))
}

func (h *Helper) Errorf (format string,msg ...interface{}) {
	h.Log(LevelError,fmt.Sprintf(format,msg...))
}

func (h *Helper) Errorw (msg ...interface{}) {
	h.Log(LevelError,msg...)
}

func (h *Helper) Fatal (msg ...interface{}) {
	h.Log(LevelFatal,fmt.Sprint(msg...))
	os.Exit(1)
}

func (h *Helper) Fatalf (format string,msg ...interface{}) {
	h.Log(LevelFatal,fmt.Sprintf(format,msg...))
	os.Exit(1)
}

func (h *Helper) Fatalw (msg ...interface{}) {
	h.Log(LevelFatal,msg...)
	os.Exit(1)
}
