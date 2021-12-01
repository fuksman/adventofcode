package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	geo := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		geo = append(geo, scanner.Text())
	}

	counters := []int{0, 0, 0, 0, 0}
	idxs := []int{0, 0, 0, 0, 0}
	rights := []int{1, 3, 5, 7, 1}
	for i, str := range geo {
		for j := 0; j < len(counters)-1; j++ {
			ch := str[idxs[j]]
			if ch == '#' {
				counters[j]++
			}
			idxs[j] = (idxs[j] + rights[j]) % len(str)
		}
		if i%2 != 1 {
			j := len(counters) - 1
			ch := str[idxs[j]]
			if ch == '#' {
				counters[j]++
			}
			idxs[j] = (idxs[j] + rights[j]) % len(str)
		}
	}
	res := 1
	for _, cnt := range counters {
		res *= cnt
	}

	fmt.Println("Part one:", counters[1])
	fmt.Println("Part two:", res)
}
