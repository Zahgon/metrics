package metrics

import (
	"io"
	"sync"
	"time"
)

// PrometheusHistogramDefaultBuckets is a list of the default bucket upper
// bounds. Those default buckets are quite generic, and it is recommended to
// pick custom buckets for improved accuracy.
var PrometheusHistogramDefaultBuckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

// PrometheusHistogram is a histogram for non-negative values with pre-defined buckets
//
// Each bucket contains a counter for values in the given range.
// Each bucket is exposed via the following metric:
//
// <metric_name>_bucket{<optional_tags>,le="upper_bound"} <counter>
//
// Where:
//
//   - <metric_name> is the metric name passed to NewPrometheusHistogram
//   - <optional_tags> is optional tags for the <metric_name>, which are passed to NewPrometheusHistogram
//   - <upper_bound> - upper bound of the current bucket. all samples <= upper_bound are in that bucket
//   - <counter> - the number of hits to the given bucket during Update* calls
//
// Next to the bucket metrics, two additional metrics track the total number of
// samples (_count) and the total sum (_sum) of all samples:
//
//   - <metric_name>_sum{<optional_tags>} <counter>
//   - <metric_name>_count{<optional_tags>} <counter>
type PrometheusHistogram struct {
	// mu guarantees synchronous update for all the counters.
	//
	// Do not use sync.RWMutex, since it has zero sense from performance PoV.
	// It only complicates the code.
	mu sync.Mutex

	// upperBounds and buckets are aligned by element position:
	// upperBounds[i] defines the upper bound for buckets[i].
	// buckets[i] contains the count of elements <= upperBounds[i]
	upperBounds []float64
	buckets     []uint64

	// count is the counter for all observations on this histogram
	count uint64

	// sum is the sum of all the values put into Histogram
	sum float64
}

// Reset resets previous observations in h.
func (h *PrometheusHistogram) Reset() { _ = "STUB: not implemented"; return }

// Update updates h with v.
//
// Negative values and NaNs are ignored.
func (h *PrometheusHistogram) Update(v float64) { _ = "STUB: not implemented"; return }

// Skip NaNs and negative values.

// +Inf, nothing to do, already accounted for in the total sum

// UpdateDuration updates request duration based on the given startTime.
func (h *PrometheusHistogram) UpdateDuration(startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// NewPrometheusHistogram creates and returns new PrometheusHistogram with the given name
// and PrometheusHistogramDefaultBuckets.
//
// name must be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
//
// The returned histogram is safe to use from concurrent goroutines.
func NewPrometheusHistogram(name string) *PrometheusHistogram {
	_ = "STUB: not implemented"
	return nil
}

// NewPrometheusHistogramExt creates and returns new PrometheusHistogram with the given name
// and given upperBounds.
//
// name must be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
//
// The returned histogram is safe to use from concurrent goroutines.
func NewPrometheusHistogramExt(name string, upperBounds []float64) *PrometheusHistogram {
	_ = "STUB: not implemented"
	return nil
}

// GetOrCreatePrometheusHistogram returns registered PrometheusHistogram with the given name
// or creates a new PrometheusHistogram if the registry doesn't contain histogram with
// the given name.
//
// name must be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
//
// The returned histogram is safe to use from concurrent goroutines.
//
// Performance tip: prefer NewPrometheusHistogram instead of GetOrCreatePrometheusHistogram.
func GetOrCreatePrometheusHistogram(name string) *PrometheusHistogram {
	_ = "STUB: not implemented"
	return nil
}

// GetOrCreatePrometheusHistogramExt returns registered PrometheusHistogram with the given name and
// upperBounds or creates new PrometheusHistogram if the registry doesn't contain histogram
// with the given name.
//
// name must be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
//
// The returned histogram is safe to use from concurrent goroutines.
//
// Performance tip: prefer NewPrometheusHistogramExt instead of GetOrCreatePrometheusHistogramExt.
func GetOrCreatePrometheusHistogramExt(name string, upperBounds []float64) *PrometheusHistogram {
	_ = "STUB: not implemented"
	return nil
}

func newPrometheusHistogram(upperBounds []float64) *PrometheusHistogram {
	_ = "STUB: not implemented"
	return nil
}

// ignore +Inf bucket as it is covered anyways

func mustValidateBuckets(upperBounds []float64) { _ = "STUB: not implemented"; return }

// ValidateBuckets validates the given upperBounds and returns an error
// if validation failed.
func ValidateBuckets(upperBounds []float64) error { _ = "STUB: not implemented"; return nil }

// LinearBuckets returns a list of upperBounds for PrometheusHistogram,
// and whose distribution is as follows:
//
//	[start, start + width, start + 2 * width, ... start + (count-1) * width]
//
// Panics if given start, width and count produce negative buckets or none buckets at all.
func LinearBuckets(start, width float64, count int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// ExponentialBuckets returns a list of upperBounds for PrometheusHistogram,
// and whose distribution is as follows:
//
//	[start, start * factor pow 1, start * factor pow 2, ... start * factor pow (count-1)]
//
// Panics if given start, width and count produce negative buckets or none buckets at all.
func ExponentialBuckets(start, factor float64, count int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (h *PrometheusHistogram) marshalTo(prefix string, w io.Writer) {
	_ = "STUB: not implemented"
	return
}

func (h *PrometheusHistogram) metricType() string { _ = "STUB: not implemented"; return "" }
