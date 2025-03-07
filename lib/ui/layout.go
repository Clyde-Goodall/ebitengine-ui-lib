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

type StickyBehavior struct {
	Sticky bool
	Anchor StickyAnchorEnum
}

type Padding struct {
	Left, Top, Right, Bottom, TopLeft, BottomRight int
}

type Anchor struct {
	X, Y int
}

type Element struct {
	Context          *Element
	ElementType      ElementTypeEnum
	AnchorX, AnchorY int
	Width, Height    int
	Padding          Padding
	AutoScale        bool
	Sticky           StickyBehavior
	Children         []Element
	Theme            ElementColorPreset
	visible          bool
	Id               string
}

func (el Element) AddElement(element Element) {
	el.Children = append(el.Children, element)
}

func autoScalingCalculation(el *Element) (float32, float32) {
	ww, wh := ebiten.WindowSize()
	proportionX := ww / el.Width
	proportionY := wh / el.Height
	var actualWidth = proportionX * el.Width
	var actualHeight = proportionY * el.Height
	return float32(actualWidth), float32(actualHeight)
}
