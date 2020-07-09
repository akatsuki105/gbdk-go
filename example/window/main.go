package main

import (
	"gbdk/api/gb"
	"gbdk/example/window/asset"
)

func main() {
	gb.SetBkgData(0, 152, asset.BGTileData)
	gb.SetBkgTiles(0, 0, 20, 18, asset.BGTileMap)
	gb.SHOW_BKG()

	gb.SetWinData(152, 9, asset.Border)
	gb.SetWinTiles(0, 0, 20, 4, asset.Window)
	gb.MoveWin(7, 112)
	gb.SHOW_WIN()
}
