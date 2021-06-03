package main

import (
	"log"
	"strings"
	"syscall"
	"unsafe"
)

const (
	MB_YESNO        = 0x00000004
	MB_ICONQUESTION = 0x00000020
)

func MsgBox(title, message string, style uintptr) int {
	titleInt, err := syscall.UTF16PtrFromString(strings.ReplaceAll(title, "\x00", "␀"))
	if err != nil {
		log.Println(err)
	}

	messageInt, err := syscall.UTF16PtrFromString(strings.ReplaceAll(message, "\x00", "␀"))
	if err != nil {
		log.Println(err)
	}

	ret, _, _ := syscall.Syscall6(messageBox.Addr(), 4,
		uintptr(0),
		uintptr(unsafe.Pointer(messageInt)),
		uintptr(unsafe.Pointer(titleInt)),
		uintptr(style),
		0,
		0)

	return int(int32(ret))
}
