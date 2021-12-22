package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	op                     string
	xRange, yRange, zRange FlatRange
}
type FlatRange [2]int64

var (
	lines   = []Operation{}
	lines50 = []Operation{}
)

func main() {
	file, err := os.Open("22.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.SplitN(scanner.Text(), " ", 2)
		coordStrs := strings.Split(tokens[1], ",")
		xStr := strings.Split(coordStrs[0][2:], "..")
		yStr := strings.Split(coordStrs[1][2:], "..")
		zStr := strings.Split(coordStrs[2][2:], "..")
		var xRange, yRange, zRange FlatRange
		xRange[0], _ = strconv.ParseInt(xStr[0], 10, 64)
		xRange[1], _ = strconv.ParseInt(xStr[1], 10, 64)
		yRange[0], _ = strconv.ParseInt(yStr[0], 10, 64)
		yRange[1], _ = strconv.ParseInt(yStr[1], 10, 64)
		zRange[0], _ = strconv.ParseInt(zStr[0], 10, 64)
		zRange[1], _ = strconv.ParseInt(zStr[1], 10, 64)
		lines = append(lines, Operation{op: tokens[0],
			xRange: xRange,
			yRange: yRange,
			zRange: zRange,
		})

		xR50 := subrange(xRange, FlatRange{-50, 50})
		yR50 := subrange(yRange, FlatRange{-50, 50})
		zR50 := subrange(zRange, FlatRange{-50, 50})
		if rangeLen(xR50) > 0 && rangeLen(yR50) > 0 && rangeLen(zR50) > 0 {
			lines50 = append(lines50, Operation{op: tokens[0],
				xRange: xR50,
				yRange: yR50,
				zRange: zR50,
			})
		}
	}

	count := int64(0)
	for idx, op := range lines50 {
		if op.op == "off" {
			continue
		}
		count += countNew(op, lines50[idx+1:])
	}
	fmt.Println("Part One: ", count)

	count = int64(0)
	for idx, op := range lines {
		if op.op == "off" {
			continue
		}
		count += countNew(op, lines[idx+1:])
	}
	fmt.Println("Part Two: ", count)
}

func countNew(item Operation, rest []Operation) int64 {
	total := rangeLen(item.xRange) * rangeLen(item.yRange) * rangeLen(item.zRange)
	conflicts := []Operation{}

	for _, i := range rest {
		xSubRange := subrange(i.xRange, item.xRange)
		ySubRange := subrange(i.yRange, item.yRange)
		zSubRange := subrange(i.zRange, item.zRange)

		if rangeLen(xSubRange) == 0 || rangeLen(ySubRange) == 0 || rangeLen(zSubRange) == 0 {
			continue
		}

		conflicts = append(conflicts, Operation{op: i.op, xRange: xSubRange, yRange: ySubRange, zRange: zSubRange})
	}

	for idx, i := range conflicts {
		total -= countNew(i, conflicts[idx+1:])
	}

	return total
}

func subrange(base, clip FlatRange) FlatRange {
	if base[1] <= clip[0] || base[0] >= clip[1] {
		return FlatRange{0, 0}
	}
	return FlatRange{max(base[0], clip[0]), min(base[1], clip[1])}
}

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

func min(i, j int64) int64 {
	if i < j {
		return i
	}
	return j
}

func rangeLen(r FlatRange) int64 {
	if r[0] == r[1] {
		return 0
	}
	return r[1] - r[0] + 1
}
