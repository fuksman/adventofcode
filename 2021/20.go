package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type cor [2]int

var (
	rules      = []bool{}
	input      = map[cor]bool{}
	output     = map[cor]bool{}
	margin     = 100
	rows, cols int
)

func main() {
	file, err := os.Open("20.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, ch := range scanner.Text() {
		rules = append(rules, ch == '#')
	}
	scanner.Scan()
	i := 0
	for scanner.Scan() {
		for j, ch := range scanner.Text() {
			input[cor{i, j}] = ch == '#'
			cols = j + 1
		}
		i++
	}
	rows = i

	c := 0
	for ; c < 2; c++ {
		enhance()
	}

	cnt := 0
	safe := c + 20
	for coord, v := range input {
		if !(coord[0] > -safe && coord[0] < rows+safe && coord[1] > -safe && coord[1] < cols+safe) {
			continue
		}
		if v {
			cnt++
		}
	}
	fmt.Println("Part One: ", cnt)

	for ; c < 50; c++ {
		enhance()
	}

	cnt = 0
	safe = c + 2
	for coord, v := range input {
		if !(coord[0] > -safe && coord[0] < rows+safe && coord[1] > -safe && coord[1] < cols+safe) {
			continue
		}
		if v {
			cnt++
		}
	}
	fmt.Println("Part One: ", cnt)
}

func enhance() {
	output = map[cor]bool{}
	for i := -margin; i < rows+margin; i++ {
		for j := -margin; j < cols+margin; j++ {
			str := ""
			for _, point := range []cor{{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1}, {i, j - 1}, {i, j}, {i, j + 1}, {i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}} {
				if input[point] {
					str += "1"
				} else {
					str += "0"
				}
			}
			idx, _ := strconv.ParseInt(str, 2, 64)
			output[cor{i, j}] = rules[idx]
		}
	}

	input = map[cor]bool{}
	for k, v := range output {
		input[k] = v
	}
}
