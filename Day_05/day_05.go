package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	instructions := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, "|")
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])
		instructions[first] = append(instructions[first], second)
	}

	//fmt.Println(instructions)

	updates := make([][]int, 0, 200)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		update := make([]int, 0, 25)
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	// fmt.Println(updates)

	invalid := Part1(instructions, updates)

	Part2(instructions, invalid)
}

func Part1(instructions map[int][]int, updates [][]int) [][]int {
	sum := 0

	invalid := make([][]int, 0, 200)

	for _, update := range updates {
		after := make([]int, 0, 255)
		for _, u := range update {
			after = append(after, instructions[u]...)
			for i := range after {
				if after[i] == u {
					after[i] = 0
				}
			}
		}

		for i, a := range after {
			if !slices.Contains(update, a) {
				after[i] = 0
			}
		}

		valid := true

		for _, a := range after {
			if a != 0 {
				valid = false
				break
			}
		}

		if valid {
			middle := len(update) / 2
			sum += update[middle]
		} else {
			invalid = append(invalid, update)
		}
	}

	fmt.Println("Part 1: ", sum)

	return invalid
}

func Part2(instructions map[int][]int, updates [][]int) {
	sum := 0

	for _, update := range updates {
		sorted := []int{0}
		valid_instructions := make(map[int][]int)
		for _, u := range update {
			for i, v := range instructions {
				if u == i || slices.Contains(v, u) {
					valid_instructions[u] = slices.Clone(instructions[u])
				}
			}
		}

		for _, v := range valid_instructions {
			for j := range v {
				if !slices.Contains(update, v[j]) {
					v[j] = 0
				}
			}
		}

		// fmt.Println(valid_instructions)

		for {
			for _, u := range update {
				all := true
				for _, i := range valid_instructions[u] {
					if !slices.Contains(sorted, i) {
						all = false
						break
					}
				}
				if all && !slices.Contains(sorted, u) {
					sorted = append(sorted, u)
					break
				}
			}

			if len(sorted) == len(update)+1 {
				break
			}
		}

		middle := len(sorted) / 2
		sum += sorted[middle]
	}

	fmt.Println("Part 2: ", sum)
}
