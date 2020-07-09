package util

import (
	"reflect"
	"strings"
)

var CType map[string]string = map[string]string{
	"uint8": "unsigned char",
}

var CFunc map[string]string = map[string]string{
	"Printf":                "printf",
	"Sprintf":               "sprintf",
	"Puts":                  "puts",
	"Abs":                   "abs",
	"Atoi":                  "atoi",
	"UINT8":                 "UINT8",
	"UINT16":                "UINT16",
	"Strcat":                "strcat",
	"Strcpy":                "strcpy",
	"Initrand":              "initrand",
	"Rand":                  "rand",
	"Randw":                 "randw",
	"Initarand":             "initarand",
	"Arand":                 "arand",
	"Gotoxy":                "gotoxy",
	"Posx":                  "posx",
	"Posy":                  "posy",
	"Setchar":               "setchar",
	"SetBkgPalette":         "set_bkg_palette",
	"SetSpritePalette":      "set_sprite_palette",
	"SetBkgPaletteEntry":    "set_bkg_palette_entry",
	"SetSpritePaletteEntry": "set_sprite_palette_entry",
	"CPUSlow":               "cpu_slow",
	"CPUFast":               "cpu_fast",
	"CGBCompatibility":      "cgb_compatibility",
	"RemoveVBL":             "remove_VBL",
	"RemoveLCD":             "remove_LCD",
	"RemoveTIM":             "remove_TIM",
	"RemoveSIO":             "remove_SIO",
	"RemoveJOY":             "remove_JOY",
	"AddVBL":                "add_VBL",
	"AddLCD":                "add_LCD",
	"AddTIM":                "add_TIM",
	"AddSIO":                "add_SIO",
	"AddJOY":                "add_JOY",
	"NowaitIntHandler":      "nowait_int_handler",
	"WaitIntHandler":        "wait_int_handler",
	"Mode":                  "get_mode",
	"SetMode":               "mode",
	"EnableInterrupts":      "enable_interrupts",
	"DisableInterrupts":     "disable_interrupts",
	"SetInterrupts":         "set_interrupts",
	"Reset":                 "reset",
	"WaitVBLDone":           "wait_vbl_done",
	"DisplayOff":            "display_off",
	"SetSpriteData":         "set_sprite_data",
	"GetSpriteData":         "get_sprite_data",
	"SetSpriteTile":         "set_sprite_tile",
	"GetSpriteTile":         "get_sprite_tile",
	"SetSpriteProp":         "set_sprite_prop",
	"GetSpriteProp":         "get_sprite_prop",
	"MoveSprite":            "move_sprite",
	"ScrollSprite":          "scroll_sprite",
	"Delay":                 "delay",
	"Joypad":                "joypad",
	"Waitpad":               "waitpad",
	"Waitpadup":             "waitpadup",
	"SetBkgData":            "set_bkg_data",
	"GetBkgData":            "get_bkg_data",
	"SetBkgTiles":           "set_bkg_tiles",
	"GetBkgTiles":           "get_bkg_tiles",
	"MoveBkg":               "move_bkg",
	"ScrollBkg":             "scroll_bkg",
	"SetWinData":            "set_win_data",
	"GetWinData":            "get_win_data",
	"SetWinTiles":           "set_win_tiles",
	"GetWinTiles":           "get_win_tiles",
	"MoveWin":               "move_win",
	"ScrollWin":             "scroll_win",
	"GPrint":                "gprint",
	"GPrintln":              "gprintln",
	"GPrintn":               "gprintn",
	"GPrintf":               "gprintf",
	"Plot":                  "plot",
	"PlotPoint":             "plot_point",
	"SwitchData":            "switch_data",
	"DrawImage":             "draw_image",
	"Line":                  "line",
	"Box":                   "box",
	"Circle":                "circle",
	"GetPix":                "getpix",
	"WriteChar":             "wrtchr",
	"GotoGXY":               "gotogxy",
	"Color":                 "color",
}

func GetCFunc(funcName string) string {
	result, ok := CFunc[funcName]
	if !ok {
		return funcName
	}
	return result
}

var GBDKPackage []string = []string{
	"github.com/Akatsuki-py/gbdk-go/api/stdio", "github.com/Akatsuki-py/gbdk-go/api/gb", "github.com/Akatsuki-py/gbdk-go/api/str", "github.com/Akatsuki-py/gbdk-go/api/stdlib", "github.com/Akatsuki-py/gbdk-go/api/macro", "github.com/Akatsuki-py/gbdk-go/api/mem", "github.com/Akatsuki-py/gbdk-go/api/drawing", "github.com/Akatsuki-py/gbdk-go/api/rand",
}

func RemoveTypePackage(path string) string {
	s := strings.ReplaceAll(path, "github.com/Akatsuki-py/gbdk-go/api/gb.", "")
	s = strings.ReplaceAll(s, "github.com/Akatsuki-py/gbdk-go/types.", "")
	return s
}

func IsArray(typeName string) (bool, string) {
	if strings.Contains(typeName, "]") {
		result := strings.Split(typeName, "]")
		return true, result[len(result)-1]
	}
	return false, typeName
}

func IsMacroStmt(funcName string) bool {
	list := []string{
		"ENABLE_RAM_MBC1",
		"DISABLE_RAM_MBC1",
		"SWITCH_16_8_MODE_MBC1",
		"SWITCH_4_32_MODE_MBC1",
		"ENABLE_RAM_MBC5",
		"DISABLE_RAM_MBC5",
		"DISPLAY_ON",
		"DISPLAY_OFF",
		"SHOW_BKG",
		"HIDE_BKG",
		"SHOW_WIN",
		"HIDE_WIN",
		"SHOW_SPRITES",
		"HIDE_SPRITES",
		"SPRITES_8x16",
		"SPRITES_8x8",
	}
	for _, f := range list {
		if funcName == f {
			return true
		}
	}
	return false
}

func IsMemFunc(funcName string) bool {
	list := []string{
		"ReadMemory", "WriteMemory",
	}
	return Contains(list, funcName)
}

func Contains(list interface{}, target interface{}) bool {
	if reflect.TypeOf(list).Kind() == reflect.Slice || reflect.TypeOf(list).Kind() == reflect.Array {
		listvalue := reflect.ValueOf(list)
		for i := 0; i < listvalue.Len(); i++ {
			if target == listvalue.Index(i).Interface() {
				return true
			}
		}
	}
	if reflect.TypeOf(target).Kind() == reflect.String && reflect.TypeOf(list).Kind() == reflect.String {
		return strings.Contains(list.(string), target.(string))
	}
	return false
}
