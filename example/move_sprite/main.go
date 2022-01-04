package main

import (
	"github.com/Akatsuki-py/gbdk-go/api/gb"
	"github.com/Akatsuki-py/gbdk-go/example/move_sprite/asset"
)

func main() {
	x := gb.UINT8(55)
	y := gb.UINT8(75)

	gb.SPRITES_8x8()
	gb.SetSpriteData(0, 0, asset.UFO)
	gb.SetSpriteTile(0, 0)
	gb.MoveSprite(0, x, y)

	gb.SHOW_SPRITES()

	for {
		switch gb.Joypad() {
		case gb.J_RIGHT:
			x++
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		case gb.J_LEFT:
			x--
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		case gb.J_UP:
			y--
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		case gb.J_DOWN:
			y++
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		}
	}
}
