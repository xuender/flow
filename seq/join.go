package seq

import (
	"fmt"
	"iter"
	"strings"
)

// Join concatenates the elements of the input sequence into a single string with a delimiter.
//
// This function takes a sequence `input` and a `delimiter` string. It concatenates the elements
// of the sequence into a single string, separated by the specified delimiter.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements.
//	delimiter string: The string to insert between elements.
//
// Returns:
//
//	string: The concatenated string with delimiters.
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
