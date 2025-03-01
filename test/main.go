package test

import (
	"github.com/Clyde-Goodall/ebitengine-ui-lib/lib/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct{}

var config lib.WindowConfig
var ui lib.UI

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.draw
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.Width, config.Height
}

func main() {
	config = lib.WindowConfig{
		Width:  1000,
		Height: 600,
	}
	// ui initialization
	ui, err := lib.InitUI(&ebiten.Image{}, config)
	if err != nil {
		panic("unable to initialize UI")
	}
	// setup
	ebiten.SetWindowSize(config.Width, config.Height)
	ebiten.SetWindowTitle("Test Window")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
