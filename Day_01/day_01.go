package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	left_nums := []int{}
	right_nums := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		left, _ := strconv.Atoi(split[0])
		left_nums = append(left_nums, left)

		right, _ := strconv.Atoi(split[1])
		right_nums = append(right_nums, right)
	}

	sort.Slice(left_nums, func(i, j int) bool { return left_nums[i] < left_nums[j] })
	sort.Slice(right_nums, func(i, j int) bool { return right_nums[i] < right_nums[j] })

	Part1(left_nums, right_nums)
	Part2(left_nums, right_nums)
}

func Part1(left_nums []int, right_nums []int) {
	sum := 0

	for i := range left_nums {
		sum += Abs(left_nums[i] - right_nums[i])
	}

	fmt.Println("Part 1: ", sum)
}

func Part2(left_nums []int, right_nums []int) {
	left_count := map[int]int{}
	right_count := map[int]int{}

	for _, val := range left_nums {
		left_count[val] += 1
	}

	for _, val := range right_nums {
		right_count[val] += 1
	}

	sum := 0

	for i := range left_count {
		sum += i * left_count[i] * right_count[i]
	}

	fmt.Println("Part 2: ", sum)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
