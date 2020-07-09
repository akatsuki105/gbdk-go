package main

import (
	"gbdk/api/gb"
	"gbdk/example/move_sprite/asset"
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
		if (gb.Joypad() & gb.J_RIGHT) > 0 {
			x++
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		}
		if (gb.Joypad() & gb.J_LEFT) > 0 {
			x--
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		}
		if (gb.Joypad() & gb.J_UP) > 0 {
			y--
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		}
		if (gb.Joypad() & gb.J_DOWN) > 0 {
			y++
			gb.MoveSprite(0, x, y)
			gb.Delay(10)
		}
	}
}
