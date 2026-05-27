package metrics

import (
	"io"
	runtimemetrics "runtime/metrics"
)

// See https://pkg.go.dev/runtime/metrics#hdr-Supported_metrics
var runtimeMetrics = [][2]string{
	{"/sched/latencies:seconds", "go_sched_latencies_seconds"},
	{"/sync/mutex/wait/total:seconds", "go_mutex_wait_seconds_total"},
	{"/cpu/classes/gc/mark/assist:cpu-seconds", "go_gc_mark_assist_cpu_seconds_total"},
	{"/cpu/classes/gc/total:cpu-seconds", "go_gc_cpu_seconds_total"},
	{"/gc/pauses:seconds", "go_gc_pauses_seconds"},
	{"/cpu/classes/scavenge/total:cpu-seconds", "go_scavenge_cpu_seconds_total"},
	{"/gc/gomemlimit:bytes", "go_memlimit_bytes"},
}

var supportedRuntimeMetrics = initSupportedRuntimeMetrics(runtimeMetrics)

func initSupportedRuntimeMetrics(rms [][2]string) [][2]string {
	_ = "STUB: not implemented"
	return nil
}

func writeGoMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

// Export build details.

func writeRuntimeMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

func writeRuntimeMetric(w io.Writer, name string, sample *runtimemetrics.Sample) {
	_ = "STUB: not implemented"
	return
}

func writeRuntimeHistogramMetric(w io.Writer, name string, h *runtimemetrics.Float64Histogram) {
	_ = "STUB: not implemented"
	return
}

// Limit the maximum bucket to 1 second, since Go runtime exposes buckets with 10K seconds,
// which have little sense. At the same time such buckets may lead to high cardinality issues
// at the scraper side.

// _sum and _count are not exposed because the Go runtime histogram lacks accurate sum data.
// Estimating the sum (as Prometheus does) could be misleading,  while exposing only `_count` without `_sum` is impractical.
// We can reconsider if precise sum data becomes available.
//
// References:
// - Go runtime histogram: https://github.com/golang/go/blob/3432c68467d50ffc622fed230a37cd401d82d4bf/src/runtime/metrics/histogram.go#L8
// - Prometheus estimate: https://github.com/prometheus/client_golang/blob/5fe1d33cea76068edd4ece5f58e52f81d225b13c/prometheus/go_collector_latest.go#L498
// - Related discussion: https://github.com/VictoriaMetrics/metrics/issues/94

// Limit the number of buckets for Go runtime histograms in order to prevent from high cardinality issues at scraper side.
const maxRuntimeHistogramBuckets = 30
