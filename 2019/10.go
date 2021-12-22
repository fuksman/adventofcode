package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Coord [2]int

var (
	astCoords = map[Coord]map[int][]Coord{}
)

func main() {
	file, err := os.Open("10.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, ch := range scanner.Text() {
			if ch == '#' {
				astCoords[Coord{x, y}] = make(map[int][]Coord)
			}
		}
		y++
	}

	max := 0
	var base Coord
	for astA := range astCoords {
		count := 0
		for astB := range astCoords {
			if astA == astB {
				continue
			}

			dx, dy := astB[0]-astA[0], astB[1]-astA[1]

			var ox, oy int
			if dx == 0 {
				oy = dy / abs(dy)
			} else if dy == 0 {
				ox = dx / abs(dx)
			} else {
				f := gcf(abs(dx), abs(dy))
				ox, oy = dx/f, dy/f
			}

			x, y := ox, oy

			v := true
			pass := 0
			for x != dx || y != dy {
				if _, ok := astCoords[Coord{astA[0] + x, astA[1] + y}]; ok {
					v = false
					pass++
				}
				x += ox
				y += oy
			}
			astCoords[astA][pass] = append(astCoords[astA][pass], astB)
			if v {
				count++
			}
		}
		if count > max {
			max = count
			base = astA
		}
	}
	fmt.Println("Part One: ", max)

	vap := astCoords[base]

	off := 0
	var p []Coord
	for i := 0; ; i++ {
		if off+len(vap[i]) < 200 {
			off += len(vap[i])
		} else {
			p = vap[i]
			break
		}
	}

	angle := func(i int) float64 {
		dx := p[i][0] - base[0]
		dy := p[i][1] - base[1]
		a := math.Atan2(float64(dy), float64(dx))
		if a < -math.Pi/2.0 {
			a += 2.0 * math.Pi
		}
		return a
	}

	sort.Slice(p, func(i, j int) bool {
		return angle(i) < angle(j)
	})

	fmt.Println("Part Two: ", p[199-off][0]*100+p[199-off][1])
}

func gcf(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcf(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
