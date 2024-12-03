package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	Part1(f)

	f.Seek(0, io.SeekStart)

	Part2(f)
}

func Part1(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	mul := []rune("mul(")
	mul_pos := 0

	num := ""

	num1 := 0

	sum := 0

	for scanner.Scan() {
		r := []rune(scanner.Text())
		if mul_pos < len(mul) && r[0] != mul[mul_pos] {
			mul_pos = 0
			continue
		}

		if mul_pos < len(mul) {
			mul_pos++
			continue
		}

		if utf8.RuneCountInString(num) > 3 {
			mul_pos = 0
			continue
		}

		if unicode.IsDigit(r[0]) {
			num += string(r[0])
			continue
		}

		if r[0] == rune(',') {
			num1, _ = strconv.Atoi(num)
			num = ""
			continue
		}

		if r[0] == rune(')') {
			num2, _ := strconv.Atoi(num)
			sum += num1 * num2
		}

		mul_pos = 0
		num1 = 0
		num = ""
	}

	fmt.Println("Part 1: ", sum)
}

func Part2(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	mul := []rune("mul(")
	mul_pos := 0

	num := ""

	num1 := 0

	sum := 0

	do_flag := true

	do := []rune("do()")
	do_pos := 0

	dont := []rune("don't()")
	dont_pos := 0

	for scanner.Scan() {
		r := []rune(scanner.Text())

		if r[0] == do[do_pos] {
			do_pos++
		} else {
			do_pos = 0
		}

		if do_pos == len(do) {
			do_flag = true
			do_pos = 0
		}

		if r[0] == dont[dont_pos] {
			dont_pos++
		} else {
			dont_pos = 0
		}

		if dont_pos == len(dont) {
			do_flag = false
			dont_pos = 0
		}

		if !do_flag {
			mul_pos = 0
			num1 = 0
			num = ""
			continue
		}
		if mul_pos < len(mul) && r[0] != mul[mul_pos] {
			mul_pos = 0
			continue
		}

		if mul_pos < len(mul) {
			mul_pos++
			continue
		}

		if utf8.RuneCountInString(num) > 3 {
			mul_pos = 0
			continue
		}

		if unicode.IsDigit(r[0]) {
			num += string(r[0])
			continue
		}

		if r[0] == rune(',') {
			num1, _ = strconv.Atoi(num)
			num = ""
			continue
		}

		if r[0] == rune(')') {
			num2, _ := strconv.Atoi(num)
			sum += num1 * num2
		}

		mul_pos = 0
		num1 = 0
		num = ""
	}

	fmt.Println("Part 1: ", sum)
}
