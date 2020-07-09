package gb

type BOOLEAN bool

type INT8 int8
type UINT8 uint8
type UINT16 uint16
type UINT32 uint32

type BYTE uint8
type WORD uint16
type DWORD uint32

// Delay NONBANKED
//
// Delays the given number of milliseconds.
// Uses no timers or interrupts, and can be called with
// interrupts disabled (why nobody knows :)
func Delay(duration UINT16) {}
