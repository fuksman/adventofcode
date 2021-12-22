package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	nums   = []int64{}
	output int64
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

	fmt.Println("Part Two: ", part2Orig())
}

// from https://github.com/tsholmes/aoc-2019/blob/master/day7/main.go
func part2Orig() int64 {
	run := func(input chan int64, output chan int64) {
		var pos int64

		next := func() int64 {
			val := nums[pos]
			pos++
			return val
		}

		val := func(p int64, fullOp int64, off int64) int64 {
			mode := (fullOp / off) % 10
			if mode == 0 {
				return nums[p]
			} else {
				return p
			}
		}

		for {
			fullOp := next()
			op := fullOp % 100
			var p1, p2, p3 int64
			switch op {
			case 1:
				p1 = next()
				p2 = next()
				p3 = next()
				nums[p3] = val(p1, fullOp, 100) + val(p2, fullOp, 1000)
			case 2:
				p1 = next()
				p2 = next()
				p3 = next()
				nums[p3] = val(p1, fullOp, 100) * val(p2, fullOp, 1000)
			case 3:
				p1 = next()
				p2 = <-input
				nums[p1] = p2
			case 4:
				p1 = next()
				output <- val(p1, fullOp, 100)
			case 5:
				p1 = next()
				p2 = next()
				if val(p1, fullOp, 100) != 0 {
					pos = val(p2, fullOp, 1000)
				}
			case 6:
				p1 = next()
				p2 = next()
				if val(p1, fullOp, 100) == 0 {
					pos = val(p2, fullOp, 1000)
				}
			case 7:
				p1 = next()
				p2 = next()
				p3 = next()
				if val(p1, fullOp, 100) < val(p2, fullOp, 1000) {
					nums[p3] = 1
				} else {
					nums[p3] = 0
				}
			case 8:
				p1 = next()
				p2 = next()
				p3 = next()
				if val(p1, fullOp, 100) == val(p2, fullOp, 1000) {
					nums[p3] = 1
				} else {
					nums[p3] = 0
				}
			case 99:
				close(output)
				return
			default:
				panic(op)
			}
		}
	}

	max := int64(0)

	for v1 := 1; v1 <= 5; v1++ {
		for v2 := 1; v2 <= 5; v2++ {
			for v3 := 1; v3 <= 5; v3++ {
				for v4 := 1; v4 <= 5; v4++ {
					for v5 := 1; v5 <= 5; v5++ {
						if v1*v2*v3*v4*v5 != 1*2*3*4*5 || v1+v2+v3+v4+v5 != 1+2+3+4+5 {
							continue
						}

						initIn := make(chan int64, 200)
						in := initIn
						var out chan int64
						for i, v := range []int{v1, v2, v3, v4, v5} {
							in <- int64(v + 4)
							if i == 0 {
								in <- 0
							}
							out = make(chan int64, 200)
							go run(in, out)
							in = out
						}

						output := int64(0)
					outer:
						for {
							select {
							case v, ok := <-out:
								if !ok {
									break outer
								}
								output = v
								initIn <- v
							}
						}
						if output > max {
							max = output
						}
					}
				}
			}
		}
	}
	return max
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
