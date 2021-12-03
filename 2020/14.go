package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("14.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	mask := []rune{}
	mem := map[int64]int64{}

	for _, line := range lines {
		tokens := strings.Split(line, " = ")
		if tokens[0] == "mask" {
			mask = []rune{}
			for _, ch := range tokens[1] {
				mask = append(mask, ch)
			}
		} else {
			adr, _ := strconv.ParseInt(tokens[0][4:len(tokens[0])-1], 10, 64)
			num, _ := strconv.ParseInt(tokens[1], 10, 64)
			numstr := []rune(strconv.FormatInt(num, 2))
			numstr = append([]rune(strings.Repeat("0", 36-len(numstr))), numstr...)
			for i, ch := range mask {
				if ch != 'X' {
					numstr[i] = ch
				}
			}
			num, _ = strconv.ParseInt(string(numstr), 2, 64)
			mem[adr] = num
		}
	}

	sum := int64(0)
	for _, v := range mem {
		sum += v
	}
	fmt.Println("Part One: ", sum)

	mem = map[int64]int64{}
	for _, line := range lines {
		tokens := strings.Split(line, " = ")
		if tokens[0] == "mask" {
			mask = []rune{}
			for _, ch := range tokens[1] {
				mask = append(mask, ch)
			}
		} else {
			adr, _ := strconv.ParseInt(tokens[0][4:len(tokens[0])-1], 10, 64)
			num, _ := strconv.ParseInt(tokens[1], 10, 64)
			adrstr := []rune(strconv.FormatInt(adr, 2))
			adrstr = append([]rune(strings.Repeat("0", 36-len(adrstr))), adrstr...)
			addresses := []string{}

			switch mask[0] {
			case '0':
				addresses = append(addresses, string(adrstr[0]))
			case '1':
				addresses = append(addresses, "1")
			case 'X':
				addresses = append(addresses, "0", "1")
			}
			for i := 1; i < len(mask); i++ {
				ch := mask[i]
				temp := []string{}
				appendix := []string{}
				switch ch {
				case '0':
					appendix = []string{string(adrstr[i])}
				case '1':
					appendix = []string{"1"}
				case 'X':
					appendix = []string{"0", "1"}
				}
				for _, a := range addresses {
					for _, app := range appendix {
						temp = append(temp, a+app)
					}
				}
				addresses = temp
			}

			for _, a := range addresses {
				adr, _ = strconv.ParseInt(string(a), 2, 64)
				mem[adr] = num
			}
		}
	}

	sum = int64(0)
	for _, v := range mem {
		sum += v
	}
	fmt.Println("Part Two: ", sum)
}
