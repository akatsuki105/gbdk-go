package gb

func SetSpriteData(firstTile, nbTiles UINT8, data []UINT8) {}
func GetSpriteData(firstTile, nbTiles UINT8, data []UINT8) {}
func SetSpriteTile(nb, tile UINT8)                         {}
func GetSpriteTile(nb UINT8)                               {}
func SetSpriteProp(nb, prop UINT8)                         {}
func GetSpriteProp(nb UINT8)                               {}
func MoveSprite(nb, x, y UINT8)                            {}
func ScrollSprite(nb, x, y UINT8)                          {}
