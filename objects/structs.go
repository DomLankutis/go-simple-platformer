package objects

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v Vector2D) getLength() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y))
}

func (v Vector2D) getNormalised() Vector2D {
	length := v.getLength()
	if length != 0 {
		return Vector2D{v.x / length, v.y / length}
	}
	return Vector2D{0,0}
}

func (v Vector2D) getTangent() Vector2D {
	return Vector2D{-v.y, v.x}
}

func (v Vector2D) add(d Vector2D) Vector2D {
	return Vector2D{v.x + d.x, v.y + d.y}
}

func (v Vector2D) div(d Vector2D) Vector2D {
	return Vector2D{v.x / d.x, v.y / d.y}
}

func (v Vector2D) sub(d Vector2D) Vector2D {
	return Vector2D{v.x - d.x, v.y - d.y}
}

func (v Vector2D) mul(d Vector2D) Vector2D {
	return Vector2D{v.x * d.x , v.y * d.y}
}

func (v Vector2D) isZero() bool {
	return v.x == 0 && v.y == 0
}

func (v Vector2D) mulToFloat(d Vector2D) float64 {
	return v.x * d.x - v.y * d.y
}

func (v Vector2D) dotProduct(d Vector2D) float64 {
	return v.x * d.x + v.y * d.y
}
