package main

import (
	"fmt"
	"time"

	"rp3keys/board"
	"rp3keys/buttons"
	"rp3keys/leds"
)

func main() {
	leds.Reset()
	time.Sleep(time.Second * 1)
	fmt.Println("started")

	left := buttons.Get(board.LEFT)
	middle := buttons.Get(board.MIDDLE)
	right := buttons.Get(board.RIGHT)

	for {
		buttons.Refresh()
		if left.Changed {
			led := leds.Off
			if left.Val {
				board.TypeMod(board.TAB, board.ALT)
				led.R = 0xff
			}
			leds.Colour(board.LEFT, led)
		}
		if middle.Changed {
			led := leds.Off
			if middle.Val {
				led.G = 0xff
				board.TypeMod(board.TAB, board.CTRL)
			}
			leds.Colour(board.MIDDLE, led)
		}
		if right.Changed {
			led := leds.Off
			if right.Val {
				board.TypeMod(board.TAB, board.SHIFT)
				led.B = 0xff
			}
			leds.Colour(board.RIGHT, led)
		}
		// TODO Fix  timing/bounce/debounce
		time.Sleep(time.Second / 25)
	}
}
