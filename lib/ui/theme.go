package lib

import "image/color"

// might want to refactor this as a user-supplied hash map but not sure yet
var LightGrey = color.Color(color.RGBA{238, 238, 238, 200})
var OffWhite = color.Color(color.RGBA{255, 255, 242, 255})
var Orange = color.Color(color.RGBA{242, 123, 45, 255})
var LightOrange = color.Color(color.RGBA{242, 238, 226, 255})
var Yellow = color.Color(color.RGBA{255, 253, 123, 255})
var Black = color.Color(color.RGBA{0, 0, 0, 255})
var Transparent = color.Color(color.RGBA{0, 0, 0, 0})

// Theme preset examplesfor UI elements.
// Currently aiming for the old Flipntoe Studio UI color palette, but will be made more agnostic later on.
type ElementColorPreset struct {
	borderColor     color.Color
	backgroundColor color.Color
	foregroundColor color.Color
	// focused/active colors
	activeBorderColor     color.Color
	activeBackgroundColor color.Color
	activeForegroundColor color.Color
}

// currently making the default transparent.
// Since UI element initialization alone don't  dictate the default sizing of the area on the screen,
// it's not completely necessary to supply a background color/Theme struct on init and should instead be handled by ebitengine draw funcs.
var DefaultContainerTheme = ElementColorPreset{
	borderColor:     Transparent,
	backgroundColor: Transparent,
	foregroundColor: Black,
	// focused/active
	activeBorderColor:     Transparent,
	activeBackgroundColor: Transparent,
	activeForegroundColor: Black,
}

var DefaultRoundedButtonTheme = ElementColorPreset{
	borderColor:     Orange,
	backgroundColor: LightOrange,
	foregroundColor: Orange,
	// focused/active
	activeBorderColor:     Black,
	activeBackgroundColor: Yellow,
	activeForegroundColor: Black,
}

var DefaultButtonTheme = ElementColorPreset{
	borderColor:     LightOrange,
	backgroundColor: OffWhite,
	foregroundColor: LightOrange,
	// focused/active
	activeBorderColor:     Orange,
	activeBackgroundColor: OffWhite,
	activeForegroundColor: Orange,
}
