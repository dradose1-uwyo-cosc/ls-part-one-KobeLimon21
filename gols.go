package main

import (
	"bufio"
	"gols/functions"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush() // flushes the buffer

	useColor := functions.IsTerminal(os.Stdout)
	functions.SimpleLS(w, os.Args[1:], useColor)
}
