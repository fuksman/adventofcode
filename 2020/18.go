package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	// "strconv"
	"strings"
)

func main() {
	file, err := os.Open("18.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]rune{}
	advlines := [][]rune{}
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
		advlines = append(advlines, []rune(scanner.Text()))
	}

	sum := int64(0)
	for i := range lines {
		openidx := strings.LastIndex(string(lines[i]), "(")
		for openidx != -1 {
			closeidx := strings.Index(string(lines[i][openidx:]), ")")
			res, _ := calc(lines[i][openidx+1 : openidx+closeidx])
			newline := append([]rune{}, lines[i][:openidx]...)
			newline = append(newline, res...)
			newline = append(newline, lines[i][openidx+closeidx+1:]...)
			lines[i] = newline
			openidx = strings.LastIndex(string(newline), "(")
		}
		_, res := calc(lines[i])
		sum += res
	}
	fmt.Println("Part One: ", sum)

	sum = int64(0)
	for i := range advlines {
		openidx := strings.LastIndex(string(advlines[i]), "(")
		for openidx != -1 {
			closeidx := strings.Index(string(advlines[i][openidx:]), ")")
			res, _ := calcAdvanced(advlines[i][openidx+1 : openidx+closeidx])
			newline := append([]rune{}, advlines[i][:openidx]...)
			newline = append(newline, res...)
			newline = append(newline, advlines[i][openidx+closeidx+1:]...)
			advlines[i] = newline
			openidx = strings.LastIndex(string(newline), "(")
		}
		_, res := calcAdvanced(advlines[i])
		sum += res
	}
	fmt.Println("Part Two: ", sum)
}

func calc(arr []rune) ([]rune, int64) {
	digs := regexp.MustCompile("[0-9]+")
	indexes := digs.FindAllStringIndex(string(arr), -1)
	res, _ := strconv.ParseInt(string(arr[indexes[0][0]:indexes[0][1]]), 10, 64)
	for i := 1; i < len(indexes); i++ {
		action := arr[indexes[i][0]-2]
		num, _ := strconv.ParseInt(string(arr[indexes[i][0]:indexes[i][1]]), 10, 64)
		switch action {
		case '+':
			res += num
		case '*':
			res *= num
		}
	}

	return []rune(strconv.FormatInt(res, 10)), res
}

func calcAdvanced(arr []rune) ([]rune, int64) {
	str := string(arr)
	var res int64
	for s, sym := range []string{" + ", " * "} {
		for substr := strings.SplitN(str, sym, 2); len(substr) == 2; substr = strings.SplitN(str, sym, 2) {
			leftFields := strings.Fields(substr[0])
			a, _ := strconv.ParseInt(leftFields[len(leftFields)-1], 10, 64)
			rightFields := strings.Fields(substr[1])
			b, _ := strconv.ParseInt(rightFields[0], 10, 64)

			switch s {
			case 0:
				res = a + b
			case 1:
				res = a * b
			}
			resstr := []rune(strconv.FormatInt(res, 10))

			left, right := "", ""
			if len(leftFields) != 1 {
				left = strings.Join(leftFields[:len(leftFields)-1], " ") + " "
			}
			if len(rightFields) != 1 {
				right = " " + strings.Join(rightFields[1:], " ")
			}
			str = left + string(resstr) + right
		}
	}

	return []rune(str), res
}
