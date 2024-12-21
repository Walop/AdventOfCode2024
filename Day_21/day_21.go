package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}

func (v *point) add(other *point) *point {
	return &point{
		v.x + other.x,
		v.y + other.y,
	}
}

func (v *point) multiply(multiplier int) *point {
	return &point{
		v.x * multiplier,
		v.y * multiplier,
	}
}

var numpad = map[byte]point{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'0': {1, 3},
	'A': {2, 3},
}

var arrowpad = map[byte]point{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

var reverse_arrowpad = map[point]byte{
	{1, 0}: '^',
	{0, 1}: '<',
	{1, 1}: 'v',
	{2, 1}: '>',
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	codes := make([][]byte, 0, 5)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		row := scanner.Bytes()
		codes = append(codes, append([]byte(nil), row...))
	}

	sum := 0

	// tmp1 := get_arrowpad_keys([]byte{'v', '<', 'v', 'A'})
	// tmp2 := get_arrowpad_keys(tmp1)
	// tmp5 := get_arrowpad_keys(tmp2)
	// tmp7 := get_arrowpad_keys(tmp5)
	// fmt.Println(len(tmp7))
	// tmp3 := get_arrowpad_keys([]byte{'v', '<', '<', 'A'})
	// tmp4 := get_arrowpad_keys(tmp3)
	// tmp6 := get_arrowpad_keys(tmp4)
	// tmp8 := get_arrowpad_keys(tmp6)
	// fmt.Println(len(tmp8))

	// return

	for _, c := range codes {
		keyseq1 := get_numpad_keys(c)

		fmt.Println(string(keyseq1))

		keyseq2 := get_arrowpad_keys(keyseq1, 2)

		fmt.Println(string(keyseq2))

		keyseq3 := get_arrowpad_keys(keyseq2, 1)

		numeric, _ := strconv.Atoi(string(c[0:3]))

		fmt.Println(len(keyseq3), numeric, string(keyseq3))
		fmt.Println()
		sum += numeric * len(keyseq3)
	}

	fmt.Println("Part 1", sum)
}

func get_numpad_keys(code []byte) []byte {
	current_pos := numpad['A']
	keyseq := make([]byte, 0, 20)

	for len(code) > 0 {
		next := numpad[code[0]]
		if current_pos == next {
			keyseq = append(keyseq, 'A')
			code = code[1:]
			continue
		}
		target_x := next.x - current_pos.x
		target_y := next.y - current_pos.y
		switch {
		case (next.x != 0 || current_pos.y != 3) && target_x < 0:
			for range -target_x {
				keyseq = append(keyseq, '<')
				current_pos = *current_pos.add(&point{-1, 0})
			}
		case (next.y != 3 || current_pos.x != 0) && target_y > 0:
			for range target_y {
				keyseq = append(keyseq, 'v')
				current_pos = *current_pos.add(&point{0, 1})
			}
		case target_x > 0:
			for range target_x {
				keyseq = append(keyseq, '>')
				current_pos = *current_pos.add(&point{1, 0})
			}
		case target_y < 0:
			for range -target_y {
				keyseq = append(keyseq, '^')
				current_pos = *current_pos.add(&point{0, -1})
			}
		case target_x < 0:
			for range -target_x {
				keyseq = append(keyseq, '<')
				current_pos = *current_pos.add(&point{-1, 0})
			}
		case target_y > 0:
			for range target_y {
				keyseq = append(keyseq, 'v')
				current_pos = *current_pos.add(&point{0, 1})
			}
		default:
			panic("All cases should be handled")
		}
	}

	return keyseq
}

func get_arrowpad_keys(input []byte, depth int) []byte {
	current_pos := arrowpad['A']
	keyseq := make([]byte, 0, 150)

	for len(input) > 0 {
		next := arrowpad[input[0]]
		if current_pos == next {
			keyseq = append(keyseq, 'A')
			input = input[1:]
			continue
		}
		target_x := next.x - current_pos.x
		target_y := next.y - current_pos.y
		switch {
		case (current_pos.y != 0 || current_pos.x != 1) && (depth%2 == 1 || next.x == 1) && target_x < 0:
			keyseq = append(keyseq, '<')
			current_pos = *current_pos.add(&point{-1, 0})
		case next.y != 0 && target_y < 0:
			keyseq = append(keyseq, '^')
			current_pos = *current_pos.add(&point{0, -1})
		case target_y > 0:
			keyseq = append(keyseq, 'v')
			current_pos = *current_pos.add(&point{0, 1})
		case target_x > 0:
			keyseq = append(keyseq, '>')
			current_pos = *current_pos.add(&point{1, 0})
		case target_x < 0:
			keyseq = append(keyseq, '<')
			current_pos = *current_pos.add(&point{-1, 0})
		case target_y < 0:
			keyseq = append(keyseq, '^')
			current_pos = *current_pos.add(&point{0, -1})
		default:
			panic("All cases should be handled")
		}
	}

	return keyseq
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
