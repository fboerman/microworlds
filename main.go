package main

import (
	"fmt"
	"github.com/fboerman/microworlds/microworlds"
	"github.com/fboerman/microworlds/sdl2canvas"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var canvas sdl2canvas.SDL2Canvas
	var forest microworlds.Forest
	p := float32(0.5)
	n := 1

	fmt.Printf("Generate forest with %d %% trees\n", int(p*100))
	canvas.Setup("Microworlds", 1600, 900)
	forest.Setup(p, 160, 90)
	forest.GetWorld().FlipBuffer()
	defer canvas.Shutdown()

	fmt.Printf("Ignite %d fires\n", n)
	forest.Ignite(n)
	forest.GetWorld().FlipBuffer()
	// initial render
	canvas.Render(forest.GetWorld())

	// tick loop
	for i := 0; i < 1000; i++ {
		forest.Tick()
		forest.GetWorld().FlipBuffer()
		canvas.Render(forest.GetWorld())
		sdl.Delay(100)
		fmt.Println(i)
	}

	sdl.Delay(10000)
}
