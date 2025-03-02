package lib

// Strictly draw functions that belong to UI and its child element

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (ui *UI) Draw(screen *ebiten.Image) error {
	if ui.context == nil {
		ui.context = screen // assign screen to context, which will generally occur in ebiten's Draw loop
	}
	if ui.elementTree == nil {
		return errors.New("no elements in elementTree to render")
	}
	for _, el := range ui.elementTree {
		ui.DrawElement(el)
	}
	return nil
}

func (ui *UI) DrawElement(el Element) {
	switch el.elementType { // some of these enums could be consolidated with extra params
	case Container:
		ui.DrawContainer(el)
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
	for _, el := range ui.elementTree {
		ui.DrawElement(el)
	}
}

// Will determine whether element will scale with window size proportionally
// If sticky, disregards scaling bool so it doesn't conflict with resizing
func (ui UI) DrawContainer(el Element) {
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
		ui.context,
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
		ui.context,
		float32(el.anchorX),
		float32(el.anchorY),
		actualWidth,
		actualHeight,
		el.theme.backgroundColor,
		true,
	)
}
