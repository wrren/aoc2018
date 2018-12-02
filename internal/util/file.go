package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadIntegers(path string) ([]int, error) {
	contents, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer contents.Close()

	scanner := bufio.NewScanner(contents)
	numbers := make([]int, 0)

	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err == nil {
			numbers = append(numbers, number)
		} else {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

func ReadStrings(path string) ([]string, error) {
	contents, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer contents.Close()

	scanner := bufio.NewScanner(contents)
	strings := make([]string, 0)

	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return strings, nil
}
