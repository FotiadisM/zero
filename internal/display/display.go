package display

import (
	"github.com/FotiadisM/zero/internal/display/split"

	"github.com/gdamore/tcell/v2"
)

type Display struct {
	screen    tcell.Screen
	splitTree split.Node
}

func NewDisplay(screen tcell.Screen) Display {
	w, h := screen.Size()
	d := Display{
		screen:    screen,
		splitTree: *split.NewRootNode(0, 0, w, h),
	}

	return d
}

func (d *Display) Draw(screen tcell.Screen) {
}
