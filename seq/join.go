package seq

import (
	"fmt"
	"iter"
	"strings"
)

// Join concatenates the elements of the input sequence into a single string with a delimiter.
//
// It takes a sequence `input` and a `delimiter` string. It concatenates the elements
// of the sequence into a single string, separated by the specified delimiter.
//
// Play: https://go.dev/play/p/QmlJIBZ2-MW
func Join[V any](input iter.Seq[V], delimiter string) string {
	buf := strings.Builder{}
	first := true

	for item := range input {
		if first {
			first = false
		} else {
			buf.WriteString(delimiter)
		}

		switch val := any(item).(type) {
		case string:
			buf.WriteString(val)
		case fmt.Stringer:
			buf.WriteString(val.String())
		default:
			buf.WriteString(fmt.Sprintf("%v", item))
		}
	}

	return buf.String()
}
