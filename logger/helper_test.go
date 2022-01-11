package logger

import "testing"

func TestNewHelper(t *testing.T) {
	var fn Option = func(helper *Helper) {
		return
	}
	h := NewHelper(DefaultLogger,fn)
	h.Debug("this is debug!")
	h.Info("this is info!")
	h.Error("this is error！")
	h.Warn("this is warn!")
}
