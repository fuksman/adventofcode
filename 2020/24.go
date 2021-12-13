package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	tiling = map[string]bool{}
	lines  = []string{}
)

func main() {
	file, err := os.Open("24.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	// i := 0
	for scanner.Scan() {
		// i++
		line := scanner.Text()
		lines = append(lines, line)
		var x, y, z int64
		for line != "" {
			switch line[0] {
			case 'e':
				y--
				x++
				line = line[1:]
			case 'w':
				y++
				x--
				line = line[1:]
			default:
				switch line[:2] {
				case "ne":
					x++
					z--
					line = line[2:]
				case "nw":
					y++
					z--
					line = line[2:]
				case "sw":
					x--
					z++
					line = line[2:]
				case "se":
					y--
					z++
					line = line[2:]
				}
			}
		}

		if !tiling[intsToString(x, y, z)] {
			count++
		} else {
			count--
		}
		tiling[intsToString(x, y, z)] = !tiling[intsToString(x, y, z)]
	}
	fmt.Println("Part One: ", count)

	count = 0
	for d := 0; d < 100; d++ {
		count = 0
		temp := map[string]bool{}
		for k := range tiling {
			nums := strings.Split(k, ",")
			x, _ := strconv.ParseInt(nums[0], 10, 64)
			y, _ := strconv.ParseInt(nums[1], 10, 64)
			z, _ := strconv.ParseInt(nums[2], 10, 64)
			neighbours := []string{intsToString(x, y+1, z-1), intsToString(x+1, y, z-1), intsToString(x+1, y-1, z), intsToString(x, y-1, z+1), intsToString(x-1, y, z+1), intsToString(x-1, y+1, z)}
			for _, neighbour := range neighbours {
				if _, ok := tiling[neighbour]; !ok {
					tiling[neighbour] = false
				}
			}
			temp[k] = tiling[k]
			if tiling[k] {
				count++
			}
		}
		for k := range tiling {
			nums := strings.Split(k, ",")
			x, _ := strconv.ParseInt(nums[0], 10, 64)
			y, _ := strconv.ParseInt(nums[1], 10, 64)
			z, _ := strconv.ParseInt(nums[2], 10, 64)
			neighbours := []string{intsToString(x, y+1, z-1), intsToString(x+1, y, z-1), intsToString(x+1, y-1, z), intsToString(x, y-1, z+1), intsToString(x-1, y, z+1), intsToString(x-1, y+1, z)}
			countBlack := 0
			for _, neighbour := range neighbours {
				if tiling[neighbour] {
					countBlack++
				}
			}
			if tiling[k] && (countBlack == 0 || countBlack > 2) {
				temp[k] = false
				count--
			}
			if !tiling[k] && countBlack == 2 {
				temp[k] = true
				count++
			}
		}
		tiling = temp
	}
	fmt.Println("Part Two: ", count)
}

func intsToString(nums ...int64) (str string) {
	strs := []string{}
	for _, num := range nums {
		str = strconv.FormatInt(num, 10)
		strs = append(strs, str)
	}
	return strings.Join(strs, ",")
}
