package main

import (
	"os"
	"strings"

	"github.com/inancgumus/screen"
)

func save() {
	contents := strings.Join(Lines, "\n")

	err := File.Truncate(0)
	throw("Unable to truncate file.", err)

	_, err = File.Seek(0, 0)
	throw("Unable to seek to beginning of file.", err)

	_, err = File.WriteString(contents)
	throw("Unable to write to file.", err)

	IsDirty = false
}

func exit() {
	ShowCursor()

	screen.MoveTopLeft()
	screen.Clear()

	_ = File.Close()

	os.Exit(0)
}
