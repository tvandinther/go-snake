package main

import (
	"sync"
	"time"
)

func SnakeGame(xSize, ySize, ticksPerSecond int) {
	snake := NewSnake(5, 5, 10, xSize, ySize)

	gridFactory := NewGridFactory(xSize, ySize)

	var renderer Renderer
	renderer = NewConsoleRenderer()

	var wg sync.WaitGroup
	wg.Add(1)

	simulation := NewSimulation(renderer, snake, gridFactory, ticksPerSecond)
	go simulation.Start(&wg)

	go func() {
		time.Sleep(20 * time.Second)
		simulation.Stop()
	}()

	wg.Wait()
}
