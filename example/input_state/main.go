package main

import (
	"gbdk/api/gb"
	"gbdk/api/stdio"
)

func main() {
	for {
		switch gb.Joypad() {
		case gb.J_LEFT:
			stdio.Printf("Left!\n")
			gb.Delay(100)
		case gb.J_RIGHT:
			stdio.Printf("Right!\n")
			gb.Delay(100)
		case gb.J_UP:
			stdio.Printf("Up!\n")
			gb.Delay(100)
		case gb.J_DOWN:
			stdio.Printf("Down!\n")
			gb.Delay(100)
		case gb.J_START:
			stdio.Printf("Start!\n")
			gb.Delay(100)
		case gb.J_SELECT:
			stdio.Printf("Select!\n")
			gb.Delay(100)
		case gb.J_A:
			stdio.Printf("A!\n")
			gb.Delay(100)
		case gb.J_B:
			stdio.Printf("B!\n")
			gb.Delay(100)
		}
	}
}
