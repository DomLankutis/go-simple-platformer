package objects

type Vector2D struct {
	x float64
	y float64
}

// Returns a new vector which is the result of the two vectors
func (v Vector2D) add(d Vector2D) Vector2D {
	return Vector2D{v.x + d.x, v.y + d.y}
}

// Returns the result of dividing the vector by another vector
func (v Vector2D) div(d Vector2D) Vector2D {
	return Vector2D{v.x / d.x, v.y / d.y}
}

func (v Vector2D) sub(d Vector2D) Vector2D {
	return Vector2D{v.x - d.x, v.y - d.y}
}

func (v Vector2D) mul(d Vector2D) Vector2D {
	return Vector2D{v.x * d.x , v.y * d.y}
}