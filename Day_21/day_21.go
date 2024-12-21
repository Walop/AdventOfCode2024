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

type keyseq_key struct {
	seq   string
	depth int
}

var keyseq_cache = map[keyseq_key]int{}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
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

		//fmt.Println(string(keyseq1))

		count := 0
		count += count_seq(keyseq1, 1, 2)
		// fmt.Println()
		fmt.Println(count)

		//fmt.Println(string(keyseq2))

		// keyseq3 := get_arrowpad_keys(keyseq2, 1)

		numeric, _ := strconv.Atoi(string(c[0:3]))

		sum += numeric * count
	}

	fmt.Println("Part 1", sum)

	sum = 0
	keyseq_cache = map[keyseq_key]int{}

	for _, c := range codes {
		keyseq1 := get_numpad_keys(c)

		count := 0
		count += count_seq(keyseq1, 1, 25)
		fmt.Println(count)

		numeric, _ := strconv.Atoi(string(c[0:3]))

		sum += numeric * count
	}

	// too high 473560765309814, 349738774242354, 307690647333324
	// not right 123433191944248
	// Should be right 307055584161760
	fmt.Println("Part 2", sum)
}

func get_numpad_keys(code []byte) []byte {
	current_pos := numpad['A']
	keyseq := make([]byte, 0, 20)

	for i := range code {
		next := numpad[code[i]]

		h, v := []byte{}, []byte{}

		target_x := next.x - current_pos.x
		target_y := next.y - current_pos.y
		for range abs(target_x) {
			if target_x > 0 {
				h = append(h, '>')
			} else {
				h = append(h, '<')
			}
		}

		for range abs(target_y) {
			if target_y > 0 {
				v = append(v, 'v')
			} else {
				v = append(v, '^')
			}
		}

		if current_pos.y == 3 && next.x == 0 {
			keyseq = append(keyseq, v...)
			keyseq = append(keyseq, h...)
		} else if current_pos.x == 0 && next.y == 3 {
			keyseq = append(keyseq, h...)
			keyseq = append(keyseq, v...)
		} else if target_x < 0 {
			keyseq = append(keyseq, h...)
			keyseq = append(keyseq, v...)
		} else {
			keyseq = append(keyseq, v...)
			keyseq = append(keyseq, h...)
		}

		current_pos = next
		keyseq = append(keyseq, 'A')
	}

	return keyseq
}

func count_seq(input []byte, depth int, max_depth int) int {
	c, found := keyseq_cache[keyseq_key{string(input), depth}]
	if found {
		return c
	}

	seq := get_arrowpad_keys(input)
	if depth == max_depth {
		return len(seq)
	}

	sum := 0
	for _, s := range split_seq(seq) {
		sum += count_seq(s, depth+1, max_depth)
	}

	keyseq_cache[keyseq_key{string(input), depth}] = sum
	return sum
}

func get_arrowpad_keys(input []byte) []byte {
	current_pos := arrowpad['A']
	keyseq := make([]byte, 0, 150)

	for i := range input {
		next := arrowpad[input[i]]

		h, v := []byte{}, []byte{}

		target_x := next.x - current_pos.x
		target_y := next.y - current_pos.y
		for range abs(target_x) {
			if target_x > 0 {
				h = append(h, '>')
			} else {
				h = append(h, '<')
			}
		}

		for range abs(target_y) {
			if target_y > 0 {
				v = append(v, 'v')
			} else {
				v = append(v, '^')
			}
		}

		if current_pos.x == 0 && next.y == 0 {
			keyseq = append(keyseq, h...)
			keyseq = append(keyseq, v...)
		} else if current_pos.y == 0 && next.x == 0 {
			keyseq = append(keyseq, v...)
			keyseq = append(keyseq, h...)
		} else if target_x < 0 {
			keyseq = append(keyseq, h...)
			keyseq = append(keyseq, v...)
		} else {
			keyseq = append(keyseq, v...)
			keyseq = append(keyseq, h...)
		}

		current_pos = next
		keyseq = append(keyseq, 'A')
	}

	return keyseq
}

func split_seq(input []byte) [][]byte {
	var result [][]byte
	var current []byte

	for _, char := range input {
		current = append(current, char)
		if char == 'A' {
			result = append(result, current)
			current = []byte{}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
