package drawing

import "github.com/Akatsuki-py/gbdk-go/api/gb"

/** Size of the screen in pixels */
const (
	GRAPHICS_WIDTH  = 160
	GRAPHICS_HEIGHT = 144
)

/** Possible drawing modes */
const (
	SOLID = iota
	OR
	XOR
	AND
)

/** Possible drawing colours */
const (
	WHITE = iota
	LTGREY
	DKGREY
	BLACK
)

/** Possible fill styles for box() and circle() */
const (
	M_NOFILL = iota
	M_FILL
)

/** Possible values for signed_value in GPrintln() and GPrintn() */
const (
	UNSIGNED = iota
	SIGNED
)

// GPrint NONBANKED
//
// Print the string 'str' with no interpretation
func GPrint(str string) {}

// GPrintln Print the long number 'number' in radix 'radix'.  signed_value should
// be set to SIGNED or UNSIGNED depending on whether the number is signed or not
func GPrintln(number int16, radix int8, signedValue int8) {}

// GPrintn Print the number 'number' as in 'GPrintln'
func GPrintn(number int16, radix int8, signedValue int8) {}

// GPrintf Print the formatted string 'fmt' with arguments '...'
func GPrintf(format string, a ...interface{}) int8 {
	return 0
}

// Plot Old style plot - try plot_point()
func Plot(x, y, color, mode gb.UINT8) {}

// PlotPoint Plot a point in the current drawing mode and colour at (x,y)
func PlotPoint(x, y gb.UINT8) {}

// SwitchData NONBANKED
//
// Exchanges the tile on screen at x,y with the tile pointed by src, original tile
// not performed in this case.
// is saved in dst. Both src and dst may be NULL - saving or copying to screen is
func SwitchData(x, y gb.UINT8, src, dst []gb.UINT8) {}

// DrawImage NONBANKED
func DrawImage(data []gb.UINT8) {}

func Line(x1, y1, x2, y2 gb.UINT8)        {}
func Box(x1, y1, x2, y2, style gb.UINT8)  {}
func Circle(x, y, radius, style gb.UINT8) {}

// GetPix Returns the current colour of the pixel at (x,y)
func GetPix(x, y gb.UINT8) gb.UINT8 {
	return 0
}

// WriteChar Prints the character 'char' in the default font at the current position
func WriteChar(char rune) {}

// GotoGXY Sets the current text position to (x,y).  Note that x and y have units of cells (8 pixels)
func GotoGXY(x, y gb.UINT8) {}

// Color Set the current foreground colour (for pixels), background colour, and draw mode
func Color(foreColor, backColor, mode gb.UINT8) {}
