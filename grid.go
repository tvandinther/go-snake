package main

type Grid [][]SquareType

func NewGrid(rows, cols int) *Grid {
	grid := make(Grid, rows)

	for r := range grid {
		grid[r] = make([]SquareType, cols)
	}

	return &grid
}

func (grid *Grid) AddSnake(snake *Snake) {
	for e := snake.Body.Front(); e != nil; e = e.Next() {
		body, ok := e.Value.(*snakeBody)
		if ok {
			grid.setSquare(body.X, body.Y, SnakeBodySquare)
		} else {
			panic("Not a snake body!")
		}
	}
	grid.setSquare(snake.head.X, snake.head.Y, SnakeHeadSquare)
}

func (grid *Grid) AddFood(food *Food) {
	grid.setSquare(food.X, food.Y, FoodSquare)
}

func (grid *Grid) setSquare(x, y int, squareType SquareType) {
	(*grid)[x][y] = squareType
}

type SquareType int

const (
	EmptySquare SquareType = iota
	SnakeHeadSquare
	SnakeBodySquare
	FoodSquare
)
