package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	neigbours map[string][]string
	smalls    map[string]bool
	paths     []string
)

func main() {
	file, err := os.Open("12.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	neigbours = map[string][]string{}
	smalls = map[string]bool{}
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		for _, node := range nodes {
			if _, ok := neigbours[node]; !ok {
				neigbours[node] = []string{}
				smalls[node] = strings.ToLower(node) == node
			}
		}
		neigbours[nodes[0]] = append(neigbours[nodes[0]], nodes[1])
		neigbours[nodes[1]] = append(neigbours[nodes[1]], nodes[0])
	}

	fmt.Println(neigbours)
	paths = []string{}
	findPaths("start", map[string]int{}, "", 1)
	fmt.Println("Part One: ", len(paths))
	paths = []string{}
	findPaths("start", map[string]int{}, "", 2)
	fmt.Println("Part Two: ", len(paths))
}

func findPaths(node string, visited map[string]int, path string, allowedVisits int) {
	al := allowedVisits
	if smalls[node] {
		visited[node]++
		if visited[node] == al {
			al = 1
		}
	}
	path += node + " "
	for _, nr := range neigbours[node] {
		if nr == "start" {
			continue
		}
		if nr == "end" {
			paths = append(paths, path+"end")
			continue
		}
		if visited[nr] < al {
			v := map[string]int{}
			for k, val := range visited {
				v[k] = val
			}
			findPaths(nr, v, path, al)
		}
	}
}
