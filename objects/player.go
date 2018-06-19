package objects

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	Npc
}

func (p *Player) Move(){
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

func (p *Player) Display(layer *ebiten.Image) {
	pos := p.GetPosition()

	if pos.X > (p.world.ViewportPosition.X + (p.world.ViewportSize.X * 0.65)) {
		p.world.ViewportPosition.X += pos.X - (p.world.ViewportPosition.X + (p.world.ViewportSize.X * 0.65))
	} else
	if pos.X < (p.world.ViewportPosition.X + (p.world.ViewportSize.X * 0.3)) {
		p.world.ViewportPosition.X -= (p.world.ViewportPosition.X + (p.world.ViewportSize.X * 0.3)) - pos.X
	}
	if pos.Y > (p.world.ViewportPosition.Y + (p.world.ViewportSize.Y * 0.5)) {
		p.world.ViewportPosition.Y += pos.Y - (p.world.ViewportPosition.Y + (p.world.ViewportSize.Y * 0.5))
	} else
	if pos.Y < (p.world.ViewportPosition.Y + (p.world.ViewportSize.Y * 0.2)) {
		p.world.ViewportPosition.Y -= (p.world.ViewportPosition.Y + (p.world.ViewportSize.Y * 0.2)) - pos.Y
	}

	p.Object.Display(layer)
}

func NewPlayer(sprite *ebiten.Image, _ error, world *World) *Player {
	p := &Player{}
	p.Sprite = sprite
	p.init()
	p.world = world
	return p
}