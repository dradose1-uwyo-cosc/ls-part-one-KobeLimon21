package functions

import "io"

type color string

const (
	reset color = "\x1b[0m"
	Blue  color = "\x1b[34m"
	Green color = "\x1b[32m"
)

func (c color) ColorPrint(w io.Writer, s string) {
	if c == Blue {
		_, _ = w.Write([]byte(Blue))
		_, _ = w.Write([]byte(s))
		_, _ = w.Write([]byte(reset))
	} else if c == Green {
		_, _ = w.Write([]byte(Green))
		_, _ = w.Write([]byte(s))
		_, _ = w.Write([]byte(reset))
	} else {
		_, _ = w.Write([]byte(s))
	}
}
