package objects


type World struct {
	FullSize			Vector2D
	ViewportSize		Vector2D
	ViewportPosition	Vector2D
}

func NewWorld(fullSize, viewportSize Vector2D) *World {
	return &World{fullSize, viewportSize, Vector2D{0,0} }
}
