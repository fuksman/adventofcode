package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("9.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := []int64{}
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		nums = append(nums, num)
	}

	invalid, idx := test(&nums, 25)
	fmt.Println("Part One: ", invalid)

	for i := idx - 1; i > 0; i-- {
		temp := []int64{nums[i]}
		sum := nums[i]
		j := i - 1
		for ; j >= 0 && sum < invalid; j-- {
			sum += nums[j]
			temp = append(temp, nums[j])
		}
		if sum == invalid {
			sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
			fmt.Println("Part Two: ", temp[0]+temp[len(temp)-1])
			break
		}
	}
}

func test(nums *[]int64, dif int) (num int64, idx int) {
	for i := dif; i < len(*nums); i++ {
		valid := false
		for j := i - dif; j < i-1 && !valid; j++ {
			for k := j + 1; k < i && !valid; k++ {
				if (*nums)[i] == (*nums)[j]+(*nums)[k] {
					valid = true
					break
				}
			}
		}
		if !valid {
			return (*nums)[i], i
		}
	}
	return 0, -1
}
