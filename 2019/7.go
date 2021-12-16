package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	nums []int64
	// input  = int64(0)
	output     int64
	phaseQueue = list.New()
	inputQueue = list.New()
)

func main() {
	file, err := os.Open("7.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")
	for _, token := range tokens {
		num, _ := strconv.ParseInt(token, 10, 64)
		nums = append(nums, num)
	}

	phasesDict := []int64{0, 1, 2, 3, 4}
	phases := permutations(phasesDict)
	maxOutput := int64(0)

	for _, phase := range phases {
		input := int64(0)
		for _, amp := range phase {
			mem := append([]int64{}, nums...)
			intcode([]int64{amp, input})
			input = output
			nums = append([]int64{}, mem...)
		}
		if output > maxOutput {
			maxOutput = output
		}
	}
	fmt.Println("Part One: ", maxOutput)

	phasesDict = []int64{5, 6, 7, 8, 9}
	phases = permutations(phasesDict)
	maxOutput = int64(0)

	for _, phase := range phases {
		phaseQueue = list.New()
		inputQueue.PushBack(int64(0))
		// input = int64(0)
		mem := append([]int64{}, nums...)
		for _, amp := range phase {
			phaseQueue.PushBack(amp)
		}
		intcodeQ()
		nums = append([]int64{}, mem...)
		// output = inputQueue.Front().Value.(int64)
		if output > maxOutput {
			maxOutput = output
		}
	}
	fmt.Println("Part Two: ", maxOutput)
}

func intcodeQ() {
	jump := 4
	inputCnt := 0
	for i := 0; i < len(nums) && nums[i] != 99; i += jump {
		operation := nums[i] % 100
		mode2 := (nums[i] % 10000) / 1000
		mode1 := (nums[i] % 1000) / 100
		switch operation {
		case 1:
			nums[nums[i+3]] = value(nums[i+1], mode1) + value(nums[i+2], mode2)
			jump = 4
		case 2:
			nums[nums[i+3]] = value(nums[i+1], mode1) * value(nums[i+2], mode2)
			jump = 4
		case 3:
			// nums[nums[i+1]] = inputSignals[inputCnt]
			switch inputCnt % 2 {
			case 0:
				phase := phaseQueue.Front().Value.(int64)
				phaseQueue.PushBack(phase)
				fmt.Println(phaseQueue.Len())
				nums[nums[i+1]] = phase
			case 1:
				input := inputQueue.Front().Value.(int64)
				inputQueue.PushBack(input)
				nums[nums[i+1]] = input
			}
			inputCnt++
			jump = 2
		case 4:
			output = nums[nums[i+1]]
			inputQueue.PushBack(output)
			jump = 2
		case 5:
			if value(nums[i+1], mode1) != 0 {
				jump = int(value(nums[i+2], mode2)) - i
			} else {
				jump = 3
			}
		case 6:
			if value(nums[i+1], mode1) == 0 {
				jump = int(value(nums[i+2], mode2)) - i
			} else {
				jump = 3
			}
		case 7:
			if value(nums[i+1], mode1) < value(nums[i+2], mode2) {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			jump = 4
		case 8:
			if value(nums[i+1], mode1) == value(nums[i+2], mode2) {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			jump = 4
		case 99:
			break
		}
	}
}

func value(v, mode int64) int64 {
	if mode == 0 {
		return nums[v]
	}
	return int64(v)
}

func intcode(inputSignals []int64) {
	jump := 4
	inputCnt := 0
	for i := 0; i < len(nums) && nums[i] != 99; i += jump {
		operation := nums[i] % 100
		mode2 := (nums[i] % 10000) / 1000
		mode1 := (nums[i] % 1000) / 100
		switch operation {
		case 1:
			nums[nums[i+3]] = value(nums[i+1], mode1) + value(nums[i+2], mode2)
			jump = 4
		case 2:
			nums[nums[i+3]] = value(nums[i+1], mode1) * value(nums[i+2], mode2)
			jump = 4
		case 3:
			nums[nums[i+1]] = inputSignals[inputCnt]
			inputCnt++
			jump = 2
		case 4:
			output = nums[nums[i+1]]
			jump = 2
		case 5:
			if value(nums[i+1], mode1) != 0 {
				jump = int(value(nums[i+2], mode2)) - i
			} else {
				jump = 3
			}
		case 6:
			if value(nums[i+1], mode1) == 0 {
				jump = int(value(nums[i+2], mode2)) - i
			} else {
				jump = 3
			}
		case 7:
			if value(nums[i+1], mode1) < value(nums[i+2], mode2) {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			jump = 4
		case 8:
			if value(nums[i+1], mode1) == value(nums[i+2], mode2) {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			jump = 4
		case 99:
			break
		}
	}
}

func permutations(arr []int64) [][]int64 {
	var helper func([]int64, int)
	res := [][]int64{}
	helper = func(arr []int64, n int) {
		if n == 1 {
			tmp := make([]int64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
