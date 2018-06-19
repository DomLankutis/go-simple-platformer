package objects

import (
	"github.com/hajimehoshi/ebiten"
)

type Enemy struct {
	Npc
	PathStart				Vector2D
	PathEnd 				Vector2D
	Direction				float64
	changeDir				bool
}

func NewEnemy(pathStart, pathEnd Vector2D, sprite *ebiten.Image, world *World) *Enemy {
	e := Enemy{PathStart: pathStart, PathEnd: pathEnd}
	e.Sprite = sprite
	e.world = world
	startpos := pathEnd.add(pathStart).div(Vector2D{2, 2})
	e.Npc.init()
	e.SetPosition(startpos)
	e.Direction = 1
	e.Speed = 2
	e.MaxVelocity = 3

	return &e
}

func (e *Enemy) Move(){
	if e.CollisionBox.GetMax().X > e.PathEnd.X {
		e.Direction = -1
	}else if e.CollisionBox.GetMin().X < e.PathStart.X{
		e.Direction = 1
	}

	e.Velocity.X = e.ApplyVelocity(e.Speed * e.Direction, e.Velocity.X, e.MaxVelocity)

	e.ApplyResistance()
	e.Opts.GeoM.Translate(e.Velocity.X, e.Velocity.Y)

}
