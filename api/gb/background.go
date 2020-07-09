package gb

func SetBkgData(firstTile, nbTiles UINT8, data []UINT8) {}
func GetBkgData(firstTile, nbTiles UINT8, data []UINT8) {}
func SetBkgTiles(x, y, w, h UINT8, tiles []UINT8)       {}
func GetBkgTiles(x, y, w, h UINT8, tiles []UINT8)       {}
func MoveBkg(x, y UINT8)                                {}
func ScrollBkg(x, y INT8)                               {}
