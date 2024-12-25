package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	keys := make([][]int, 0, 300)
	locks := make([][]int, 0, 300)

	schematic := make([]int, 5)
	schematic_type := 0

	for scanner.Scan() {
		row := scanner.Text()
		if len(row) == 0 {
			switch schematic_type {
			case 1:
				keys = append(keys, schematic)
			case 2:
				locks = append(locks, schematic)
			default:
				panic("No schematic type")
			}
			schematic_type = 0
			schematic = make([]int, 5)
			continue
		}
		if schematic_type == 0 {
			switch row {
			case ".....":
				schematic_type = 1
			case "#####":
				schematic_type = 2
			default:
				panic("Could not determine schematic type")
			}
		}
		for i, b := range row {
			if b == '#' {
				schematic[i] += 1
			}
		}
	}

	switch schematic_type {
	case 1:
		keys = append(keys, schematic)
	case 2:
		locks = append(locks, schematic)
	default:
		panic("No schematic type")
	}

	fmt.Println("Found", len(keys), "keys and", len(locks), "locks")

	pairs_fit := 0

	for i := range keys {
		for j := range locks {
			fits := 0
			for k := range len(keys[0]) {
				if keys[i][k]+locks[j][k] > 7 {
					break
				}
				fits++
			}
			if fits == 5 {
				pairs_fit++
			}
		}
	}

	fmt.Println("Key-lock pairs fit", pairs_fit)
}
