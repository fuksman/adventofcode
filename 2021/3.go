package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]rune{}
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	gamma, epsilon := []rune{}, []rune{}
	counters := count(&lines)

	for _, counter := range counters {
		if counter['0'] > counter['1'] {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		} else {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		}
	}
	gammaD, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonD, _ := strconv.ParseInt(string(epsilon), 2, 64)
	fmt.Println("Part One: ", gammaD*epsilonD)

	oxygen, co2 := append([][]rune{}, lines...), append([][]rune{}, lines...)

	for i := 0; len(oxygen) != 1 && i < len(lines[0]); i++ {
		temp := [][]rune{}
		counters := count(&oxygen)
		var max rune
		if counters[i]['0'] > counters[i]['1'] {
			max = '0'
		} else {
			max = '1'
		}
		for l := 0; l < len(oxygen); l++ {
			if oxygen[l][i] == max {
				temp = append(temp, oxygen[l])
			}
		}
		oxygen = append([][]rune{}, temp...)
	}

	for i := 0; len(co2) != 1 && i < len(lines[0]); i++ {
		temp := [][]rune{}
		counters := count(&co2)
		var max rune
		if counters[i]['0'] > counters[i]['1'] {
			max = '1'
		} else {
			max = '0'
		}
		for l := 0; l < len(co2); l++ {
			if co2[l][i] == max {
				temp = append(temp, co2[l])
			}
		}
		co2 = append([][]rune{}, temp...)
	}
	oxygenD, _ := strconv.ParseInt(string(oxygen[0]), 2, 64)
	co2D, _ := strconv.ParseInt(string(co2[0]), 2, 64)
	fmt.Println("Part Two: ", oxygenD*co2D)
}

func count(lines *[][]rune) []map[rune]int {
	counters := []map[rune]int{}
	for range (*lines)[0] {
		m := map[rune]int{'0': 0, '1': 0}
		counters = append(counters, m)
	}
	for _, line := range *lines {
		for i, sym := range line {
			counters[i][sym] += 1
		}
	}
	return counters
}
