//go:build !linux && !windows && !solaris && !darwin

package metrics

import (
	"io"
)

func writeProcessMetrics(w io.Writer) {
	_ = "STUB: not implemented"
	// TODO: implement it
	return
}

func writeFDMetrics(w io.Writer) {
	_ = "STUB: not implemented"
	// TODO: implement it.
	return
}
