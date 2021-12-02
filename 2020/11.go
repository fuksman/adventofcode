package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("11.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]byte{}
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	fmt.Println("Part One: ", move(&lines, around, 4))
	fmt.Println("Part Two: ", move(&lines, seenAround, 5))
}

func move(layout *[][]byte, checker func(int, int, int, int, byte, *[][]byte) int, threshold int) int {
	occupied := 0
	changed := true
	lines := *layout
	for changed {
		temp := [][]byte{}
		occupied = 0
		changed = false
		for i, line := range lines {
			temp = append(temp, []byte{})
			for j, ch := range line {
				temp[i] = append(temp[i], ch)
				switch ch {
				case '.':
					continue
				case 'L':
					if checker(i, j, len(lines), len(line), '#', &lines) == 0 {
						temp[i][j] = '#'
						occupied++
						changed = true
						continue
					}
				case '#':
					occupied++
					if checker(i, j, len(lines), len(line), '#', &lines) >= threshold {
						temp[i][j] = 'L'
						occupied--
						changed = true
						continue
					}
				}
			}
		}
		lines = temp
	}
	return occupied
}

func around(r, c, maxR, maxC int, sym byte, lines *[][]byte) int {
	count := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i >= 0 && i < maxR && j >= 0 && !(i == r && j == c) && j < maxC && (*lines)[i][j] == sym {
				count++
			}
		}
	}
	return count
}

func seenAround(r, c, maxR, maxC int, sym byte, lines *[][]byte) int {
	count := 0
	// x ? ?
	// ? + ?
	// ? ? ?
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	// ? ? x
	// ? + ?
	// ? ? ?
	for i, j := r-1, c+1; i >= 0 && j < maxC; i, j = i-1, j+1 {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	// ? ? ?
	// ? + ?
	// ? ? x
	for i, j := r+1, c+1; i < maxR && j < maxC; i, j = i+1, j+1 {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	// ? ? ?
	// ? + ?
	// x ? ?
	for i, j := r+1, c-1; i < maxR && j >= 0; i, j = i+1, j-1 {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	// ? x ?
	// ? + ?
	// ? ? ?
	for i, j := r-1, c; i >= 0; i-- {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	// ? ? ?
	// ? + x
	// ? ? ?
	for i, j := r, c+1; j < maxC; j++ {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	// ? ? ?
	// ? + ?
	// ? x ?
	for i, j := r+1, c; i < maxR; i++ {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	// ? ? ?
	// x + ?
	// ? ? ?
	for i, j := r, c-1; j >= 0; j-- {
		if (*lines)[i][j] == '.' {
			continue
		} else if (*lines)[i][j] == sym {
			count++
			break
		} else {
			break
		}
	}

	return count
}
