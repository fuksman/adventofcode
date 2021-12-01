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
	nums, _ := twoSum(measurements, 2020)
	fmt.Println(measurements[nums[0]] * measurements[nums[1]])
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
