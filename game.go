package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	tileSize = 32
	gridW    = 20
	gridH    = 15
)

type Game struct {
	playerX, playerY int //player coordinates
	floorImg         *ebiten.Image
	playerImg        *ebiten.Image
}

func NewGame() *Game {
	floor := ebiten.NewImage(tileSize-1, tileSize-1)
	floor.Fill(color.RGBA{60, 60, 70, 255})

	player := ebiten.NewImage(tileSize-8, tileSize-8)
	player.Fill(color.RGBA{200, 200, 60, 255})

	return &Game{
		playerX:   gridW / 2,
		playerY:   gridH / 2,
		floorImg:  floor,
		playerImg: player,
	}

}

func (g *Game) Update() error {
	dx, dy := 0, 0

	//grid movement
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft):
		dx = -1
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowRight):
		dx = 1
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowUp):
		dy = -1
	case inpututil.IsKeyJustPressed(ebiten.KeyArrowDown):
		dy = 1
	}

	if dx != 0 || dy != 0 {
		nx, ny := g.playerX+dx, g.playerY+dy
		if g.isInBounds(nx, ny) {
			g.playerX, g.playerY = nx, ny
		}
	}

	return nil
}

func (g *Game) isInBounds(x, y int) bool {
	return x >= 0 && x < gridW && y >= 0 && y < gridH
}
func (g *Game) Draw(screen *ebiten.Image) {

	//floor
	for y := 0; y < gridH; y++ {
		for x := 0; x < gridW; x++ {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
			screen.DrawImage(g.floorImg, options)
		}
	}

	//player
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(g.playerX*tileSize+4), float64(g.playerY*tileSize+4))
	screen.DrawImage(g.playerImg, options)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gridW * tileSize, gridH * tileSize
}
