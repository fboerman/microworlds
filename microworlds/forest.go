package microworlds

import (
	"errors"
	"github.com/fboerman/microworlds/utils"
	"image/color"
)

type Forest struct {
	w MicroWorld
}

func (f *Forest) Setup(p float32, width int, heigth int) error {
	if p > 1 {
		return errors.New("p should be less then 1")
	}
	f.w.Setup(width, heigth)

	// initialize a random forest with p% of total cells as trees
	Ntot := width * heigth
	Ntrees := int(p * float32(Ntot))
	indices := utils.RandomListN(Ntrees, Ntot)
	for _, n := range indices {
		//fmt.Printf("%d: (%d,%d)\n", n, w.cells[n].Q, w.cells[n].R)
		f.w.hexs_W[n].Active = true
		f.w.hexs_W[n].C = color.RGBA{0, 128, 0, 255}
	}

	return nil
}

func (f *Forest) GetWorld() *MicroWorld {
	return &f.w
}

func (f *Forest) Ignite(n int) {
	//ignite n fires
	for i := 0; i < n; i++ {
		c := f.w.GetRandomActiveHex()
		c_ := f.w.GetHexW(c.Q, c.R)
		c_.C = color.RGBA{255, 0, 0, 255}
	}
}

func (f *Forest) Tick() bool {
	// spread the fire and also give fire effect
	// returns true if something changed, false otherwise
	change := false
	cells := f.w.GetActiveHexs()
	for _, c := range cells {
		if c.Active {
			c_ := f.w.GetHexW(c.Q, c.R)
			if c.C.R == 255 {
				// fire has been ignited just now so check all neighbors if they are active and a tree then ignite them
				// since all trees get ignited by a fire we only need to check this once
				// so when fire is at full 255 (just ignited)
				neighbors := f.w.GetNeighborHexs(c.Q, c.R)
				for _, n := range neighbors {
					if n != nil {
						if n.Active && n.C.G != 0 {
							// ignite!
							n_ := f.w.GetHexW(n.Q, n.R)
							n_.C.G = 0
							n_.C.R = 255
							change = true
						}
					}
				}
			}
			if c.C.R > 0 {
				// be decreasing the red step by step you get a nice dying fire effect
				if c.C.R > 25 {
					c_.C.R -= 25
				} else {
					c_.Active = false
				}
				change = true
			}
		}
	}

	return change
}
