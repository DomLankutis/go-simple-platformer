package objects

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type StaticShape struct {
	Collision Collidable
	Position  Vector2D
	Sprite    *ebiten.Image
	Opts      *ebiten.DrawImageOptions
}

func CreateStaticShape(x, y float64, sprite *ebiten.Image) *StaticShape {
	s := &StaticShape{}
	s.Sprite = sprite
	s.Position.x = x
	s.Position.y = y
	s.Opts = &ebiten.DrawImageOptions{}
	s.Opts.GeoM.Translate(x, y)
	s.Collision.setStaticCollisionVector(s)
	return s
}

func (s *StaticShape) Display(layer *ebiten.Image) {
	s.Sprite.Fill(color.NRGBA{0xff, 0x0a, 0xa0,0xff})
	layer.DrawImage(s.Sprite, s.Opts)
}

type Collidable struct {
	VectorMin Vector2D
	VectorMax Vector2D
}

func (c *Collidable) setStaticCollisionVector(s *StaticShape) {
	maxX, maxY := s.Sprite.Size()
	c.VectorMax = Vector2D{s.Position.x + float64(maxX), s.Position.y + float64(maxY)}
	c.VectorMin = Vector2D{s.Position.x, s.Position.y}
}

func (c *Collidable) setCollisionVector(p *Player) {
	maxX, maxY := p.Sprite.Size()
	c.VectorMax = Vector2D{p.Position.x + float64(maxX), p.Position.y + float64(maxY)}
	c.VectorMin = Vector2D{p.Position.x, p.Position.y}
}

func (c *Collidable) IsColliding(obj *Collidable) bool{

	if c.VectorMax.x < obj.VectorMin.x  || c.VectorMin.x > obj.VectorMax.x{
		return false
	}
	if c.VectorMax.y < obj.VectorMin.y  || c.VectorMin.y > obj.VectorMax.y{
		return false
	}

	return true
}
