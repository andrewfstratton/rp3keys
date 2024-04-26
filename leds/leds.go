package leds

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const LED_GPIO_PIN = machine.GPIO18

var led ws2812.Device

var colours [3]color.RGBA
var Off = color.RGBA{R: 0x00, G: 0x00, B: 0x00}

func init() {
	led = ws2812.NewWS2812(LED_GPIO_PIN)
	led.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		time.Sleep(time.Millisecond * 10)
		All(Off)
	}()
}

func Colour(id int, colour color.RGBA) {
	colours[id] = colour
}

func All(colour color.RGBA) {
	for id := range colours {
		Colour(id, colour)
	}
	led.WriteColors(colours[:])
}

func Reset() {
	All(Off)
}
