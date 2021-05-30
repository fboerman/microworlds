package microworlds

import (
	"fmt"
	"github.com/fboerman/microworlds/utils"
	"image/color"
)

type Cell struct {
	Active bool
	C      color.RGBA
	X      int
	Y      int
}

type MicroWorld struct {
	Width   int
	Heigth  int
	cells_R []Cell
	cells_W []Cell
}

func (w *MicroWorld) Setup(width int, heigth int) {
	w.Width = width
	w.Heigth = heigth
	w.cells_R = make([]Cell, width*heigth)
	w.cells_W = make([]Cell, width*heigth)

	fmt.Printf("cells: %d, Width:%d, Heigth: %d\n", len(w.cells_R), w.Width, w.Heigth)
	for y := 0; y < heigth; y++ {
		for x := 0; x < width; x++ {
			cell := w.GetCellW(x, y)
			cell.X = x
			cell.Y = y
			//fmt.Println(cell.X, cell.Y)
		}
	}
}

func (w *MicroWorld) FlipBuffer() {
	copy(w.cells_R, w.cells_W)
}

func (w *MicroWorld) GetCell(x int, y int) *Cell {
	// returns pointer to the cell at the coordinate from current read buffer
	if x < 0 || x >= w.Width || y < 0 || y >= w.Heigth {
		return nil
	}
	return &w.cells_R[y*w.Width+x]
}

func (w *MicroWorld) GetCellW(x int, y int) *Cell {
	// returns pointer to the cell at the coordinate from current write buffer
	if x < 0 || x >= w.Width || y < 0 || y >= w.Heigth {
		return nil
	}
	return &w.cells_W[y*w.Width+x]
}

func (w *MicroWorld) GetNeighborCells(x int, y int) []*Cell {
	directions := [8][2]int{
		{-1, -1},
		{0, -1},
		{+1, -1},
		{-1, 0},
		{+1, 0},
		{-1, +1},
		{0, +1},
		{+1, +1},
	}
	cells := []*Cell{}
	for _, dir := range directions {
		c := w.GetCell(x+dir[0], y+dir[1])
		if c != nil {
			cells = append(cells, c)
		}
	}

	return cells
}

func (w *MicroWorld) NumActiveCells() int {
	num := 0
	for _, c := range w.cells_R {
		if c.Active {
			num++
		}
	}

	return num
}

func (w *MicroWorld) GetActiveCells() []*Cell {
	cells := []*Cell{}
	for i, c := range w.cells_R {
		if c.Active {
			cells = append(cells, &w.cells_R[i])
		}
	}

	return cells
}

func (w *MicroWorld) GetRandomActiveCell() *Cell {
	cells := w.GetActiveCells()
	if len(cells) == 0 {
		return nil
	}
	return cells[utils.Randomn(len(cells))]
}
