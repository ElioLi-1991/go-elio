package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

type stdLogger struct {
	logger *log.Logger
	pool *sync.Pool
}

func NewStdLogger(w io.Writer) Logger{
	return &stdLogger{
		logger:log.New(w,"",0),
		pool: &sync.Pool{New: func() interface{} {return new(bytes.Buffer)}},
	}
}

func (l *stdLogger) Log(level Level,msg ...interface{}) error {
	if len(msg) < 1 {
		return nil
	}
	buf := l.pool.Get().(*bytes.Buffer)
	buf.WriteString(fmt.Sprintf("[%s] %s ",level,time.Now().Format("2006/1/2 15:04:05")))
	for _,val := range msg {
		_,_ = fmt.Fprintf(buf,"%v",val)
	}
	_ = l.logger.Output(4,buf.String())
	buf.Reset()
	l.pool.Put(buf)
	return nil
}
