package sdl2canvas

import (
	"fmt"
	"github.com/fboerman/microworlds/microworlds"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type SDL2Canvas struct {
	windowWidth  int
	windowHeight int
	window       *sdl.Window
	renderer     *sdl.Renderer
	texture      *sdl.Texture
	pixels       []byte
	event        sdl.Event
	err          error
	Running      bool
}

const CELLSIZE int = 10
const SPACING int = 0

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

	s.texture, s.err = s.renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING,
		int32(windowWidth), int32(windowHeight))
	if s.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", s.texture)
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

func (s *SDL2Canvas) SetSquare(x_ int, y_ int, spacing int, c sdl.Color) {
	x_ *= CELLSIZE
	y_ *= CELLSIZE
	for y := 0; y < CELLSIZE-spacing; y++ {
		for x := 0; x < CELLSIZE-spacing; x++ {
			s.SetPixel(x_+x, y_+y, c)
		}
	}
}

func (s *SDL2Canvas) Render(w *microworlds.MicroWorld) {
	for y := 0; y < w.Heigth; y++ {
		for x := 0; x < w.Width; x++ {
			cell := w.GetCell(x, y)
			if cell.Active {
				s.SetSquare(x, y, SPACING, sdl.Color{cell.C.R, cell.C.G, cell.C.B, cell.C.A})
			}
		}
	}

	s.texture.Update(nil, s.pixels, int(s.windowWidth*4))
	s.renderer.SetDrawColor(0, 0, 0, 255)
	s.renderer.Clear()
	s.renderer.Copy(s.texture, nil, nil)
	s.renderer.Present()
}

func (s *SDL2Canvas) Shutdown() {
	s.texture.Destroy()
	s.renderer.Destroy()
	s.window.Destroy()
	sdl.Quit()
}
