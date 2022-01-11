package logger

import "log"

var DefaultLogger Logger = NewStdLogger(log.Writer())


type Logger interface {
	Log(level Level,msg ...interface{}) error
}

type logger struct {

}
