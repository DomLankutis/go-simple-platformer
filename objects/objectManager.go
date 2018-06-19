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
	Collide(object Object) bool
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

func (manager *ObjectManager) collideBetweenManagers(collidableManager *ObjectManager) (bool, []int) {
	var collidingObjects []int
	state := false
	for _, objectOfCollidable := range collidableManager.Objects{
		newState, newCollidingObjects := manager.collideBetweenObjects(objectOfCollidable.(Collider).GetObject())
		if !state {
			state = newState
			collidingObjects = append(collidingObjects, newCollidingObjects...)
		}
	}
	return state, collidingObjects
}

func (manager *ObjectManager) collideBetweenObjects(object *Object) (bool, []int) {
	var collidingObjects []int
	state := false
	for index, objectOfManager := range manager.Objects {
		newState := object.Collide(*objectOfManager.(Collider).GetObject())
		if !state {
			state = newState
			collidingObjects = append(collidingObjects, index)
		}
	}
	return state, collidingObjects
}

func (manager *ObjectManager) Collide(o interface{}) (bool, []int) {
	switch o.(type) {
	case *ObjectManager:
		return manager.collideBetweenManagers(o.(*ObjectManager))
	default:
		return manager.collideBetweenObjects(o.(Collider).GetObject())
	}
	return false, nil
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