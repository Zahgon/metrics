//go:build darwin && !ios && cgo

package metrics

/*
int vm_get_memory_info(unsigned long long *rss, unsigned long long *vs);
*/
import "C"

func getMemory() (*memoryInfo, error) { _ = "STUB: not implemented"; return nil, nil }
