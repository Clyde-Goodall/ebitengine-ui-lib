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
	BorderColor     color.Color
	BackgroundColor color.Color
	ForegroundColor color.Color
	// focused/active colors
	ActiveBorderColor     color.Color
	ActiveBackgroundColor color.Color
	ActiveForegroundColor color.Color
}

// currently making the default transparent.
// Since UI element initialization alone don't  dictate the default sizing of the area on the screen,
// it's not completely necessary to supply a background color/Theme struct on init and should instead be handled by ebitengine draw funcs.
var DefaultContainerTheme = ElementColorPreset{
	BorderColor:     LightGrey,
	BackgroundColor: LightGrey,
	ForegroundColor: Black,
	// focused/active
	ActiveBorderColor:     Transparent,
	ActiveBackgroundColor: Transparent,
	ActiveForegroundColor: Black,
}

var DefaultRoundedButtonTheme = ElementColorPreset{
	BorderColor:     Orange,
	BackgroundColor: LightOrange,
	ForegroundColor: Orange,
	// focused/active
	ActiveBorderColor:     Black,
	ActiveBackgroundColor: Yellow,
	ActiveForegroundColor: Black,
}

var DefaultButtonTheme = ElementColorPreset{
	BorderColor:     LightOrange,
	BackgroundColor: OffWhite,
	ForegroundColor: LightOrange,
	// focused/active
	ActiveBorderColor:     Orange,
	ActiveBackgroundColor: OffWhite,
	ActiveForegroundColor: Orange,
}
