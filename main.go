package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zedzee37/znoise/noise"
)


const WIDTH = 800
const HEIGHT = 800

func main() {
	n := noise.NewPerlinNoise(100, 3, 2.0, 0.1, 0.5, 1.0)

	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(WIDTH, HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()


	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for x := range WIDTH {
			for y := range HEIGHT {
				fx, fy := float64(x) / float64(WIDTH), float64(y) / float64(HEIGHT)
				noiseValue, err := n.Get(fx, fy)

				if err != nil {
					panic(err)
				}

				rl.DrawPixel(int32(x), int32(y), rl.Color{
					R: uint8(noiseValue * 255),
					G: uint8(noiseValue * 255),
					B: uint8(noiseValue * 255),
					A: 255,
				})	
			}
		}
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
