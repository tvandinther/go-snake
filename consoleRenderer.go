package main

import (
	"fmt"
	"os"
	"strings"
)

type consoleRenderer struct {
	empty     string
	snakeHead string
	snakeBody string
	food      string
	border    string
}

func NewConsoleRenderer() *consoleRenderer {
	return &consoleRenderer{
		empty:     "  ",
		snakeBody: "()",
		snakeHead: "{}",
		food:      "><",
		border:    "##",
	}
}

func (renderer *consoleRenderer) Render(gridPtr *Grid, score int) {
	var gridString string

	width := len((*gridPtr)[0])

	gridString += strings.Repeat(renderer.border, width+2)
	gridString += "\n" + renderer.border

	for i := 0; i < len(*gridPtr); i++ {
		for j := 0; j < len((*gridPtr)[i]); j++ {
			var square string
			switch (*gridPtr)[j][i] {
			case EmptySquare:
				square = renderer.empty
			case SnakeHeadSquare:
				square = renderer.snakeHead
			case SnakeBodySquare:
				square = renderer.snakeBody
			case FoodSquare:
				square = renderer.food
			}
			gridString += square
		}

		gridString += renderer.border + "\n" + renderer.border
	}
	gridString += strings.Repeat(renderer.border, width+1)
	gridString += "\n"

	scoreString := fmt.Sprintf("            Score: %d", score)

	finalPrint := gridString + scoreString
	renderer.Clear()
	fmt.Println(finalPrint)
}

func (renderer *consoleRenderer) Clear() {
	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run()

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
