//go:build windows

package metrics

import (
	"io"
	"syscall"
)

var (
	modpsapi    = syscall.NewLazyDLL("psapi.dll")
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	// https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getprocessmemoryinfo
	procGetProcessMemoryInfo  = modpsapi.NewProc("GetProcessMemoryInfo")
	procGetProcessHandleCount = modkernel32.NewProc("GetProcessHandleCount")
)

// https://learn.microsoft.com/en-us/windows/win32/api/psapi/ns-psapi-process_memory_counters_ex
type processMemoryCounters struct {
	_                          uint32
	PageFaultCount             uint32
	PeakWorkingSetSize         uintptr
	WorkingSetSize             uintptr
	QuotaPeakPagedPoolUsage    uintptr
	QuotaPagedPoolUsage        uintptr
	QuotaPeakNonPagedPoolUsage uintptr
	QuotaNonPagedPoolUsage     uintptr
	PagefileUsage              uintptr
	PeakPagefileUsage          uintptr
	PrivateUsage               uintptr
}

func writeProcessMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

func writeFDMetrics(w io.Writer) { _ = "STUB: not implemented"; return }

// it seems to be hard-coded limit for 64-bit systems
// https://learn.microsoft.com/en-us/archive/blogs/markrussinovich/pushing-the-limits-of-windows-handles#maximum-number-of-handles
