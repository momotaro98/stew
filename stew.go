package stew

import (
	"fmt"
	"runtime"
)

func Wrap(err error) error {
	pc := make([]uintptr, 1)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	s := fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
	return fmt.Errorf("%w %s", err, s)
}
