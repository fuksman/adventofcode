package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("6.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counters := map[int64]int{}
	scanner.Scan()
	line := scanner.Text()
	nums := strings.Split(line, ",")
	for _, sym := range nums {
		num, _ := strconv.ParseInt(sym, 10, 64)
		counters[num]++
	}

	sum := 0
	for c := 0; c < 80; c++ {
		sum = 0
		nextState := map[int64]int{}
		for days, count := range counters {
			sum += count
			if days == 0 {
				sum += count
				nextState[6] += count
				nextState[8] += count
			} else {
				nextState[days-1] += count
			}
		}
		counters = nextState
	}

	fmt.Println("Part One: ", sum)

	for c := 80; c < 256; c++ {
		sum = 0
		nextState := map[int64]int{}
		for days, count := range counters {
			sum += count
			if days == 0 {
				sum += count
				nextState[6] += count
				nextState[8] += count
			} else {
				nextState[days-1] += count
			}
		}
		counters = nextState
	}

	fmt.Println("Part Two: ", sum)
}
