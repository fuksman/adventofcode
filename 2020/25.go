package main

import (
	"fmt"
)

var (
	cardPublic = int64(9789649)
	doorPublic = int64(3647239)
	cardId     = int64(7)
	doorId     = int64(7)
)

func main() {
	i := 1
	res := cardId
	for res != cardPublic {
		i++
		res *= cardId
		res = res % 20201227
	}

	res = doorPublic
	for c := 1; c < i; c++ {
		res *= doorPublic
		res = res % 20201227
	}
	fmt.Println("Part One: ", res)
}
