package leds

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const ws2812Pin = machine.GPIO18

var led ws2812.Device

var colours [3]color.RGBA
var Off = color.RGBA{R: 0x00, G: 0x00, B: 0x00}

func init() {
	led = ws2812.NewWS2812(ws2812Pin)
	led.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		time.Sleep(time.Millisecond * 10)
		All(Off)
	}()
}

func Colour(id int, colour color.RGBA) {
	colours[id] = colour
	led.WriteColors(colours[:])
}

func All(colour color.RGBA) {
	for id := range colours {
		colours[id] = colour
	}
	led.WriteColors(colours[:])
}

func Reset() {
	All(Off)
}
