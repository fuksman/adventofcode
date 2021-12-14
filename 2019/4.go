package main

import (
	"fmt"
	"strconv"
)

var (
	inputMin = int64(197487)
	inputMax = int64(673251)
)

func main() {
	cnt := 0
	for i := inputMin; i <= inputMax; i++ {
		str := strconv.FormatInt(i, 10)
		pair := false
		order := true
		for j := 0; j < len(str)-1; j++ {
			if str[j] > str[j+1] {
				order = false
				break
			}
			if !pair && str[j] == str[j+1] {
				if (j == 0 || str[j-1] != str[j]) && (j == len(str)-2 || str[j+2] != str[j]) { // Part Two
					pair = true
				}
			}
		}
		if order && pair {
			cnt++
		}
	}
	fmt.Println(cnt)
}
