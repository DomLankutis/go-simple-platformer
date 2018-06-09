package objects

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v Vector2D) getLength() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

func (v Vector2D) getNormalised() Vector2D {
	length := v.getLength()
	if length != 0 {
		return Vector2D{v.X / length, v.Y / length}
	}
	return Vector2D{0,0}
}

func (v Vector2D) getTangent() Vector2D {
	return Vector2D{-v.Y, v.X}
}

func (v Vector2D) add(d Vector2D) Vector2D {
	return Vector2D{v.X + d.X, v.Y + d.Y}
}

func (v Vector2D) div(d Vector2D) Vector2D {
	return Vector2D{v.X / d.X, v.Y / d.Y}
}

func (v Vector2D) sub(d Vector2D) Vector2D {
	return Vector2D{v.X - d.X, v.Y - d.Y}
}

func (v Vector2D) mul(d Vector2D) Vector2D {
	return Vector2D{v.X * d.X, v.Y * d.Y}
}

func (v Vector2D) isZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vector2D) mulToFloat(d Vector2D) float64 {
	return v.X* d.X - v.Y* d.Y
}

func (v Vector2D) dotProduct(d Vector2D) float64 {
	return (v.X * d.X) + (v.Y * d.Y)
}
