package lib

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
)

type WindowConfig struct {
	Width, Height int
}

type UI struct {
	context      *ebiten.Image
	elementTree  []Element
	windowConfig WindowConfig
}

func InitUI(context *ebiten.Image, config WindowConfig) (UI, error) {
	if config.Width == 0 && config.Height == 0 {
		return UI{}, errors.New("window size cannot have zero or negative values")
	}
	return UI{
		context:      context,
		elementTree:  make([]Element, 0),
		windowConfig: config,
	}, nil
}

func (ui *UI) append(old Element) {
	ui.elementTree = append(ui.elementTree, Element{})
}

func (el *Element) append(child Element) {
	el.children = append(el.children, child)
}

func (ui *UI) getScreenResizeInfo() (int, int) {
	return 0, 0
}

func (ui *UI) Draw(screen *ebiten.Image) {
	if ui.elementTree == nil {
		return
	}
	
}
