package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type port struct {
	operation string
	input1    string
	input2    string
	output    string
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	wires := make(map[string]bool, 200)

	for scanner.Scan() {
		row := scanner.Text()
		if len(row) == 0 {
			break
		}
		split := strings.Split(row, ": ")
		wires[split[0]] = split[1] == "1"
	}

	ports := make([]port, 0, 200)

	for scanner.Scan() {
		row := scanner.Text()
		split1 := strings.Split(row, " -> ")
		split2 := strings.Split(row, " ")
		port1 := split2[0]
		port2 := split2[2]
		if port1 < port2 {
			port1 = split2[2]
			port2 = split2[0]
		}
		ports = append(ports, port{
			split2[1],
			port1,
			port2,
			split1[1],
		})
	}

	slices.SortFunc(ports, compare_ports)

	part1(copy_map(wires), append([]port{}, ports...))
	part2(copy_map(wires), append([]port{}, ports...))
	//produce_dot(ports)
}

func compare_ports(a port, b port) int {
	if len(a.input1)+len(a.input2) > len(b.input1)+len(b.input2) {
		return 1
	}
	if len(a.input1)+len(a.input2) < len(b.input1)+len(b.input2) {
		return -1
	}
	if a.input1 > b.input1 {
		return -1
	}
	if a.input1 < b.input1 {
		return 1
	}
	if a.input2 > b.input2 {
		return -1
	}
	if a.input2 < b.input2 {
		return 1
	}
	if a.operation > b.operation {
		return -1
	}
	return 1
}

func part1(wires map[string]bool, ports []port) {

	evaluate_all_wires(wires, ports)
	// fmt.Println(wires)
	// fmt.Println(ports)

	result := build_number('z', wires)

	fmt.Println("Part 1: ", result)
}

