package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instuction struct {
	dir string
	val int64
}

var (
	lines = [][]*instuction{}
)

func main() {
	file, err := os.Open("3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	max := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, ",")
		temp := []*instuction{}
		for _, token := range tokens {
			num, _ := strconv.ParseInt(token[1:], 10, 64)
			temp = append(temp, &instuction{dir: token[0:1], val: num})
			if num > max {
				max = num
			}
		}
		lines = append(lines, temp)
	}

	pointDestinations := map[[2]int64]map[int]int{}
	panel := make([][]bool, max*100)
	for i := 0; i < int(max)*100; i++ {
		panel[i] = make([]bool, max*100)
	}
	center := []int64{max * 50, max * 50}
	crossings := [][2]int64{}

	calcPoint := func(i, j int64, l, dest int, newLine *[][2]int64) {
		point := [2]int64{i, j}
		if panel[i][j] {
			crossings = append(crossings, point)
		}
		*newLine = append(*newLine, point)
		if pointDestinations[point] == nil {
			pointDestinations[point] = map[int]int{}
		}
		if _, ok := pointDestinations[point][l]; !ok {
			pointDestinations[point][l] = dest
		}
	}

	for l, line := range lines {
		i, j := center[0], center[1]
		newLine := [][2]int64{}
		dest := 0
		for _, instP := range line {
			switch instP.dir {
			case "R":
				for k := int64(0); k < instP.val; k++ {
					j++
					dest++
					calcPoint(i, j, l, dest, &newLine)
				}
			case "L":
				for k := int64(0); k < instP.val; k++ {
					j--
					dest++
					calcPoint(i, j, l, dest, &newLine)
				}
			case "U":
				for k := int64(0); k < instP.val; k++ {
					i--
					dest++
					calcPoint(i, j, l, dest, &newLine)
				}
			case "D":
				for k := int64(0); k < instP.val; k++ {
					i++
					dest++
					calcPoint(i, j, l, dest, &newLine)
				}
			}
		}
		for _, point := range newLine {
			panel[point[0]][point[1]] = true
		}
	}

	distance := int64(1000000)
	steps := 1000000
	for _, crossing := range crossings {
		dis := abs(center[0]-crossing[0]) + abs(center[1]-crossing[1])
		if dis < distance {
			distance = dis
		}
		if len(pointDestinations[crossing]) == 2 {
			d := 0
			for _, v := range pointDestinations[crossing] {
				d += v
			}
			if d < steps {
				steps = d
			}
		}
	}
	fmt.Println("Part One: ", distance)
	fmt.Println("Part Two: ", steps)
}

func abs(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}
