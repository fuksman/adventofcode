package main

import "fmt"

func main() {
	nums := []int{15, 12, 0, 14, 3, 1}
	game := map[int][]int{}
	for i, num := range nums {
		game[num] = append(game[num], i)
	}
	for i := len(nums); i < 30000000; i++ {
		prev := nums[i-1]
		newNum := 0
		if len(game[prev]) != 1 {
			newNum = i - game[prev][len(game[prev])-2] - 1
		}
		if i == 4 {
			fmt.Println(nums, newNum)
			fmt.Println(game, prev, game[prev])
		}
		nums = append(nums, newNum)
		game[newNum] = append(game[newNum], i)
	}

	fmt.Println("Part One: ", nums[2019])
	fmt.Println("Part Two: ", nums[len(nums)-1])
}
