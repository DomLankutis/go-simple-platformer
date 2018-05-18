package objects

import (
	"github.com/hajimehoshi/ebiten"
	"math"
	"image/color"
)

type Player struct {
	Collidable
	Velocity                                Vector2D
	Position                                Vector2D
	PlayerColour							color.NRGBA
	Speed, MaxVelocity, Gravity, Resistance float64
	Sprite                                  *ebiten.Image
	Opts                                    *ebiten.DrawImageOptions
}

func NewPlayer(sprite *ebiten.Image, err error) *Player {
	p := &Player{}
	p.Sprite = sprite
	p.init()
	return p

}

func (p *Player) init() {
	p.Speed = 0.3
	p.Gravity = 0.1
	p.Resistance = 0.1
	p.MaxVelocity = 5
	p.Position.x = 50
	p.Position.y = 50
	p.PlayerColour = color.NRGBA{0x00, 0xfa, 0x9f, 0xff}
	p.Opts = &ebiten.DrawImageOptions{}
	p.Opts.GeoM.Translate(p.Position.x, p.Position.y)
	p.Collidable.setCollisionVector(p)
}

func (p *Player) ApplyResistance() {
	if p.Velocity.x < 0 {
		if p.Velocity.x > -p.Speed {
			p.Velocity.x = 0
		} else {
			p.Velocity.x += p.Resistance
		}
	} else {
		if p.Velocity.x < p.Speed {
			p.Velocity.x = 0
		} else {
			p.Velocity.x -= p.Resistance
		}
	}
	if p.Velocity.y < 0 {
		if p.Velocity.y > -p.Speed {
			p.Velocity.y = 0
		} else {
			p.Velocity.y += p.Gravity
		}
	} else {
		if p.Velocity.y < p.Speed {
			p.Velocity.y = 0
		} else {
			p.Velocity.y -= p.Gravity
		}
	}
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

func (p *Player) SetPosition() {
	p.Position.x += p.Velocity.x
	p.Position.y += p.Velocity.y
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
	p.SetPosition()
	p.setCollisionVector(p)
}

func (p *Player) Display(layer *ebiten.Image) {
	p.Sprite.Fill(p.PlayerColour)
	p.Opts.GeoM.Translate(p.Velocity.x, p.Velocity.y)
	layer.DrawImage(p.Sprite, p.Opts)
}
