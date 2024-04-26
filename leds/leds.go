package leds

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

var led ws2812.Device

var colours [3]color.RGBA
var Off = color.RGBA{R: 0x00, G: 0x00, B: 0x00}

func init() {
	led_pin := machine.GPIO18
	led_pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// go func() {
	// 	time.Sleep(time.Second / 100)
	led = ws2812.New(led_pin)
	Reset()
	// fmt.Print("Off")
	// }()
}

func setColour(id int, colour color.RGBA) {
	colours[id] = colour
}

func All(colour color.RGBA) {
	for id := range colours {
		setColour(id, colour)
	}
	led.WriteColors(colours[:])
}

func Reset() {
	All(Off)
}
