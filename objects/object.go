package objects

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
)

type Object struct {
	world									*World
	CollisionBox							*AABB
	Sprite                                  *ebiten.Image
	Colour                                  color.NRGBA
	Opts                                    *ebiten.DrawImageOptions
	Size 									Vector2D
	Velocity                                Vector2D
	CanJump 								bool
	MaxVelocity								float64
}

func (o *Object) GetPosition() Vector2D {
	x := o.Opts.GeoM.Element(0, 2)
	y := o.Opts.GeoM.Element(1, 2)

	return Vector2D{x, y}
}

func (o *Object) SetPosition(d Vector2D) {
	currentPos := o.GetPosition()
	d1 := d.sub(currentPos)
	o.Opts.GeoM.Translate(d1.X, d1.Y)
}

func (o *Object) AppendPosition(d Vector2D) {
	newPos := o.GetPosition()
	newPos.add(d)
	o.Opts.GeoM.Translate(newPos.X, newPos.Y)
}

func (o *Object) GetSprite() *ebiten.Image {
	return o.Sprite
}

func (n *Object) GetVelocity() Vector2D {
	return n.Velocity
}

func (o *Object) ApplyDirectVelocity(d Vector2D) {
	o.Opts.GeoM.Translate(d.X, d.Y)
}

func (n *Object) ApplyVelocity(toBe, current, limit float64) float64 {
	if math.Abs(toBe + current) < limit {
		current += toBe
	} else {
		if toBe < 0 {
			current = -limit
		} else {
			current = limit
		}
	}
	return current
}

func (o *Object) Display(layer *ebiten.Image) {
	o.CollisionBox.center = o.GetPosition().add(o.Size.div(Vector2D{2, 2}))
	o.Sprite.Fill(o.Colour)
	cameraOpts := *o.Opts
	cameraOpts.GeoM.Translate(-o.world.ViewportPosition.X, -o.world.ViewportPosition.Y)
	layer.DrawImage(o.Sprite, &cameraOpts)
}

func (n *Object) GetObject() *Object {
	return n
}

func NewObject(colour color.NRGBA, posx, posy float64, sprite *ebiten.Image, world *World) *Object{
	o := &Object{}
	o.Sprite = sprite
	o.Colour = colour
	o.Opts = &ebiten.DrawImageOptions{}
	pos := Vector2D{posx, posy}
	o.Opts.GeoM.Translate(posx, posy)
	w, h := o.Sprite.Size()
	o.Size = Vector2D{float64(w), float64(h)}
	o.CollisionBox = newAABB(pos.add(o.Size.div(Vector2D{2, 2})), o.Size.div(Vector2D{2, 2}))
	o.world = world
	return o
}