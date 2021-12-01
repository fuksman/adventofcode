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
			ch       string
			password string
		)
		str := scanner.Text()
		tokens := strings.Fields(str)
		numsStr := tokens[0]
		nums := strings.Split(numsStr, "-")

		min, _ = strconv.ParseInt(nums[0], 10, 32)
		max, _ = strconv.ParseInt(nums[1], 10, 32)
		ch = tokens[1][:len(tokens[1])-1]
		password = tokens[2]
		if cnt := int64(strings.Count(password, ch)); cnt >= min && cnt <= max {
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
