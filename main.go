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

	left := buttons.Get(buttons.LEFT)
	middle := buttons.Get(buttons.MIDDLE)
	right := buttons.Get(buttons.RIGHT)

	for {
		// led := color.RGBA{R: 0x00, G: 0x00, B: 0x00} // default to switch off when changed
		left.Refresh()
		if left.Changed {
			if left.Val {
				kbd.Down(keyboard.KeyModifierShift)
				// led = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
			} else {
				kbd.Up(keyboard.KeyModifierShift)
			}
			// leds := []color.RGBA{led, led, led}
			// ws.WriteColors(leds[:])
		}
		middle.Refresh()
		if middle.Changed {
			if middle.Val {
				kbd.Down('a')
				// led = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
			} else {
				kbd.Up('a')
			}
			// leds := []color.RGBA{led, led, led}
			// ws.WriteColors(leds[:])
		}
		right.Refresh()
		if right.Changed {
			if right.Val {
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
