package student

import "github.com/01-edu/z01"

func Raid1a(w, h int) {
	if w == 0 || h == 0 {
		z01.PrintRune(10)
	} else {
		for a := 0; a < h; a++ {
			for b := 0; b < w; b++ {
				if a == 0 || a == h-1 {
					if b == 0 || b == w-1 {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('-')
					}
				} else {
					if b == 0 || b == w-1 {
						z01.PrintRune('|')
					} else {
						z01.PrintRune(' ')
					}
				}
				if b == w-1 {
					z01.PrintRune(10)
				}
			}
		}
	}
}
