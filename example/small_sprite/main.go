package main

import (
	"gbdk/api/gb"
	"gbdk/example/small_sprite/asset"
)

func main() {
	gb.SPRITES_8x8()
	gb.SetSpriteData(0, 8, asset.Sprite)
	gb.SetSpriteTile(0, 0)
	gb.MoveSprite(0, 50, 50)
	gb.SHOW_SPRITES()
}
