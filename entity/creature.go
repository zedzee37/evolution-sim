package entity

import (
	"evolution-sim/genetics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Position rl.Vector2
	StatsPool genetics.Stats
	genes []genetics.Gene
}

func NewCreature(position rl.Vector2) *Creature {
	creature := new(Creature)
	creature.Position = position
	creature.genes = make([]genetics.Gene, 0)
	return creature
}

func (creature *Creature) ApplyDefaultGenes() {
}

func (creature *Creature) ApplyRandomGenes(n int) {
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
