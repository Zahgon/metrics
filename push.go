package metrics

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"compress/gzip"
)

// PushOptions is the list of options, which may be applied to InitPushWithOptions().
type PushOptions struct {
	// ExtraLabels is an optional comma-separated list of `label="value"` labels, which must be added to all the metrics before pushing them to pushURL.
	ExtraLabels string

	// Headers is an optional list of HTTP headers to add to every push request to pushURL.
	//
	// Every item in the list must have the form `Header: value`. For example, `Authorization: Custom my-top-secret`.
	Headers []string

	// Whether to disable HTTP request body compression before sending the metrics to pushURL.
	//
	// By default the compression is enabled.
	DisableCompression bool

	// Method is HTTP request method to use when pushing metrics to pushURL.
	//
	// By default the Method is GET.
	Method string

	// Optional WaitGroup for waiting until all the push workers created with this WaitGroup are stopped.
	WaitGroup *sync.WaitGroup
}

// InitPushWithOptions sets up periodic push for globally registered metrics to the given pushURL with the given interval.
//
// The periodic push is stopped when ctx is canceled.
// It is possible to wait until the background metrics push worker is stopped on a WaitGroup passed via opts.WaitGroup.
//
// If pushProcessMetrics is set to true, then 'process_*' and `go_*` metrics are also pushed to pushURL.
//
// opts may contain additional configuration options if non-nil.
//
// The metrics are pushed to pushURL in Prometheus text exposition format.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
//
// It is OK calling InitPushWithOptions multiple times with different pushURL -
// in this case metrics are pushed to all the provided pushURL urls.
func InitPushWithOptions(ctx context.Context, pushURL string, interval time.Duration, pushProcessMetrics bool, opts *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// InitPushProcessMetrics sets up periodic push for 'process_*' metrics to the given pushURL with the given interval.
//
// extraLabels may contain comma-separated list of `label="value"` labels, which will be added
// to all the metrics before pushing them to pushURL.
//
// The metrics are pushed to pushURL in Prometheus text exposition format.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
//
// It is OK calling InitPushProcessMetrics multiple times with different pushURL -
// in this case metrics are pushed to all the provided pushURL urls.
func InitPushProcessMetrics(pushURL string, interval time.Duration, extraLabels string) error {
	_ = "STUB: not implemented"
	return nil
}

// InitPush sets up periodic push for globally registered metrics to the given pushURL with the given interval.
//
// extraLabels may contain comma-separated list of `label="value"` labels, which will be added
// to all the metrics before pushing them to pushURL.
//
// If pushProcessMetrics is set to true, then 'process_*' and `go_*` metrics are also pushed to pushURL.
//
// The metrics are pushed to pushURL in Prometheus text exposition format.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
//
// It is OK calling InitPush multiple times with different pushURL -
// in this case metrics are pushed to all the provided pushURL urls.
func InitPush(pushURL string, interval time.Duration, extraLabels string, pushProcessMetrics bool) error {
	_ = "STUB: not implemented"
	return nil
}

// PushMetrics pushes globally registered metrics to pushURL.
//
// If pushProcessMetrics is set to true, then 'process_*' and `go_*` metrics are also pushed to pushURL.
//
// opts may contain additional configuration options if non-nil.
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
func PushMetrics(ctx context.Context, pushURL string, pushProcessMetrics bool, opts *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// InitPushWithOptions sets up periodic push for metrics from s to the given pushURL with the given interval.
//
// The periodic push is stopped when the ctx is canceled.
// It is possible to wait until the background metrics push worker is stopped on a WaitGroup passed via opts.WaitGroup.
//
// opts may contain additional configuration options if non-nil.
//
// The metrics are pushed to pushURL in Prometheus text exposition format.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
//
// It is OK calling InitPushWithOptions multiple times with different pushURL -
// in this case metrics are pushed to all the provided pushURL urls.
func (s *Set) InitPushWithOptions(ctx context.Context, pushURL string, interval time.Duration, opts *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// InitPush sets up periodic push for metrics from s to the given pushURL with the given interval.
//
// extraLabels may contain comma-separated list of `label="value"` labels, which will be added
// to all the metrics before pushing them to pushURL.
//
// The metrics are pushed to pushURL in Prometheus text exposition format.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
//
// It is OK calling InitPush multiple times with different pushURL -
// in this case metrics are pushed to all the provided pushURL urls.
func (s *Set) InitPush(pushURL string, interval time.Duration, extraLabels string) error {
	_ = "STUB: not implemented"
	return nil
}

// PushMetrics pushes s metrics to pushURL.
//
// opts may contain additional configuration options if non-nil.
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
func (s *Set) PushMetrics(ctx context.Context, pushURL string, opts *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// InitPushExt sets up periodic push for metrics obtained by calling writeMetrics with the given interval.
//
// extraLabels may contain comma-separated list of `label="value"` labels, which will be added
// to all the metrics before pushing them to pushURL.
//
// The writeMetrics callback must write metrics to w in Prometheus text exposition format without timestamps and trailing comments.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
//
// It is OK calling InitPushExt multiple times with different pushURL -
// in this case metrics are pushed to all the provided pushURL urls.
//
// It is OK calling InitPushExt multiple times with different writeMetrics -
// in this case all the metrics generated by writeMetrics callbacks are written to pushURL.
func InitPushExt(pushURL string, interval time.Duration, extraLabels string, writeMetrics func(w io.Writer)) error {
	_ = "STUB: not implemented"
	return nil
}

// InitPushExtWithOptions sets up periodic push for metrics obtained by calling writeMetrics with the given interval.
//
// The writeMetrics callback must write metrics to w in Prometheus text exposition format without timestamps and trailing comments.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// The periodic push is stopped when the ctx is canceled.
// It is possible to wait until the background metrics push worker is stopped on a WaitGroup passed via opts.WaitGroup.
//
// opts may contain additional configuration options if non-nil.
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
//
// It is OK calling InitPushExtWithOptions multiple times with different pushURL -
// in this case metrics are pushed to all the provided pushURL urls.
//
// It is OK calling InitPushExtWithOptions multiple times with different writeMetrics -
// in this case all the metrics generated by writeMetrics callbacks are written to pushURL.
func InitPushExtWithOptions(ctx context.Context, pushURL string, interval time.Duration, writeMetrics func(w io.Writer), opts *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// validate interval

// PushMetricsExt pushes metrics generated by wirteMetrics to pushURL.
//
// The writeMetrics callback must write metrics to w in Prometheus text exposition format without timestamps and trailing comments.
// See https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#text-based-format
//
// opts may contain additional configuration options if non-nil.
//
// It is recommended pushing metrics to /api/v1/import/prometheus endpoint according to
// https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/#how-to-import-data-in-prometheus-exposition-format
func PushMetricsExt(ctx context.Context, pushURL string, writeMetrics func(w io.Writer), opts *PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type pushContext struct {
	pushURL            *url.URL
	method             string
	pushURLRedacted    string
	extraLabels        string
	headers            http.Header
	disableCompression bool

	client *http.Client

	pushesTotal      *Counter
	bytesPushedTotal *Counter
	pushBlockSize    *Histogram
	pushDuration     *Histogram
	pushErrors       *Counter
}

func newPushContext(pushURL string, opts *PushOptions) (*pushContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate pushURL

// validate ExtraLabels

// validate Headers

func (pc *pushContext) pushMetrics(ctx context.Context, writeMetrics func(w io.Writer)) error {
	_ = "STUB: not implemented"
	return nil
}

// Update metrics

// Prepare the request to sent to pc.pushURL

// Set the needed headers, and `Content-Type` allowed be overwrited.

// Perform the request

var pushMetricsSet = NewSet()

func writePushMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

func addExtraLabels(dst, src []byte, extraLabels string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Skip empy lines

// Copy comments as is

var bashBytes = []byte("#")

func getBytesBuffer() *bytesBuffer { _ = "STUB: not implemented"; return nil }

func putBytesBuffer(bb *bytesBuffer) { _ = "STUB: not implemented"; return }

var bytesBufferPool sync.Pool

type bytesBuffer struct {
	B []byte
}

func (bb *bytesBuffer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func getGzipWriter(w io.Writer) *gzip.Writer { _ = "STUB: not implemented"; return nil }

func putGzipWriter(zw *gzip.Writer) { _ = "STUB: not implemented"; return }

var gzipWriterPool sync.Pool
