package main

import (
	"fmt"
	"github.com/Clyde-Goodall/ebitengine-ui-lib/lib/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

type Game struct{}

var config lib.WindowConfig
var ui lib.UI

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.Width, config.Height
}

func UITesting(t *testing.T) {
	config = lib.WindowConfig{
		Width:  1000,
		Height: 600,
	}
	// ui initialization
	ui, err := lib.InitUI(&ebiten.Image{}, config)
	if err != nil {
		t.Fatalf("unable to initialize UI")
	}
	fmt.Print(ui)

	ebiten.SetWindowSize(config.Width, config.Height)
	ebiten.SetWindowTitle("Test Window")

	if err := ebiten.RunGame(&Game{}); err != nil {
		t.Fatalf("unable to run the game")
	}

}
