//go:build darwin && !ios

package metrics

import (
	"errors"
	"io"
)

// errNotImplemented is returned by stub functions that replace cgo functions, when cgo
// isn't available.
var errNotImplemented = errors.New("not implemented")

func writeProcessMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

// The proc structure returned by kern.proc.pid above has an Rusage member,
// but it is not filled in, so it needs to be fetched by getrusage(2).  For
// that call, the UTime, STime, and Maxrss members are filled out, but not
// Ixrss, Idrss, or Isrss for the memory usage.  Memory stats will require
// access to the C API to call task_info(TASK_BASIC_INFO).

func writeFDMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

func getOpenFileCount() (float64, error) {
	_ = "STUB: not implemented"
	// Alternately, the undocumented proc_pidinfo(PROC_PIDLISTFDS) can be used to
	// return a list of open fds, but that requires a way to call C APIs.  The
	// benefits, however, include fewer system calls and not failing when at the
	// open file soft limit.
	return 0, nil
}

// Avoid ReadDir(), as it calls stat(2) on each descriptor.  Not only is
// that info not used, but KQUEUE descriptors fail stat(2), which causes
// the whole method to fail.

// Subtract 1 to ignore the open /dev/fd descriptor above.

func getSoftLimit(which int) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func getProcessStartTime() (float64, error) {
	_ = "STUB: not implemented"
	// Call sysctl to get kinfo_proc for current process
	return 0, nil
}

/* CTL_KERN */ /* KERN_PROC */ /* KERN_PROC_PID */

// First call to get the size

// Second call to get the actual data

// The kinfo_proc struct layout on Darwin has p_starttime (struct timeval) at specific offset
// For amd64 and arm64, the offset is at 0x60 (96 bytes)
// struct timeval has tv_sec (int64) and tv_usec (int32)

// Read tv_sec (8 bytes) and tv_usec (4 bytes)

type memoryInfo struct {
	vsize uint64 // Virtual memory size in bytes
	rss   uint64 // Resident memory size in bytes
}
