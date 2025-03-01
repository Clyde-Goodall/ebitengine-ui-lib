package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ElementSizing struct {
	padding int
	width   int
	height  int
	pack    bool
}

func (el Element) RoundedButton(
	screen *ebiten.Image,
	label string,
	callback func(),
	sizing ElementSizing,
) Element {
	// TODO
	return Element{}
}
