package seq

import (
	"fmt"
	"iter"
	"strings"
)

// Join function concatenates the elements of the given sequence 'input' into a single string,
// separated by the specified 'delimiter'.
// This function works with elements of any type E, converting non-string types to strings using fmt.Sprintf.
//
// Parameters:
//
//	input: The sequence of elements to join, of type iter.Seq[E].
//	delimiter: The string to insert between each element.
//
// Returns:
//
//	A single string containing all the elements of the sequence, separated by the delimiter.
func Join[E any](input iter.Seq[E], delimiter string) string {
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
