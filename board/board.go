package board

import (
	"machine/usb/hid/keyboard"
)

const (
	LEFT = iota
	MIDDLE
	RIGHT
)

const (
	CTRL = 1 << iota
	SHIFT
	ALT
	// WIN
	NONE = 0
)

const (
	COPY  keyboard.Keycode = 'c'
	CUT   keyboard.Keycode = 'x'
	PASTE keyboard.Keycode = 'v'
	TAB                    = keyboard.KeyTab
)

var kbd = keyboard.Port()

func Type(code keyboard.Keycode) {
	TypeMod(code, NONE)
}

func TypeMod(code keyboard.Keycode, modifier int) {
	if (modifier & CTRL) != 0 {
		kbd.Down(keyboard.KeyLeftCtrl)
		defer kbd.Up(keyboard.KeyLeftCtrl)
	}
	if (modifier & SHIFT) != 0 {
		kbd.Down(keyboard.KeyLeftShift)
		defer kbd.Up(keyboard.KeyLeftShift)
	}
	if (modifier & ALT) != 0 {
		kbd.Down(keyboard.KeyLeftAlt)
		defer kbd.Up(keyboard.KeyLeftAlt)
	}
	kbd.Press(code)
}
