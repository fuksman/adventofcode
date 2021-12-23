package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord [3]int

type Moon struct {
	position Coord
	velosity Coord
}

func (m *Moon) energy() int {
	pot, kin := 0, 0
	for i := 0; i < 3; i++ {
		pot += abs(m.position[i])
		kin += abs(m.velosity[i])
	}
	return pot * kin
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

var (
	defaultVel = Coord{0, 0, 0}
	moons      = []*Moon{}
	initial    = []Coord{}
)

func main() {
	file, err := os.Open("12.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y, z int
		fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, &Moon{position: Coord{x, y, z}, velosity: defaultVel})
		initial = append(initial, Coord{x, y, z})
	}

	found := [3]bool{}
	cycles := Coord{}

	for c := 1; !(found[0] && found[1] && found[2]); c++ {
		if c == 1001 {
			energy := 0
			for _, m := range moons {
				energy += m.energy()
			}
			fmt.Println("Part One: ", energy)
		}
		for i, m1 := range moons {
			for j := i + 1; j < len(moons); j++ {
				m2 := moons[j]
				for pos := 0; pos < 3; pos++ {
					if m1.position[pos] > m2.position[pos] {
						m1.velosity[pos]--
						m2.velosity[pos]++
					}
					if m1.position[pos] < m2.position[pos] {
						m1.velosity[pos]++
						m2.velosity[pos]--
					}
				}
			}
			for pos := 0; pos < 3; pos++ {
				m1.position[pos] += m1.velosity[pos]
			}
		}

		for i := 0; i < 3; i++ {
			if found[i] {
				continue
			}
			eq := true
			for j := range moons {
				if moons[j].position[i] != initial[j][i] || moons[j].velosity[i] != defaultVel[i] {
					eq = false
					break
				}
			}
			if eq {
				found[i] = true
				cycles[i] = c
			}
		}
	}
	fmt.Println("Part Two: ", LCM(cycles[0], cycles[1], cycles[2]))
}
