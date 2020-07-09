package main

import (
	"github.com/Akatsuki-py/gbdk-go/api/drawing"
	"github.com/Akatsuki-py/gbdk-go/api/gb"
)

func main() {
	gb.NR50_REG = 0xff
	gb.NR51_REG = 0xff
	gb.NR52_REG = 0x80

	drawing.GotoGXY(1, 1)
	drawing.GPrintf("====== Beep ======")

	drawing.GotoGXY(2, 3)
	drawing.GPrintf("Press any button")

	for {
		joypadState := gb.Joypad()
		if joypadState > 0 {
			gb.NR10_REG = 0x38
			gb.NR11_REG = 0x70
			gb.NR12_REG = 0xe0
			gb.NR13_REG = 0x0a
			gb.NR14_REG = 0xc6

			gb.NR51_REG |= 0x11

			gb.Delay(200)
		}
	}
}
