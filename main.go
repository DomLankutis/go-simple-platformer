package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"./objects"
	"fmt"
)

var player objects.Player
var col objects.Object

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xa0, 0x01, 0xfa, 0xff})

	player.Move(screen)
	player.Display(screen)
	col.Display(screen)

	player.Colour = color.NRGBA{0x0f, 0xff, 0xaf, 0xff}
	player.Collide(col)
	player.CollisionBox.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Position: %0.f", player.GetPosition()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nVelocity: %.0f", player.Velocity))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nFPS: %.1f", ebiten.CurrentFPS()))


	return nil
}

func main() {
	objImage, _ := ebiten.NewImage(90, 80, ebiten.FilterNearest)
	col = *objects.NewObject(color.NRGBA{0xf0, 0xff,0xaf, 0xff}, 640 / 2, 360 / 2, objImage)
	player = *objects.NewPlayer(ebiten.NewImage(50, 50, ebiten.FilterNearest))
	if err := ebiten.Run(update, 640, 360, 3, "Game"); err != nil {
		fmt.Println(err)
	}
}
