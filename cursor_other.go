//go:build !windows
// +build !windows

package main

import "fmt"

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

func HideCursor() {
	fmt.Print("\x1b[?25l")
}
