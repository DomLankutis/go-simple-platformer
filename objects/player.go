package objects

import (
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Player struct {
	Npc
}

func applyVelocity(toBe, current, max float64) float64 {
	if math.Abs(toBe + current) < max {
		current += toBe
	} else {
		if toBe < 0 {
			current = -max
		} else {
			current = max
		}
	}
	return current
}

func (p *Player) Move(layer *ebiten.Image){
	toBeVelocity := Vector2D{0,0};
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		toBeVelocity.y = -p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		toBeVelocity.y = p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		toBeVelocity.x = -p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		toBeVelocity.x = p.Speed
	}

	p.Velocity.x = applyVelocity(toBeVelocity.x, p.Velocity.x, p.MaxVelocity)
	p.Velocity.y = applyVelocity(toBeVelocity.y, p.Velocity.y, p.MaxVelocity)

	p.ApplyResistance()
}

func NewPlayer(sprite *ebiten.Image, err error) *Player {
	p := &Player{}
	p.Sprite = sprite
	p.init()
	return p
}