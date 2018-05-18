package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"fmt"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"./objects"
)

var player objects.Player
var floor objects.StaticShape

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xa0, 0x01, 0xfa, 0xff})

	player.Move(screen)
	player.Display(screen)
	floor.Display(screen)

	if player.IsColliding(&floor.Collision) {
		player.PlayerColour = color.NRGBA{0xff, 0x00, 0x00, 0xff}
	} else {
		player.PlayerColour = color.NRGBA{0xaf, 0xf0, 0x0a, 0xff}
	}


	ebitenutil.DebugPrint(screen, fmt.Sprintf("Position: %0.f", player.Position))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nVelocity: %.0f", player.Velocity))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nFPS: %.1f", ebiten.CurrentFPS()))


	return nil
}

func main() {
	floorImage, _ := ebiten.NewImage(640, 20, ebiten.FilterNearest)
	floor = *objects.CreateStaticShape(0, 340, floorImage)
	player = *objects.NewPlayer(ebiten.NewImage(50, 50, ebiten.FilterNearest))
	if err := ebiten.Run(update, 640, 360, 3, "Game Of Life"); err != nil {
		fmt.Println(err)
	}
}
