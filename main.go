package main

import (
	"image/color"
	"machine"
	"machine/usb/hid/keyboard"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var (
	kbd = keyboard.Port()

// uart = machine.Serial
)

const (
	BUTTON_LEFT = iota
	BUTTON_MIDDLE
	BUTTON_RIGHT
)

// func kbd_send(str string) {
// 	for i, runeChar := range str {
// 		substr := fmt.Sprintf("<%d:", runeChar)
// 		kbd.Write([]byte(substr))
// 		kbd.WriteByte(str[i])
// 		kbd.WriteByte('>')
// 	}
// 	kbd.Release()
// }

func init() {
	// uart.Configure(machine.UARTConfig{TX: machine.UART_TX_PIN, RX: machine.UART_RX_PIN})
}

func main() {
	var neo machine.Pin = machine.GPIO18
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)
	var leds [3]color.RGBA

	button_left := NewButton(BUTTON_LEFT)
	button_middle := NewButton(BUTTON_MIDDLE)

	for {
		button_left.Refresh()
		if button_left.Changed {
			if button_left.Val {
				kbd.Down(keyboard.KeyModifierShift)
				leds[0] = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
			} else {
				kbd.Up(keyboard.KeyModifierShift)
				leds[0] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			}
			ws.WriteColors(leds[:])
		}
		button_middle.Refresh()
		if button_middle.Changed {
			if button_middle.Val {
				kbd.Down('a')
				leds[1] = color.RGBA{R: 0xff, G: 0xff, B: 0xff}
			} else {
				kbd.Up('a')
				leds[1] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
			}
			ws.WriteColors(leds[:])
		}
		time.Sleep(time.Second / 10)
	}
	// for {
	// 	// b, _ := uart.ReadByte()
	// 	var leds [1]color.RGBA
	// 	leds[0] = color.RGBA{R: 0xff, G: 0xff, B: 0x00}
	// 	ws.WriteColors(leds[:])
	// 	time.Sleep(time.Second / 4)
	// 	leds[0] = color.RGBA{R: 0x00, G: 0x00, B: 0xff}
	// 	ws.WriteColors(leds[:])
	// 	time.Sleep(time.Second / 4)
	// }
}

type Button struct {
	id      int
	Val     bool
	Changed bool
}

var lookup = []machine.Pin{14, 13, 12}

func NewButton(id int) Button {
	button := Button{id, false, false}
	pin := lookup[id]
	pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	return button
}

func (button *Button) Refresh() {
	button.Changed = false
	val := button.Pressed()
	if val != button.Val {
		button.Val = val
		button.Changed = true
	}
}

func (button *Button) Pressed() bool {
	pin := lookup[button.id]
	result := !pin.Get()
	return result
}
