package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	possibleIngridients map[string]map[string]int
	ingridientsCount    map[string]int
	maxCount            map[string]int
)

func main() {
	file, err := os.Open("21.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	possibleIngridients = map[string]map[string]int{}
	ingridientsCount = map[string]int{}
	maxCount = map[string]int{}
	for _, line := range lines {
		parts := strings.Split(line, " (contains ")
		ingridients := strings.Fields(parts[0])
		allergens := strings.Split(parts[1][:len(parts[1])-1], ", ")
		for _, allergen := range allergens {
			if _, ok := possibleIngridients[allergen]; !ok {
				possibleIngridients[allergen] = map[string]int{}
			}
			for _, ingridient := range ingridients {
				possibleIngridients[allergen][ingridient]++
				if possibleIngridients[allergen][ingridient] > maxCount[allergen] {
					maxCount[allergen] = possibleIngridients[allergen][ingridient]
				}
			}
		}
		for _, ingridient := range ingridients {
			ingridientsCount[ingridient]++
		}
	}

	possibilities := map[string]bool{}
	for allergen, ingridients := range possibleIngridients {
		for ingingridient, count := range ingridients {
			if count < maxCount[allergen] {
				delete(possibleIngridients[allergen], ingingridient)
			} else {
				possibilities[ingingridient] = true
			}
		}
	}

	sum := 0
	for ingridient, count := range ingridientsCount {
		if !possibilities[ingridient] {
			sum += count
		}
	}

	fmt.Println("Part One: ", sum)

	sure := map[string]string{}
	sureAlgs := []string{}
	sureIngs := []string{}
	for len(possibleIngridients) != 0 {
		for allergen, ingridients := range possibleIngridients {
			if len(ingridients) == 1 {
				for ingridient := range ingridients {
					sureAlgs = append(sureAlgs, allergen)
					sure[ingridient] = allergen
					delete(possibleIngridients, allergen)
				}
			}
			for ingridient := range ingridients {
				if _, ok := sure[ingridient]; ok {
					delete(possibleIngridients[allergen], ingridient)
				}
			}
		}
	}
	sort.Strings(sureAlgs)
	for _, alg := range sureAlgs {
		for i, a := range sure {
			if a == alg {
				sureIngs = append(sureIngs, i)
			}
		}
	}
	fmt.Println("Part Two: ", strings.Join(sureIngs, ","))
}
