package main

type Snake struct {
	Xsize int
	Ysize int
	head *snakeBody
	justAte bool
	Body []*snakeBody
}

type snakeBody struct {
	X int
	Y int
	nextMovement CoordinateDelta
}

func NewSnake(x, y, length, xSize, ySize int) *Snake {
	snake := Snake{
		Body: []*snakeBody{},
		head: newSnakeBody(x, y, MoveRight),
		Xsize: xSize,
		Ysize: ySize,
	}
	generateSnakeBody(&snake, length)

	return &snake
}

func (snake *Snake) Length() int {
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

func (snake *Snake) SetMovement(delta CoordinateDelta) {
	snake.head.nextMovement = delta
}

func (snake *Snake) Move() {
	lastMovement := snake.head.nextMovement

	iterations := len(snake.Body)

	// MOVE HEAD
	snake.head.X += snake.head.nextMovement.X
	snake.head.Y += snake.head.nextMovement.Y

	snake.head.X = Mod(snake.head.X, snake.Xsize)
	snake.head.Y = Mod(snake.head.Y, snake.Ysize)

	if snake.justAte {
		snake.justAte = false

		return
	}

	// MOVE BODY
	for i := 0; i < iterations; i++ {
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

//func (snakeBody *snakeBody) Move(delta CoordinateDelta) {
//	snakeBody.X += snakeBody.nextMovement.X
//	snakeBody.Y += snakeBody.nextMovement.Y
//}

func (snake *Snake) eat() {
	snake.Body = append(snake.Body, newSnakeBody(
		snake.head.X,
		snake.head.Y,
		snake.head.nextMovement,
		))
	snake.justAte = true
}

func newSnakeBody(x, y int, nextMovement CoordinateDelta) *snakeBody {
	return &snakeBody{
		X: x,
		Y: y,
		nextMovement: nextMovement,
	}
}

func generateSnakeBody(snake *Snake, length int) {
	for i := length; i > 0; i-- {
		snakeBody := newSnakeBody(Mod(snake.head.X - i, snake.Xsize), Mod(snake.head.Y, snake.Ysize), MoveRight)
		snake.Body = append(snake.Body, snakeBody)
	}
}
