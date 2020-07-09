package gb

func SetWinData(firstTile, nbTiles UINT8, data []UINT8) {}
func GetWinData(firstTile, nbTiles UINT8, data []UINT8) {}
func SetWinTiles(x, y, w, h UINT8, tiles []UINT8)       {}
func GetWinTiles(x, y, w, h UINT8, tiles []UINT8)       {}
func MoveWin(x, y UINT8)                                {}
func ScrollWin(x, y INT8)                               {}
