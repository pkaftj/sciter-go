package window

import (
	"github.com/iwctwbai/sciter-go"
	"runtime"
)

type Window struct {
	*sciter.Sciter
	creationFlags sciter.WindowCreationFlag
}

func (w *Window) run() {
	// runtime.LockOSThread()
}

// https://github.com/golang/go/wiki/LockOSThread
// https://github.com/iwctwbai/sciter-go/issues/201
func init() {
	runtime.LockOSThread()
}
