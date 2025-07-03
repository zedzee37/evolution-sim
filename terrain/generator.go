package terrain

import (
	"fmt"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zedzee37/znoise/noise"
)

const FALL_OFF_RADIUS = 350.0
const FALL_OFF_SPEED = 0.005

type Map struct {
	Grid [][]rl.Color	
	Width int
	Height int
}

func NewMap(width int, height int) *Map {
	grid := make([][]rl.Color, width)
	
	for row := range width {
		grid[row] = make([]rl.Color, height)
	}

	worldMap := new(Map)
	*worldMap = Map{
		Grid: grid,
		Width: width,
		Height: height,
	}
 
	return worldMap
}

func (worldMap *Map) ApplyHeightMap(n noise.Noise) error {
	orderedThresholds := make([]float64, 0)
	for threshold := range TileSet {
		orderedThresholds = append(orderedThresholds, threshold)
	}
	slices.Sort(orderedThresholds)

	for x := range worldMap.Width {
		for y := range worldMap.Height {
			fx, fy := float64(x) / float64(worldMap.Width), float64(y) / float64(worldMap.Height)
			noiseValue, err := n.Get(fx, fy)

			center := rl.Vector2{
				X: float32(worldMap.Width) / 2.0,
				Y: float32(worldMap.Height) / 2.0,
			}

			distance := rl.Vector2Distance(center, rl.Vector2{
				X: float32(x),
				Y: float32(y),
			})

			if distance > FALL_OFF_RADIUS {
				fallOffAmt := FALL_OFF_SPEED * (distance - FALL_OFF_RADIUS)
				
				noiseValue -= float64(fallOffAmt)
				noiseValue = max(noiseValue, 0)
			}
			
			if err != nil {
				return err
			}

			prevThreshold := 0.0
			for _, threshold := range orderedThresholds {
				if noiseValue <= threshold {
					tileType := TileSet[threshold]

					adjustedNoiseValue := noiseValue - prevThreshold 
					adjustedEndThreshold := threshold - prevThreshold
					color := rl.ColorLerp(
						tileType.StartColor,
						tileType.EndColor,
						float32(adjustedNoiseValue) / float32(adjustedEndThreshold,
					))
					worldMap.Grid[x][y] = color
					break
				}
			}
		}
	}

	return nil
}

func (worldMap *Map) DrawMap(from rl.Vector2, tileSize rl.Vector2) error {	
	if tileSize.X <= 0 || tileSize.Y <= 0 {
		return fmt.Errorf("Expected a tile size greater than 0.")
	}

	for x := range worldMap.Width {
		for y := range worldMap.Height {
			color := worldMap.Grid[x][y]
			
			pos := rl.Vector2Add(from, rl.Vector2{
				X: float32(x) * float32(tileSize.X),
				Y: float32(y) * float32(tileSize.Y),
			})
			
			rl.DrawRectangleV(pos, tileSize, color)
		}
	}

	return nil
}

