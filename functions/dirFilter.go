package functions

import "os"

func dirFilter(entries []os.DirEntry) []os.DirEntry {
	out := make([]os.DirEntry, 0)

	for i := 0; i < len(entries); i++ {
		n := entries[i].Name()
		if len(n) > 0 && n[0] != '.' { // filters out the hidden files
			out = append(out, entries[i])
		}
	}

	return out
}
