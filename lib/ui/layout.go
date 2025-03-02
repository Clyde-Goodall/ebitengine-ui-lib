package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ElementTypeEnum int

const (
	Container = iota
	DecoratedContainer
	Button
	RoundedButton
	Text
)

func (e ElementTypeEnum) String() string {
	return [...]string{
		"container",
		"decorated_container",
		"button",
		"rounded_button",
		"text",
	}[e]
}

type StickyAnchorEnum int

const (
	TopLeft = iota
	TopRight
	BottomLeft
	BottomRight
)

func (e StickyAnchorEnum) String() string {
	return [...]string{
		"top_left",
		"top_right",
		"bottom_left",
		"bottom_right",
	}[e]
}

type Element struct {
	context            *Element
	elementType        ElementTypeEnum
	anchorX, anchorY   int
	width, height      int
	paddingX, paddingY int
	autoScale          bool
	sticky             bool
	children           []Element
	theme              ElementColorPreset
	stickyAnchor       StickyAnchorEnum
}

func (el Element) AddElement(element Element) {
	el.children = append(el.children, element)
}

func autoScalingCalculation(el Element) (float32, float32) {
	ww, wh := ebiten.WindowSize()
	proportionX := ww / el.width
	proportionY := wh / el.height
	var actualWidth = proportionX * el.width
	var actualHeight = proportionY * el.height
	return float32(actualWidth), float32(actualHeight)
}
