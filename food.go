package main

import "math/rand"

type Food struct {
	X int
	Y int
}

func NewFood(xSize, ySize int) *Food {
	return &Food{
		X: Mod(rand.Int(), xSize),
		Y: Mod(rand.Int(), ySize),
	}
}
