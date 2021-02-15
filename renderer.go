package main

type Renderer interface {
	Render(simulation *Simulation)
	Clear()
}
