package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var claims = [...]string{
	"#1 @ 1,3: 4x4",
	"#2 @ 3,1: 4x4",
	"#3 @ 5,5: 2x2",
}

func TestNewClaimFromString(t *testing.T) {
	c1, err := NewClaimFromString(claims[0])
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, c1.ID, 1)
	assert.Equal(t, c1.X, 1)
	assert.Equal(t, c1.Y, 3)
	assert.Equal(t, c1.Width, 4)
	assert.Equal(t, c1.Length, 4)

	c2, err := NewClaimFromString(claims[1])
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, c2.ID, 2)
	assert.Equal(t, c2.X, 3)
	assert.Equal(t, c2.Y, 1)
	assert.Equal(t, c2.Width, 4)
	assert.Equal(t, c2.Length, 4)

	c3, err := NewClaimFromString(claims[2])
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, c3.ID, 3)
	assert.Equal(t, c3.X, 5)
	assert.Equal(t, c3.Y, 5)
	assert.Equal(t, c3.Width, 2)
	assert.Equal(t, c3.Length, 2)
}
