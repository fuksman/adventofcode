package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	risks         = map[[2]int]int64{}
	lengths       = map[[2]int]int64{}
	unvisited     = map[[2]int]bool{}
	max           = 0
	start, target [2]int
	pq            PriorityQueue
)

func main() {
	file, err := os.Open("15.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "")
		for j, token := range tokens {
			num, _ := strconv.ParseInt(token, 10, 64)
			risks[[2]int{i, j}] = num
		}
		i++
	}
	max = i

	risks2 := map[[2]int]int64{}
	for ri := 0; ri <= 4; ri++ {
		for rj := 0; rj <= 4; rj++ {
			for coord, risk := range risks {
				i, j := coord[0], coord[1]
				nextRisk := risk + int64(ri) + int64(rj)
				if nextRisk > 9 {
					nextRisk -= 9
				}
				risks2[[2]int{i + ri*max, j + rj*max}] = nextRisk
				lengths[[2]int{i + ri*max, j + rj*max}] = int64(1000000000)
				unvisited[[2]int{i + ri*max, j + rj*max}] = true
			}
		}
	}

	risks = risks2
	start = [2]int{0, 0}
	lengths[start] = 0

	target1 := [2]int{max - 1, max - 1}
	max *= 5
	target2 := [2]int{max - 1, max - 1}
	target = target1

	pq = make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		value:    start,
		priority: 0,
	})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item).value

		if _, ok := unvisited[target]; !ok {
			break
		}
		neigbours := [][2]int{}
		i := current[0]
		j := current[1]
		for _, n := range [][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
			if _, ok := unvisited[n]; n[0] >= 0 && n[1] >= 0 && n[0] < max && n[1] < max && ok {
				neigbours = append(neigbours, n)
			}
		}
		if len(neigbours) == 0 {
			delete(unvisited, current)
			continue
		}
		for _, n := range neigbours {
			newLength := lengths[current] + risks[n]
			if newLength < lengths[n] {
				lengths[n] = newLength
				heap.Push(&pq, &Item{
					value:    n,
					priority: newLength,
				})
			}
		}
		delete(unvisited, current)
	}

	fmt.Println("Part One: ", lengths[target1])
	fmt.Println("Part Two: ", lengths[target2])
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    [2]int // The value of the item; arbitrary.
	priority int64  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
