package main

type TileType int

const (
	Floor TileType = 1
	Wall  TileType = 2
)

// place walls
func NewMap() [][]TileType {
	tiles := make([][]TileType, gridH)
	for y := range tiles {
		tiles[y] = make([]TileType, gridW)
	}

	//test walls
	tiles[5][7] = Wall
	tiles[3][8] = Wall
	tiles[4][9] = Wall
	tiles[6][7] = Wall

	return tiles
}
