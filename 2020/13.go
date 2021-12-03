package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("13.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	eststr := scanner.Text()
	est, _ := strconv.ParseInt(eststr, 10, 64)
	scanner.Scan()
	sch := scanner.Text()

	schedule := strings.Split(sch, ",")
	buses := []int64{}
	for _, entry := range schedule {
		if entry != "x" {
			bus, _ := strconv.ParseInt(entry, 10, 64)
			buses = append(buses, bus)
		}
	}
	diff := make([]int64, len(buses))
	min := int64(math.MaxInt64)
	minidx := -1
	for i, d := range diff {
		for d-est < 0 {
			d += buses[i]
		}
		diff[i] = d - est
		if diff[i] < min {
			min = diff[i]
			minidx = i
		}
	}

	fmt.Println("Part One: ", min*buses[minidx])

	offset := []int64{}
	for o, entry := range schedule {
		if entry != "x" {
			offset = append(offset, int64(o))
		}
	}

	i := int64(0)
	runningProduct := int64(1)
	for b, bus := range buses {
		for (i+offset[b])%bus != 0 {
			i += runningProduct
		}
		runningProduct *= bus
	}

	fmt.Println("Part Two: ", i)
}
