package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("12.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	directions := map[byte][2]byte{
		'E': {'N', 'S'},
		'S': {'E', 'W'},
		'W': {'S', 'N'},
		'N': {'W', 'E'},
	}

	dir := byte('E')
	var i, j int64
	for _, line := range lines {
		instr := line[0]
		num, _ := strconv.ParseInt(line[1:], 10, 64)
		switch instr {
		case 'L':
			for i := num / 90; i > 0; i-- {
				dir = directions[dir][0]
			}
		case 'R':
			for i := num / 90; i > 0; i-- {
				dir = directions[dir][1]
			}
		case 'F':
			i, j = move(dir, i, j, num)
		default:
			i, j = move(instr, i, j, num)
		}
	}
	fmt.Println("Part One: ", abs(i)+abs(j))

	dir = byte('E')
	var wi, wj int64
	i, j = 0, 0
	wi, wj = 10, 1
	for _, line := range lines {
		instr := line[0]
		num, _ := strconv.ParseInt(line[1:], 10, 64)
		switch instr {
		case 'L':
			for k := num / 90; k > 0; k-- {
				wi, wj = -wj, wi
			}
		case 'R':
			for k := num / 90; k > 0; k-- {
				wi, wj = wj, -wi
			}
		case 'F':
			i += wi * num
			j += wj * num
		default:
			wi, wj = move(instr, wi, wj, num)
		}
	}

	fmt.Println("Part Two: ", abs(i)+abs(j))
}

func move(instr byte, i, j, num int64) (int64, int64) {
	switch instr {
	case 'W':
		i -= num
	case 'E':
		i += num
	case 'N':
		j += num
	case 'S':
		j -= num
	}
	return i, j
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
