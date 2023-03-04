package piscine

import "github.com/01-edu/z01"

// func main() {
// 	QuadB(1, 1)
// }

func QuadB(x, y int) {
	//check width and height greater than zero
	if x > 0 && y > 0 {
		// iterate over pattern row
		for a := 0; a < y; a++ {
			//First Row
			if a == 0 {
				//start with
				z01.PrintRune('/')
				//keep up with before the end
				for i := 1; i < x-1; i++ {
					z01.PrintRune('*')
				}
				// at the end of row
				if x > 1 {
					z01.PrintRune('\\')
				}
				z01.PrintRune('\n')
				//last row example (5,3) a:3 == y:4-1=3
			} else if a == y-1 {
				z01.PrintRune('\\')
				//fill stars between exept first and last
				for i := 1; i < x-1; i++ {
					z01.PrintRune('*')
				}
				if x > 1 || a == y-1 {
					z01.PrintRune('/')
				}
				z01.PrintRune('\n')

			} else {
				//fill between with star for first and and space for the rest stop before the end
				z01.PrintRune('*')
				for i := 1; i < x-1; i++ {
					z01.PrintRune(' ')
				}
				if x > 1 {
					z01.PrintRune('*')
				}
				z01.PrintRune('\n')
			}

		}

	}
}
