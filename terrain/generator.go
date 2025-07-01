package terrain

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zedzee37/znoise/noise"
)

type Map struct {
	Grid [][]*TileType	
	Width int
	Height int
}

func NewMap(width int, height int) *Map {
	grid := make([][]*TileType, width)
	
	for row := range width {
		grid[row] = make([]*TileType, height)
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
	for x := range worldMap.Width {
		for y := range worldMap.Height {
			fx, fy := float64(x) / float64(worldMap.Width), float64(y) / float64(worldMap.Height)
			noiseValue, err := n.Get(fx, fy)
			
			if err != nil {
				return err
			}

			for threshold, tileType := range TileSet {
				if noiseValue <= threshold {
					worldMap.Grid[x][y] = &tileType
				}
			}
		}
	}

	return nil
}

func (worldMap *Map) DrawMap(from rl.Vector2, tileSize rl.Vector2) {
}

