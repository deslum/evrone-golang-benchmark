package benchstat_example

import (
	"bytes"
	"strings"
)

// WSCustom WriteString with grow.
func WSCustom(s []string) string {
	var b strings.Builder
	l := len(s)
	b.Grow(l)
	for _, el := range s {
		b.WriteString(el)
	}

	return b.String()
}

// WSDefault WriteString without grow.
func WSDefault(s []string) string {
	var buffer = new(bytes.Buffer)
	for _, el := range s {
		buffer.WriteString(el)
	}

	return buffer.String()
}
