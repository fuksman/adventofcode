package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	name  string
	orbit *object
}

var (
	objects = map[string]*object{}
)

func main() {
	file, err := os.Open("6.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		objs := strings.Split(scanner.Text(), ")")
		center, ok := objects[objs[0]]
		if !ok {
			center = &object{name: objs[0]}
			objects[objs[0]] = center
		}
		obj, ok := objects[objs[1]]
		if !ok {
			obj = &object{name: objs[1], orbit: center}
		}
		obj.orbit = center
		objects[objs[1]] = obj
	}

	cnt := 0
	for _, obj := range objects {
		for orbit := obj.orbit; orbit != nil; orbit = orbit.orbit {
			cnt++
		}
	}
	fmt.Println("Part One: ", cnt)

	myPath := []string{}
	obj := objects["YOU"]
	for orbit := obj.orbit; orbit != nil; orbit = orbit.orbit {
		myPath = append(myPath, orbit.name)
	}

	santaPath := []string{}
	obj = objects["SAN"]
	for orbit := obj.orbit; orbit != nil; orbit = orbit.orbit {
		santaPath = append(santaPath, orbit.name)
	}

	path := []string{}
	for _, str := range myPath {
		path = append(path, str)
		found := false
		i := 0
		for ; i < len(santaPath); i++ {
			if santaPath[i] == str {
				found = true
				break
			}
		}
		if found {
			for ; i-1 >= 0; i-- {
				path = append(path, santaPath[i-1])
			}
			break
		}
	}
	fmt.Println("Part Two: ", len(path)-1)
}
