package buffer

import (
	"fmt"
	"os"

	"github.com/aretext/aretext/text"
)

type BufferType int

const (
	File BufferType = iota
	Scratch
)

type BufferOptions struct {
	Filetype string
}

type Buffer struct {
	Type BufferType

	Opts BufferOptions
	AbsPath string
	File    *os.File

	TextTree *text.Tree

	Cursors []Cursor

	Dirty bool
}

func NewBuffer(absPath string) *Buffer {
	b := &Buffer{
		AbsPath:  absPath,
		TextTree: text.NewTree(),
	}
	return b
}

func NewBufferFromFile(absPath string, f *os.File) *Buffer {
	tt, err := text.NewTreeFromReader(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	b := &Buffer{
		AbsPath:  absPath,
		File:     f,
		TextTree: tt,
	}
	return b
}
