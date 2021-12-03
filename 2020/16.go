package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("16.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := map[string][][]int64{}
	ruleNames := []string{}
	for scanner.Scan() && scanner.Text() != "" {
		rule := scanner.Text()
		tokens := strings.Split(rule, ": ")
		ranges := strings.Split(tokens[1], " or ")
		for _, r := range ranges {
			borders := strings.Split(r, "-")
			b1, _ := strconv.ParseInt(borders[0], 10, 64)
			b2, _ := strconv.ParseInt(borders[1], 10, 64)
			rules[tokens[0]] = append(rules[tokens[0]], []int64{b1, b2})
		}
		ruleNames = append(ruleNames, tokens[0])
	}

	myticket := []int64{}
	scanner.Scan()
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")
	for _, token := range tokens {
		field, _ := strconv.ParseInt(token, 10, 64)
		myticket = append(myticket, field)
	}

	tickets := [][]int64{}
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		ticket := []int64{}
		tokens := strings.Split(scanner.Text(), ",")
		for _, token := range tokens {
			field, _ := strconv.ParseInt(token, 10, 64)
			ticket = append(ticket, field)
		}
		tickets = append(tickets, ticket)
	}

	sum := int64(0)
	for i := 0; i < len(tickets); i++ {
		ticket := tickets[i]
		for _, field := range ticket {
			valid := false
			for _, rule := range rules {
				if (field >= rule[0][0] && field <= rule[0][1]) || (field >= rule[1][0] && field <= rule[1][1]) {
					valid = true
					break
				}
			}
			if !valid {
				sum += field
				tickets = append(tickets[:i], tickets[i+1:]...)
				i--
			}
		}
	}

	fmt.Println("Part One: ", sum)

	rulesPositions := map[string][]int{}
	cleared := map[string]bool{}
	tickets = append(tickets, myticket)
	for name, rule := range rules {
		// cleared[name] = false
		for j := range tickets[0] {
			matched := true
			for _, ticket := range tickets {
				field := ticket[j]
				if !((field >= rule[0][0] && field <= rule[0][1]) || (field >= rule[1][0] && field <= rule[1][1])) {
					matched = false
					break
				}
			}
			if matched {
				rulesPositions[name] = append(rulesPositions[name], j)
			}
		}
		if len(rulesPositions[name]) == 1 {
			cleared[name] = true
		}
	}

	// allCleared := false
	for len(cleared) < len(ruleNames) {
		for name, clear := range cleared {
			if clear {
				position := rulesPositions[name][0]
				for r, positions := range rulesPositions {
					if !cleared[r] {
						for i, p := range positions {
							if p == position {
								rulesPositions[r] = append(rulesPositions[r][:i], rulesPositions[r][i+1:]...)
							}
						}
						if len(rulesPositions[r]) == 1 {
							cleared[r] = true
						}
					}
				}
			}
		}
	}

	prod := int64(1)
	for _, name := range ruleNames {
		if strings.HasPrefix(name, "departure") {
			prod *= myticket[rulesPositions[name][0]]
		}
	}

	fmt.Println("Part Two: ", prod)

	// fmt.Println(rules)
	// fmt.Println(myticket)
	// fmt.Println(tickets)
	// fmt.Println(rulesPositions)
	// fmt.Println(cleared)
}
