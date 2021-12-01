package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("6.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	type group struct {
		answer map[rune]int
		size   int
	}
	answers := []group{}
	g := group{}
	g.answer = map[rune]int{}
	countA, countE := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			g.size++
			for _, ch := range line {
				g.answer[ch]++
			}
		} else {
			countA += len(g.answer)
			for _, ans := range g.answer {
				if ans == g.size {
					countE++
				}
			}
			answers = append(answers, g)
			g = group{}
			g.answer = map[rune]int{}
		}
	}
	countA += len(g.answer)
	for _, ans := range g.answer {
		if ans == g.size {
			countE++
		}
	}
	answers = append(answers, g)

	fmt.Println("Part One: ", countA)
	fmt.Println("Part Two: ", countE)
}
