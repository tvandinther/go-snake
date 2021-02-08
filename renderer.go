package main

type Renderer interface {
	Render(grid *Grid, score int)
	Clear()
}
