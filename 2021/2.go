package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var x, y int64
	for _, line := range lines {
		fields := strings.Fields(line)
		num, _ := strconv.ParseInt(fields[1], 10, 64)
		switch fields[0] {
		case "forward":
			x += num
		case "up":
			y -= num
		case "down":
			y += num
		}
	}
	fmt.Println("Part One: ", x*y)

	x, y = 0, 0
	var aim int64
	for _, line := range lines {
		fields := strings.Fields(line)
		num, _ := strconv.ParseInt(fields[1], 10, 64)
		switch fields[0] {
		case "forward":
			x += num
			y += aim * num
		case "up":
			aim -= num
		case "down":
			aim += num
		}
	}
	fmt.Println("Part Two: ", x*y)
}
