package editor

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/FotiadisM/zero/internal/buffer"
	"github.com/FotiadisM/zero/internal/display"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	screen tcell.Screen

	// tree
	// buffers
	// display

	// buffers holds oppened documents
	buffers []*buffer.Buffer

	display display.Display
}

func NewEditor(screen tcell.Screen, paths []string) Editor {
	e := Editor{
		screen:  screen,
		buffers: []*buffer.Buffer{},
		display: display.NewDisplay(screen),
	}

	for _, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fi, err := os.Stat(absPath)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			e.buffers = append(e.buffers, buffer.NewBuffer(absPath))
			continue
		}

		if fi.IsDir() {
			// TODO: open zero in the specific dir if passed as arg
			// note: in case of more that 1 dirs passed as args, ignore the rest
			continue
		}

		f, err := os.Open(absPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		e.buffers = append(e.buffers, buffer.NewBufferFromFile(absPath, f))
	}

	return e
}

func (e *Editor) RunEventLoop() {
	for {
		e.display.Draw(e.screen)
		e.screen.Show()

		event := e.screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlQ {
				e.screen.Fini()
				return
			}
			if event.Key() == tcell.KeyDown || event.Rune() == 'j' {
			}
			if event.Key() == tcell.KeyUp || event.Rune() == 'k' {
			}
			if event.Key() == tcell.KeyLeft || event.Rune() == 'h' {
			}
			if event.Key() == tcell.KeyRight || event.Rune() == 'l' {
			}
		}
	}
}
