package main

import (
	"fmt"
	"strings"
)

var asciiFont = map[rune][]string{
	'A': {
		"   A   ",
		"  A A  ",
		" A   A ",
		" AAAAA ",
		" A   A ",
	},
	'B': {
		" BBBB  ",
		" B   B ",
		" BBBB  ",
		" B   B ",
		" BBBB  ",
	},
	'C': {
		"  CCCC ",
		" C     ",
		" C     ",
		" C     ",
		"  CCCC ",
	},
	'D': {
		" DDDD  ",
		" D   D ",
		" D   D ",
		" D   D ",
		" DDDD  ",
	},
	'E': {
		" EEEEE ",
		" E     ",
		" EEEE  ",
		" E     ",
		" EEEEE ",
	},
	'F': {
		" FFFFF ",
		" F     ",
		" FFFF  ",
		" F     ",
		" F     ",
	},
	'G': {
		"  GGGG ",
		" G     ",
		" G  GG ",
		" G   G ",
		"  GGGG ",
	},
	'H': {
		" H   H ",
		" H   H ",
		" HHHHH ",
		" H   H ",
		" H   H ",
	},
	'I': {
		" IIIII ",
		"   I   ",
		"   I   ",
		"   I   ",
		" IIIII ",
	},
	'J': {
		"  JJJJ ",
		"    J  ",
		"    J  ",
		" J  J  ",
		"  JJ   ",
	},
	'K': {
		" K   K ",
		" K  K  ",
		" KKK   ",
		" K  K  ",
		" K   K ",
	},
	'L': {
		" L     ",
		" L     ",
		" L     ",
		" L     ",
		" LLLLL ",
	},
	'M': {
		" M   M ",
		" MM MM ",
		" M M M ",
		" M   M ",
		" M   M ",
	},
	'N': {
		" N   N ",
		" NN  N ",
		" N N N ",
		" N  NN ",
		" N   N ",
	},
	'O': {
		"  OOO  ",
		" O   O ",
		" O   O ",
		" O   O ",
		"  OOO  ",
	},
	'P': {
		" PPPP  ",
		" P   P ",
		" PPPP  ",
		" P     ",
		" P     ",
	},
	'Q': {
		"  QQQ  ",
		" Q   Q ",
		" Q   Q ",
		" Q  QQ ",
		"  QQQQ ",
	},
	'R': {
		" RRRR  ",
		" R   R ",
		" RRRR  ",
		" R  R  ",
		" R   R ",
	},
	'S': {
		"  SSSS ",
		" S     ",
		"  SSS  ",
		"     S ",
		" SSSS  ",
	},
	'T': {
		" TTTTT ",
		"   T   ",
		"   T   ",
		"   T   ",
		"   T   ",
	},
	'U': {
		" U   U ",
		" U   U ",
		" U   U ",
		" U   U ",
		"  UUU  ",
	},
	'V': {
		" V   V ",
		" V   V ",
		" V   V ",
		"  V V  ",
		"   V   ",
	},
	'W': {
		" W   W ",
		" W   W ",
		" W W W ",
		" WW WW ",
		" W   W ",
	},
	'X': {
		" X   X ",
		"  X X  ",
		"   X   ",
		"  X X  ",
		" X   X ",
	},
	'Y': {
		" Y   Y ",
		"  Y Y  ",
		"   Y   ",
		"   Y   ",
		"   Y   ",
	},
	'Z': {
		" ZZZZZ ",
		"    Z  ",
		"   Z   ",
		"  Z    ",
		" ZZZZZ ",
	},
	' ': {
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
	},
}

func generateASCIIArt(text string) string {
	lines := make([]string, 5)

	for _, char := range strings.ToUpper(text) {
		if asciiChar, ok := asciiFont[char]; ok {
			for i := 0; i < 5; i++ {
				lines[i] += asciiChar[i] + " "
			}
		} else {
			for i := 0; i < 5; i++ {
				lines[i] += "      " //for unsupported characters
			}
		}
	}

	return strings.Join(lines, "\n")
}

func main() {
	text := "HELLO"
	asciiArt := generateASCIIArt(text)
	fmt.Println(asciiArt)
}
