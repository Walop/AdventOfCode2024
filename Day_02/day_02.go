package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rows := make([][]int, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		nums := make([]int, len(split))
		for i, v := range split {
			nums[i], _ = strconv.Atoi(v)
		}
		rows = append(rows, nums)
	}

	Part1(rows)
	Part2(rows)
}

func Part1(rows [][]int) {
	safe_count := 0

	for _, row := range rows {
		if IsSafe(row) {
			safe_count++
		}
	}

	fmt.Println("Part 1: ", safe_count)
}

func Part2(rows [][]int) {
	safe_count := 0

	for _, row := range rows {

		if IsSafe(row) {
			safe_count++
			continue
		}

		for i := range row {
			new_row := make([]int, 0, len(row)-1)
			new_row = append(new_row, row[:i]...)
			new_row = append(new_row, row[i+1:]...)
			if IsSafe(new_row) {
				safe_count++
				break
			}
		}
	}

	fmt.Println("Part 2: ", safe_count)
}

func IsSafe(row []int) bool {
	prev := row[0]
	safe := true
	asc := row[0] < row[1]
	for _, n := range row[1:] {
		change := n - prev
		if change < 0 && asc {
			safe = false
			break
		}
		if change > 0 && !asc {
			safe = false
			break
		}
		magnitude := Abs(change)
		if magnitude < 1 || magnitude > 3 {
			safe = false
			break
		}
		prev = n
	}
	return safe
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
