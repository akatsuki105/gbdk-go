package gb

const (
	RGB_RED UINT16 = iota
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

func RGB(r, g, b UINT8) UINT16 {
	return 0
}

// SetBkgPalette Set bkg palette(s).
func SetBkgPalette(firstPalette, nbPalettes UINT8, rgbData []UINT16) {}

// SetSpritePalette Set sprite palette(s).
func SetSpritePalette(firstPalette, nbPalettes UINT8, rgbData []UINT16) {}

// SetBkgPaletteEntry Set a bkg palette entry.
func SetBkgPaletteEntry(palette, entry UINT8, rgbData UINT16) {}

// SetSpritePaletteEntry Set a sprite palette entry.
func SetSpritePaletteEntry(palette, entry UINT8, rgbData UINT16) {}

// CPUSlow Set CPU speed to slow operation. Make sure interrupts are disabled before call.
func CPUSlow() {}

// CPUFast Set CPU speed to fast operation. Make sure interrupts are disabled before call.
func CPUFast() {}

// CGBCompatibility Set defaults compatible with normal GameBoy.
func CGBCompatibility() {}
