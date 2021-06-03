package main

import (
	"syscall"
	"unsafe"
)

const (
	ACCENT_DISABLED                   = iota // [1] Default value. Background is black.
	ACCENT_ENABLE_GRADIENT            = iota // [2] Background is GradientColor, alpha channel ignored.
	ACCENT_ENABLE_TRANSPARENTGRADIENT = iota // [3] Background is GradientColor.
	ACCENT_ENABLE_BLURBEHIND          = iota // [4] Background is GradientColor, with blur effect.
	ACCENT_ENABLE_ACRYLICBLURBEHIND   = iota // [5] Background is GradientColor, with acrylic blur effect.
	ACCENT_INVALID_STATE              = iota // [6] Unknown.
	ACCENT_ENABLE_TRANSPARENT         = iota // [7] Unknown. Seems to draw background fully transparent.
)

type ACCENT_POLICY struct {
	ACCENT_STATE  int
	AccentFlags   uint32
	GradientColor uint32
	AnimationId   uint32
}

const (
	WCA_UNDEFINED                     = iota
	WCA_NCRENDERING_ENABLED           = iota
	WCA_NCRENDERING_POLICY            = iota
	WCA_TRANSITIONS_FORCEDISABLED     = iota
	WCA_ALLOW_NCPAINT                 = iota
	WCA_CAPTION_BUTTON_BOUNDS         = iota
	WCA_NONCLIENT_RTL_LAYOUT          = iota
	WCA_FORCE_ICONIC_REPRESENTATION   = iota
	WCA_EXTENDED_FRAME_BOUNDS         = iota
	WCA_HAS_ICONIC_BITMAP             = iota
	WCA_THEME_ATTRIBUTES              = iota
	WCA_NCRENDERING_EXILED            = iota
	WCA_NCADORNMENTINFO               = iota
	WCA_EXCLUDED_FROM_LIVEPREVIEW     = iota
	WCA_VIDEO_OVERLAY_ACTIVE          = iota
	WCA_FORCE_ACTIVEWINDOW_APPEARANCE = iota
	WCA_DISALLOW_PEEK                 = iota
	WCA_CLOAK                         = iota
	WCA_CLOAKED                       = iota
	WCA_ACCENT_POLICY                 = iota
	WCA_FREEZE_REPRESENTATION         = iota
	WCA_EVER_UNCLOAKED                = iota
	WCA_VISUAL_OWNER                  = iota
	WCA_LAST                          = iota
)

type WINCOMPATTRDATA struct {
	nAttribute int
	pData      *ACCENT_POLICY
	ulDataSize uint32
}

func SetWindowCompositionAttribute(hwnd HWND, accent_state int, dword uint32) bool {
	var accent = ACCENT_POLICY{}
	accent.ACCENT_STATE = accent_state
	data := WINCOMPATTRDATA{
		nAttribute: WCA_ACCENT_POLICY,
		pData:      &accent,
	}
	data.ulDataSize = uint32(unsafe.Sizeof(data))

	ret, _, _ := syscall.Syscall(setWindowCompositionAttribute.Addr(), 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&data)),
		0)
	return ret != 0
}
