package main

import (
	"AdventOfCode2024/util/timer"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer timer.Timer("Main")()

	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(f)

	initial := make([]int, 0, 10)

	var nums = strings.Split(input, " ")
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		initial = append(initial, n)
	}

	part1(initial)
	part2(initial)
}

func part1(initial []int) {
	arr := make([]int, 0, 200_000)

	arr = append(arr, initial...)

	for range 25 {
		for i, elem := range arr {
			res := process(&elem)
			arr[i] = res[0]
			if len(res) > 1 {
				arr = append(arr, res[1])
			}
		}
	}

	fmt.Println("Part 1: ", len(arr))
}

func part2(initial []int) {
	counts := make(map[int]int64)
	for _, elem := range initial {
		counts[elem] += 1
	}

	for range 75 {
		next := make(map[int]int64)
		for num, count := range counts {
			res := process(&num)
			for _, val := range res {
				next[val] += count
			}
		}
		counts = next
	}

	sum := int64(0)

	for _, v := range counts {
		sum += v
	}

	fmt.Println("Part 2: ", sum)
}

func process(num *int) []int {
	if *num == 0 {
		return []int{1}
	}
	magnitude := 0
	temp := *num
	for temp > 0 {
		magnitude += 1
		temp /= 10
	}
	if magnitude%2 == 1 {
		return []int{*num * 2024}
	}

	divider := math.Pow10(magnitude / 2)
	return []int{
		*num / int(divider),
		*num % int(divider),
	}
}
