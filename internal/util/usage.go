package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func Usage() {
	fmt.Println("usage:")
	fmt.Printf("\t%s <part1|part2> <input_file>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}
