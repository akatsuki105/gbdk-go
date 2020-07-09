package main

import (
	"gbdk/api/drawing"
	"gbdk/api/gb"
	"gbdk/api/rand"
)

var badguyAI = []gb.UINT8{
	32, 32, 33, 34, 35, 35, 36, 37,
	38, 38, 39, 40, 41, 41, 42, 43,
	44, 44, 45, 46, 46, 47, 48, 48,
	49, 50, 50, 51, 51, 52, 53, 53,
	54, 54, 55, 55, 56, 56, 57, 57,
	58, 58, 58, 59, 59, 60, 60, 60,
	61, 61, 61, 61, 62, 62, 62, 62,
	62, 63, 63, 63, 63, 63, 63, 63,
	63, 63, 63, 63, 63, 63, 63, 63,
	62, 62, 62, 62, 62, 61, 61, 61,
	61, 60, 60, 60, 59, 59, 59, 58,
	58, 57, 57, 56, 56, 55, 55, 54,
	54, 53, 53, 52, 52, 51, 50, 50,
	49, 49, 48, 47, 47, 46, 45, 44,
	44, 43, 42, 42, 41, 40, 39, 39,
	38, 37, 36, 36, 35, 34, 33, 33,
	32, 31, 30, 29, 29, 28, 27, 26,
	26, 25, 24, 23, 23, 22, 21, 20,
	20, 19, 18, 18, 17, 16, 16, 15,
	14, 14, 13, 12, 12, 11, 11, 10,
	9, 9, 8, 8, 7, 7, 6, 6,
	6, 5, 5, 4, 4, 4, 3, 3,
	3, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 1, 2, 2,
	2, 3, 3, 3, 4, 4, 4, 5,
	5, 5, 6, 6, 7, 7, 8, 8,
	9, 9, 10, 11, 11, 12, 12, 13,
	14, 14, 15, 16, 16, 17, 18, 18,
	19, 20, 20, 21, 22, 23, 23, 24,
	25, 26, 26, 27, 28, 29, 29, 30,
}

var playing gb.UINT8
var joypadState gb.UINT8
var PlayerX gb.UINT8
var PlayerY gb.UINT8
var BadguyX gb.UINT8
var BadguyY gb.UINT8
var BadguyZ gb.UINT8
var BadguyOffset gb.UINT8
var PlayerShotX gb.UINT8
var PlayerShotY gb.UINT8
var PlayerShotZ gb.UINT8

func UpdateGraphics() {
	gb.DisableInterrupts()

	gb.ScrollBkg(1, 0)
	gb.MoveSprite(0, PlayerX, PlayerY)
	gb.MoveSprite(1, BadguyX, BadguyY)
	gb.MoveSprite(2, PlayerShotX, PlayerShotY)

	gb.EnableInterrupts()
}

func UpdateJoypad() {
	joypadState = gb.Joypad()

	if joypadState&gb.J_LEFT > 0 {
		if PlayerX > 8 {
			PlayerX--
		}
	}
	if joypadState&gb.J_RIGHT > 0 {
		if PlayerX < 160 {
			PlayerX++
		}
	}
	if joypadState&gb.J_UP > 0 {
		if PlayerY > 16 {
			PlayerY--
		}
	}
	if joypadState&gb.J_DOWN > 0 {
		if PlayerY < 152 {
			PlayerY++
		}
	}

	if joypadState&gb.J_A > 0 {
		if PlayerShotZ == 0 {
			PlayerShotZ = 1
			PlayerShotX = PlayerX
			PlayerShotY = PlayerY
		}
	}
	if PlayerShotZ == 1 {
		PlayerShotX += 2
		if PlayerShotX > 240 {
			PlayerShotX = 250
			PlayerShotY = 250
			PlayerShotZ = 0
		}
	}
}

func UpdateBadguy() {
	BadguyX -= 1
	if BadguyX > 240 {
		BadguyOffset = gb.UINT8(rand.Rand())
		for BadguyOffset > 134 {
			BadguyOffset = gb.UINT8(rand.Rand())
		}
		BadguyX = 239
	}

	BadguyY = BadguyOffset + badguyAI[BadguyZ]
	BadguyZ++
}

func CollisionDetection() {
	if PlayerShotY > BadguyY-4 {
		if PlayerShotY < BadguyY+4 {
			if PlayerShotX > BadguyX-4 {
				if PlayerShotX < BadguyX+4 {
					PlayerShotZ = 0
					PlayerShotX = 250
					PlayerShotY = 250
					BadguyX = 255
				}
			}
		}
	}

	if PlayerY > BadguyY-4 {
		if PlayerY < BadguyY+4 {
			if PlayerX > BadguyX-4 {
				if PlayerX < BadguyX+4 {
					playing = 0
					gb.Delay(1000)
				}
			}
		}
	}
}

func DoGameplay() {
	playing = 1

	PlayerX = drawing.GRAPHICS_WIDTH / 2
	PlayerY = drawing.GRAPHICS_HEIGHT / 2

	for playing == 1 {
		UpdateJoypad()
		UpdateBadguy()
		CollisionDetection()
		gb.Delay(10)
	}
}
