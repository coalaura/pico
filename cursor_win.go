//go:build windows
// +build windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleCursorInfo = kernel32.NewProc("GetConsoleCursorInfo")
	procSetConsoleCursorInfo = kernel32.NewProc("SetConsoleCursorInfo")
)

type dword uint32

type consoleCursorInfo struct {
	size    dword
	visible int32
}

func ShowCursor() {
	handle := syscall.Handle(os.Stdout.Fd())

	var cci consoleCursorInfo
	_, _, _ = procGetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
	cci.visible = 1

	_, _, _ = procSetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
}

func HideCursor() {
	handle := syscall.Handle(os.Stdout.Fd())

	var cci consoleCursorInfo
	_, _, _ = procGetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
	cci.visible = 0

	_, _, _ = procSetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
}
