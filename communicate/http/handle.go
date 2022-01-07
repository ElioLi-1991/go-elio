package http

import "net/http"

type HandleFuncComplex []HandleFunc

type HandleFunc struct {
	Path string
	Func http.HandlerFunc
}
