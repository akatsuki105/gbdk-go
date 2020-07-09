package asset

import (
	"github.com/Akatsuki-py/gbdk-go/api/gb"
	"github.com/Akatsuki-py/gbdk-go/api/macro"
)

var SpriteTilesBank = macro.Define(0)

var SpriteTiles = []gb.UINT8{
	0x40, 0x40, 0xE0, 0xA0, 0xFE, 0x80, 0xC8, 0xB8,
	0x8C, 0xF4, 0x96, 0xEA, 0xA9, 0xD7, 0xFF, 0xFF,
	0x02, 0x40, 0xC5, 0xA3, 0xE5, 0xBB, 0xF5, 0xAB,
	0xF5, 0xAB, 0xF5, 0xAB, 0xFD, 0xBB, 0x42, 0x42,
	0x00, 0x00, 0x00, 0x18, 0x18, 0x3C, 0x3C, 0x66,
	0x3C, 0x66, 0x18, 0x3C, 0x00, 0x18, 0x00, 0x00,
}
