package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	actual := count(b, false, false)
	assert.Equal(t, 4, actual)
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	actual := count(b, false, true)
	assert.Equal(t, 3, actual)
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("Hello, 🌎")
	actual := count(b, true, false)
	assert.Equal(t, 11, actual)
}
