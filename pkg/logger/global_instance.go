package logger

import "sync"

var gl Logger
var lMutex sync.RWMutex

func SetGlobalLogger(l Logger) {
	lMutex.Lock()
	defer lMutex.Unlock()
	gl = l
}

func L() Logger {
	lMutex.RLock()
	g := gl
	lMutex.RUnlock()
	return g
}

var GL Logger = &NopLogger{}
