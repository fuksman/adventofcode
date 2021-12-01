package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("5.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var max int64
	seats := make([][]bool, 128)
	for i := range seats {
		seats[i] = make([]bool, 8)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row, column := convert(scanner.Text())
		seats[row][column] = true
		id := row*8 + column
		if id > max {
			max = id
		}
	}
	fmt.Println("Part One: ", max)

	for i := range seats {
		for j := range seats[i] {
			if !seats[i][j] && ((j < 7 && seats[i][j+1]) || (j > 0 && seats[i][j-1])) {
				fmt.Println("Part Two: ", i*8+j)
				break
			}
		}
	}
}

// FBFBBFF == 0101100
// RLR == 101
func convert(str string) (row int64, column int64) {
	var r, c string
	for _, ch := range str {
		switch ch {
		case 'F':
			r += "0"
		case 'B':
			r += "1"
		case 'L':
			c += "0"
		case 'R':
			c += "1"
		}
	}
	row, _ = strconv.ParseInt(r, 2, 64)
	column, _ = strconv.ParseInt(c, 2, 64)
	return
}
