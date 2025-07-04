package entity

import (
	"evolution-sim/genetics"
	"evolution-sim/terrain"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Acceleration rl.Vector2
	StatsPool genetics.Stats
	worldMap *terrain.Map
	genes [genetics.GenoTypeCount]genetics.Gene
	rng *rand.Rand
}

func NewCreature(position rl.Vector2, worldMap *terrain.Map) *Creature {
	creature := new(Creature)
	creature.Position = position
	creature.Velocity = rl.Vector2Zero()
	creature.Acceleration = rl.Vector2Zero()
	creature.worldMap = worldMap
	
	source := rand.NewSource(time.Now().Unix())
	creature.rng = rand.New(source)

	return creature
}

func (creature *Creature) Draw(offsetX int32, offsetY int32) {
	drawX := offsetX + int32(creature.Position.X)
	drawY := offsetY + int32(creature.Position.Y)
	rl.DrawCircle(drawX, drawY, 5.0, rl.Red)
}

func (_ *Creature) Load() {
}

func (creature *Creature) getSurroundingTiles() map[rl.Vector2]*terrain.TileType {
	directionMap := make(map[rl.Vector2]*terrain.TileType)

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			pos := rl.Vector2Add(creature.Position, rl.Vector2{
				X: float32(x),
				Y: float32(y),
			})	

			tileType, err := creature.worldMap.GetTile(int(pos.X), int(pos.Y))
			
			if err != nil {
				continue
			}

			direction := rl.Vector2Normalize(rl.Vector2Subtract(pos, creature.Position))
			directionMap[direction] = tileType
		}
	}

	return directionMap
}

func (creature *Creature) Tick(dt float64) {
	creature.Position = rl.Vector2Add(creature.Position, creature.Velocity)
	creature.Velocity = rl.Vector2Add(creature.Velocity, creature.Acceleration)

	surroundingTiles := creature.getSurroundingTiles()

	if len(surroundingTiles) <= 0 {
		return
	}

	possibleDirections := make([]rl.Vector2, 0)
	for direction, tileType := range surroundingTiles {
		if tileType != &terrain.Grass {
			continue
		}
		
		possibleDirections = append(possibleDirections, direction)
	}

	directionCount := len(possibleDirections)
	if directionCount == 0 {
		return
	}
	
	directionIndex := creature.rng.Int() % (directionCount - 1)
	direction := possibleDirections[directionIndex]
	creature.Velocity = rl.Vector2Scale(direction, 2)
}
