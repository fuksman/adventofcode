package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord [2]int64

type Reaction struct {
	output      int64
	ingridients map[string]int64
}

var (
	reactions = map[string]*Reaction{}
	list      = []string{}
)

func main() {
	file, err := os.Open("14.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), " => ")
		output := strings.Fields(sides[1])
		num, _ := strconv.ParseInt(output[0], 10, 64)
		reaction := Reaction{output: num, ingridients: map[string]int64{}}
		reactions[output[1]] = &reaction
		ingridients := strings.Split(sides[0], ", ")
		for _, ingridient := range ingridients {
			input := strings.Fields(ingridient)
			num, _ := strconv.ParseInt(input[0], 10, 64)
			reaction.ingridients[input[1]] = num
		}
	}

	fmt.Println("Part One: ", ore(map[string]int64{"FUEL": 1}))
	fmt.Println("Part Two: ", sort.Search(1000000000000, func(n int) bool {
		return ore(map[string]int64{"FUEL": int64(n)}) > 1000000000000
	})-1)
}

func ore(needed map[string]int64) int64 {
	for ingridient := range needed {
		if ingridient != "ORE" && needed[ingridient] > 0 {
			amount := (needed[ingridient]-1)/reactions[ingridient].output + 1
			needed[ingridient] -= reactions[ingridient].output * amount

			for r := range reactions[ingridient].ingridients {
				needed[r] += reactions[ingridient].ingridients[r] * amount
			}
			ore(needed)
			break
		}
	}
	return needed["ORE"]
}
