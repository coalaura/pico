package main

import (
	"fmt"
	"strings"

	"github.com/inancgumus/screen"
)

func draw(lines []string) {
	screen.MoveTopLeft()

	title := fmt.Sprintf("Pico - %s%s ", FileName, DirtyIndicator())
	footer := fmt.Sprintf("Ln %d, Col %d ", ActiveCursor.AbsoluteY()+1, ActiveCursor.AbsoluteX()+1)

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
	textLen := len([]rune(text))

	// Append padding
	prefixLen := len([]rune(prefix)) - 2
	suffixLen := len([]rune(suffix)) - 2

	paddingLen := ActiveCursor.BodyWidth - textLen - prefixLen - suffixLen

	if paddingLen > 0 {
		text += strings.Repeat(padding, paddingLen)
	}

	// -1 is footer and header (non-content) rows
	if y != -1 {
		// Clamp the text length and scroll horizontally
		if textLen > 0 {
			from := ScrollX
			to := ScrollX + ActiveCursor.BodyWidth

			for textLen+1 < to {
				text += " "
				textLen++
			}

			text = string([]rune(text)[from:to])
		}

		// Insert cursor
		if !ActiveCursor.Disabled && ActiveCursor.Y == y {
			char := ActiveCursor.CurrentChar()

			if ActiveCursor.X == ActiveCursor.BodyWidth {
				text = text[1:]
			}

			text = ReplaceCharacterAt(text, InvertColors(char), ActiveCursor.X)
		}
	}

	sb.WriteString(text)

	sb.WriteString(suffix)

	return sb.String()
}
