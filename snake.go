package main

import "container/list"

type Snake struct {
	Xsize int
	Ysize int
	head *snakeBody
	justAte bool
	Body *list.List
}

type snakeBody struct {
	X int
	Y int
	nextMovement CoordinateDelta
}

func NewSnake(x, y, length, xSize, ySize int) *Snake {
	snake := Snake{
		Body: list.New(),
		head: newSnakeBody(x, y, MoveRight),
		Xsize: xSize,
		Ysize: ySize,
	}
	generateSnakeBody(&snake, length)

	return &snake
}

func (snake *Snake) Length() int {
	return snake.Body.Len()
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
	//lastMovement := snake.head.nextMovement

	// MOVE HEAD
	snake.head.X += snake.head.nextMovement.X
	snake.head.Y += snake.head.nextMovement.Y

	snake.head.X = Mod(snake.head.X, snake.Xsize)
	snake.head.Y = Mod(snake.head.Y, snake.Ysize)

	if snake.justAte {
		snake.Body.PushFront(
			newSnakeBody(
				snake.head.X,
				snake.head.Y,
				snake.head.nextMovement,
			),
		)
		snake.justAte = false

		return
	} else {
		endOfSnakeElement := snake.Body.Back()
		endOfSnake, ok := endOfSnakeElement.Value.(*snakeBody)
		if ok {
			endOfSnake.X = snake.head.X
			endOfSnake.Y = snake.head.Y
			endOfSnake.nextMovement = snake.head.nextMovement

			snake.Body.MoveToFront(endOfSnakeElement)
		} else {
			panic("Not a snake body!")
		}

		snake.Body.MoveToFront(endOfSnakeElement)
	}

	// MOVE BODY
}

//func (snakeBody *snakeBody) Move(delta CoordinateDelta) {
//	snakeBody.X += snakeBody.nextMovement.X
//	snakeBody.Y += snakeBody.nextMovement.Y
//}

func (snake *Snake) eat() {
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
	for i := 0; i < length; i++ {
		snakeBody := newSnakeBody(Mod(snake.head.X - i, snake.Xsize), Mod(snake.head.Y, snake.Ysize), MoveRight)
		snake.Body.PushBack(snakeBody)
	}
}
