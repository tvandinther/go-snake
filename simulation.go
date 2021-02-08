package main

import (
	"fmt"
	"sync"
	"time"
)

type simulation struct {
	renderer            Renderer
	snake               *Snake
	food                *Food
	gridFactory         *GridFactory
	millisecondsPerTick time.Duration
	tick                int
	score               int
	done                chan bool
}

func NewSimulation(renderer Renderer, snake *Snake, gridFactory *GridFactory, ticksPerSecond int) simulation {
	return simulation{
		renderer:            renderer,
		snake:               snake,
		gridFactory:         gridFactory,
		millisecondsPerTick: time.Duration(1000 / ticksPerSecond),
		done:                make(chan bool, 1),
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
		case <-simulation.done:
			return
		case _ = <-ticker.C:
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

	movementNumber := simulation.tick%len(movements) + 3

	if movementNumber < len(movements) {
		snake.SetMovement(movements[movementNumber])
	}

	snake.Move()
	simulation.checkFood()
	simulation.display()
}

func (simulation *simulation) setFood() {
	var food *Food
	ok := true

	for ok {
		food = NewFood(simulation.gridFactory.xSize, simulation.gridFactory.ySize)
		ok = isBodyCollision(food, simulation.snake)
	}

	simulation.food = food
}

func isBodyCollision(food *Food, snake *Snake) bool {
	fmt.Println(food.X, food.Y)
	for e := snake.Body.Front(); e != nil; e = e.Next() {
		body, ok := e.Value.(*snakeBody)
		if ok {
			if food.X == body.X && food.Y == body.Y {
				return true
			}
		} else {
			panic(e.Value)
		}
	}
	return false
}

func (simulation *simulation) checkFood() {
	snakeHead := simulation.snake.head
	food := simulation.food

	if snakeHead.X == food.X && snakeHead.Y == food.Y {
		simulation.score++
		simulation.snake.eat()
		simulation.setFood()
	}
}

func (simulation *simulation) display() {
	grid := simulation.gridFactory.GenerateBlankGrid()
	grid.AddSnake(simulation.snake)
	grid.AddFood(simulation.food)

	simulation.renderer.Render(grid, simulation.score)
}
