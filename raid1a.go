package student

import "github.com/01-edu/z01"

func Raid1a(w, h int) {
	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			if x == 0 || x == w-1 {
				if y == 0 || y == h-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if y == 0 || y == h-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune(10)
	}
}
