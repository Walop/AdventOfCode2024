package main

import (
	"AdventOfCode2024/util/timer"
	"fmt"
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
		b = a & 7                                  // B MOD 8
		b ^= 2                                     // B XOR 2
		c = a >> b                                 // A / pow(2,b)
		a >>= 3                                    // A / 8
		b ^= c                                     // B XOR C
		b ^= 7                                     // B XOR 7
		output = append(output, strconv.Itoa(b&7)) // output B MOD 8
	}

	fmt.Println("Part 1 :", strings.Join(output, ","))
}
