package main

type GridFactory struct {
	xSize int
	ySize int
}

func NewGridFactory(xSize, ySize int) *GridFactory {
	return &GridFactory{
		xSize: xSize,
		ySize: ySize,
	}
}

func (gridFactory *GridFactory) GenerateBlankGrid() *Grid {
	return NewGrid(gridFactory.xSize, gridFactory.ySize)
}
