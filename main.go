// Copyright 2021 Siôn le Roux.  All rights reserved.
// Use of this source code is subject to an MIT-style
// licence which can be found in the LICENSE file.

package main

import (
	"errors"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func main() {
	gameWidth, gameHeight := 640, 480

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("ebitengine-game-template")

	game := &Game{
		Width:  gameWidth,
		Height: gameHeight,
		Player: &Player{image.Pt(gameWidth/2, gameHeight/2)},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// Game represents the main game state
type Game struct {
	Width  int
	Height int
	Player *Player
}

// Layout is hardcoded for now, may be made dynamic in future
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.Width, g.Height
}

// Update calculates game logic
func (g *Game) Update() error {

	// Pressing Q any time quits immediately
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return errors.New("game quit by player")
	}

	// Pressing F toggles full-screen
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}

	// Movement controls
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Player.Move()
	}

	// XXX: Write game logic here

	return nil
}

// Draw draws the game screen by one frame
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		float64(g.Player.Coords.X),
		float64(g.Player.Coords.Y),
		20,
		20,
		color.White,
	)
}

// Player is the player character in the game
type Player struct {
	Coords image.Point
}

// Move moves the player upwards
func (p *Player) Move() {
	p.Coords.Y--
}
