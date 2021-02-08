package main

type Renderer interface {
	Render(grid *Grid)
	Clear()
}
