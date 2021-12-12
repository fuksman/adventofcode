package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	tiles                    map[string]*tile
	commonBorders            map[*tile]map[*tile][]rune
	photo, photoWithMonsters *tile
	monsterMask              [][]rune
)

type tile struct {
	num    string
	data   [][]rune
	placed bool
}

func newTile(num string, data [][]rune) *tile {
	t := new(tile)
	t.num = num
	t.data = data
	return t
}
func (t *tile) top() []rune {
	return t.data[0]
}
func (t *tile) right() []rune {
	side := []rune{}
	n := len(t.data)
	for _, line := range t.data {
		side = append(side, line[n-1])
	}
	return side
}
func (t *tile) bottom() []rune {
	return t.data[len(t.data)-1]
}
func (t *tile) left() []rune {
	side := []rune{}
	for _, line := range t.data {
		side = append(side, line[0])
	}
	return side
}

func (t *tile) rotate() {
	for i, j := 0, len(t.data)-1; i < j; i, j = i+1, j-1 {
		t.data[i], t.data[j] = t.data[j], t.data[i]
	}
	for i := 0; i < len(t.data); i++ {
		for j := 0; j < i; j++ {
			t.data[i][j], t.data[j][i] = t.data[j][i], t.data[i][j]
		}
	}
}
func (t *tile) flipV() {
	for i, j := 0, len(t.data)-1; i < j; i, j = i+1, j-1 {
		t.data[i], t.data[j] = t.data[j], t.data[i]
	}
}
func (t *tile) flipH() {
	for i := 0; i < len(t.data); i++ {
		for j, l := 0, len(t.data[i])-1; j < l; j, l = j+1, l-1 {
			t.data[i][j], t.data[i][l] = t.data[i][l], t.data[i][j]
		}
	}
}
func (t *tile) print() {
	fmt.Println(t.num)
	for _, line := range t.data {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	file, err := os.Open("20.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tiles = map[string]*tile{}
	commonBorders = map[*tile]map[*tile][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == 'T' {
			num := line[5 : len(line)-1]
			data := [][]rune{}
			for scanner.Scan() && scanner.Text() != "" {
				data = append(data, []rune(scanner.Text()))
			}
			tiles[num] = newTile(num, data)
			commonBorders[tiles[num]] = map[*tile][]rune{}
		}
	}

	n := int(math.Sqrt(float64(len(tiles))))
	picture := make([][]string, n)
	for i := 0; i < n; i++ {
		picture[i] = make([]string, n)
	}

	topLeft := ""
	prod := int64(1)
	for t1 := range tiles {
		for t2 := range tiles {
			if t1 == t2 {
				continue
			}
			findBorders(tiles[t1], tiles[t2])
		}
		if len(commonBorders[tiles[t1]]) == 2 {
			topLeft = tiles[t1].num
			picture[0][0] = topLeft
			n, _ := strconv.ParseInt(topLeft, 10, 64)
			prod *= n
		}
	}
	fmt.Println("Part One: ", prod)

	t := tiles[topLeft]
	l := t
	nrs := neighbours(t)
	for weakEqual(t.top(), commonBorders[t][nrs[0]]) || weakEqual(t.top(), commonBorders[t][nrs[1]]) || weakEqual(t.left(), commonBorders[t][nrs[0]]) || weakEqual(t.left(), commonBorders[t][nrs[1]]) {
		t.rotate()
	}
	t.placed = true

	l.print()
	for nr := range commonBorders[t] {
		fmt.Println(nr.num, string(commonBorders[t][nr]), string(t.bottom()))
	}

	for i := 1; i < n; i++ {
		var d, r *tile
		for neighbour := range commonBorders[t] {
			fmt.Println(neighbour.num, neighbour.placed)
			if neighbour.placed {
				continue
			}
			if hasCommonBorder(neighbour, t.bottom()) {
				d = neighbour
				picture[i][0] = d.num
			}
		}

		for neighbour := range commonBorders[l] {
			if neighbour.placed {
				continue
			}
			if hasCommonBorder(neighbour, l.right()) {
				r = neighbour
				picture[0][i] = r.num
			}
		}

		for !weakEqual(t.bottom(), d.top()) {
			d.rotate()
		}
		if !equal(d.top(), t.bottom()) {
			d.flipH()
		}
		d.placed = true
		t = d
		for !weakEqual(r.left(), l.right()) {
			r.rotate()
		}
		if !equal(r.left(), l.right()) {
			r.flipV()
		}
		r.placed = true
		l = r
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			lefter := tiles[picture[i][j-1]]
			upper := tiles[picture[i-1][j]]
			t := common(neighbours(lefter), neighbours(upper))
			picture[i][j] = t.num
			for !weakEqual(upper.bottom(), t.top()) {
				t.rotate()
			}
			if !equal(upper.bottom(), t.top()) {
				t.flipH()
			}
			t.placed = true
		}
	}

	printPicture(&picture)
	data := [][]rune{}
	n = len(tiles[picture[0][0]].data)
	for i, line := range picture {
		for r := 1; r < n-1; r++ {
			data = append(data, []rune{})
			for _, num := range line {
				data[(n-2)*i+r-1] = append(data[(n-2)*i+r-1], tiles[num].data[r][1:n-1]...)
			}
		}
	}
	photo = newTile("Photo", data)
	photo.print()
	photoWithMonsters = newTile("Photo with monsters", data)

	monsterMask = [][]rune{[]rune("                  # "), []rune("#    ##    ##    ###"), []rune(" #  #  #  #  #  #   ")}
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()

	photoWithMonsters.flipH()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()

	photoWithMonsters.flipV()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()

	photoWithMonsters.flipH()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()
	photo.rotate()
	photoWithMonsters.rotate()
	findMonster()

	photoWithMonsters.print()
	count := 0
	for _, line := range photoWithMonsters.data {
		for _, sym := range line {
			if sym == '#' {
				count++
			}
		}
	}
	fmt.Println("Part Two: ", count)
}

func findMonster() {
	for i := 0; i < len(photo.data)-3; i++ {
		for j := 0; j < len(photo.data)-20; j++ {
			found := true
			for k := 0; found && k < len(monsterMask); k++ {
				for l := 0; found && l < len(monsterMask[k]); l++ {
					if monsterMask[k][l] == '#' {
						if photo.data[i+k][j+l] != '#' {
							found = false
						}
					}
				}
			}
			if found {
				for k := 0; found && k < len(monsterMask); k++ {
					for l := 0; found && l < len(monsterMask[k]); l++ {
						if monsterMask[k][l] == '#' {
							photoWithMonsters.data[i+k][j+l] = 'O'
						}
					}
				}
			}
		}
	}
}

func hasCommonBorder(t *tile, b []rune) bool {
	return weakEqual(t.top(), b) || weakEqual(t.right(), b) || weakEqual(t.bottom(), b) || weakEqual(t.left(), b)
}

func findBorders(t1, t2 *tile) {
	if _, ok := commonBorders[t1][t2]; ok {
		return
	}
	for c1 := 0; c1 < 4; c1++ {
		top, right, bottom, left := t1.top(), t1.right(), t1.bottom(), t1.left()
		for c2 := 0; c2 < 4; c2++ {
			if weakEqual(top, t2.bottom()) || weakEqual(top, t2.top()) {
				commonBorders[t1][t2] = top
				commonBorders[t2][t1] = top
				return
			}
			if weakEqual(right, t2.left()) || weakEqual(right, t2.right()) {
				commonBorders[t1][t2] = right
				commonBorders[t2][t1] = right
				return
			}
			if weakEqual(bottom, t2.top()) || weakEqual(bottom, t2.bottom()) {
				commonBorders[t1][t2] = bottom
				commonBorders[t2][t1] = bottom
				return
			}
			if weakEqual(left, t2.right()) || weakEqual(left, t2.left()) {
				commonBorders[t1][t2] = left
				commonBorders[t2][t1] = left
				return
			}
			t2.rotate()
		}
		t1.rotate()
	}
}

func reversed(border []rune) (res []rune) {
	for i := len(border) - 1; i >= 0; i-- {
		res = append(res, border[i])
	}
	return
}

func equal(b1, b2 []rune) bool {
	if len(b1) != len(b2) {
		return false
	}
	eq := true
	for i, el := range b1 {
		if el != b2[i] {
			eq = false
			break
		}
	}
	return eq
}

func weakEqual(b1, b2 []rune) bool {
	return equal(b1, b2) || equal(b1, reversed(b2))
}

func printPicture(p *[][]string) {
	picture := *p
	n := len(tiles[picture[0][0]].data)
	for _, line := range picture {
		for r := 0; r < n; r++ {
			for _, num := range line {
				if num == "" {
					continue
				}
				fmt.Printf("%s   ", string(tiles[num].data[r]))
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func neighbours(t *tile) (res []*tile) {
	for nr := range commonBorders[t] {
		res = append(res, nr)
	}
	return
}

func common(s1, s2 []*tile) *tile {
	for _, el1 := range s1 {
		if el1.placed {
			continue
		}
		for _, el2 := range s2 {
			if el1 == el2 {
				return el1
			}
		}
	}
	return nil
}
