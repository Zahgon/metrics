package metrics

import (
	"io"
	"log"
	"os"
	"time"
)

// See https://github.com/prometheus/procfs/blob/a4ac0826abceb44c40fc71daed2b301db498b93e/proc_stat.go#L40 .
const userHZ = 100

// Different environments may have different page size.
//
// See https://github.com/VictoriaMetrics/VictoriaMetrics/issues/6457
var pageSizeBytes = uint64(os.Getpagesize())

// See http://man7.org/linux/man-pages/man5/proc.5.html
type procStat struct {
	State       byte
	Ppid        int
	Pgrp        int
	Session     int
	TtyNr       int
	Tpgid       int
	Flags       uint
	Minflt      uint
	Cminflt     uint
	Majflt      uint
	Cmajflt     uint
	Utime       uint
	Stime       uint
	Cutime      int
	Cstime      int
	Priority    int
	Nice        int
	NumThreads  int
	ItrealValue int
	Starttime   uint64
	Vsize       uint
	Rss         int
}

func writeProcessMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

// Search for the end of command.

// It is expensive obtaining `process_open_fds` when big number of file descriptors is opened,
// so don't do it here.
// See writeFDMetrics instead.

// Calculate totalTime by dividing the sum of p.Utime and p.Stime by userHZ.
// This reduces possible floating-point precision loss

var procSelfIOErrLogged uint32

func writeIOMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

// Do not spam the logs with errors - this error cannot be fixed without process restart.
// See https://github.com/VictoriaMetrics/metrics/issues/42

var startTimeSeconds = time.Now().Unix()

// writeFDMetrics writes process_max_fds and process_open_fds metrics to w.
func writeFDMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

func getOpenFDsCount(path string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func getMaxFilesLimit(path string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Extract soft limit.

// https://man7.org/linux/man-pages/man5/procfs.5.html
type memStats struct {
	vmPeak   uint64
	rssPeak  uint64
	rssAnon  uint64
	rssFile  uint64
	rssShmem uint64
}

func writeProcessMemMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

func getMemStats(path string) (*memStats, error) { _ = "STUB: not implemented"; return nil, nil }

// Extract key value.

// writePSIMetrics writes PSI total metrics for the current process to w.
//
// See https://docs.kernel.org/accounting/psi.html
func writePSIMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

// Failed to initialize PSI metrics

func psiTotalSecs(microsecs uint64) float64 {
	_ = "STUB: not implemented"
	// PSI total stats is in microseconds according to https://docs.kernel.org/accounting/psi.html
	// Convert it to seconds.
	return 0
}

// psiMetricsStart contains the initial PSI metric values on program start.
// it is needed in order to make sure the exposed PSI metrics start from zero.
var psiMetricsStart = func() *psiMetrics {
	m, err := getPSIMetrics()
	if err != nil {
		log.Printf("INFO: metrics: disable exposing PSI metrics because of failed init: %s", err)
		return nil
	}
	return m
}()

type psiMetrics struct {
	cpuSome uint64
	cpuFull uint64
	ioSome  uint64
	ioFull  uint64
	memSome uint64
	memFull uint64
}

func getPSIMetrics() (*psiMetrics, error) { _ = "STUB: not implemented"; return nil, nil }

// Do nothing, since PSI requires cgroup v2, and the process doesn't run under cgroup v2.

func readPSITotals(cgroupPath, statsName string) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getCgroupV2Path() string { _ = "STUB: not implemented"; return "" }

// Drop trailing slash if it exsits. This prevents from '//' in the constructed paths by the caller.
