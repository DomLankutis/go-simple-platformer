package objects

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	Npc
}

func (p *Player) Move(){
	//toBeVelocity := Vector2D{0,0};
	if ebiten.IsKeyPressed(ebiten.KeyW) && p.CanJump{
		p.Velocity.Y = p.ApplyVelocity(-p.JumpForce, p.Velocity.Y, p.JumpForce)
		p.CanJump = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Velocity.Y = p.ApplyVelocity(p.Speed, p.Velocity.Y, p.MaxVelocity)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Velocity.X = p.ApplyVelocity(-p.Speed, p.Velocity.X, p.MaxVelocity)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Velocity.X = p.ApplyVelocity(p.Speed, p.Velocity.X, p.MaxVelocity)
	}

	p.ApplyResistance()
	p.Opts.GeoM.Translate(p.Velocity.X, p.Velocity.Y)

}

func NewPlayer(sprite *ebiten.Image, err error) *Player {
	p := &Player{}
	p.Sprite = sprite
	p.init()
	return p
}