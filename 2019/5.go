package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	nums   []int64
	input  = int64(5)
	output int64
)

func main() {
	file, err := os.Open("5.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")
	for _, token := range tokens {
		num, _ := strconv.ParseInt(token, 10, 64)
		nums = append(nums, num)
	}
	jump := 4
	for i := 0; i < len(nums) && nums[i] != 99; i += jump {
		operation := nums[i] % 100
		mode2 := (nums[i] % 10000) / 1000
		mode1 := (nums[i] % 1000) / 100
		switch operation {
		case 1:
			nums[nums[i+3]] = value(nums[i+1], mode1) + value(nums[i+2], mode2)
			jump = 4
		case 2:
			nums[nums[i+3]] = value(nums[i+1], mode1) * value(nums[i+2], mode2)
			jump = 4
		case 3:
			nums[nums[i+1]] = input
			jump = 2
		case 4:
			output = nums[nums[i+1]]
			jump = 2
		case 5:
			if value(nums[i+1], mode1) != 0 {
				jump = int(value(nums[i+2], mode2)) - i
			} else {
				jump = 3
			}
		case 6:
			if value(nums[i+1], mode1) == 0 {
				jump = int(value(nums[i+2], mode2)) - i
			} else {
				jump = 3
			}
		case 7:
			if value(nums[i+1], mode1) < value(nums[i+2], mode2) {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			jump = 4
		case 8:
			if value(nums[i+1], mode1) == value(nums[i+2], mode2) {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			jump = 4
		case 99:
			break
		}
	}
	fmt.Println("Part One: ", output)
}

func value(v, mode int64) int64 {
	if mode == 0 {
		return nums[v]
	}
	return int64(v)
}
