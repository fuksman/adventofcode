package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		var (
			min, max int64
			ch       byte
			password string
		)
		str := scanner.Text()
		tokens := strings.Fields(str)
		numsStr := tokens[0]
		nums := strings.Split(numsStr, "-")

		min, _ = strconv.ParseInt(nums[0], 10, 32)
		max, _ = strconv.ParseInt(nums[1], 10, 32)
		ch = tokens[1][0]
		password = tokens[2]
		if (password[min-1] == ch) != (password[max-1] == ch) {
			count++
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}
}
