//Danny Radosevich

//Re-writing ls command in Go
//check if writing to terminal

package functions

import "os"

func IsTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}
