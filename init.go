package main

import (
	"flag"
	"log"
	"os"

	"golang.org/x/sys/windows"
)

type HANDLE uintptr
type HWND HANDLE

const zeroHandle = HWND(0)

var (
	// Library
	libuser32                     *windows.LazyDLL
	setWindowCompositionAttribute *windows.LazyProc
	messageBox                    *windows.LazyProc
	findWindow                    *windows.LazyProc
	findWindowEx                  *windows.LazyProc
	registerWindowMessage         *windows.LazyProc

	FLAG_ACCENT int
)

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile) // https://ispycode.com/GO/Logging/Setting-output-flags

	// Library
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	setWindowCompositionAttribute = libuser32.NewProc("SetWindowCompositionAttribute")
	messageBox = libuser32.NewProc("MessageBoxW")
	findWindow = libuser32.NewProc("FindWindowW")
	findWindowEx = libuser32.NewProc("FindWindowExW")
	registerWindowMessage = libuser32.NewProc("RegisterWindowMessageW")

	flag.IntVar(&FLAG_ACCENT, "accent", 0, `1 Default value. Background is black.
2 Background is GradientColor, alpha channel ignored.
3 Background is GradientColor.
4 Background is GradientColor, with blur effect.
5 Background is GradientColor, with acrylic blur effect.
6 Unknown.
7 Unknown. Seems to draw background fully transparent.`)
	flag.Parse()

	if len(os.Args) < 2 {
		result := MsgBox("Transparent Taskbar", "Change Taskbar style to Transparent?", MB_ICONQUESTION|MB_YESNO)
		switch result {
		case 6:
			FLAG_ACCENT = ACCENT_ENABLE_TRANSPARENT
		default:
			FLAG_ACCENT = ACCENT_ENABLE_GRADIENT
		}
	}
}
