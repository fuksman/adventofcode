package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("22.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	p1 := []int64{}
	for scanner.Scan() && scanner.Text() != "" {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		p1 = append(p1, num)
	}
	p12 := append([]int64{}, p1...)
	scanner.Scan()
	p2 := []int64{}
	for scanner.Scan() && scanner.Text() != "" {
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		p2 = append(p2, num)
	}
	p22 := append([]int64{}, p2...)

	for len(p1) != 0 && len(p2) != 0 {
		p1card := p1[0]
		p2card := p2[0]
		if p1card > p2card {
			temp := p1[1:]
			p1 = append(temp, p1card, p2card)
			p2 = p2[1:]
		} else {
			temp := p2[1:]
			p2 = append(temp, p2card, p1card)
			p1 = p1[1:]
		}
	}

	var winner *[]int64
	if len(p1) == 0 {
		winner = &p2
	} else {
		winner = &p1
	}

	fmt.Println("Part One: ", calcScore(winner))

	_, winner = recursiveCombat(p12, p22)
	fmt.Println("Part Two: ", calcScore(winner))

}

func intsToString(s []int64) (str string) {
	for _, num := range s {
		str += strconv.FormatInt(num, 10) + ","
	}
	return str
}

func calcScore(winner *[]int64) (score int64) {
	n := len(*winner)
	for i := 0; i < n; i++ {
		score += (*winner)[i] * int64(n-i)
	}
	return
}

func recursiveCombat(p1, p2 []int64) (idx int, winner *[]int64) {
	history1 := map[string]bool{}
	history2 := map[string]bool{}
	for len(p1) != 0 && len(p2) != 0 {
		if history1[intsToString(p1)] || history2[intsToString(p2)] {
			winner = &p1
			return 1, winner
		}
		history1[intsToString(p1)] = true
		history2[intsToString(p2)] = true

		p1card := p1[0]
		p2card := p2[0]
		if len(p1[1:]) >= int(p1card) && len(p2[1:]) >= int(p2card) {
			winnerIdx, _ := recursiveCombat(append([]int64{}, p1[1:int(p1card)+1]...), append([]int64{}, p2[1:int(p2card)+1]...))
			switch winnerIdx {
			case 1:
				temp := p1[1:]
				p1 = append(temp, p1card, p2card)
				p2 = p2[1:]
			case 2:
				temp := p2[1:]
				p2 = append(temp, p2card, p1card)
				p1 = p1[1:]
			}
		} else {
			switch max(p1card, p2card) {
			case p1card:
				temp := p1[1:]
				p1 = append(temp, p1card, p2card)
				p2 = p2[1:]
			case p2card:
				temp := p2[1:]
				p2 = append(temp, p2card, p1card)
				p1 = p1[1:]
			}
		}
	}
	if len(p1) == 0 {
		return 2, &p2
	} else {
		return 1, &p1
	}
}

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}
