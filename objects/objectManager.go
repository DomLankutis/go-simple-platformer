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

func (manager *ObjectManager) collideBetweenManagers(collidableManager *ObjectManager) bool {
	var newState bool
	state := false
	for _, objectOfCollideble := range collidableManager.Objects{
		if !state {
			newState = manager.collideBetweenObjects(objectOfCollideble.(Collider).GetObject())
			state = newState
		}
	}
	return state
}

func (manager *ObjectManager) collideBetweenObjects(object *Object) bool {
	var newState bool
	state := false
	for _, objectOfManager := range manager.Objects {
		if !state {
			newState = object.Collide(*objectOfManager.(*Object))
			state = newState
		}
	}
	return state
}

func (manager *ObjectManager) Collide(o interface{}) bool {
	switch o.(type) {
	case *ObjectManager:
		return manager.collideBetweenManagers(o.(*ObjectManager))
	default:
		return manager.collideBetweenObjects(o.(*Object))
	}
	return false
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