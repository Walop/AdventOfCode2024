package main

import (
	"AdventOfCode2024/util/timer"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	defer timer.Timer("Main")()

	a := 30878003
	b := 0
	c := 0

	output := make([]string, 0, 16)

	for a > 0 {
		b = a & 7                                  // A MOD 8
		b ^= 2                                     // B XOR 2
		c = a >> b                                 // A / pow(2,b)
		a >>= 3                                    // A / 8
		b ^= c                                     // B XOR C
		b ^= 7                                     // B XOR 7
		output = append(output, strconv.Itoa(b&7)) // output B MOD 8
	}

	fmt.Println("Part 1 :", strings.Join(output, ","))

	program := []uint64{2, 4, 1, 2, 7, 5, 0, 3, 4, 7, 1, 7, 5, 5, 3, 0}

	slices.Reverse(program)

	a_long := find_quine_input(0, program)

	fmt.Println("Part 2: ", a_long)
}

func find_quine_input(a uint64, program []uint64) uint64 {
	if len(program) == 0 {
		return a
	}
	val := program[0]
	next_program := program[1:]
	a <<= 3
	for x := range 8 {
		next_a := uint64(0)
		if run(a+uint64(x)) == val {
			fmt.Println("Found", val, next_program, "possible next octal", x)
			fmt.Printf("%b\n", a+uint64(x))
			fmt.Println()
			next_a = find_quine_input(a+uint64(x), next_program)
			if next_a != 0xFFFFFFFFFFFFFFFF {
				return next_a
			}
		}
		if x == 7 && next_a == 0xFFFFFFFFFFFFFFFF {
			fmt.Println(val, program)
			panic("Not found")
		}
	}
	return 0xFFFFFFFFFFFFFFFF
}

func run(a uint64) uint64 {
	b := a & 7   // A MOD 8
	b ^= 2       // B XOR 2
	c := a >> b  // A / pow(2,b)
	b ^= c       // B XOR C
	b ^= 7       // B XOR 7
	return b & 7 // output B MOD 8
}
