package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	measurements := []int64{}
	for scanner.Scan() {
		current, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		measurements = append(measurements, current)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	counter := 0
	for i := 0; i < len(measurements)-1; i++ {
		if measurements[i] < measurements[i+1] {
			counter++
		}
	}

	fmt.Println(counter)
}
