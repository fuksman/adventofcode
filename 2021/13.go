package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	dots           = [][]int{}
	paper          = [][]bool{}
	instruction    = []string{}
	lenght, height int
)

func main() {
	file, err := os.Open("13.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		tokens := strings.Split(scanner.Text(), ",")
		temp := []int{}
		for _, token := range tokens {
			num, _ := strconv.ParseInt(token, 10, 64)
			temp = append(temp, int(num))
		}
		if temp[0] > lenght {
			lenght = temp[0]
		}
		if temp[1] > height {
			height = temp[1]
		}
		dots = append(dots, temp)
	}

	for scanner.Scan() {
		instruction = append(instruction, scanner.Text())
	}

	for j := 0; j <= height; j++ {
		paper = append(paper, []bool{})
		for i := 0; i <= lenght; i++ {
			paper[j] = append(paper[j], false)
		}
	}

	for _, dot := range dots {
		paper[dot[1]][dot[0]] = true
	}

	fold := 0
	cnt := 0
	for _, inst := range instruction {
		fold++
		line, _ := strconv.ParseInt(inst[13:], 10, 64)
		switch inst[11] {
		case 'x':
			foldLeft(int(line))
		case 'y':
			foldUp(int(line))
		}
		if fold == 1 {
			for _, line := range paper {
				for _, dot := range line {
					if dot {
						cnt++
					}
				}
			}
			fmt.Println("Part One: ", cnt)
		}
	}
	fmt.Println("Part Two: ")
	for _, line := range paper {
		for _, dot := range line {
			if dot {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func foldUp(y int) {
	for j := y + 1; j < len(paper); j++ {
		for i := 0; i < len(paper[0]); i++ {
			paper[y-(j-y)][i] = paper[y-(j-y)][i] || paper[j][i]
		}
	}
	paper = paper[:y]
}

func foldLeft(x int) {
	for j := 0; j < len(paper); j++ {
		for i := x + 1; i < len(paper[j]); i++ {
			paper[j][x-(i-x)] = paper[j][x-(i-x)] || paper[j][i]
		}
		paper[j] = paper[j][:x]
	}
}
