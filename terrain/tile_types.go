package terrain

import rl "github.com/gen2brain/raylib-go/raylib"

type TileType struct {
	StartColor rl.Color
	EndColor rl.Color
	Threshold float64
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
	EndColor: rl.DarkGreen,
}
var Mountain = TileType{
	StartColor: rl.DarkGray,
	EndColor: rl.White,
}

var TileSet = map[float64]TileType{
	0.3: Water,
	0.4: Sand,
	0.6: Grass,
	1.0: Mountain,
}

