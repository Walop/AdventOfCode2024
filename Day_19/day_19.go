package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	design_cache = make(map[string]uint64)
)

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	possible := make(map[string]bool)
	not_possible := make(map[string]bool)
	// possible_max_len := 0

	scanner.Scan()
	for _, t := range strings.Split(scanner.Text(), ", ") {
		possible[t] = true
		// possible_max_len = len(t)
	}

	for p := range possible {
		for i := 1; i < len(p); i++ {
			fragment := p[:i]
			if !possible[fragment] {
				not_possible[fragment] = true
			}
		}
		for i := 0; i < len(p)-1; i++ {
			fragment := p[i:]
			if !possible[fragment] {
				not_possible[fragment] = true
			}
		}

	}

	designs := make([]string, 0, 400)

	scanner.Scan()

	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}

	// fmt.Println(possible)
	// fmt.Println(designs)

	possible_count := 0
	total_ways := uint64(0)

	for i := range designs {
		fmt.Println("Trying design ", i)
		different_ways := build_design_stolen(designs[i], possible)
		if different_ways > 0 {
			possible_count++
		}
		total_ways += different_ways
	}

	fmt.Println(possible_count)
	fmt.Println(total_ways)
}

func build_design_stolen(design string, possible map[string]bool) uint64 {
	if len(design) == 0 {
		return 1
	}
	count, found := design_cache[design]
	if found {
		return count
	}

	for p := range possible {
		if strings.HasPrefix(design, p) {
			count += build_design_stolen(design[len(p):], possible)
		}
	}
	design_cache[design] = count
	return count
}

func build_design(design string, possible map[string]bool, possible_max_len *int, not_possible map[string]bool, built string) uint64 {
	count, found := design_cache[design]
	if found {
		return count
	}
	if len(design) == 0 {
		return 1
	}
	if not_possible[design] {
		return 0
	}

	loop_len := *possible_max_len
	if len(design) < loop_len {
		loop_len = len(design)
	}

	possible_count := uint64(0)

	for i := loop_len; i > 0; i-- {
		fragment := string(design[:i])
		if possible[fragment] {
			new_built := built + fragment
			possible[built] = true
			if len(built) > *possible_max_len {
				*possible_max_len = len(built)
			}
			next := string(design[i:])
			count, found := design_cache[next]
			if found {
				possible_count = count
				break
			}
			possible_count += build_design(next, possible, possible_max_len, not_possible, new_built)
		}
	}
	if possible_count == 0 {
		not_possible[design] = true
	}
	design_cache[design] = possible_count
	return possible_count
}
