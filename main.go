package main

import (
	"github.com/fboerman/microworlds/sdl2canvas"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var canvas sdl2canvas.SDL2Canvas

	canvas.Setup("Microworlds", 800, 600)
	defer canvas.Shutdown()

	canvas.Render()

	sdl.Delay(10000)
}
