package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	nums    [][]int64
	flashed [][]bool
)

func main() {
	file, err := os.Open("11.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums = [][]int64{}
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "")
		line := []int64{}
		for _, token := range tokens {
			num, _ := strconv.ParseInt(token, 10, 64)
			line = append(line, num)
		}
		nums = append(nums, line)
	}

	counter := 0
	synced := -1
	for c := 0; synced == -1; c++ {
		flashed = make([][]bool, len(nums))
		for i := range flashed {
			flashed[i] = make([]bool, len(nums[i]))
		}
		for i, line := range nums {
			for j := range line {
				counter += flash(i, j)
			}
		}
		if c == 99 {
			fmt.Println("Part One: ", counter)
		}

		synced = c
		for i := range flashed {
			if synced == -1 {
				break
			}
			for j := range flashed[i] {
				if !flashed[i][j] {
					synced = -1
					break
				}
			}
		}
	}
	fmt.Println("Part Two: ", synced+1)
}

func flash(i, j int) (cnt int) {
	if flashed[i][j] {
		return
	}
	nums[i][j]++
	if nums[i][j] > 9 {
		flashed[i][j] = true
		cnt++
		nums[i][j] = 0
		for k := i - 1; k <= i+1; k++ {
			for l := j - 1; l <= j+1; l++ {
				if !(k == i && l == j) && k >= 0 && k < len(nums) && l >= 0 && l < len(nums[i]) {
					cnt += flash(k, l)
				}
			}
		}
	}
	return
}
