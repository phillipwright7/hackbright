package logger

import (
	"fmt"
	"io"
)

// Logger is an interface that represents a logger used to log messages, errors,
// and other information.
type Logger interface {
	Log(...interface{})
}

type logger struct {
	out io.Writer
}

func (lg *logger) Log(args ...interface{}) {
	fmt.Fprint(lg.out, args...)
	fmt.Fprintln(lg.out)
}

func New(w io.Writer) Logger {
	return &logger{out: w}
}
