package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"./objects"
	"fmt"
)

var player objects.Player
var col, col1 objects.Object

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xa0, 0x01, 0xfa, 0xff})

	player.Move(screen)
	player.Collide(col)
	player.Collide(col1)

	col.Display(screen)
	col1.Display(screen)
	player.Display(screen)


	ebitenutil.DebugPrint(screen, fmt.Sprintf("Position: %0.f", player.GetPosition()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nVelocity: %.0f", player.Velocity))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nFPS: %.1f", ebiten.CurrentFPS()))

	return nil
}

func main() {
	objImage, _ := ebiten.NewImage(90, 80, ebiten.FilterNearest)
	col = *objects.NewObject(color.NRGBA{0xf0, 0xff,0xaf, 0xff}, 640 / 2, 360 / 2, objImage)
	col1 = *objects.NewObject(color.NRGBA{0xf0, 0xff, 0xaf, 0xff}, 640 / 2 +90, 360 / 2 + 50, objImage)
	player = *objects.NewPlayer(ebiten.NewImage(50, 50, ebiten.FilterNearest))
	if err := ebiten.Run(update, 640, 360, 3, "Game"); err != nil {
		fmt.Println(err)
	}
}
