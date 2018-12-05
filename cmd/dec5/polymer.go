package main

import (
	"bufio"
	"io"
	"math"
	"unicode"
)

type Buffer []rune

type Polymer struct {
	Buffer          Buffer
	FilteredBuffers map[rune]Buffer
}

func NewPolymer() Polymer {
	return Polymer{Buffer: make(Buffer, 0), FilteredBuffers: map[rune]Buffer{}}
}

func PairReacts(a, b rune) bool {
	return (int(a) == (int(b) - 32)) || (int(b) == (int(a) - 32))
}

func (p Polymer) Length() int {
	return len(p.Buffer)
}

func (p Polymer) MinLength() int {
	min := math.MaxInt32
	for _, v := range p.FilteredBuffers {
		if len(v) < min {
			min = len(v)
		}
	}
	return min
}

func Combine(buffer Buffer, r rune) Buffer {
	if len(buffer) > 0 {
		last := buffer[len(buffer)-1]
		if PairReacts(last, r) {
			buffer = buffer[:len(buffer)-1]
		} else {
			buffer = append(buffer, r)
		}
	} else {
		buffer = append(buffer, r)
	}
	return buffer
}

func (p *Polymer) Chain(r rune) {
	if int(r) < 65 || int(r) > 122 {
		return
	}
	p.Buffer = Combine(p.Buffer, r)
	key := unicode.ToLower(r)

	_, ok := p.FilteredBuffers[key]
	if ok == false {
		p.FilteredBuffers[key] = make([]rune, 0)
	}

	for k, v := range p.FilteredBuffers {
		if k != key {
			p.FilteredBuffers[k] = Combine(v, r)
		}
	}
}

func (p *Polymer) ReadAndReact(reader io.Reader) error {
	bufferedReader := bufio.NewReader(reader)
	for {
		r, _, err := bufferedReader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		p.Chain(r)
	}
}
