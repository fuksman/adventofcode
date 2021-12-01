package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	parents  map[string]int64
	children map[string]int64
}

func main() {
	file, err := os.Open("7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	rules := map[string]*node{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// faded purple bags contain 4 pale beige bags, 2 striped violet bags, 3 muted olive bags, 4 vibrant chartreuse bags.
		line := scanner.Text()
		line = line[:len(line)-1]
		line = strings.ReplaceAll(line, " bags", "")
		line = strings.ReplaceAll(line, " bag", "")
		tokens := strings.Split(line, " contain ")
		contents := strings.Split(tokens[1], ", ")
		if _, ok := rules[tokens[0]]; !ok {
			rules[tokens[0]] = &node{parents: make(map[string]int64), children: make(map[string]int64)}
		}
		for _, entry := range contents {
			if entry == "no other" {
				continue
			}
			weight, chtype := int64(entry[0])-int64('0'), entry[2:]
			rules[tokens[0]].children[chtype] += weight
			if _, ok := rules[chtype]; !ok {
				rules[chtype] = &node{parents: make(map[string]int64), children: make(map[string]int64)}
			}
			rules[chtype].parents[tokens[0]] += weight
		}
	}

	find := "shiny gold"
	parents := findBagParents(rules, find)
	fmt.Println("Part One: ", len(parents))
	fmt.Println("Part Two: ", countBagChildren(rules, find))
}

func findBagParents(rules map[string]*node, find string) map[string]int64 {
	n := rules[find]
	result := map[string]int64{}
	for name, weight := range n.parents {
		result[name] += weight
		for n, w := range findBagParents(rules, name) {
			result[n] += w
		}
	}
	return result
}

func countBagChildren(rules map[string]*node, find string) int64 {
	n := rules[find]
	var res int64
	for name, weight := range n.children {
		res += weight + weight*countBagChildren(rules, name)
	}
	return res
}
