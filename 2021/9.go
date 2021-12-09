package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	location [][]int64
	counted  [][]bool
)

func main() {
	file, err := os.Open("9.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		temp := []int64{}
		cnt := []bool{}
		for _, loc := range line {
			num, _ := strconv.ParseInt(string(loc), 10, 64)
			temp = append(temp, num)
			cnt = append(cnt, false)
		}
		location = append(location, temp)
		counted = append(counted, cnt)
	}

	lows := []int64{}
	sum := int64(0)
	n := len(location)
	l := len(location[0])
	for i, line := range location {
		for j, num := range line {
			if !((i-1 >= 0 && location[i-1][j] <= num) || (i+1 < n && location[i+1][j] <= num) || (j-1 >= 0 && location[i][j-1] <= num) || (j+1 < l && location[i][j+1] <= num)) {
				lows = append(lows, int64(basin(i, j)))
				sum += num + 1
			}
		}
	}
	fmt.Println("Part One: ", sum)
	sort.Slice(lows, func(i, j int) bool { return lows[i] > lows[j] })
	prod := int64(1)
	for i := 0; i < 3; i++ {
		prod *= lows[i]
	}
	fmt.Println("Part Two: ", prod)
}

func basin(i, j int) int {
	counted[i][j] = true
	size := 1
	n := len(location)
	l := len(location[0])
	if i-1 >= 0 && location[i-1][j] != 9 && !counted[i-1][j] && location[i-1][j] > location[i][j] {
		size += basin(i-1, j)
	}
	if i+1 < n && location[i+1][j] != 9 && !counted[i+1][j] && location[i+1][j] > location[i][j] {
		size += basin(i+1, j)
	}
	if j-1 >= 0 && location[i][j-1] != 9 && !counted[i][j-1] && location[i][j-1] > location[i][j] {
		size += basin(i, j-1)
	}
	if j+1 < l && location[i][j+1] != 9 && !counted[i][j+1] && location[i][j+1] > location[i][j] {
		size += basin(i, j+1)
	}
	return size
}
