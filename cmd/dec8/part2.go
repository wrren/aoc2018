package main

import "io"

func part2(reader io.Reader) (int, error) {
	root, err := ReadNode(reader)
	if err != nil {
		return 0, err
	}

	return root.GetValue(), nil
}
