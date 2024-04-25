package main

import (
	"fmt"
	"machine/usb/hid/keyboard"
	"time"

	"rp3keys/buttons"
)

var (
	kbd = keyboard.Port()

// uart = machine.Serial
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
	time.Sleep(time.Second * 1)
	fmt.Println("started")
	// var neo machine.Pin = machine.GPIO18
	// neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// ws := ws2812.New(neo)

	button_left := buttons.Get(buttons.LEFT)
	button_middle := buttons.Get(buttons.MIDDLE)
	button_right := buttons.Get(buttons.RIGHT)

	for {
		// led := color.RGBA{R: 0x00, G: 0x00, B: 0x00} // default to switch off when changed
		button_left.Refresh()
		if button_left.Changed {
			if button_left.Val {
				kbd.Down(keyboard.KeyModifierShift)
				// led = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
			} else {
				kbd.Up(keyboard.KeyModifierShift)
			}
			// leds := []color.RGBA{led, led, led}
			// ws.WriteColors(leds[:])
		}
		button_middle.Refresh()
		if button_middle.Changed {
			if button_middle.Val {
				kbd.Down('a')
				// led = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
			} else {
				kbd.Up('a')
			}
			// leds := []color.RGBA{led, led, led}
			// ws.WriteColors(leds[:])
		}
		button_right.Refresh()
		if button_right.Changed {
			if button_right.Val {
				// led = color.RGBA{R: 0x00, G: 0x00, B: 0xff}
			} else {
			}
			// leds := []color.RGBA{led, led, led}
			// ws.WriteColors(leds[:])
		}
		// TODO Fix  timing/bounce/debounce
		time.Sleep(time.Second / 10)
	}
}
