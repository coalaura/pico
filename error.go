package main

import (
	"fmt"

	"github.com/inancgumus/screen"
)

func throw(msg string, err error) {
	if err == nil {
		return
	}

	screen.Clear()

	fmt.Println(msg)

	panic(err)
}
