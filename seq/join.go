package seq

import (
	"fmt"
	"iter"
	"strings"
)

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
