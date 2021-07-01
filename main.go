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
	var p float32
	var n int
	fmt.Println("Please input float [0-1] for p% tree density")
	fmt.Scanf("%f\n", &p)
	fmt.Println("Please input integer for number of starting fires")
	fmt.Scanf("%d\n", &n)

	fmt.Printf("Generate forest with %d %% trees\n", int(p*100))
	canvas.Setup("Microworlds", 1920, 1080)
	forest.Setup(p, 192, 108)
	forest.GetWorld().FlipBuffer()
	defer canvas.Shutdown()

	fmt.Printf("Ignite %d fires\n", n)
	forest.Ignite(n)
	forest.GetWorld().FlipBuffer()
	// initial render
	canvas.Render(forest.GetWorld())

	//// tick loop
	//for i := 0; ; i++ {
	//	if !forest.Tick() {
	//		break
	//	}
	//	fmt.Printf("Tick: %d\n", i)
	//	start := time.Now()
	//	forest.GetWorld().FlipBuffer()
	//	elapsed1 := time.Since(start)
	//	canvas.Render(forest.GetWorld())
	//	elapsed2 := time.Since(start)
	//	fmt.Printf("Took %s, of which %s was buffer flip\n", elapsed2, elapsed1)
	//	sdl.Delay(100)
	//}

	sdl.Delay(5000)
}
