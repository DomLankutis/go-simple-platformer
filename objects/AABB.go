package objects

import (
	"math"
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

// Aligned Axis Bounding Box
type AABB struct {
	center 							Vector2D
	extents 						Vector2D
	velocity						Vector2D
	acceleration					Vector2D
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

func newAABB(center, extents Vector2D, physics ...Vector2D) *AABB {
	a := new(AABB)
	a.center = center
	a.extents = extents
	if len(physics) > 0 {
		a.velocity = physics[0]
		a.acceleration = physics[1]
	}
	return a
}

func (a *AABB) Draw(layer *ebiten.Image) {
	img, _ := ebiten.NewImage(int(math.Abs(a.GetSize().X)), int(math.Abs(a.GetSize().Y)), ebiten.FilterNearest)
	temp := ebiten.DrawImageOptions{}
	img.Fill(color.NRGBA{0xaf, 0x00, 0x0b, 0x5f})
	temp.GeoM.Translate(a.GetMin().X, a.GetMin().Y)
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

	minDist := math.Abs(point.X - min.X)
	boundsPoint := Vector2D{min.X, point.Y}

	if math.Abs(max.X - point.X) < minDist {
		minDist = math.Abs(max.X - point.X)
		boundsPoint = Vector2D{max.X, point.Y}
	}

	if math.Abs(max.Y - point.Y) < minDist {
		minDist = math.Abs(max.Y - point.Y)
		boundsPoint = Vector2D{point.X, max.Y}
	}

	if math.Abs(min.Y - point.Y) < minDist {
		minDist = math.Abs(min.Y - point.Y)
		boundsPoint = Vector2D{point.X, min.Y}
	}

	return boundsPoint
}

func (a *AABB) getRayIntersectionFractionOfFirstRay(originA, endA, originB, endB Vector2D) float64 {
	r := endA.sub(originA)
	s := endB.sub(originB)

	numerator := (originB.sub(originA)).mulToFloat(r)
	denominator := r.mulToFloat(s)

	if numerator == 0 && denominator == 0 {
		return math.Inf(1)
	}
	if denominator == 0 {
		return math.Inf(1)
	}

	u := numerator / denominator
	t := ((originB.sub(originA)).mulToFloat(s)) / denominator

	if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
		return t
	}
	return math.Inf(1)
}

func (a *AABB) getRayIntersectionFraction(origin, direction Vector2D) float64 {
	end := origin.add(direction)

	min := a.GetMin()
	max := a.GetMax()
	minT := a.getRayIntersectionFractionOfFirstRay(origin, end, min, max)
	x := a.getRayIntersectionFractionOfFirstRay(origin, end, Vector2D{min.X, max.Y}, Vector2D{max.X, max.Y})
	if x < minT {
		minT = x
	}
	x = a.getRayIntersectionFractionOfFirstRay(origin, end, Vector2D{max.X, max.Y}, Vector2D{max.X, min.Y})
	if x < minT {
		minT = x
	}
	x = a.getRayIntersectionFractionOfFirstRay(origin, end, Vector2D{max.X, max.Y}, Vector2D{max.X, min.Y})
	if x < minT {
		minT = x
	}

	return minT
}

func (o *Object) IsOverlapping(obj Object) (bool, *AABB){
	diff := obj.CollisionBox.minkowskiDifference(o.CollisionBox)
	return diff.GetMin().X <= 0 && diff.GetMax().X >= 0 && diff.GetMin().Y <= 0 && diff.GetMax().Y >= 0, diff
}

func (o *Object) Collide(obj Object) bool {

	overlapping, diff := o.IsOverlapping(obj)

	if overlapping {
		penetrationVector := diff.closestPointOnBoundsToPoint(Vector2D{0, 0})

		o.ApplyDirectVelocity(penetrationVector)

		if !penetrationVector.isZero() {
			tangent := (penetrationVector.getNormalised()).getTangent()

			oDotProduct := o.Velocity.dotProduct(tangent)
			objDotProduct := obj.Velocity.dotProduct(tangent)

			o.Velocity = Vector2D{oDotProduct, oDotProduct}.mul(tangent)
			obj.Velocity = Vector2D{objDotProduct, objDotProduct}.mul(tangent)

			if penetrationVector.Y < 0 || penetrationVector.X != 0 {
				o.CanJump = true
			}
			return true
		}
	}else {
		relativeMotion := o.Velocity.sub(obj.Velocity)
		h := diff.getRayIntersectionFraction(Vector2D{0, 0}, relativeMotion)

		if h < math.Inf(1) {

			o.ApplyDirectVelocity(o.Velocity.mul(Vector2D{h,h}))
			obj.ApplyDirectVelocity(obj.Velocity.mul(Vector2D{h,h}))


			tangent := (relativeMotion.getNormalised()).getTangent()

			oDotProduct := o.Velocity.dotProduct(tangent)
			objDotProduct := obj.Velocity.dotProduct(tangent)

			o.ApplyDirectVelocity(Vector2D{oDotProduct, oDotProduct}.mul(tangent))
			obj.ApplyDirectVelocity(Vector2D{objDotProduct, objDotProduct}.mul(tangent))
			o.CanJump = true

			return true
		}
	}
	return false
}