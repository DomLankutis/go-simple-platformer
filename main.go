package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"./objects"
	"fmt"
)

var player objects.Player
var collidables *objects.ObjectManager
var collectables *objects.ObjectManager

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0xa0, 0x01, 0xfa, 0xff})

	player.Move()
	collidables.Collide(player.GetObject())
	collidables.Display(screen)

	if collectables.Collide(player.GetObject()) {
		player.SetPosition(objects.Vector2D{50,50})
	}
	collectables.Display(screen)
	collectables.Collide(player.GetObject())

	player.Display(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Position: %0.f", player.GetPosition()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nVelocity: %.0f", player.Velocity))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nFPS: %.1f", ebiten.CurrentFPS()))

	return nil
}

func main() {
	collidables = objects.NewObjectManager()
	collectables = objects.NewObjectManager()


	objImage, _ := ebiten.NewImage(640, 80, ebiten.FilterNearest)
	objImage1, _ := ebiten.NewImage(10, 10, ebiten.FilterNearest)
	collidables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff,0xaf, 0xff}, -200, 80, objImage))
	collidables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff, 0xaf, 0xff}, 640 / 2 +90, 260, objImage))
	collidables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff, 0xaf, 0xff}, 0, 320, objImage))
	collectables.AddObject(objects.NewObject(color.NRGBA{0xf0, 0xff, 0xaf, 0xff}, 50, 280, objImage1))
	player = *objects.NewPlayer(ebiten.NewImage(50, 50, ebiten.FilterNearest))
	if err := ebiten.Run(update, 640, 360, 3, "Game"); err != nil {
		fmt.Println(err)
	}
}
