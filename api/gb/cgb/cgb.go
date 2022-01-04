package cgb

import "github.com/Akatsuki-py/gbdk-go/api/gb"

const (
	RGB_RED gb.UINT16 = iota
	RGB_DARKRED
	RGB_GREEN
	RGB_DARKGREEN
	RGB_BLUE
	RGB_DARKBLUE
	RGB_YELLOW
	RGB_DARKYELLOW
	RGB_CYAN
	RGB_AQUA
	RGB_PINK
	RGB_PURPLE
	RGB_BLACK
	RGB_DARKGRAY
	RGB_LIGHTGRAY
	RGB_WHITE
	RGB_LIGHTFLESH
	RGB_BROWN
	RGB_ORANGE
	RGB_TEAL
)

func RGB(r, g, b gb.UINT8) gb.UINT16 {
	return 0
}

// SetBkgPalette Set bkg palette(s).
func SetBkgPalette(firstPalette, nbPalettes gb.UINT8, rgbData []gb.UINT16) {}

// SetSpritePalette Set sprite palette(s).
func SetSpritePalette(firstPalette, nbPalettes gb.UINT8, rgbData []gb.UINT16) {}

// SetBkgPaletteEntry Set a bkg palette entry.
func SetBkgPaletteEntry(palette, entry gb.UINT8, rgbData gb.UINT16) {}

// SetSpritePaletteEntry Set a sprite palette entry.
func SetSpritePaletteEntry(palette, entry gb.UINT8, rgbData gb.UINT16) {}

// CPUSlow Set CPU speed to slow operation. Make sure interrupts are disabled before call.
func CPUSlow() {}

// CPUFast Set CPU speed to fast operation. Make sure interrupts are disabled before call.
func CPUFast() {}

// CGBCompatibility Set defaults compatible with normal GameBoy.
func CGBCompatibility() {}
