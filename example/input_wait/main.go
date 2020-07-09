package main

import (
	"github.com/Akatsuki-py/gbdk-go/api/gb"
	"github.com/Akatsuki-py/gbdk-go/api/stdio"
)

func main() {
	for {
		stdio.Printf("Press Start\n\n")
		gb.Waitpad(gb.J_START)

		stdio.Printf("Please hold down A!\n\n")
		gb.Waitpad(gb.J_A)
		stdio.Printf("Holding down A!\n\n")
		gb.Waitpadup()
		stdio.Printf("Tired already?\n\n\n")
	}
}
