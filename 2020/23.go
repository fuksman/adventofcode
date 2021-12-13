package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	input    = "496138527"
	addr     = map[int64]*node{}
	min, max int64
)

type node struct {
	val   int64
	next1 *node
	next2 *node
}

func main() {
	min = 1
	max = 9
	tokens := strings.Split(input, "")
	digits := []int64{}
	for _, token := range tokens {
		num, _ := strconv.ParseInt(token, 10, 64)
		digits = append(digits, num)
		n := node{val: num}
		addr[num] = &n
	}
	for i := range digits {
		addr[digits[i]].next1 = addr[digits[(i+1)%len(digits)]]
	}

	current := addr[digits[0]]
	for c := 0; c < 100; c++ {
		// fmt.Println(current.val, current.next.val)
		pickUp := []*node{current.next1, current.next1.next1, current.next1.next1.next1}
		current.next1 = pickUp[2].next1
		pickUp[2].next1 = nil

		// fmt.Println(current.val, current.next.val)

		dest := &node{val: current.val - 1}
		if dest.val < min {
			dest.val = max
		}
		for (addr[dest.val] == pickUp[0]) || (addr[dest.val] == pickUp[1]) || (addr[dest.val] == pickUp[2]) {
			dest.val--
			if dest.val < min {
				dest.val = max
			}
		}
		// fmt.Println(dest.val)
		dest = addr[dest.val]
		// fmt.Println(addr, dest)

		pickUp[2].next1 = dest.next1
		dest.next1 = pickUp[0]
		current = current.next1
	}

	p1 := addr[1].next1
	fmt.Print("Part One: ")
	for p1 != addr[1] {
		fmt.Print(p1.val)
		p1 = p1.next1
	}
	fmt.Println()

	max = 1000000
	for num := int64(10); num <= max; num++ {
		digits = append(digits, num)
		n := node{val: num}
		addr[num] = &n
	}

	for i := range digits {
		addr[digits[i]].next2 = addr[digits[(i+1)%len(digits)]]
	}

	current = addr[digits[0]]
	for c := 0; c < 10000000; c++ {
		// fmt.Println(current.val, current.next.val)
		pickUp := []*node{current.next2, current.next2.next2, current.next2.next2.next2}
		current.next2 = pickUp[2].next2
		pickUp[2].next2 = nil

		// fmt.Println(current.val, current.next.val)

		dest := &node{val: current.val - 1}
		if dest.val < min {
			dest.val = max
		}
		for (addr[dest.val] == pickUp[0]) || (addr[dest.val] == pickUp[1]) || (addr[dest.val] == pickUp[2]) {
			dest.val--
			if dest.val < min {
				dest.val = max
			}
		}
		// fmt.Println(dest.val)
		dest = addr[dest.val]
		// fmt.Println(addr, dest)

		pickUp[2].next2 = dest.next2
		dest.next2 = pickUp[0]
		current = current.next2
	}

	p2 := addr[1].next2
	fmt.Println("Part Two: ", p2.val*p2.next2.val)
}
