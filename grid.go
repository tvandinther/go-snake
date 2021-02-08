package main

type Grid [][]SquareType

func NewGrid(rows, cols int) *Grid {
	grid := make(Grid, rows)

	for r := range grid {
		grid[r] = make([]SquareType, cols)
	}

	return &grid
}

func (grid Grid) AddSnake(snake *snake) {
	for i := 0; i < len(snake.Body); i++ {
		snakeBody := *snake.Body[i]
		grid.setSquare(snakeBody.X, snakeBody.Y, SnakeSquare)
	}
}

func (grid Grid) setSquare(x, y int, squareType SquareType) {
	grid[x][y] = squareType
}

type SquareType int

const (
	EmptySquare SquareType = iota
	SnakeSquare
	FoodSquare
)
