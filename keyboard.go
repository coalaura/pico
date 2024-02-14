package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func readSingleKey() (rune, keyboard.Key) {
	char, key, err := keyboard.GetKey()
	if err != nil {
		fmt.Println("Failed to read key:", err)

		os.Exit(1)
	}

	return char, key
}

func readAndHandleKey() {
	char, key := readSingleKey()

	switch key {
	/* Controls */
	case keyboard.KeyCtrlO:
		save()
	case keyboard.KeyCtrlX:
		save()
		exit()
	case keyboard.KeyBackspace, keyboard.KeyBackspace2:
		deleteChar()
	case keyboard.KeyDelete:
		ActiveCursor.Right() // 100000 IQ move
		deleteChar()

	/* Movement */
	case keyboard.KeyArrowUp:
		ActiveCursor.Up()
	case keyboard.KeyArrowDown:
		ActiveCursor.Down()
	case keyboard.KeyArrowLeft:
		ActiveCursor.Left()
	case keyboard.KeyArrowRight:
		ActiveCursor.Right()
	case keyboard.KeyHome:
		ActiveCursor.MoveStartOfLine()
	case keyboard.KeyEnd:
		ActiveCursor.MoveEndOfLine()

	/* Typing */
	case keyboard.KeyEnter:
		insertNewLine()
	case keyboard.KeySpace:
		insertChars(' ')
	case keyboard.KeyTab:
		// Tab is 4 spaces, change my mind
		insertChars(' ', ' ', ' ', ' ')
	default:
		if char != 0 {
			insertChars(char)
		}
	}
}

func insertNewLine() {
	x := ActiveCursor.AbsoluteX()
	y := ActiveCursor.AbsoluteY()

	currentLine := Lines[y]

	linesBefore := Lines[:y]
	linesAfter := Lines[y+1:]

	// Copy
	newLines := append([]string{}, linesBefore...)

	if x == 0 {
		newLines = append(newLines, "", currentLine)
	} else if x == len(currentLine) {
		newLines = append(newLines, currentLine, "")
	} else {
		before := currentLine[:x]
		after := currentLine[x:]

		newLines = append(newLines, before, after)
	}

	// Paste
	Lines = append(newLines, linesAfter...)

	ActiveCursor.MoveStartOfLine()
	ActiveCursor.Down()

	IsDirty = true
}

func insertChars(chars ...rune) {
	x := ActiveCursor.AbsoluteX()
	y := ActiveCursor.AbsoluteY()

	line := Lines[y]

	before := line[:x]
	after := line[x:]

	Lines[y] = before + string(chars) + after

	ActiveCursor.SetAbsX(ActiveCursor.AbsoluteX() + len(chars))

	IsDirty = true
}

func deleteChar() {
	x := ActiveCursor.AbsoluteX()
	y := ActiveCursor.AbsoluteY()

	if x == 0 && y == 0 {
		return
	}

	current := Lines[y]

	if x == 0 {
		previous := Lines[y-1]

		Lines[y-1] = previous + current

		Lines = append(Lines[:y], Lines[y+1:]...)

		ActiveCursor.Up()
		ActiveCursor.MoveEndOfLine()
	} else {
		before := current[:x-1]
		after := current[x:]

		Lines[y] = before + after

		ActiveCursor.Left()
	}

	IsDirty = true
}
