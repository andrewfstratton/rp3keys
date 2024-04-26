package main

import (
	"machine/usb/hid/keyboard"
	"time"

	"rp3keys/buttons"
	"rp3keys/leds"
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
	leds.Reset()
	// time.Sleep(time.Second * 1)
	// fmt.Println("started")

	left := buttons.Get(buttons.LEFT)
	middle := buttons.Get(buttons.MIDDLE)
	right := buttons.Get(buttons.RIGHT)

	for {
		buttons.Refresh()
		if left.Changed {
			led := leds.Off
			if left.Val {
				kbd.Down(keyboard.KeyModifierShift)
				led.R = 0xff
			} else {
				kbd.Up(keyboard.KeyModifierShift)
			}
			leds.All(led)
		}
		if middle.Changed {
			led := leds.Off
			if middle.Val {
				kbd.Down('a')
				led.G = 0xff
			} else {
				kbd.Up('a')
			}
			leds.All(led)
		}
		if right.Changed {
			led := leds.Off
			if right.Val {
				led.B = 0xff
			} else {
			}
			leds.All(led)
		}
		// TODO Fix  timing/bounce/debounce
		time.Sleep(time.Second / 10)
	}
}
