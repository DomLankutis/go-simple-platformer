package objects

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	Npc
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

	p.Velocity.x = p.ApplyVelocity(toBeVelocity.x, p.Velocity.x)
	p.Velocity.y = p.ApplyVelocity(toBeVelocity.y, p.Velocity.y)

	p.ApplyResistance()
	p.Opts.GeoM.Translate(p.Velocity.x, p.Velocity.y)

}

func NewPlayer(sprite *ebiten.Image, err error) *Player {
	p := &Player{}
	p.Sprite = sprite
	p.init()
	return p
}