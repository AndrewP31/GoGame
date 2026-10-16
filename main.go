package main

import(
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
	ebiten.SetWindowSize(gridW * tileSize, gridH * tileSize)
	ebiten.SetWindowTitle("Roguelike game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}