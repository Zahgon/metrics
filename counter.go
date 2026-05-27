package metrics

import (
	"io"
)

// NewCounter registers and returns new counter with the given name.
//
// name must be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
//
// The returned counter is safe to use from concurrent goroutines.
func NewCounter(name string) *Counter { _ = "STUB: not implemented"; return nil }

// Counter is a counter.
//
// It may be used as a gauge if Dec and Set are called.
type Counter struct {
	n uint64
}

// Inc increments c.
func (c *Counter) Inc() { _ = "STUB: not implemented"; return }

// Dec decrements c.
func (c *Counter) Dec() { _ = "STUB: not implemented"; return }

// Add adds n to c.
func (c *Counter) Add(n int) { _ = "STUB: not implemented"; return }

// AddInt64 adds n to c.
func (c *Counter) AddInt64(n int64) { _ = "STUB: not implemented"; return }

// Get returns the current value for c.
func (c *Counter) Get() uint64 { _ = "STUB: not implemented"; return 0 }

// Set sets c value to n.
func (c *Counter) Set(n uint64) { _ = "STUB: not implemented"; return }

// marshalTo marshals c with the given prefix to w.
func (c *Counter) marshalTo(prefix string, w io.Writer) { _ = "STUB: not implemented"; return }

func (c *Counter) metricType() string {
	_ = "STUB: not implemented"

	// GetOrCreateCounter returns registered counter with the given name
	// or creates new counter if the registry doesn't contain counter with
	// the given name.
	//
	// name must be valid Prometheus-compatible metric with possible labels.
	// For instance,
	//
	//   - foo
	//   - foo{bar="baz"}
	//   - foo{bar="baz",aaa="b"}
	//
	// The returned counter is safe to use from concurrent goroutines.
	//
	// Performance tip: prefer NewCounter instead of GetOrCreateCounter.
	return ""
}

func GetOrCreateCounter(name string) *Counter { _ = "STUB: not implemented"; return nil }
