package main

import (
	"github.com/Akatsuki-py/gbdk-go/api/gb"
	"github.com/Akatsuki-py/gbdk-go/example/background/asset"
)

func main() {
	gb.SetBkgData(0, 131, asset.TileData)

	gb.VBK_REG = 1
	gb.VBK_REG = 0

	gb.SetBkgTiles(0, 0, 20, 18, asset.TileMap)

	gb.SHOW_BKG()
	gb.DISPLAY_ON()
}
