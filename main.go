package main

import (
	"evolution-sim/entity"
	"evolution-sim/terrain"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zedzee37/znoise/noise"
)


const WIDTH = 800
const HEIGHT = 800

func main() {
	heightMap := noise.NewPerlinNoise(100, 5, 2.0, 0.03, 0.5, 1.0)
	worldMap := terrain.NewMap(WIDTH, HEIGHT)
	err := worldMap.ApplyHeightMap(heightMap)

	if err != nil {
		panic(err)
	}

	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(WIDTH, HEIGHT, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	creature := entity.NewCreature(rl.Vector2{
		X: float32(WIDTH) / 2.0,
		Y: float32(HEIGHT) / 2.0,
	})

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		
		err := worldMap.DrawMap(rl.Vector2Zero(), rl.Vector2{X: 1, Y: 1})
		if err != nil {
			panic(err)
		}

		creature.Draw(0, 0)

		rl.EndDrawing()
	}
}
