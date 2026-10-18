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
	playerX, playerY int
	floorImg         *ebiten.Image
	playerImg        *ebiten.Image
	tiles            [][]TileType
	wallImg          *ebiten.Image
}

func NewGame() *Game {
	floor := ebiten.NewImage(tileSize-1, tileSize-1)
	floor.Fill(color.RGBA{60, 60, 70, 255})

	wall := ebiten.NewImage(tileSize-1, tileSize-1)
	wall.Fill(color.RGBA{0, 0, 0, 0})

	player := ebiten.NewImage(tileSize-8, tileSize-8)
	player.Fill(color.RGBA{200, 200, 60, 255})

	return &Game{
		playerX:   gridW / 2,
		playerY:   gridH / 2,
		floorImg:  floor,
		playerImg: player,
		wallImg:   wall,
		tiles:     NewMap(),
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

// edges and walls have collision
func (g *Game) isInBounds(x, y int) bool {
	return x >= 0 && x < gridW && y >= 0 && y < gridH && g.tiles[y][x] != Wall
}

func (g *Game) Draw(screen *ebiten.Image) {

	//floor or wall
	for y := 0; y < gridH; y++ {
		for x := 0; x < gridW; x++ {
			img := g.floorImg
			if g.tiles[y][x] == Wall {
				img = g.wallImg
			}

			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
			screen.DrawImage(img, options)
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
