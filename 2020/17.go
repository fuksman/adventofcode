package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)

func main() {
	file, err := os.Open("17.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	numOfCycles := 6
	active := '#'
	inactive := '.'

	scanner := bufio.NewScanner(file)
	lines := [][]rune{}
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	x := numOfCycles*2 + len(lines[0])
	y := numOfCycles*2 + len(lines)
	z := numOfCycles*2 + 1
	w := z

	cubes := initState3(x, y, z)
	for j := 0; j < len(lines); j++ {
		for i := 0; i < len(lines[0]); i++ {
			cubes[numOfCycles][numOfCycles+j][numOfCycles+i] = lines[j][i]
		}
	}

	activeCount := 0
	for cycle := 0; cycle < numOfCycles; cycle++ {
		nextState := initState3(x, y, z)
		activeCount = 0
		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				for k := 0; k < z; k++ {
					activeNeighbours := countNeighbours3(cubes, i, j, k, x, y, z, active)
					switch cubes[k][j][i] {
					case active:
						if activeNeighbours == 2 || activeNeighbours == 3 {
							nextState[k][j][i] = active
							activeCount++
						}
					case inactive:
						if activeNeighbours == 3 {
							nextState[k][j][i] = active
							activeCount++
						}
					}
				}
			}
		}
		cubes = nextState
	}

	fmt.Println("Part One: ", activeCount)

	// fmt.Printf("After %d cycle: ", cycle+1)
	// fmt.Println()
	// for k := 0; k < z; k++ {
	// 	fmt.Printf("z=%d: ", k-numOfCycles)
	// 	fmt.Println()
	// 	for j := 0; j < y; j++ {
	// 		fmt.Println(string(cubes[k][j]))
	// 	}
	// 	fmt.Println()
	// }
	hypercubes := initState4(x, y, z, w)
	for j := 0; j < len(lines); j++ {
		for i := 0; i < len(lines[0]); i++ {
			hypercubes[numOfCycles][numOfCycles][numOfCycles+j][numOfCycles+i] = lines[j][i]
		}
	}

	activeCount = 0
	for cycle := 0; cycle < numOfCycles; cycle++ {
		nextState := initState4(x, y, z, w)
		activeCount = 0
		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				for k := 0; k < z; k++ {
					for l := 0; l < w; l++ {
						activeNeighbours := countNeighbours4(hypercubes, i, j, k, l, x, y, z, w, active)
						switch hypercubes[l][k][j][i] {
						case active:
							if activeNeighbours == 2 || activeNeighbours == 3 {
								nextState[l][k][j][i] = active
								activeCount++
							}
						case inactive:
							if activeNeighbours == 3 {
								nextState[l][k][j][i] = active
								activeCount++
							}
						}
					}
				}
			}
		}
		hypercubes = nextState
	}
	fmt.Println("Part Two: ", activeCount)
}

func initState3(x, y, z int) [][][]rune {
	state := make([][][]rune, z)
	for k := 0; k < z; k++ {
		state[k] = make([][]rune, y)
		for j := 0; j < y; j++ {
			state[k][j] = make([]rune, x)
			for i := 0; i < x; i++ {
				state[k][j][i] = '.'
			}
		}
	}
	return state
}

func countNeighbours3(state [][][]rune, i, j, k int, x, y, z int, active rune) (activeCount int) {
	for l := i - 1; l <= i+1; l++ {
		for m := j - 1; m <= j+1; m++ {
			for n := k - 1; n <= k+1; n++ {
				if !(l == i && m == j && n == k) && l > 0 && l < x && m > 0 && m < y && n > 0 && n < z && state[n][m][l] == active {
					activeCount++
				}

			}
		}
	}
	return
}

func initState4(x, y, z, w int) [][][][]rune {
	state := make([][][][]rune, w)
	for l := 0; l < w; l++ {
		state[l] = make([][][]rune, z)
		for k := 0; k < z; k++ {
			state[l][k] = make([][]rune, y)
			for j := 0; j < y; j++ {
				state[l][k][j] = make([]rune, x)
				for i := 0; i < x; i++ {
					state[l][k][j][i] = '.'
				}
			}
		}
	}
	return state
}

func countNeighbours4(state [][][][]rune, i, j, k, o int, x, y, z, w int, active rune) (activeCount int) {
	for l := i - 1; l <= i+1; l++ {
		for m := j - 1; m <= j+1; m++ {
			for n := k - 1; n <= k+1; n++ {
				for p := o - 1; p <= o+1; p++ {
					if !(l == i && m == j && n == k && p == o) && l > 0 && l < x && m > 0 && m < y && n > 0 && n < z && p > 0 && p < w && state[p][n][m][l] == active {
						activeCount++
					}
				}
			}
		}
	}
	return
}
