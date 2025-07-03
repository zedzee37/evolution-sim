package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Creature struct {
	Position rl.Vector2
}

func NewCreature(position rl.Vector2) *Creature {
	creature := new(Creature)
	creature.Position = position
	return creature
}

func (creature *Creature) Draw(offsetX int32, offsetY int32) {
	drawX := offsetX + int32(creature.Position.X)
	drawY := offsetY + int32(creature.Position.Y)
	rl.DrawCircle(drawX, drawY, 5.0, rl.Red)
}

func (_ *Creature) Load() {
}

func (_ *Creature) Tick(dt float64) {
}
