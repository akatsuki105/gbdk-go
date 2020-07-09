package gb

const (
	M_DRAWING    = 0x01
	M_TEXT_OUT   = 0x02
	M_TEXT_INOUT = 0x03
	M_NO_SCROLL  = 0x04
	M_NO_INTERP  = 0x08
	S_PALETTE    = 0x10
	S_FLIPX      = 0x20
	S_FLIPY      = 0x40
	S_PRIORITY   = 0x80
)

// Mode NONBANKED
//
// gbdk: UINT8 get_mode(void) NONBANKED;
//
// Mode Returns the current mode
func Mode() UINT8 {
	return 0
}

// SetMode NONBANKED
//
// gbdk: void mode(UINT8 m) NONBANKED;
//
// Set the current mode - one of M_* defined above
func SetMode(m UINT8) {}
