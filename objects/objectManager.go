package objects

import (
	"github.com/hajimehoshi/ebiten"
)

type Displayer interface {
	Display(layer *ebiten.Image)
}

type Mover interface {
	Move()
}

type Collider interface {
	Collide(object Object) (bool, Vector2D)
	GetObject() *Object
}

type ObjectManager struct {
	Objects 			[]interface{}
}

func NewObjectManager() *ObjectManager {
	o := ObjectManager{make([]interface{}, 0)}
	return &o
}

func (manager *ObjectManager) AddObject(object interface{}) {
	manager.Objects = append(manager.Objects, object)
}

func (manager *ObjectManager) collideBetweenManagers(collidableManager *ObjectManager) (bool, []Vector2D, []int) {
	var collidingObjects []int
	var penetrationVectors []Vector2D
	state := false
	for _, objectOfCollidable := range collidableManager.Objects{
		newState, newpenetrationVectors, newCollidingObjects := manager.collideBetweenObjects(objectOfCollidable.(Collider).GetObject())
		if !state {
			state = newState
			collidingObjects = append(collidingObjects, newCollidingObjects...)
			penetrationVectors = append(penetrationVectors, newpenetrationVectors...)
		}
	}
	return state && penetrationVectors[0].IsZero(), penetrationVectors, collidingObjects
}

func (manager *ObjectManager) collideBetweenObjects(object *Object) (bool, []Vector2D, []int) {
	var collidingObjects []int
	var penetrationVectors []Vector2D
	state := false
	for index, objectOfManager := range manager.Objects {
		newState, penetrationVector := object.Collide(*objectOfManager.(Collider).GetObject())
		if !state {
			state = newState
			collidingObjects = append(collidingObjects, index)
			penetrationVectors = append(penetrationVectors, penetrationVector)
		}
	}
	return state && !penetrationVectors[0].IsZero(), penetrationVectors, collidingObjects
}

func (manager *ObjectManager) Collide(o interface{}) (bool, []Vector2D, []int) {
	switch o.(type) {
	case *ObjectManager:
		return manager.collideBetweenManagers(o.(*ObjectManager))
	default:
		return manager.collideBetweenObjects(o.(Collider).GetObject())
	}
	return false, nil, nil
}

func (manger *ObjectManager) Move() {
	for _, objectOfManager := range manger.Objects {
		objectOfManager.(Mover).Move()
	}
}

func (manager *ObjectManager) Display(layer *ebiten.Image) {
	for _, objectOfManager := range manager.Objects {
		objectOfManager.(Displayer).Display(layer)
	}
}