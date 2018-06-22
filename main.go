package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"./objects"
	"fmt"
)

var world *objects.World

var player objects.Player
var collidables *objects.ObjectManager
var collectables *objects.ObjectManager
var enemies *objects.ObjectManager

func update(screen *ebiten.Image) error {

	player.Move()
	enemies.Move()

	if statement, _, _:= collectables.Collide(player.GetObject()); statement {
		player.SetPosition(objects.Vector2D{50,50})
	}
	collidables.Collide(player.GetObject())
	collidables.Collide(enemies)
	enemies.Collide(player.GetObject())

	collidables.Display(screen)
	collectables.Display(screen)
	enemies.Display(screen)
	player.Display(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Position: %0.f", player.GetPosition()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nVelocity: %.0f", player.Velocity))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nFPS: %.1f", ebiten.CurrentFPS()))

	return nil
}

func main() {

	width, height := 640, 360
	world = objects.NewWorld(objects.Vector2D{float64(width * 5), float64(height * 5)}, objects.Vector2D{float64(width), float64(height)})

	collidables = objects.NewObjectManager()
	collectables = objects.NewObjectManager()
	enemies = objects.NewObjectManager()

	objImage, _ := ebiten.NewImage(640, 80, ebiten.FilterNearest)
	objImage1, _ := ebiten.NewImage(10, 10, ebiten.FilterNearest)
	playerImage, _ := ebiten.NewImage(50, 50, ebiten.FilterNearest)

	collidables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff,0xaf, 0xff}, -200, 80, objImage, world))
	collidables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff, 0xaf, 0xff}, 640 / 2 +90, 220, objImage, world))
	collidables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff, 0xaf, 0xff}, 0, 320, objImage, world))
	collectables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff, 0xaf, 0xff}, 50, 280, objImage1, world))

	enemies.AddObject(objects.NewEnemy(objects.Vector2D{-200, 80}, objects.Vector2D{-200+640, 80}, playerImage, world))

	player = *objects.NewPlayer(playerImage, nil, world)

	if err := ebiten.Run(update, width, height, 3, "Game"); err != nil {
		fmt.Println(err)
	}
}
