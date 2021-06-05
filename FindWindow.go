package main

import (
	"syscall"
	"unsafe"
)

func FindWindow(lpClassName, lpWindowName *uint16) HWND {
	ret, _, _ := syscall.Syscall(findWindow.Addr(), 2,
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0)

	return HWND(ret)
}

func FindWindowEx(parentHandle, childAfter HWND, lpClassName, lpWindowName *uint16) HWND {
	ret, _, _ := syscall.Syscall6(findWindowEx.Addr(), 4,
		uintptr(parentHandle),
		uintptr(childAfter),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0, 0)

	return HWND(ret)
}
