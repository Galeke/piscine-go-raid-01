package student

import "fmt"

func Raid1a(w, h int) {
	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			if x == 0 || x == w-1 {
				if y == 0 || y == h-1 {
					fmt.Print('o')
				} else {
					fmt.Print('-')
				}
			} else {
				if y == 0 || y == h-1 {
					fmt.Print('|')
				} else {
					fmt.Print(' ')
				}
			}
		}
	}
}
