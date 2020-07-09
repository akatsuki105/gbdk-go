package main

import (
	"gbdk/api/gb"
	"gbdk/example/big_sprite_animation/asset"
)

func main() {
	gb.SPRITES_8x16()
	gb.SetSpriteData(0, 8, asset.Cards)
	gb.SetSpriteTile(0, 0)
	gb.MoveSprite(0, 75, 75)
	gb.SetSpriteTile(1, 2)
	gb.MoveSprite(1, 75+8, 75)
	gb.SHOW_SPRITES()

	for {
		gb.SetSpriteTile(0, 4)
		gb.SetSpriteTile(1, 6)
		gb.Delay(500)
		gb.SetSpriteTile(0, 0)
		gb.SetSpriteTile(1, 2)
		gb.Delay(500)
	}
}
