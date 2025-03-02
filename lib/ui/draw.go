package lib

// Strictly draw functions that belong to UI and its child element

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (ui *UI) DrawUI(screen *ebiten.Image) error {
	if ui.Context == nil {
		ui.Context = screen // assign screen to context, which will generally occur in ebiten's Draw loop
	}
	// if len(ui.ElementTree) == 0 {
	// 	fmt.Println("No element tree")
	// 	return errors.New("no elements in elementTree to render")
	// }
	for _, el := range ui.ElementTree {
		fmt.Printf("Element: %v\n", el)
		ui.DrawElement(&el)
	}
	// fmt.Println("fuck fuck fuck")
	return nil
}

func (ui *UI) DrawElement(el *Element) {
	fmt.Printf("Element type: %v\n", el.ElementType)
	switch el.ElementType { // some of these enums could be consolidated with extra params
	case Container:
		ui.DrawContainer(el)
		break
	case DecoratedContainer:
		ui.DrawContainer(el)
		break
	case Button:
		ui.DrawButton(el)
		break
	case RoundedButton:
		break
	case Text:
		break
	default:
		panic("Invalid element type")
	}
	// Loop through child elements
	if len(el.Children) > 0 {
		for _, childEl := range el.Children {
			ui.DrawElement(&childEl)
		}
	}
}

// Will determine whether element will scale with window size proportionally
// If sticky, disregards scaling bool so it doesn't conflict with resizing
func (ui *UI) DrawContainer(el *Element) {
	actualWidth := float32(el.Width)   // anchorX
	actualHeight := float32(el.Height) // anchorY
	if el.AutoScale && !el.Sticky.Sticky {
		actualWidth, actualHeight = autoScalingCalculation(el)
	}
	x, y := el.AnchorX, el.AnchorY
	if el.Sticky.Sticky { // no need to check for autoScale, as it should override autoScale
		switch el.Sticky.Anchor { // anchor anchorX, anchorY, width, height to whatever corner it's sticking to.
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
		ui.Context,
		float32(el.AnchorX),
		float32(el.AnchorY),
		actualWidth,
		actualHeight,
		DefaultContainerTheme.BackgroundColor,
		true,
	)
	fmt.Println("Drawing container")
	// if el.paddingX > 0 { // inset whole
	//
	// }
	// vector.DrawFilledRect(
	// 	ui.Context,
	// 	float32(el.AnchorX),
	// 	float32(el.AnchorY),
	// 	actualWidth,
	// 	actualHeight,
	// 	el.Theme.backgroundColor,
	// 	true,
	// )
}
func (ui *UI) DrawButton(el *Element) {

}
