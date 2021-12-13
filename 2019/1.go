package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	lines []int64
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := int64(0)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, num)
		add := num
		for add = add/3 - 2; add > 0; add = add/3 - 2 {
			sum += add
			fmt.Println(add)
		}
	}
	fmt.Println("Part Two: ", sum)
}
