package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	openToClose map[byte]byte
)

func main() {
	file, err := os.Open("10.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	points := map[byte]int{
		')': 3,
		']': 57,
		'>': 25137,
		'}': 1197,
	}
	scorePoints := map[byte]int{
		')': 1,
		']': 2,
		'>': 4,
		'}': 3,
	}
	openToClose = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	corrupted := []string{}
	illegal := map[byte]int{}
	// endings := []string{}
	scores := []int{}
	for _, line := range lines {
		opt := optimize(line)
		corrupted = append(corrupted, opt)
		if idx := strings.IndexAny(opt, ")]}>"); idx != -1 {
			illegal[opt[idx]]++
		} else {
			// ending := ""
			score := 0
			for i := len(opt) - 1; i >= 0; i-- {
				// ending += string(openToClose[opt[i]])
				score = score*5 + scorePoints[openToClose[opt[i]]]
			}
			scores = append(scores, score)
		}
	}
	sum := 0
	for ch, cnt := range illegal {
		sum += points[ch] * cnt
	}
	fmt.Println("Part One: ", sum)

	sort.Ints(scores)
	fmt.Println("Part Two: ", scores[(len(scores)-1)/2])
}

func optimize(s string) string {
	for i := 0; i < len(s)-1; {
		if openToClose[s[i]] == s[i+1] {
			s = s[:i] + s[i+2:]
			if i != 0 {
				i--
			}
		} else {
			i++
		}
	}
	return s
}
