package gb

const (
	J_START  = 0x80
	J_SELECT = 0x40
	J_B      = 0x20
	J_A      = 0x10
	J_DOWN   = 0x08
	J_UP     = 0x04
	J_LEFT   = 0x02
	J_RIGHT  = 0x01
)

// Joypad Reads and returns the current state of the joypad.
// Follows Nintendo's guidelines for reading the pad.
// Return value is an OR of J_*
//
// @see J_START
func Joypad() UINT8 {
	return 0
}

// Waitpad Waits until all the keys given in mask are pressed.
// Normally only used for checking one key, but it will
// support many, even J_LEFT at the same time as J_RIGHT :)
//
// @see joypad, J_START
func Waitpad(mask UINT8) UINT8 {
	return 0
}

// Waitpadup Waits for the pad and all buttons to be released.
func Waitpadup() {}
