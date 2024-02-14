package main

import (
	"fmt"
	"strings"

	"github.com/inancgumus/screen"
)

func draw(lines []string) {
	screen.MoveTopLeft()

	title := fmt.Sprintf("Pico - %s%s ", FileName, DirtyIndicator())
	footer := fmt.Sprintf("Ln %d, Col %d ", ActiveCursor.Y+1, ActiveCursor.X+1)

	// Draw the header
	fmt.Println(renderRow(-1, "┌ ", title, "─", "─┐"))

	// Draw the body
	for y := 0; y < ActiveCursor.BodyHeight; y++ {
		absoluteY := y + ScrollY

		line := ""

		if absoluteY < len(lines) {
			line = lines[absoluteY]
		}

		fmt.Println(renderRow(y, "│ ", line, " ", " │"))
	}

	// Draw the footer
	fmt.Print(renderRow(-1, "└ ", footer, "─", " ^O Save ─ ^C Exit ─ ^X Save & Exit ┘"))

	screen.MoveTopLeft()
}

func renderRow(y int, prefix, text, padding, suffix string) string {
	var sb strings.Builder

	sb.WriteString(prefix)

	// Text length (with support for unicode)
	runes := []rune(text)
	textLen := len(runes)

	// Insert cursor
	if !ActiveCursor.Disabled && ActiveCursor.Y == y {
		x := ActiveCursor.X

		if textLen == 0 {
			sb.WriteString("█")

			textLen++ // Add an extra space for the cursor
		} else if x == 0 {
			sb.WriteString("█")
			sb.WriteString(string(runes[1:]))
		} else if x == textLen {
			sb.WriteString(text)
			sb.WriteString("█")

			textLen++ // Add an extra space for the cursor
		} else {
			sb.WriteString(string(runes[:x]))
			sb.WriteString("█")
			sb.WriteString(string(runes[x+1:]))
		}
	} else {
		sb.WriteString(text)
	}

	prefixLen := len([]rune(prefix)) - 2
	suffixLen := len([]rune(suffix)) - 2

	paddingLen := ActiveCursor.BodyWidth - textLen - prefixLen - suffixLen

	if paddingLen > 0 {
		sb.WriteString(strings.Repeat(padding, paddingLen))
	}

	sb.WriteString(suffix)

	return sb.String()
}
