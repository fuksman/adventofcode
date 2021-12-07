package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := ""
	scanner.Scan()
	input = scanner.Text()
	nums := []int64{}
	sum := int64(0)
	fields := strings.Split(input, ",")
	for _, field := range fields {
		num, _ := strconv.ParseInt(field, 10, 64)
		sum += num
		nums = append(nums, num)
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	mean := sum / int64(len(nums))
	median := int64(0)
	if n := len(nums); n%2 == 0 {
		median = (nums[n/2-1] + nums[n/2]) / 2
	} else {
		median = nums[(n-1)/2]
	}

	fuelMedian, fuelMean := int64(0), int64(0)
	for _, num := range nums {
		fuelMedian += abs(num - median)
		for i := 1; i <= int(abs(num-mean)); i++ {
			fuelMean += int64(i)
		}

	}

	fmt.Println("Part One: ", fuelMedian)
	fmt.Println("Part Two: ", mean, fuelMean)

}

func abs(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}
