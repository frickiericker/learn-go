package main

import (
	"syscall"
	"unsafe"
)

func setTerminalSize(handle uintptr, rows, columns int) error {
	type winsizeT struct {
		Row, Col, Xpixel, Ypixel uint16
	}
	winsize := winsizeT{
		Row: uint16(rows),
		Col: uint16(columns),
	}
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		handle,
		syscall.TIOCSWINSZ,
		uintptr(unsafe.Pointer(&winsize)))
	if errno != 0 {
		return errno
	}
	return nil
}
