package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	nums   = []int64{}
	width  = 25
	height = 6
)

func main() {
	file, err := os.Open("8.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, str := range strings.Split(scanner.Text(), "") {
		num, _ := strconv.ParseInt(str, 10, 64)
		nums = append(nums, num)
	}

	counters := []map[int64]int{}
	layers := [][]int64{}
	for i := 0; i < len(nums); i += width * height {
		layer := append([]int64{}, nums[i:i+width*height]...)
		layers = append(layers, layer)
		counter := map[int64]int{}
		for _, num := range layer {
			counter[num]++
		}
		counters = append(counters, counter)
	}

	minZero := 0
	minZeroVal := width*height + 1
	for layer, counter := range counters {
		if counter[0] < minZeroVal {
			minZeroVal = counter[0]
			minZero = layer
		}
	}
	fmt.Println("Part One: ", counters[minZero][1]*counters[minZero][2])

	image := []int64{}
	l := 0
	for i, num := range layers[0] {
		for num == 2 {
			l++
			num = layers[l][i]
		}
		image = append(image, num)
		l = 0
	}
	fmt.Println("Part Two: ")
	for i := 0; i < len(image); i += width {
		for j := 0; j < width; j++ {
			if image[i+j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
