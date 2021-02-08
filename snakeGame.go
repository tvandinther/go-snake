package main

import (
	//"github.com/gdamor/tcell"
	"time"
)

func SnakeGame(xSize, ySize, ticksPerSecond int)  {
	snake := NewSnake(5, 5, 10, xSize, ySize)

	grid := NewGrid(xSize, ySize)
	grid.AddSnake(snake)
	var renderer Renderer
	renderer = NewConsoleRenderer()
	renderer.Render(grid)

	millisecondsPerTick := time.Duration(1000 / ticksPerSecond)
	ticker := time.NewTicker(millisecondsPerTick * time.Millisecond)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	i := 0
	movements := []CoordinateDelta{
		MoveDown,
		MoveDown,
		MoveRight,
		MoveRight,
		MoveRight,
		MoveDown,
		MoveDown,
		MoveDown,
		MoveLeft,
		MoveLeft,
		MoveLeft,
		MoveUp,
		MoveUp,
	}

	for {
		select {

		case <- done:
		return

		case _ = <-ticker.C:
		if i < len(movements) {snake.SetMovement(movements[i])}

		snake.Move()

		grid := NewGrid(xSize, ySize)
		grid.AddSnake(snake)
		var renderer Renderer
		renderer = NewConsoleRenderer()
		renderer.Render(grid)

		i++
		}
	}
}