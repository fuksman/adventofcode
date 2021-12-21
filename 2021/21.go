package main

import (
	"fmt"
)

var (
	pos1start = 3
	pos2start = 7
	positions = []int{pos1start, pos2start}
	score     = []int{0, 0}

	diceSize = 100
	rolls    = 0
	lastRoll = 0
)

const trackSize int = 10

func main() {
	current := 0
	for score[0] < 1000 && score[1] < 1000 {
		newPos := positions[current] + rollDie()
		for newPos > trackSize {
			newPos -= trackSize
		}
		positions[current] = newPos
		score[current] += newPos
		current++
		current %= 2
	}
	res := rolls
	if score[0] > score[1] {
		res *= score[1]
	} else {
		res *= score[0]
	}
	fmt.Println("Part One: ", res)

	const winScore int = 21
	dp := [winScore][winScore][trackSize + 1][trackSize + 1][2]int{}
	dp[0][0][pos1start][pos2start][0] = 1

	rolls := make([]int, trackSize)
	for roll1 := 1; roll1 <= 3; roll1++ {
		for roll2 := 1; roll2 <= 3; roll2++ {
			for roll3 := 1; roll3 <= 3; roll3++ {
				rolls[roll1+roll2+roll3]++
			}
		}
	}

	win1, win2 := 0, 0

	for score1 := 0; score1 < winScore; score1++ {
		for score2 := 0; score2 < winScore; score2++ {
			for pos1 := 1; pos1 <= trackSize; pos1++ {
				for pos2 := 1; pos2 <= trackSize; pos2++ {
					for roll := 3; roll <= 9; roll++ { // from 1+1+1 to 3+3+3
						if dp[score1][score2][pos1][pos2][0] > 0 {
							newPos1 := pos1 + roll
							for newPos1 > trackSize {
								newPos1 -= trackSize
							}
							newScore1 := score1 + newPos1
							newCount := dp[score1][score2][pos1][pos2][0] * rolls[roll]
							if newScore1 >= winScore {
								win1 += newCount
							} else {
								dp[newScore1][score2][newPos1][pos2][1] += newCount
							}
						}
						if dp[score1][score2][pos1][pos2][1] > 0 {
							newPos2 := pos2 + roll
							for newPos2 > trackSize {
								newPos2 -= trackSize
							}
							newScore2 := score2 + newPos2
							newCount := dp[score1][score2][pos1][pos2][1] * rolls[roll]
							if newScore2 >= winScore {
								win2 += newCount
							} else {
								dp[score1][newScore2][pos1][newPos2][0] += newCount
							}
						}
					}
				}
			}
		}
	}

	if win1 > win2 {
		fmt.Println("Part Two: ", win1)
	} else {
		fmt.Println("Part Two: ", win2)
	}
}

func rollDie() int {
	sum := 0
	for i := 0; i < 3; i++ {
		lastRoll++
		if lastRoll > diceSize {
			lastRoll -= diceSize
		}
		sum += lastRoll
		rolls++
	}
	return sum
}
