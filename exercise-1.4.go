// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, fns := range counts {
		if len(fns) > 1 {
			fmt.Printf("\"%s\" duplicate in:\n", line)
			for fn := range fns {
				fmt.Printf("\tfile:\t%s\n", fn)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, fn string, counts map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		dict, ok := counts[input.Text()]
		if !ok {
			dict = make(map[string]bool)
			counts[input.Text()] = dict
		}
		dict[fn] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
