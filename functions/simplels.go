package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func SimpleLS(w io.Writer, args []string, useColor bool) { // if no parameters, list the directory
	if len(args) == 0 {
		listDir(w, ".", useColor)
		return
	}

	files, dirs := splitTargets(args)

	sort.Strings(files) // sorts files alphabetically
	sort.Strings(dirs)  // sorts directories alphabetically

	for _, f := range files { // iterates over files and prints them
		printFileTarget(w, f, useColor)
	}

	if len(dirs) == 0 { // no more directories to list
		return
	}

	showHeaders := len(dirs) > 1 // only prints directory headers if there are multiple
	for i, d := range dirs {     // iterates over the dirs
		if showHeaders {
			fmt.Fprintf(w, "%s:\n", d)
		}
		listDir(w, d, useColor) // lists the  contents
		if i != len(dirs)-1 {
			fmt.Fprintln(w)
		}
	}
}
func splitTargets(args []string) (files []string, dirs []string) { // seperates files and directories into two slices
	files = make([]string, 0, len(args))
	dirs = make([]string, 0, len(args))

	for _, p := range args { // checks each arg to see if it's a file or directory and adds it to the appropriate slice
		info, err := os.Lstat(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: %s: %v\n", p, err)
			continue
		}
		if info.IsDir() {
			dirs = append(dirs, p) // adds to dir
		} else {
			files = append(files, p) //adds to file
		}
	}

	return files, dirs
}

func printFileTarget(w io.Writer, path string, useColor bool) { // prints a file target, colors it if it's executable and useColor is true
	info, err := os.Lstat(path) // grabs file info for path
	if err != nil {
		fmt.Fprintf(os.Stderr, "gols: %s: %v\n", path, err)
		return
	}

	name := filepath.Base(path)

	if useColor {
		mode := info.Mode()
		if mode.IsRegular() && (mode&0111) != 0 { // checks if its regular and execute permissions are set
			Green.ColorPrint(w, name) // reg file so print green
			fmt.Fprintln(w)
			return
		}
	}

	fmt.Fprintln(w, name) // print names if not colored
}

func listDir(w io.Writer, dir string, useColor bool) { // lists contents of a directory
	entries, err := os.ReadDir(dir) // reads the dir
	if err != nil {
		fmt.Fprintf(os.Stderr, "gols: %s: %v\n", dir, err)
		return
	}

	entries = dirFilter(entries) // filters out hidden entries

	sort.Slice(entries, func(i, j int) bool { // sorts entries alphabetically by name
		return entries[i].Name() < entries[j].Name()
	})

	for _, e := range entries { // loops through entries
		name := e.Name()
		full := filepath.Join(dir, name) // joins dir and name for the path

		info, err := os.Lstat(full)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: %s: %v\n", full, err)
			continue
		}

		if useColor {
			mode := info.Mode()
			if info.IsDir() {
				Blue.ColorPrint(w, name) // directory so print blue
				fmt.Fprintln(w)
				continue
			}
			if mode.IsRegular() && (mode&0111) != 0 {
				Green.ColorPrint(w, name) // regular file with execute permissions so print green
				fmt.Fprintln(w)
				continue
			}
		}

		fmt.Fprintln(w, name) // print name if not colored
	}
}
