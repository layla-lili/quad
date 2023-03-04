package piscine

import "github.com/01-edu/z01"

// func main() {
// 	QuadA(1, 5)
// }

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for a := 0; a < y; a++ {

			if a == 0 || a == y-1 {
				z01.PrintRune('o')
				for i := 1; i < x-1; i++ {
					z01.PrintRune('-')
				}
				if x > 1 {
					z01.PrintRune('o')
				}
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('|')
				for i := 1; i < x-1; i++ {
					z01.PrintRune(' ')
				}
				if x > 1 {
					z01.PrintRune('|')
				}
				z01.PrintRune('\n')
			}
		}
	}
}
