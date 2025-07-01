package main

import (
	"evolution-sim/terrain"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zedzee37/znoise/noise"
)


const WIDTH = 800
const HEIGHT = 800

func main() {
	heightMap := noise.NewPerlinNoise(100, 3, 2.0, 0.1, 0.5, 1.0)
	worldMap := terrain.NewMap(WIDTH, HEIGHT)
	err := worldMap.ApplyHeightMap(heightMap)

	if err != nil {
		panic(err)
	}

	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(WIDTH, HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()


	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
