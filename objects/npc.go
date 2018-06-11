package objects

import (
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

type Npc struct {
	Object
	Speed, Gravity, Resistance, JumpForce 	float64
}

func (n *Npc) init() {
	n.Speed = 0.3
	n.Gravity = 1
	n.JumpForce = n.Gravity * 10
	n.Resistance = 0.1
	n.MaxVelocity = 5
	n.Colour = color.NRGBA{0x00, 0xfa, 0x9f, 0xff}
	n.Opts = &ebiten.DrawImageOptions{}
	n.Opts.GeoM.Translate(10, 10)
	w, h := n.Sprite.Size()
	n.Size = Vector2D{float64(w), float64(h)}
	n.CollisionBox = newAABB(Vector2D{0,0}, n.Size.div(Vector2D{2, 2}))
}

func (n *Npc) ApplyResistance() {
	if n.Velocity.X < 0 {
		if n.Velocity.X > -n.Speed {
			n.Velocity.X = 0
		} else {
			n.Velocity.X += n.Resistance
		}
	} else {
		if n.Velocity.X < n.Speed {
			n.Velocity.X = 0
		} else {
			n.Velocity.X -= n.Resistance
		}
	}
	n.Velocity.Y += n.Gravity
}

func (n *Npc) GetObject() *Object {
	return &n.Object
}

