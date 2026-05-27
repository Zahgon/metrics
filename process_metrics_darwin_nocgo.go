//go:build darwin && !ios && !cgo

package metrics

func getMemory() (*memoryInfo, error) { _ = "STUB: not implemented"; return nil, nil }
