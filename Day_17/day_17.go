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

	a_long := uint64(0)

	for _, val := range program {
		a_long <<= 3
		for x := range 8 {
			if run(a_long+uint64(x)) == val {
				a_long += uint64(x)
				break
			}
		}
	}

	fmt.Println("Part 2: ", a_long)
}

func run(a uint64) uint64 {
	b := a & 7   // A MOD 8
	b ^= 2       // B XOR 2
	c := a >> b  // A / pow(2,b)
	b ^= c       // B XOR C
	b ^= 7       // B XOR 7
	return b & 7 // output B MOD 8
}
