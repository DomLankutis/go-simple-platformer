package objects

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	Npc
}

func (p *Player) Move(layer *ebiten.Image){
	//toBeVelocity := Vector2D{0,0};
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.CanJump{
		p.Velocity.y = p.ApplyVelocity(-p.JumpForce, p.Velocity.y, p.JumpForce)
		p.CanJump = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Velocity.y = p.ApplyVelocity(p.Speed, p.Velocity.y, p.MaxVelocity)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Velocity.x = p.ApplyVelocity(-p.Speed, p.Velocity.x, p.MaxVelocity)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Velocity.x = p.ApplyVelocity(p.Speed, p.Velocity.x, p.MaxVelocity)
	}

	p.ApplyResistance()
	p.Opts.GeoM.Translate(p.Velocity.x, p.Velocity.y)

}

func NewPlayer(sprite *ebiten.Image, err error) *Player {
	p := &Player{}
	p.Sprite = sprite
	p.init()
	return p
}