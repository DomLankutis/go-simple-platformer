package objects

import (
	"math"
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

//type CollidableInterface interface {
//	GetPosition() Vector2D
//	GetVelocity() Vector2D
//	GetSprite() *ebiten.Image
//	ApplyPosition(x, y float64)
//}
//
//type StaticShape struct {
//	Collision Collidable
//	Position  Vector2D
//	Sprite    *ebiten.Image
//	Opts      *ebiten.DrawImageOptions
//}
//
//func CreateStaticShape(x, y float64, sprite *ebiten.Image) *StaticShape {
//	s := &StaticShape{}
//	s.Sprite = sprite
//	s.Position.x = x
//	s.Position.y = y
//	s.Opts = &ebiten.DrawImageOptions{}
//	s.Opts.GeoM.Translate(x, y)
//	s.Collision.setStaticCollisionVector(s)
//	return s
//}
//
//func (s *StaticShape) Display(layer *ebiten.Image) {
//	s.Sprite.Fill(color.NRGBA{0xff, 0x0a, 0xa0,0xff})
//	layer.DrawImage(s.Sprite, s.Opts)
//}
//
//type Collidable struct {
//	VectorMin Vector2D
//	VectorMax Vector2D
//	Velocity Vector2D
//}
//
//func (c *Collidable) setStaticCollisionVector(s *StaticShape) {
//	maxX, maxY := s.Sprite.Size()
//	c.VectorMax = Vector2D{s.Position.x + float64(maxX), s.Position.y + float64(maxY)}
//	c.VectorMin = Vector2D{s.Position.x, s.Position.y}
//}
//
//func (c *Collidable) setCollisionVector(p CollidableInterface) {
//	positionVector := p.GetPosition()
//	maxX, maxY := p.GetSprite().Size()
//	c.VectorMax = Vector2D{positionVector.x + float64(maxX), positionVector.y + float64(maxY)}
//	c.VectorMin = Vector2D{positionVector.x, positionVector.y}
//}
//
////func (p *Collidable) Overlaping(obj *Collidable) (Vector2D, float64) {
////
////
////}
////
////
////
////func (c *Npc) Colliding(obj *StaticShape) {
////
////
////}

// Aligned Axis Bounding Box
type AABB struct {
	center 							Vector2D
	extents 						Vector2D
}

func (a AABB) GetMin() Vector2D{
	return a.center.sub(a.extents)
}

func (a AABB) GetMax() Vector2D{
	return a.center.add(a.extents)
}

func (a AABB) GetSize() Vector2D{
	return a.extents.mul(Vector2D{2, 2 })
}

func newAABB(center, extents Vector2D) *AABB {
	a := new(AABB)
	a.center = center
	a.extents = extents
	return a
}

func (a *AABB) Draw(layer *ebiten.Image) {
	img, _ := ebiten.NewImage(int(a.GetSize().x), int(a.GetSize().y), ebiten.FilterNearest)
	temp := ebiten.DrawImageOptions{}
	img.Fill(color.White)
	temp.GeoM.Translate(a.GetMin().x + a.GetSize().x, a.GetMin().y + a.GetSize().y)
	layer.DrawImage(img, &temp)
}

func (a *AABB) minkowskiDifference(obj *AABB) *AABB {
	topLeft := a.GetMin().sub(obj.GetMax())
	fullSize := a.GetSize().add(obj.GetSize())

	return newAABB(topLeft.add(fullSize.div(Vector2D{2, 2})), fullSize.div(Vector2D{2, 2}))
}

func (a *AABB) closestPointOnBoundsToPoint(point Vector2D) Vector2D{
	max := a.GetMax()
	min := a.GetMin()
	minDist := math.Abs(point.x - min.x)
	boundsPoint := Vector2D{min.x, point.y}

	if math.Abs(max.x- point.x) < minDist {
		minDist = math.Abs(max.x - point.x)
		boundsPoint = Vector2D{max.x, point.y}
	}

	if math.Abs(max.y- point.y) < minDist {
		minDist = math.Abs(max.y - point.y)
		boundsPoint = Vector2D{point.x, max.y}
	}

	if math.Abs(min.y- point.y) < minDist {
		minDist = math.Abs(min.y - point.y)
		boundsPoint = Vector2D{point.x, min.x}
	}

	return boundsPoint
}

func (o *Object) Collide(obj Object, layer *ebiten.Image){
	diff := obj.CollisionBox.minkowskiDifference(o.CollisionBox)
	diff.Draw(layer)
	//o.CollisionBox.Draw(layer)
	if diff.GetMin().x <= 0 && diff.GetMax().x >= 0 &&
		diff.GetMin().y <= 0 && diff.GetMax().y >= 0 {
			o.Colour = color.NRGBA{0xff, 0x00, 0x0a, 0xff}
	}
}