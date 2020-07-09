package main

import (
	"gbdk/api/gb"
	"gbdk/example/small_sprite/asset"
)

func main() {
	gb.SPRITES_8x16()
	gb.SetSpriteData(0, 4, asset.Sprite)
	gb.SetSpriteTile(0, 0)
	gb.MoveSprite(0, 75, 75)
	gb.SetSpriteTile(1, 2)
	gb.MoveSprite(1, 75+8, 75)
	gb.SHOW_SPRITES()
}
