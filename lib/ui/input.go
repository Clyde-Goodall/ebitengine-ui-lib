package lib

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type ElementSizing struct {
	padding int
	width   int
	height  int
	pack    bool
}

func (ui *UI) Poll() {
	if len(ui.ElementTree) == 0 {
		return
	}
	for _, el := range ui.ElementTree {
		fmt.Printf("Element: %v\n", el)
		ui.inputListener(&el)
	}
}

func (ui *UI) inputListener(el *Element) {

	if len(el.Children) == 0 {
		return
	}
	for _, childEl := range el.Children {
		childEl.processCallback()
	}
}

func (el *Element) processCallback() {

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
