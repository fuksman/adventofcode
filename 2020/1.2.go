package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	measurements, err := readInts("1.txt")
	if err != nil {
		log.Fatal()
	}
	for i := 1; i < len(measurements); i++ {
		nums, err := twoSum(measurements[i:], 2020-measurements[i-1])
		if err == nil {
			fmt.Println(measurements[i-1] * measurements[i:][nums[0]] * measurements[i:][nums[1]])
			break
		}
	}
}

func readInts(filename string) ([]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ints := []int64{}
	for scanner.Scan() {
		current, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		ints = append(ints, current)
	}
	return ints, scanner.Err()
}

func twoSum(nums []int64, target int64) ([]int64, error) {
	hashmap := map[int64]int64{}
	for i, val := range nums {
		complement := target - val
		if res, ok := hashmap[complement]; ok {
			return []int64{int64(i), res}, nil
		}
		hashmap[val] = int64(i)
	}
	return nil, errors.New("haven't found nums for the target")
}
