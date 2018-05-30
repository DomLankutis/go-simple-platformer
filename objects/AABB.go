package objects

import (
	"math"
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"fmt"
)

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
	img, _ := ebiten.NewImage(int(math.Abs(a.GetSize().x)), int(math.Abs(a.GetSize().y)), ebiten.FilterNearest)
	temp := ebiten.DrawImageOptions{}
	img.Fill(color.NRGBA{0x1f, 0xff, 0x0b, 0x5f})
	temp.GeoM.Translate(a.GetMin().x, a.GetMin().y)
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

	if math.Abs(max.x - point.x) < minDist {
		minDist = math.Abs(max.x - point.x)
		boundsPoint = Vector2D{max.x, point.y}
	}

	if math.Abs(max.y - point.y) < minDist {
		minDist = math.Abs(max.y - point.y)
		boundsPoint = Vector2D{point.x, max.y}
	}

	if math.Abs(min.y - point.y) < minDist {
		minDist = math.Abs(min.y - point.y)
		boundsPoint = Vector2D{point.x, min.y}
	}

	fmt.Println(boundsPoint)
	return boundsPoint
}

func (o *Object) Collide(obj Object){
	diff := obj.CollisionBox.minkowskiDifference(o.CollisionBox)
	if diff.GetMin().x <= 0 && diff.GetMax().x >= 0 &&
		diff.GetMin().y <= 0 && diff.GetMax().y >= 0 {
			o.Colour = color.NRGBA{0xff, 0x00, 0x0a, 0xff}
			o.SetPosition(diff.closestPointOnBoundsToPoint(Vector2D{0,0}))
	}
}