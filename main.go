package main

import (
	"log"
	"syscall"
)

func main() {
	go WinKeyTrigger()
	changeTaskbar()
	select {}
}

func changeTaskbar() {
	szShellTray, err := syscall.UTF16PtrFromString("Shell_TrayWnd")
	if err != nil {
		log.Println(err)
	}
	firstTaskbarHWND := FindWindow(szShellTray, nil) // https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww
	if !SetWindowCompositionAttribute(firstTaskbarHWND, FLAG_ACCENT, 1) {
		log.Println("could not change the taskbar")
	}

	szShellSecondaryTray, err := syscall.UTF16PtrFromString("Shell_SecondaryTrayWnd")
	if err != nil {
		log.Println(err)
	}

	var otherBars = HWND(0)
	for {
		otherBars = FindWindowEx(0, otherBars, szShellSecondaryTray, nil)
		if otherBars == zeroHandle {
			break
		}
		if !SetWindowCompositionAttribute(otherBars, FLAG_ACCENT, 1) {
			log.Println("could not change the other taskbar")
		}
	}
}
