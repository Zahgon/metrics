package metrics

import (
	"regexp"
)

// ValidateMetric validates provided string
// to be valid Prometheus-compatible metric with possible labels.
// For instance,
//
//   - foo
//   - foo{bar="baz"}
//   - foo{bar="baz",aaa="b"}
func ValidateMetric(s string) error { _ = "STUB: not implemented"; return nil }

func validateTags(s string) error { _ = "STUB: not implemented"; return nil }

func skipSpace(s string) string { _ = "STUB: not implemented"; return "" }

func validateIdent(s string) error { _ = "STUB: not implemented"; return nil }

var identRegexp = regexp.MustCompile("^[a-zA-Z_:.][a-zA-Z0-9_:.]*$")
