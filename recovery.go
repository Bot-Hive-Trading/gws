package gws

import (
	"log"
	"runtime"
	"unsafe"
)

type Logger interface {
	Error(v ...any)
}

type stdLogger struct{}

func (c *stdLogger) Error(v ...any) {
	log.Println(v...)
}

func Recovery(logger Logger) {
	if e := recover(); e != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		msg := *(*string)(unsafe.Pointer(&buf))
		logger.Error("fatal error:", e, msg)
	}
}
