package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	screen             *ebiten.Image
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

func (el Element) Draw() {
	// Should first draw the element type as described in el and then
	// recursively draw all children of this and subsequent elements.
	// On its face, this seems like significantly DX to implement than in zig but we'll see
	switch el.elementType {
	case Container:
		el.DrawContainer()
		break
	case DecoratedContainer:
		break
	case Button:
		break
	case RoundedButton:
		break
	case Text:
		break
	default:
		panic("Invalid element type")
	}
}

func (el Element) AddElement(element Element) {
	el.children = append(el.children, element)
}

// Will determine whether element will scale with window size proportionally
// If sticky, disregards scaling bool so it doesn't conflict with resizing
func (el Element) DrawContainer() {
	actualWidth := float32(el.width)   // anchorX
	actualHeight := float32(el.height) // anchorY
	if el.autoScale && !el.sticky {
		actualWidth, actualHeight = autoScalingCalculation(el)
	}

	if el.sticky { // no need to check for autoScale, as it should override autoScale
		switch el.stickyAnchor { // anchor anchorX, anchorY, width, height to whatever corner it's sticking to.
		case TopLeft: // technically, top left wouldn't matter since x,y should already be relative to the window's position
			break
		case TopRight: // corresponds to x expansion

			break
		case BottomLeft: // corresponds to y expansion

			break
		case BottomRight: // corresponds to both x and y expansion

			break
		}
	}

	vector.DrawFilledRect(
		el.screen,
		float32(el.anchorX),
		float32(el.anchorY),
		actualWidth,
		actualHeight,
		el.theme.backgroundColor,
		true,
	)
	if el.paddingX > 0 { // inset whole

	}
	vector.DrawFilledRect(
		el.screen,
		float32(el.anchorX),
		float32(el.anchorY),
		actualWidth,
		actualHeight,
		el.theme.backgroundColor,
		true,
	)
}

func autoScalingCalculation(el Element) (float32, float32) {
	ww, wh := ebiten.WindowSize()
	proportionX := ww / el.width
	proportionY := wh / el.height
	var actualWidth = proportionX * el.width
	var actualHeight = proportionY * el.height
	return float32(actualWidth), float32(actualHeight)
}
