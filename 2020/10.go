package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("10.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := []int64{0}
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		nums = append(nums, num)
	}
	sort.Slice(nums, func(i int, j int) bool { return nums[i] < nums[j] })
	nums = append(nums, nums[len(nums)-1]+3)
	jolts := map[int64]int{}
	for i := 1; i < len(nums); i++ {
		jolts[nums[i]-nums[i-1]] += 1
	}
	fmt.Println("Part One: ", jolts[1]*jolts[3])

	arr := map[string]int{}
	fmt.Println("Part Two: ", countOptions(nums[1:], 0, arr))
}

func countOptions(nums []int64, lastJolt int, arr map[string]int) int {
	str := makeMemoKey(nums, lastJolt)
	if v, ok := arr[str]; ok {
		return v
	}

	if len(nums) == 0 {
		return 1
	}

	var count int
	for i, v := range nums {
		if int(v)-lastJolt <= 3 {
			count += countOptions(nums[i+1:], int(v), arr)
		} else {
			break
		}
	}
	arr[str] = count

	return count
}

func makeMemoKey(nums []int64, lastJolt int) string {
	ans := strconv.Itoa(lastJolt) + "x"
	for _, v := range nums {
		ans += strconv.Itoa(int(v))
	}
	return ans
}
