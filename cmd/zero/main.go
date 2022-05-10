package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FotiadisM/zero/internal/editor"

	"github.com/gdamore/tcell/v2"
)

func main() {
	flag.Parse()

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e := editor.NewEditor(screen, flag.Args())
	e.RunEventLoop()
}
