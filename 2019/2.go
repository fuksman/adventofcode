package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	nums []int64
)

func main() {
	file, err := os.Open("2.txt")
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
	p1, p2 := int64(65), int64(77)
	nums[1], p1 = p1, nums[1]
	nums[2], p2 = p2, nums[2]
	for i := 0; i < len(nums) && nums[i] != 99; i += 4 {
		switch nums[i] {
		case 1:
			nums[nums[i+3]] = nums[nums[i+1]] + nums[nums[i+2]]
		case 2:
			nums[nums[i+3]] = nums[nums[i+1]] * nums[nums[i+2]]
		case 99:
			break
		}
	}
	fmt.Println("Part One: ", nums[0])
}
