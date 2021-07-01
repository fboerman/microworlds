package microworlds

import (
	"fmt"
	"github.com/fboerman/microworlds/utils"
	"image/color"
)

// code for hexagon stuff inspired by:
// https://www.redblobgames.com/grids/hexagons/
// https://www.redblobgames.com/grids/hexagons/implementation.html
// even-q offset coordinate system with flat tops is used

type Hex struct {
	Active bool
	C      color.RGBA
	Q      int
	R      int
}

type MicroWorld struct {
	Width  int
	Heigth int
	hexs_R []Hex
	hexs_W []Hex
}

func (w *MicroWorld) Setup(width int, heigth int) {
	w.Width = width
	w.Heigth = heigth
	w.hexs_R = make([]Hex, width*heigth)
	w.hexs_W = make([]Hex, width*heigth)

	fmt.Printf("cells: %d, Width:%d, Heigth: %d\n", len(w.hexs_R), w.Width, w.Heigth)
	for y := 0; y < heigth; y++ {
		for x := 0; x < width; x++ {
			cell := w.GetHexW(x, y)
			cell.Q = x
			cell.R = y
			//fmt.Println(cell.Q, cell.R)
		}
	}
}

func (w *MicroWorld) FlipBuffer() {
	copy(w.hexs_R, w.hexs_W)
}

func (w *MicroWorld) GetHex(q int, r int) *Hex {
	// returns pointer to the cell at the coordinate from current read buffer
	if q < 0 || q >= w.Width || r < 0 || r >= w.Heigth {
		return nil
	}
	return &w.hexs_R[r*w.Width+q]
}

func (w *MicroWorld) GetHexW(q int, r int) *Hex {
	// returns pointer to the cell at the coordinate from current write buffer
	if q < 0 || q >= w.Width || r < 0 || r >= w.Heigth {
		return nil
	}
	return &w.hexs_W[r*w.Width+q]
}

func (w *MicroWorld) GetNeighborHexs(q int, r int) []*Hex {
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
	cells := []*Hex{}
	for _, dir := range directions {
		c := w.GetHex(q+dir[0], r+dir[1])
		if c != nil {
			cells = append(cells, c)
		}
	}

	return cells
}

func (w *MicroWorld) NumActiveHex() int {
	num := 0
	for _, c := range w.hexs_R {
		if c.Active {
			num++
		}
	}

	return num
}

func (w *MicroWorld) GetActiveHexs() []*Hex {
	cells := []*Hex{}
	for i, c := range w.hexs_R {
		if c.Active {
			cells = append(cells, &w.hexs_R[i])
		}
	}

	return cells
}

func (w *MicroWorld) GetRandomActiveHex() *Hex {
	cells := w.GetActiveHexs()
	if len(cells) == 0 {
		return nil
	}
	return cells[utils.Randomn(len(cells))]
}
