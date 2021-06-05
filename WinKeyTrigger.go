package main

import hook "github.com/robotn/gohook"

func WinKeyTrigger() {
	EvChan := hook.Start()
	defer hook.End()
	for ev := range EvChan {
		if ev.Rawcode == 92 && ev.Kind == hook.KeyUp {
			changeTaskbar()
		}
	}
}
