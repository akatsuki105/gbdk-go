package main

import (
	"github.com/Akatsuki-py/gbdk-go/api/gb"
	"github.com/Akatsuki-py/gbdk-go/api/rand"
)

// bank0
func vblint() {
	gb.SWITCH_ROM_MBC1(2)
	UpdateGraphics()
}

func main() {
	gb.SWITCH_ROM_MBC1(3)
	DrawTitleScreen()
	rand.Initrand(6444)
	LoadGameTiles()
	gb.DisableInterrupts()
	gb.AddVBL(vblint)
	gb.SWITCH_ROM_MBC1(2)
	gb.EnableInterrupts()
	DoGameplay()
	gb.Reset()
}
