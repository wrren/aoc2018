package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Header struct {
	ChildCount    int
	MetaDataCount int
}

func ReadInteger(reader io.Reader) (int, error) {
	buffered := bufio.NewReader(reader)
	intString, err := buffered.ReadString(' ')
	if err != nil && err != io.EOF {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(intString))
}

func ReadHeader(reader io.Reader) (Header, error) {
	header := Header{}

	childCount, err := ReadInteger(reader)
	if err != nil {
		return header, err
	}
	header.ChildCount = childCount

	metaDataCount, err := ReadInteger(reader)
	if err != nil {
		return header, err
	}
	header.MetaDataCount = metaDataCount

	return header, nil
}
