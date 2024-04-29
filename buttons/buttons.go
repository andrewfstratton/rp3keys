package buttons

import (
	"machine"
)

type Button struct {
	Val     bool
	Changed bool
	pin     machine.Pin
}

var buttons = []*Button{}

func init() {
	pins := []machine.Pin{14, 13, 12} // pins are in reverse order
	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
		button := Button{pin: pin}
		buttons = append(buttons, &button)
	}
}

func Get(id int) *Button {
	return buttons[id]
}

func Refresh() {
	for id := range buttons {
		button := buttons[id]
		button.Changed = false
		val := button.Pressed()
		if val != button.Val {
			button.Val = val
			button.Changed = true
		}
	}
}

func (button *Button) Pressed() bool {
	result := !button.pin.Get() // pressed is false/low
	return result
}
