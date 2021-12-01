package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	acc, _ := execute(lines)
	fmt.Println("Part One: ", acc)

	for i := 0; i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		var acc int64
		var valid bool
		switch fields[0] {
		case "jmp":
			mod := make([]string, len(lines))
			copy(mod, lines)
			mod[i] = "nop " + fields[1]
			acc, valid = execute(mod)
		case "nop":
			mod := make([]string, len(lines))
			copy(mod, lines)
			mod[i] = "jmp " + fields[1]
			acc, valid = execute(mod)
		default:
			continue
		}
		if valid {
			fmt.Println("Part Two: ", acc)
			break
		}
	}
}

func execute(lines []string) (acc int64, valid bool) {
	visited := map[int]bool{}
	for i := 0; i < len(lines); {
		fields := strings.Fields(lines[i])
		if _, ok := visited[i]; !ok {
			visited[i] = true
		} else {
			return acc, false
		}

		diff, _ := strconv.ParseInt(fields[1], 10, 32)
		switch fields[0] {
		case "acc":
			acc += diff
			i++
		case "jmp":
			i += int(diff)
		case "nop":
			i++
		}
	}
	return acc, true
}
