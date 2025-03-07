package lib

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type WindowConfig struct {
	Width, Height int
}

type UI struct {
	Context      *ebiten.Image
	ElementTree  []Element
	WindowConfig WindowConfig
}

func InitUI(context *ebiten.Image, config WindowConfig) (UI, error) {
	if config.Width == 0 && config.Height == 0 {
		return UI{}, errors.New("window size cannot have zero or negative values")
	}
	return UI{
		Context:      context,
		ElementTree:  make([]Element, 0),
		WindowConfig: config,
	}, nil
}

func (ui *UI) AddChild(elType ElementTypeEnum, sticky StickyBehavior, anchor Anchor, sizing ElementSizing, id string) {
	el := Element{ // still need a method of positioning elements based on parent container position
		ElementType: elType,
		AnchorX:     anchor.X,
		AnchorY:     anchor.Y,
		Width:       sizing.Width,
		Height:      sizing.Height,
		// Padding
		// Packed
		Sticky:  sticky,
		visible: true,
		Id:      id,
	}
	fmt.Println("Added element to tree")
	ui.ElementTree = append(ui.ElementTree, el)
}

func (el *Element) AddChild(child Element) {
	el.Children = append(el.Children, child)
}

func (el *Element) RemoveChild(ref int) {

}

func (ui *UI) getScreenResizeInfo() (int, int) {
	ww, wh := ebiten.WindowSize()
	differenceX, differenceY := ui.WindowConfig.Width-ww, ui.WindowConfig.Height-wh
	return differenceX, differenceY
}
