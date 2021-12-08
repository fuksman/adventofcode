package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("8.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	count := 0
	lenToNum := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}
	for _, line := range lines {
		tokens := strings.Split(line, " | ")
		fields := strings.Fields(tokens[1])
		for _, field := range fields {
			n := len(field)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				count++
			}
		}
	}
	fmt.Println("Part One: ", count)

	sum := 0
	for _, line := range lines {
		tokens := strings.Split(line, " | ")
		signals := strings.Fields(tokens[0])
		digits := strings.Fields(tokens[1])
		numExamples := map[int][]rune{}
		maybe5, maybe6 := [][]rune{}, [][]rune{}
		for _, signal := range signals {
			n := len(signal)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				numExamples[lenToNum[n]] = sortRunes([]rune(signal))
			}
			if n == 5 {
				maybe5 = append(maybe5, sortRunes([]rune(signal)))
			}
			if n == 6 {
				maybe6 = append(maybe6, sortRunes([]rune(signal)))
			}
		}

		runesOfFour := diff(numExamples[4], numExamples[1])
		for _, len6 := range maybe6 {
			segs := diff(len6, numExamples[4])
			switch len(segs) {
			case 2:
				numExamples[9] = len6
			case 3:
				if strings.ContainsRune(string(len6), runesOfFour[0]) && strings.ContainsRune(string(len6), runesOfFour[1]) {
					numExamples[6] = len6
				} else {
					numExamples[0] = len6
				}
			}
		}

		for _, len5 := range maybe5 {
			if strings.ContainsRune(string(len5), numExamples[1][0]) && strings.ContainsRune(string(len5), numExamples[1][1]) {
				numExamples[3] = len5
			} else if strings.ContainsRune(string(len5), runesOfFour[0]) && strings.ContainsRune(string(len5), runesOfFour[1]) {
				numExamples[5] = len5
			} else {
				numExamples[2] = len5
			}
		}

		res := 0
		for i, digit := range digits {
			pos := int(math.Pow10(3 - i))
			num := -1
			r := sortRunes([]rune(digit))
			for n, ex := range numExamples {
				if string(ex) == string(r) {
					num = n
					break
				}
			}
			res += num * pos
		}
		sum += res
	}
	fmt.Println("Part Two: ", sum)
}

func sortRunes(str []rune) []rune {
	sort.Slice(str, func(i, j int) bool { return str[i] < str[j] })
	return str
}

func diff(str1, str2 []rune) (res []rune) {
	for _, r1 := range str1 {
		found := false
		for _, r2 := range str2 {
			if r1 == r2 {
				found = true
				break
			}
		}
		if !found {
			res = append(res, r1)
		}
	}
	return sortRunes(res)
}
