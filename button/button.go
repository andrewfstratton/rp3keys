package button

import (
	"machine"
)

const (
	LEFT = iota
	MIDDLE
	RIGHT
)

type Button struct {
	id      int
	Val     bool
	Changed bool
}

var lookup = []machine.Pin{14, 13, 12}

func New(id int) Button {
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
	result := !pin.Get() // false/low is pressed
	return result
}
