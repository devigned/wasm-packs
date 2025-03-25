// Package stringio contains helpers for string I/O.
package stringio

// Writer is the interface implemented by anything that can write strings,
// such as [bytes.Buffer] or [strings.Builder].
type Writer interface {
	WriteString(string) (int, error)
}

// Write writes strings to w, returning the total bytes written and/or an error.
func Write(w Writer, strings ...string) (int, error) {
	var total int
	for _, s := range strings {
		n, err := w.WriteString(s)
		total += n
		if err != nil {
			break
		}
	}
	return total, nil
}
