package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"os"
)

type vec2 struct {
	x float64
	y float64
}

type claw_machine struct {
	button_a vec2
	button_b vec2
	prize    vec2
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	machines := make([]claw_machine, 0, 500)

	for scanner.Scan() {
		machine := claw_machine{}

		var x, y int
		fmt.Sscanf(scanner.Text(), "Button A: X+%d, Y+%d", &x, &y)
		machine.button_a.x = float64(x)
		machine.button_a.y = float64(y)

		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Button B: X+%d, Y+%d", &x, &y)
		machine.button_b.x = float64(x)
		machine.button_b.y = float64(y)

		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Prize: X=%d, Y=%d", &x, &y)
		machine.prize.x = float64(x)
		machine.prize.y = float64(y)

		machines = append(machines, machine)
		scanner.Scan()
	}

	ch := make(chan int64)

	for _, machine := range machines {
		go calculate_cost(&machine, ch)
	}

	sum := int64(0)

	for range machines {
		sum += <-ch
	}

	fmt.Println("Part 1: ", sum)

	for _, machine := range machines {
		machine.prize.x += 10_000_000_000_000
		machine.prize.y += 10_000_000_000_000
		go calculate_cost(&machine, ch)
	}

	sum = int64(0)

	for range machines {
		sum += <-ch
	}

	fmt.Println("Part 2: ", sum)
}

func calculate_cost(machine *claw_machine, ch chan int64) {
	a, b, c, d := machine.button_a.x, machine.button_a.y, machine.button_b.x, machine.button_b.y
	rx, ry := machine.prize.x, machine.prize.y
	m := (d*rx - c*ry) / (a*d - b*c)
	n := (b*rx - a*ry) / (b*c - a*d)

	i_m := int64(m)
	i_n := int64(n)
	if i_m <= 0 || i_n <= 0 || m != float64(i_m) || n != float64(i_n) {
		ch <- 0
		return
	}

	tokens := 3*i_m + i_n

	ch <- tokens
}
