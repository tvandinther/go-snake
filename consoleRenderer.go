package main

import (
	"github.com/gdamore/tcell/v2"
	"os"
	"strings"
)

type consoleRenderer struct {
	screen tcell.Screen
	style tcell.Style
	empty     string
	snakeHead string
	snakeBody string
	food      string
	border    string
}

func NewConsoleRenderer() *consoleRenderer {
	screen, _ := tcell.NewScreen()
	style := tcell.Style{}
	style.Background(tcell.Color18)
	err := screen.Init()
	if err != nil {
		panic(err)
	}
	return &consoleRenderer{
		screen: screen,
		style: style,
		empty:     "  ",
		snakeBody: "()",
		snakeHead: "{}",
		food:      "><",
		border:    "##",
	}
}

func (renderer *consoleRenderer) setContentString(x, y int, content string) {
	screen := renderer.screen
	mainc := rune(content[0])
	combc := []rune(content[1:])
	screen.SetContent(x, y, mainc, combc, renderer.style)
}

func (renderer *consoleRenderer) Render(simulation *Simulation) {
	screen := renderer.screen
	screen.Clear()
	//var gridString string
	//
	renderer.drawBorder(simulation.gridFactory.xSize, simulation.gridFactory.ySize)
	renderer.drawSnake(simulation.snake)

	//gridString += "\n" + renderer.border
	//
	//for i := 0; i < len(*gridPtr); i++ {
	//	for j := 0; j < len((*gridPtr)[i]); j++ {
	//		var square string
	//		switch (*gridPtr)[j][i] {
	//		case EmptySquare:
	//			square = renderer.empty
	//		case SnakeHeadSquare:
	//			square = renderer.snakeHead
	//		case SnakeBodySquare:
	//			square = renderer.snakeBody
	//		case FoodSquare:
	//			square = renderer.food
	//		}
	//		gridString += square
	//	}
	//
	//	gridString += renderer.border + "\n" + renderer.border
	//}
	//gridString += strings.Repeat(renderer.border, width+1)
	//gridString += "\n"
	//
	//scoreString := fmt.Sprintf("            Score: %d", score)

	//finalPrint := gridString + scoreString
	screen.Show()
	//fmt.Println(finalPrint)
}

func (renderer *consoleRenderer) drawBorder(innerWidth, innerHeight int) {
	borderWidth := len(renderer.border)

	renderer.setContentString(0, 0, strings.Repeat(renderer.border, innerWidth + borderWidth))
	for i := 0; i < innerHeight; i++ {
		renderer.setContentString(0, i + 1, renderer.border)
		renderer.setContentString((innerWidth * borderWidth) + 1, i + 1, renderer.border)
	}
	renderer.setContentString(0, innerHeight + 1, strings.Repeat(renderer.border, innerWidth + borderWidth))
}

func (renderer *consoleRenderer) drawSnake(snake *Snake) {
	snakeWidth := len(renderer.snakeHead)
	borderWidth := len(renderer.border)

	renderer.setContentString((snake.head.X * snakeWidth)+ borderWidth, snake.head.Y + 1, renderer.snakeHead)

	for e := snake.Body.Front(); e != nil; e = e.Next() {
		body, ok := e.Value.(*snakeBody)
		if ok {
			renderer.setContentString((body.X * snakeWidth)+ borderWidth, body.Y + 1, renderer.snakeBody)
		}
	}
}

//func (renderer *consoleRenderer) _Render(gridPtr *Grid, score int) {
//	var gridString string
//
//	width := len((*gridPtr)[0])
//
//	gridString += strings.Repeat(renderer.border, width+2)
//	gridString += "\n" + renderer.border
//
//	for i := 0; i < len(*gridPtr); i++ {
//		for j := 0; j < len((*gridPtr)[i]); j++ {
//			var square string
//			switch (*gridPtr)[j][i] {
//			case EmptySquare:
//				square = renderer.empty
//			case SnakeHeadSquare:
//				square = renderer.snakeHead
//			case SnakeBodySquare:
//				square = renderer.snakeBody
//			case FoodSquare:
//				square = renderer.food
//			}
//			gridString += square
//		}
//
//		gridString += renderer.border + "\n" + renderer.border
//	}
//	gridString += strings.Repeat(renderer.border, width+1)
//	gridString += "\n"
//
//	scoreString := fmt.Sprintf("            Score: %d", score)
//
//	finalPrint := gridString + scoreString
//	renderer.Clear()
//	fmt.Println(finalPrint)
//}

func (renderer *consoleRenderer) Clear() {
	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run()

	_, err := os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	if err != nil {
		panic(err)
	}
}
