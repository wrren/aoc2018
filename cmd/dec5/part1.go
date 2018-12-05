package main

import "io"

func part1(reader io.Reader) (int, error) {
	polymer := NewPolymer()

	err := polymer.ReadAndReact(reader)
	if err != nil {
		return 0, err
	}

	return polymer.Length(), nil
}
