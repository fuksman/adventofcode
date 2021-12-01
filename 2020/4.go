package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passports := []map[string]string{}
	pass := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			pairs := strings.Fields(line)
			for _, pair := range pairs {
				p := strings.Split(pair, ":")
				pass[p[0]] = p[1]
			}
		} else {
			passports = append(passports, pass)
			pass = map[string]string{}
		}
	}
	passports = append(passports, pass)

	type validator func(string) bool
	fields := map[string]validator{"byr": byr,
		"iyr": iyr,
		"eyr": eyr,
		"hgt": hgt,
		"hcl": hcl,
		"ecl": ecl,
		"pid": pid,
	}

	countKeys, countFull := 0, 0
	for _, pass := range passports {
		validKeys, validFull := true, true
		for key, f := range fields {
			if val, ok := pass[key]; ok {
				if !f(val) {
					validFull = false
				}
			} else {
				validKeys, validFull = false, false
				break
			}
		}
		if validKeys {
			countKeys++
		}
		if validFull {
			fmt.Println(pass)
			countFull++
		}
	}

	fmt.Println("Part One: ", countKeys)
	fmt.Println("Part Two: ", countFull)
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.

func byr(str string) bool {
	num, err := strconv.ParseInt(str, 10, 32)
	if err == nil && num >= 1920 && num <= 2002 && len(str) == 4 {
		return true
	}
	return false
}
func iyr(str string) bool {
	num, err := strconv.ParseInt(str, 10, 32)
	if err == nil && num >= 2010 && num <= 2020 && len(str) == 4 {
		return true
	}
	return false
}
func eyr(str string) bool {
	num, err := strconv.ParseInt(str, 10, 32)
	if err == nil && num >= 2020 && num <= 2030 && len(str) == 4 {
		return true
	}
	return false
}
func hgt(str string) bool {
	if len(str) < 4 {
		return false
	}
	unit := str[len(str)-2:]
	num, err := strconv.ParseInt(str[:len(str)-2], 10, 32)
	if err == nil && unit == "cm" && num >= 150 && num <= 193 {
		return true
	}
	if err == nil && unit == "in" && num >= 59 && num <= 76 {
		return true
	}
	return false
}
func hcl(str string) bool {
	return regexp.MustCompile("^#[0-9a-f]{6}$").Match([]byte(str))
}
func ecl(str string) bool {
	opts := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, col := range opts {
		if col == str {
			return true
		}
	}
	return false
}
func pid(str string) bool {
	return regexp.MustCompile("^[0-9]{9}$").Match([]byte(str))
}
