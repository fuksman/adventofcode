package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	mapping = [][]rune{}
)

func main() {
	file, err := os.Open("25.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mapping = append(mapping, []rune(scanner.Text()))
	}

	step := 0
	for moved := true; moved; step++ {
		moved = false
		tempE := make([][]rune, len(mapping))
		for t := range tempE {
			tempE[t] = []rune(strings.Repeat(".", len(mapping[0])))
		}

		for i := 0; i < len(mapping); i++ {
			for j := 0; j < len(mapping[0]); j++ {
				sym := mapping[i][j]
				switch sym {
				case 'v':
					tempE[i][j] = sym
				case '>':
					next := j + 1
					if next-len(mapping[0]) == 0 {
						next = 0
					}
					if mapping[i][next] == '.' {
						tempE[i][next] = sym
						moved = true
					} else {
						tempE[i][j] = sym
					}
				}
			}
		}
		mapping = tempE

		tempS := make([][]rune, len(mapping))
		for t := range tempS {
			tempS[t] = []rune(strings.Repeat(".", len(mapping[0])))
		}

		for i := 0; i < len(mapping); i++ {
			for j := 0; j < len(mapping[0]); j++ {
				sym := mapping[i][j]
				switch sym {
				case '>':
					tempS[i][j] = sym
				case 'v':
					next := i + 1
					if next-len(mapping) == 0 {
						next = 0
					}
					if mapping[next][j] == '.' {
						tempS[next][j] = sym
						moved = true
					} else {
						tempS[i][j] = sym
					}
				}
			}
		}
		mapping = tempS
	}
	fmt.Println("Part One: ", step)
}
