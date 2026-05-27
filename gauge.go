package metrics

import (
	"io"
)

// NewGauge registers and returns gauge with the given name, which calls f to obtain gauge value.
//
// name must be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
//
// f must be safe for concurrent calls.
// if f is nil, then it is expected that the gauge value is changed via Set(), Inc(), Dec() and Add() calls.
//
// The returned gauge is safe to use from concurrent goroutines.
//
// See also FloatCounter for working with floating-point values.
func NewGauge(name string, f func() float64) *Gauge { _ = "STUB: not implemented"; return nil }

// Gauge is a float64 gauge.
type Gauge struct {
	// valueBits contains uint64 representation of float64 passed to Gauge.Set.
	valueBits uint64

	// f is a callback, which is called for returning the gauge value.
	f func() float64
}

// Get returns the current value for g.
func (g *Gauge) Get() float64 { _ = "STUB: not implemented"; return 0 }

// Set sets g value to v.
//
// The g must be created with nil callback in order to be able to call this function.
func (g *Gauge) Set(v float64) { _ = "STUB: not implemented"; return }

// Inc increments g by 1.
//
// The g must be created with nil callback in order to be able to call this function.
func (g *Gauge) Inc() {
	_ = "STUB: not implemented"

	// Dec decrements g by 1.
	//
	// The g must be created with nil callback in order to be able to call this function.
	return
}

func (g *Gauge) Dec() {
	_ = "STUB: not implemented"

	// Add adds fAdd to g. fAdd may be positive and negative.
	//
	// The g must be created with nil callback in order to be able to call this function.
	return
}

func (g *Gauge) Add(fAdd float64) { _ = "STUB: not implemented"; return }

func (g *Gauge) marshalTo(prefix string, w io.Writer) { _ = "STUB: not implemented"; return }

// Marshal integer values without scientific notation

func (g *Gauge) metricType() string {
	_ = "STUB: not implemented"

	// GetOrCreateGauge returns registered gauge with the given name
	// or creates new gauge if the registry doesn't contain gauge with
	// the given name.
	//
	// name must be valid Prometheus-compatible metric with possible labels.
	// For instance,
	//
	//   - foo
	//   - foo{bar="baz"}
	//   - foo{bar="baz",aaa="b"}
	//
	// The returned gauge is safe to use from concurrent goroutines.
	//
	// Performance tip: prefer NewGauge instead of GetOrCreateGauge.
	//
	// See also FloatCounter for working with floating-point values.
	return ""
}

func GetOrCreateGauge(name string, f func() float64) *Gauge { _ = "STUB: not implemented"; return nil }
