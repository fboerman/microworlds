package sdl2canvas

import (
	"fmt"
	"github.com/fboerman/microworlds/microworlds"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type SDL2Canvas struct {
	windowWidth  int
	windowHeight int
	window       *sdl.Window
	renderer     *sdl.Renderer
	pixels       []byte
	event        sdl.Event
	err          error
	Running      bool
}

const CELLSIZE int16 = 10

// Setup Window / renderer / texture
func (s *SDL2Canvas) Setup(title string, windowWidth int, windowHeight int) {
	sdl.Init(sdl.INIT_EVERYTHING)

	var flags uint32 = sdl.WINDOW_SHOWN

	s.windowWidth = windowWidth
	s.windowHeight = windowHeight

	s.window, s.err = sdl.CreateWindow(title,
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		int32(windowWidth), int32(windowHeight),
		flags)
	if s.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Window: %s\n", s.err)
		os.Exit(1)
	}

	s.renderer, s.err = sdl.CreateRenderer(s.window, -1, sdl.RENDERER_ACCELERATED)
	if s.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", s.err)
		os.Exit(1)
	}

	s.pixels = make([]byte, windowWidth*windowHeight*4)

	s.Running = true
}

func (s *SDL2Canvas) HandleEvents() {
	for s.event = sdl.PollEvent(); s.event != nil; s.event = sdl.PollEvent() {
		switch t := s.event.(type) {
		case *sdl.QuitEvent:
			s.Running = false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				s.Running = false
			}
		}
	}
}

func (s *SDL2Canvas) SetPixel(x int, y int, c sdl.Color) {
	index := (y*s.windowWidth + x) * 4

	if index < len(s.pixels)-4 && index >= 0 {
		s.pixels[index] = c.R
		s.pixels[index+1] = c.G
		s.pixels[index+2] = c.B
	}
}

func (s *SDL2Canvas) SetSquare(x_ int, y_ int, c sdl.Color) {
	x := int16(x_)
	y := int16(y_)
	var x_array = []int16{x * CELLSIZE, (x + 1) * CELLSIZE, (x + 1) * CELLSIZE, x * CELLSIZE}
	var y_array = []int16{y * CELLSIZE, y * CELLSIZE, (y + 1) * CELLSIZE, (y + 1) * CELLSIZE}
	gfx.FilledPolygonColor(s.renderer, x_array, y_array, c)
}

func (s *SDL2Canvas) Render(w *microworlds.MicroWorld) {
	s.renderer.SetDrawColor(0, 0, 0, 255)
	s.renderer.Clear()

	// // if you want gridlines activate below function.
	// // if you want gridlines overwriting the cell surface then put these below the for loop
	//for i := int16(0); i < int16(s.windowHeight)/CELLSIZE; i++ {
	//	gfx.HlineColor(s.renderer, 0, int32(s.windowWidth), int32(i*CELLSIZE), sdl.Color{255, 255, 255, 255})
	//}
	//for i := int16(0); i < int16(s.windowWidth)/CELLSIZE; i++ {
	//	gfx.VlineColor(s.renderer, int32(i*CELLSIZE), 0, int32(s.windowHeight), sdl.Color{255, 255, 255, 255})
	//}

	for y := 0; y < w.Heigth; y++ {
		for x := 0; x < w.Width; x++ {
			cell := w.GetCell(x, y)
			if cell.Active {
				s.SetSquare(x, y, sdl.Color{cell.C.R, cell.C.G, cell.C.B, cell.C.A})
			}
		}
	}

	s.renderer.Present()
}

func (s *SDL2Canvas) Shutdown() {
	s.renderer.Destroy()
	s.window.Destroy()
	sdl.Quit()
}
