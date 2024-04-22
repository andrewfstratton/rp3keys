package main

import (
	"machine/usb/hid/keyboard"
	"time"
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
	// var neo machine.Pin = machine.NEOPIXEL
	// neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// ws := ws2812.New(neo)

	for {
		kbd.Press(keyboard.KeyModifierShift)
		time.Sleep(time.Second / 2)
		kbd.Release()
		time.Sleep(time.Second / 2)
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
