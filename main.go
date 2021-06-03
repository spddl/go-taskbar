package main

import (
	"log"
	"syscall"
)

func main() {
	szShellTray, err := syscall.UTF16PtrFromString("Shell_TrayWnd")
	if err != nil {
		log.Println(err)
	}
	firstTaskbarHWND := FindWindow(szShellTray, nil) // https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-findwindoww

	szShellSecondaryTray, err := syscall.UTF16PtrFromString("Shell_SecondaryTrayWnd")
	if err != nil {
		log.Println(err)
	}
	secondTaskbarHWND := FindWindow(szShellSecondaryTray, nil)

	if firstTaskbarHWND != zeroHandle {
		result := SetWindowCompositionAttribute(firstTaskbarHWND, FLAG_ACCENT, 1)
		if !result {
			log.Println("could not change the taskbar")
		}
	}
	if secondTaskbarHWND != zeroHandle {
		result := SetWindowCompositionAttribute(secondTaskbarHWND, FLAG_ACCENT, 1)
		if !result {
			log.Println("could not change the second taskbar")
		}
	}
}
