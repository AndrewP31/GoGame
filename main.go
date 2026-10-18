//Andrew Pfohl
//CSC 461
//DSU

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(gridW*tileSize, gridH*tileSize)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GoGame")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
