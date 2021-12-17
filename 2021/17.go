package main

import "fmt"

var (
	xmin, xmax = 138, 184
	ymin, ymax = -125, -71
	highest    = 0
)

func main() {
	success := 0
	for initdx := 0; initdx <= xmax; initdx++ {
		for initdy := ymin; initdy <= -ymax*10; initdy++ {
			x, y := 0, 0
			dx, dy := initdx, initdy
			maxy := 0
			for !(dx > 0 && x > xmax) && !(dx == 0 && (x < xmin || x > xmax)) && !(dy < 0 && y < ymin) && !(x >= xmin && x <= xmax && y >= ymin && y <= ymax) {
				x += dx
				y += dy
				if dx > 0 {
					dx--
				}
				if dx < 0 {
					dx++
				}
				dy--
				if y > maxy {
					maxy = y
				}
			}
			if x >= xmin && x <= xmax && y >= ymin && y <= ymax {
				success++
				if maxy > highest {
					highest = maxy
				}
			}
		}
	}
	fmt.Println("Part One: ", highest)
	fmt.Println("Part Two: ", success)
}
