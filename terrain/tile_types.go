package terrain

import rl "github.com/gen2brain/raylib-go/raylib"

type TileType struct {
	StartColor rl.Color
	EndColor rl.Color
}

var Water = TileType{
	StartColor: rl.DarkBlue,
	EndColor: rl.SkyBlue,
}
var Sand = TileType{
	StartColor: rl.Yellow,
	EndColor: rl.Yellow,
}
var Grass = TileType{
	StartColor: rl.Lime,
	EndColor: rl.Green,
}
var Mountain = TileType{
	StartColor: rl.DarkGray,
	EndColor: rl.DarkGray,
}

var TileSet = map[float64]TileType{
	0.3: Water,
	0.4: Sand,
	0.6: Grass,
	1.0: Mountain,
}

