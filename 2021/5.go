package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("5.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]int{}
	maxpoint := 0
	for scanner.Scan() {
		nums := []int{}
		line := scanner.Text()
		points := strings.Split(line, " -> ")
		for _, point := range points {
			coords := strings.Split(point, ",")
			for _, coord := range coords {
				num, _ := strconv.ParseInt(coord, 10, 64)
				nums = append(nums, int(num))
				if int(num) > maxpoint {
					maxpoint = int(num)
				}
			}
		}
		lines = append(lines, nums)
	}

	counters := make([][]int, maxpoint+1)
	for i := range counters {
		counters[i] = make([]int, maxpoint+1)
	}
	for _, line := range lines {
		if line[0] == line[2] {
			for j := min(line[1], line[3]); j <= max(line[1], line[3]); j++ {
				counters[line[0]][j]++
			}
		}
		if line[1] == line[3] {
			for i := min(line[0], line[2]); i <= max(line[0], line[2]); i++ {
				counters[i][line[1]]++
			}
		}
	}

	counter := 0
	for i := range counters {
		for j := range counters[i] {
			if counters[i][j] > 1 {
				counter++
			}
		}
	}
	fmt.Println("Part One: ", counter)

	for _, line := range lines {
		if line[0] < line[2] {
			if line[1] < line[3] {
				for i, j := line[0], line[1]; i <= line[2] && j <= line[3]; i, j = i+1, j+1 {
					counters[i][j]++
				}
			}
			if line[1] > line[3] {
				for i, j := line[0], line[1]; i <= line[2] && j >= line[3]; i, j = i+1, j-1 {
					counters[i][j]++
				}
			}
		}
		if line[0] > line[2] {
			if line[1] < line[3] {
				for i, j := line[0], line[1]; i >= line[2] && j <= line[3]; i, j = i-1, j+1 {
					counters[i][j]++
				}
			}
			if line[1] > line[3] {
				for i, j := line[0], line[1]; i >= line[2] && j >= line[3]; i, j = i-1, j-1 {
					counters[i][j]++
				}
			}
		}
	}

	counter = 0
	for i := range counters {
		for j := range counters[i] {
			if counters[i][j] > 1 {
				counter++
			}
		}
	}
	fmt.Println("Part Two: ", counter)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
