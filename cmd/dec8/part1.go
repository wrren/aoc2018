package main

import (
	"io"
)

func part1(reader io.Reader) (int, error) {
	root, err := ReadNode(reader)
	if err != nil {
		return 0, err
	}

	return root.SumMetaData(), nil
}
