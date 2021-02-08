package main

import (
	"sync"
	"time"
)

type simulation struct {
	renderer Renderer
	snake    *Snake
	food *Food
	gridFactory *GridFactory
	millisecondsPerTick time.Duration
	tick int
	score int
	done chan bool
}

func NewSimulation(renderer Renderer, snake *Snake, gridFactory *GridFactory, ticksPerSecond int) simulation {
	return simulation{
		renderer: renderer,
		snake: snake,
		gridFactory: gridFactory,
		millisecondsPerTick: time.Duration(1000 / ticksPerSecond),
		done: make(chan bool, 1),
	}
}

func (simulation *simulation) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	simulation.food = NewFood(simulation.gridFactory.xSize, simulation.gridFactory.ySize)
	simulation.display()

	ticker := time.NewTicker(simulation.millisecondsPerTick * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <- simulation.done:
			return
		case _ = <- ticker.C:
			simulation.Tick()
		}
	}
}

func (simulation *simulation) Stop() {
	simulation.done <- true
}

func (simulation *simulation) Tick() {
	simulation.tick++

	snake := simulation.snake

	movementNumber := simulation.tick % len(movements) + 3

	if movementNumber < len(movements) {snake.SetMovement(movements[movementNumber])}

	simulation.checkFood()
	snake.Move()
	simulation.display()
}

func (simulation *simulation) setFood() {
	var ok bool
	for ok {
		food := NewFood(simulation.gridFactory.xSize, simulation.gridFactory.ySize)
		ok = !isBodyCollision(food, simulation.snake)
	}
}

func isBodyCollision(food *Food, snake *Snake) bool {
	snakeBody := snake.Body

	for i := 0; i < len(snakeBody); i++ {
		snakeBodySegment := snakeBody[i]
		if food.X == snakeBodySegment.X && food.Y == snakeBodySegment.Y {
			return true
		}
	}
	return false
}

func (simulation *simulation) checkFood() {
	snakeHead := simulation.snake.head
	food := simulation.food

	if snakeHead.X == food.X && snakeHead.Y == food.Y {
		simulation.score++
		simulation.setFood()
	}
}

func (simulation *simulation) display() {
	grid := simulation.gridFactory.GenerateBlankGrid()
	grid.AddSnake(simulation.snake)
	grid.AddFood(simulation.food)

	simulation.renderer.Render(grid, simulation.score)
}