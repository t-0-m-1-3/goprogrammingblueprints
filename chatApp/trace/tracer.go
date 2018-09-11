package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes the object capable of
// tracing events throughout the code/
type Tracer interface {
	Trace(...interface{})
}

// New Creates a new Tracer that will write the output to io.writer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// tracer is a Tracer that writes to an io.Writer
type tracer struct {
	out io.Writer
}

// Trace writes the arguments to this Tracers io.Writer
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// nilTracer will go here
type nilTracer struct{}

// Trace for nil tracer does nada
func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a tracer that will ignore calls to Trace
func Off() Tracer {
	return &nilTracer{}
}