func copy_map(m map[string]bool) map[string]bool {
	m2 := make(map[string]bool, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

func evaluate_all_wires(wires map[string]bool, ports []port) {
	for len(ports) > 0 {
		i := 0
		for {
			p := ports[i]
			i1, exists1 := wires[p.input1]
			i2, exists2 := wires[p.input2]
			if exists1 && exists2 {
				wires[p.output] = evaluate_gate(p.operation, i1, i2)
				ports = append(ports[:i], ports[i+1:]...)
			} else {
				i++
			}
			if i >= len(ports) {
				break
			}
		}
	}
}

func evaluate_gate(operation string, input1 bool, input2 bool) bool {
	switch {
	case operation == "OR":
		return input1 || input2
	case operation == "AND":
		return input1 && input2
	case operation == "XOR":
		return !(input1 == input2)
	}

	panic("Undefined operation")
}

func part2(wires map[string]bool, ports []port) {
	for i := range 45 {
		num := strconv.Itoa(i)
		for len(num) < 2 {
			num = "0" + num
		}
		wires["x"+num] = true
		wires["y"+num] = true
	}

	nxz := make([]int, 0, 3)
	xnz := make([]int, 0, 3)
	swapped := make([]string, 0, 8)
	for i, p := range ports {
		if p.output[0] == 'z' && p.output != "z45" && p.operation != "XOR" {
			nxz = append(nxz, i)
			swapped = append(swapped, p.output)
		}
		if p.input1[0] < 'x' && p.output[0] != 'z' && p.operation == "XOR" {
			xnz = append(xnz, i)
			swapped = append(swapped, p.output)
		}
	}

	for _, i := range xnz {
		a := first_z_using_output(ports, ports[i].output)
		b := 0
		for _, j := range nxz {
			if ports[j].output == a {
				b = j
				break
			}
		}
		temp := ports[i].output
		ports[i].output = ports[b].output
		ports[b].output = temp

		fmt.Println(temp, ports[i].output)
	}

	ports_temp := append([]port{}, ports...)
	wires_temp := copy_map(wires)

	x := build_number('x', wires)
	y := build_number('y', wires)

	evaluate_all_wires(wires_temp, ports_temp)

	z := build_number('z', wires_temp)

	fmt.Println()
	fmt.Printf("%046b %d\n", x, x)
	fmt.Printf("%046b %d\n", y, y)
	fmt.Printf("%046b %d\n", z, z)

	fmt.Println()

	diff := z - x - y
	if diff < 0 {
		diff = -diff
	}
	log_diff := 0
	if diff > 0 {
		log_diff = int(math.Log2(float64(diff)))
	}
	fmt.Printf("%d + %d = %d, diff %d log2 %d\n", x, y, z, diff, log_diff)

	false_carry := strconv.Itoa(log_diff)
	for len(false_carry) < 2 {
		false_carry = "0" + false_carry
	}
	carry_swap := make([]int, 0, 2)
	for i, p := range ports {
		if strings.HasSuffix(p.input1, false_carry) || strings.HasSuffix(p.input2, false_carry) {
			carry_swap = append(carry_swap, i)
			swapped = append(swapped, p.output)
		}
	}

	temp := ports[carry_swap[0]].output
	ports[carry_swap[0]].output = ports[carry_swap[1]].output
	ports[carry_swap[1]].output = temp

	evaluate_all_wires(wires, ports)

	z = build_number('z', wires)

	fmt.Println()
	fmt.Printf("%046b %d\n", x, x)
	fmt.Printf("%046b %d\n", y, y)
	fmt.Printf("%046b %d\n", z, z)

	fmt.Println()

	diff = z - x - y
	log_diff = 0
	if diff > 0 {
		log_diff = int(math.Log2(float64(diff)))
	}
	fmt.Printf("%d + %d = %d, diff %d log2 %d\n", x, y, z, diff, log_diff)

	slices.Sort(swapped)

	fmt.Println("Part 2: ", strings.Join(swapped, ","))
}

func first_z_using_output(ports []port, output string) string {
	next_output := ""
	for _, p := range ports {
		if p.input1 == output || p.input2 == output {
			next_output = p.output
			break
		}
	}
	if next_output[0] == 'z' {
		num, _ := strconv.Atoi(string(next_output[1:]))
		str := strconv.Itoa(num - 1)
		for len(str) < 2 {
			str = "0" + str
		}
		return "z" + str
	}
	return first_z_using_output(ports, next_output)
}

func build_number(start byte, wires map[string]bool) int64 {
	ouput_wires := make([]string, 0, 45)
	for k := range wires {
		if k[0] == start {
			ouput_wires = append(ouput_wires, k)
		}
	}

	slices.Sort(ouput_wires)
	slices.Reverse(ouput_wires)

	result := int64(0)
	for _, v := range ouput_wires {
		result <<= 1
		if wires[v] {
			result ^= 1
		}
	}
	return result
}

func produce_dot(ports []port) {
	fmt.Println("digraph {")

	for i := range 46 {
		if i < 45 {
			fmt.Printf("x%02d [fillcolor=\"blue\" style=\"filled\"];\n", i)
			fmt.Printf("y%02d [fillcolor=\"green\" style=\"filled\"];\n", i)
		}
		fmt.Printf("z%02d [fillcolor=\"red\" style=\"filled\"];\n", i)
	}

	for i, p := range ports {
		attr := ""
		switch p.operation {
		case "XOR":
			attr += "[fillcolor=\"yellow\" style=\"filled\"]"
		case "AND":
			attr += "[fillcolor=\"orange\" style=\"filled\"]"
		case "OR":
			attr += "[fillcolor=\"pink\" style=\"filled\"]"
		}

		fmt.Printf("%s %s;\n", p.operation+strconv.Itoa(i), attr)

		fmt.Printf("%s -> %s;\n", p.input1, p.operation+strconv.Itoa(i))
		fmt.Printf("%s -> %s;\n", p.input2, p.operation+strconv.Itoa(i))
		fmt.Printf("%s -> %s;\n", p.operation+strconv.Itoa(i), p.output)
	}

	fmt.Println("}")
}
