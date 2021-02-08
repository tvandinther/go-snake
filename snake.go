package main

type snake struct {
	Xsize int
	Ysize int
	head *snakeBody
	Body []*snakeBody
}

type snakeBody struct {
	X int
	Y int
	nextMovement CoordinateDelta
}

func NewSnake(x, y, length, xSize, ySize int) *snake {
	snake := snake{
		Body: []*snakeBody{},
		head: newSnakeBody(x, y, MoveRight),
		Xsize: xSize,
		Ysize: ySize,
	}
	generateSnakeBody(&snake, length)

	return &snake
}

func (snake snake) Length() int {
	return len(snake.Body)
}

type CoordinateDelta struct {
	X int
	Y int
}

var (
	MoveUp = CoordinateDelta{
		Y: -1,
	}
	MoveDown = CoordinateDelta{
		Y: 1,
	}
	MoveLeft = CoordinateDelta{
		X: -1,
	}
	MoveRight = CoordinateDelta{
		X: 1,
	}
)

func (snake *snake) SetMovement(delta CoordinateDelta) {
	snake.head.nextMovement = delta
}

func (snake *snake) Move() {
	lastMovement := snake.head.nextMovement

	for i := 0; i < len(snake.Body); i++ {
		snakeBody := snake.Body[i]
		snakeBody.X += snakeBody.nextMovement.X
		snakeBody.Y += snakeBody.nextMovement.Y

		nextLastMovement := snakeBody.nextMovement
		snakeBody.nextMovement = lastMovement
		lastMovement = nextLastMovement

		snakeBody.X = Mod(snakeBody.X, snake.Xsize)
		snakeBody.Y = Mod(snakeBody.Y, snake.Ysize)
	}
}

func newSnakeBody(x, y int, nextMovement CoordinateDelta) *snakeBody {
	return &snakeBody{
		X: x,
		Y: y,
		nextMovement: nextMovement,
	}
}

func generateSnakeBody(snake *snake, length int) {
	for i := 0; i < length; i++ {
		snakeBody := newSnakeBody(Mod(snake.head.X - i, snake.Xsize), Mod(snake.head.Y, snake.Ysize), MoveRight)
		if i == 0 {snake.head = snakeBody}
		snake.Body = append(snake.Body, snakeBody)
	}
}
