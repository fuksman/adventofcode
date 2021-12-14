package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	polymer     string
	instuctions = map[string]string{}
)

func main() {
	file, err := os.Open("14.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer = scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " -> ")
		instuctions[tokens[0]] = tokens[1]
	}

	pairs := map[string]int{}
	for i := 0; i < len(polymer)-1; i++ {
		pairs[polymer[i:i+2]]++
	}
	for d := 0; d < 10; d++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			newPairs[string(k[0])+instuctions[k]] += v
			newPairs[instuctions[k]+string(k[1])] += v
		}
		pairs = newPairs
	}
	counters := map[rune]int{}
	counters[rune(polymer[0])]++
	for k, v := range pairs {
		counters[rune(k[1])] += v
	}
	cnts := []int{}
	for _, v := range counters {
		cnts = append(cnts, v)
	}
	sort.Ints(cnts)
	fmt.Println("Part One: ", cnts[len(cnts)-1]-cnts[0])

	for d := 10; d < 40; d++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			newPairs[string(k[0])+instuctions[k]] += v
			newPairs[instuctions[k]+string(k[1])] += v
		}
		pairs = newPairs
	}
	counters = map[rune]int{}
	counters[rune(polymer[0])]++
	for k, v := range pairs {
		counters[rune(k[1])] += v
	}
	cnts = []int{}
	for _, v := range counters {
		cnts = append(cnts, v)
	}
	sort.Ints(cnts)
	fmt.Println("Part Two: ", cnts[len(cnts)-1]-cnts[0])
}
