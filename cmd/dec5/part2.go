package main

import "io"

func part2(reader io.Reader) (int, error) {
	polymer := NewPolymer()

	err := polymer.ReadAndReact(reader)
	if err != nil {
		return 0, err
	}

	return polymer.MinLength(), nil
}
