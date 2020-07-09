package gb

const (
	VBL_IFLAG = 0x01
	LCD_IFLAG = 0x02
	TIM_IFLAG = 0x04
	SIO_IFLAG = 0x08
	JOY_IFLAG = 0x10
)

// IntHandler NONBANKED
//
// Interrupt handlers
type IntHandler func()

// RemoveVBL NONBANKED
//
// The remove functions will remove any interrupt handler.  A handler of NULL will cause bad things to happen.
func RemoveVBL(h IntHandler) {}

// RemoveLCD NONBANKED
//
// The remove functions will remove any interrupt handler.  A handler of NULL will cause bad things to happen.
func RemoveLCD(h IntHandler) {}

// RemoveTIM NONBANKED
//
// The remove functions will remove any interrupt handler.  A handler of NULL will cause bad things to happen.
func RemoveTIM(h IntHandler) {}

// RemoveSIO NONBANKED
//
// The remove functions will remove any interrupt handler.  A handler of NULL will cause bad things to happen.
func RemoveSIO(h IntHandler) {}

// RemoveJOY NONBANKED
//
// The remove functions will remove any interrupt handler.  A handler of NULL will cause bad things to happen.
func RemoveJOY(h IntHandler) {}

// AddVBL NONBANKED
func AddVBL(h IntHandler) {}

// AddLCD NONBANKED
func AddLCD(h IntHandler) {}

// AddTIM NONBANKED
func AddTIM(h IntHandler) {}

// AddSIO NONBANKED
func AddSIO(h IntHandler) {}

// AddJOY NONBANKED
func AddJOY(h IntHandler) {}

// NowaitIntHandler NONBANKED
//
// Interrupt handler chain terminator that don't wait for .STAT
//
// You must add this handler the last in every interrupt handler
// chain if you want to change the default interrupt handler
// behaviour that waits for LCD controller mode to become 1 or 0
// before return from the interrupt.
func NowaitIntHandler() {}

// WaitIntHandler NONBANKED
//
// Interrupt handler chain terminator that waits for .STAT and
// returns in the BEGINNING of mode0 or mode1 ONLY
func WaitIntHandler() {}

// EnableInterrupts NONBANKED
//
// Enables unmasked interrupts
// @see disable_interrupts
func EnableInterrupts() {}

// DisableInterrupts NONBANKED
//
// This function may be called as many times as you like;
// however the first call to enable_interrupts will re-enable them.
// @see enable_interrupts
func DisableInterrupts() {}

// SetInterrupts NONBANKED
//
// Clears any pending interrupts and sets the interrupt mask
// register IO to flags.
//
// @see VBL_IFLAG
//
// @param flags	A logical OR of *_IFLAGS
func SetInterrupts(flags UINT8) {}

// Reset NONBANKED
//
// Performs a warm reset by reloading the CPU value then jumping to the start of crt0 (0x0150)
func Reset() {}

// WaitVBLDone NONBANKED
//
// Waits for the vertical blank interrupt (VBL) to finish.
// This can be used to sync animation with the screen
// re-draw.  If VBL interrupt is disabled, this function will
// never return.  If the screen is off this function returns
// immediately.
func WaitVBLDone() {}

// DisplayOff NONBANKED
//
// Turns the display off.
// Waits until the VBL interrupt before turning the display off.
//
// @see DISPLAY_ON
func DisplayOff() {}
