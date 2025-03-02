package main

import (
	"fmt"
	"github.com/Clyde-Goodall/ebitengine-ui-lib/lib/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var config lib.WindowConfig
var ui lib.UI

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	err := ui.Draw(screen)
	if err != nil {
		return
	}
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
	ui.AddChild(lib.Element{
		lib.Container,
	})
	fmt.Print(ui)

	ebiten.SetWindowSize(config.Width, config.Height)
	ebiten.SetWindowTitle("Test Window")

	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println("unable to run the game")
	}

}
