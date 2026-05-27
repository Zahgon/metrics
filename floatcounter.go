package metrics

import (
	"io"
	"sync"
)

// NewFloatCounter registers and returns new counter of float64 type with the given name.
//
// name must be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
//
// The returned counter is safe to use from concurrent goroutines.
func NewFloatCounter(name string) *FloatCounter { _ = "STUB: not implemented"; return nil }

// FloatCounter is a float64 counter guarded by RWmutex.
//
// It may be used as a gauge if Add and Sub are called.
type FloatCounter struct {
	mu sync.Mutex
	n  float64
}

// Add adds n to fc.
func (fc *FloatCounter) Add(n float64) { _ = "STUB: not implemented"; return }

// Sub substracts n from fc.
func (fc *FloatCounter) Sub(n float64) { _ = "STUB: not implemented"; return }

// Get returns the current value for fc.
func (fc *FloatCounter) Get() float64 { _ = "STUB: not implemented"; return 0 }

// Set sets fc value to n.
func (fc *FloatCounter) Set(n float64) { _ = "STUB: not implemented"; return }

// marshalTo marshals fc with the given prefix to w.
func (fc *FloatCounter) marshalTo(prefix string, w io.Writer) { _ = "STUB: not implemented"; return }

func (fc *FloatCounter) metricType() string {
	_ = "STUB: not implemented"

	// GetOrCreateFloatCounter returns registered FloatCounter with the given name
	// or creates new FloatCounter if the registry doesn't contain FloatCounter with
	// the given name.
	//
	// name must be valid Prometheus-compatible metric with possible labels.
	// For instance,
	//
	//   - foo
	//   - foo{bar="baz"}
	//   - foo{bar="baz",aaa="b"}
	//
	// The returned FloatCounter is safe to use from concurrent goroutines.
	//
	// Performance tip: prefer NewFloatCounter instead of GetOrCreateFloatCounter.
	return ""
}

func GetOrCreateFloatCounter(name string) *FloatCounter { _ = "STUB: not implemented"; return nil }
