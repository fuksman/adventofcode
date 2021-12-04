package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	input := []int64{}
	markings := []map[int64]bool{}
	mappings := []map[int64][]int{}
	rowcounters := [][]int{}
	columncounters := [][]int{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")
	for _, token := range tokens {
		num, _ := strconv.ParseInt(token, 10, 64)
		input = append(input, num)
	}

	scanner.Scan()
	for scanner.Scan() {
		ticket := [][]int64{}
		marking := map[int64]bool{}
		mapping := map[int64][]int{}
		i := 0
		for scanner.Text() != "" {
			fields := strings.Fields(scanner.Text())
			row := []int64{}
			for j, field := range fields {
				num, _ := strconv.ParseInt(field, 10, 64)
				row = append(row, num)
				marking[num] = false
				mapping[num] = []int{i, j}
			}
			i++
			ticket = append(ticket, row)
			scanner.Scan()
		}
		markings = append(markings, marking)
		mappings = append(mappings, mapping)
		rowcounters = append(rowcounters, make([]int, len(ticket)))
		columncounters = append(columncounters, make([]int, len(ticket)))
	}

	win := []int{}
	won := map[int]bool{}
	winnum := []int{}
	for n := 0; len(won) != len(mappings) && n < len(input); n++ {
		num := input[n]
		for t := 0; t < len(mappings); t++ {
			if won[t] {
				continue
			}
			ticket := mappings[t]
			if mapping, ok := ticket[num]; ok {
				markings[t][num] = true
				rowcounters[t][mapping[0]]++
				columncounters[t][mapping[1]]++
				if rowcounters[t][mapping[0]] == 5 || columncounters[t][mapping[1]] == 5 {
					win = append(win, t)
					winnum = append(winnum, int(num))
					won[t] = true
				}
			}
		}
	}

	sum := 0
	for num, check := range markings[win[0]] {
		if !check {
			sum += int(num)
		}
	}

	fmt.Println("Part One: ", sum*winnum[0])

	sum = 0
	for num, check := range markings[win[len(win)-1]] {
		if !check {
			sum += int(num)
		}
	}

	fmt.Println("Part Two: ", sum*winnum[len(winnum)-1])
}
