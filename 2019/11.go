package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord [2]int

var (
	picture = map[Coord]int64{}
	in      = make(chan int64)
	out     = make(chan int64)
)

func main() {
	file, err := os.Open("11.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	ic := LoadIntCode(scanner.Text(), ChanInput(in), ChanOutput(out))

	pos := Coord{0, 0}
	dir := Coord{-1, 0}
	picture[Coord{0, 0}] = 1

	done := make(chan struct{})
	go func() {
		defer close(done)
		ic.RunToEnd()
	}()

outer:
	for {
		select {
		case <-done:
			break outer
		case in <- picture[pos]:
		case color := <-out:
			picture[pos] = color
			d := <-out
			if d == 0 {
				dir = Coord{-dir[1], dir[0]}
			} else {
				dir = Coord{dir[1], -dir[0]}
			}
			pos = Coord{pos[0] + dir[0], pos[1] + dir[1]}
		}
	}
	fmt.Println("Part One: ", len(picture))

	topLeft := Coord{}
	bottomRight := Coord{}
	for coor, color := range picture {
		if color == 0 {
			continue
		}
		if coor[0] < topLeft[0] {
			topLeft[0] = coor[0]
		}
		if coor[0] > bottomRight[0] {
			bottomRight[0] = coor[0]
		}
		if coor[1] < topLeft[1] {
			topLeft[1] = coor[1]
		}
		if coor[1] > bottomRight[1] {
			bottomRight[1] = coor[1]
		}
	}
	fmt.Println("Part Two:")
	for i := topLeft[0]; i <= bottomRight[0]; i++ {
		for j := topLeft[1]; j <= bottomRight[1]; j++ {
			if picture[Coord{i, j}] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func LoadIntCode(src string, input func() int64, output func(int64)) IntCode {
	opstr := strings.Split(src, ",")
	mem := make(map[int64]int64, len(opstr))
	for i, s := range opstr {
		mem[int64(i)], _ = strconv.ParseInt(s, 10, 64)
	}

	return IntCode{
		Mem:    mem,
		Input:  input,
		Output: output,
	}
}

func FixedInput(input ...int64) func() int64 {
	return func() int64 {
		var v int64
		v, input = input[0], input[1:]
		return v
	}
}

func PrintOutput(str string) func(int64) {
	return func(v int64) {
		fmt.Println(str+": ", v)
	}
}

func ChanInput(ch <-chan int64) func() int64 {
	return func() int64 {
		return <-ch
	}
}

func ChanOutput(ch chan<- int64) func(int64) {
	return func(v int64) {
		ch <- v
	}
}

type IntCode struct {
	Mem    map[int64]int64
	Base   int64
	Pos    int64
	Input  func() int64
	Output func(int64)
	Done   bool

	modes [3]int64
	modep int
}

func (i *IntCode) RunToEnd() {
	for !i.Done {
		i.Step()
	}
}

func (i *IntCode) Step() {
	if i.Done {
		panic("done")
	}
	op := i.nextOp()
	var p1, p2, p3 int64
	switch op {
	case 1:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		i.writeMem(p3, p1+p2)
	case 2:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		i.writeMem(p3, p1*p2)
	case 3:
		p1 = i.nextPtr()
		p2 = i.Input()
		i.writeMem(p1, p2)
	case 4:
		p1 = i.nextVal()
		i.Output(p1)
	case 5:
		p1 = i.nextVal()
		p2 = i.nextVal()
		if p1 != 0 {
			i.Pos = p2
		}
	case 6:
		p1 = i.nextVal()
		p2 = i.nextVal()
		if p1 == 0 {
			i.Pos = p2
		}
	case 7:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		if p1 < p2 {
			i.writeMem(p3, 1)
		} else {
			i.writeMem(p3, 0)
		}
	case 8:
		p1 = i.nextVal()
		p2 = i.nextVal()
		p3 = i.nextPtr()
		if p1 == p2 {
			i.writeMem(p3, 1)
		} else {
			i.writeMem(p3, 0)
		}
	case 9:
		p1 = i.nextVal()
		i.Base += p1
	case 99:
		i.Done = true
	default:
		panic(op)
	}
}

func (i *IntCode) next() int64 {
	v := i.readMem(i.Pos)
	i.Pos++
	return v
}

func (i *IntCode) nextOp() int64 {
	fullOp := i.next()
	op := fullOp % 100
	i.modes[0] = (fullOp / 100) % 10
	i.modes[1] = (fullOp / 1000) % 10
	i.modes[2] = (fullOp / 10000) % 10
	i.modep = 0
	return op
}

func (i *IntCode) nextVal() int64 {
	v := i.next()
	mode := i.modes[i.modep]
	i.modep++

	switch mode {
	case 0:
		return i.readMem(v)
	case 1:
		return v
	case 2:
		return i.readMem(i.Base + v)
	default:
		panic(mode)
	}
}

func (i *IntCode) nextPtr() int64 {
	mode := i.modes[i.modep]
	i.modep++

	v := i.next()
	if mode == 2 {
		v += i.Base
	}
	return v
}

func (i *IntCode) readMem(v int64) int64 {
	return i.Mem[v]
}

func (i *IntCode) writeMem(p int64, v int64) {
	i.Mem[p] = v
}
