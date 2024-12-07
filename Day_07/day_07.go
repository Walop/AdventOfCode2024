package main

import (
	"bufio"
	"fmt"
	"os"
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

	equations := make([][]uint64, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, ": ")
		result, _ := strconv.ParseUint(split[0], 10, 64)
		values := strings.Split(split[1], " ")
		ivals := make([]uint64, 0, 16)
		for i := range values {
			ix := len(values) - i - 1
			ival, _ := strconv.ParseUint(values[ix], 10, 64)
			ivals = append(ivals, ival)
		}
		ivals = append(ivals, result)
		equations = append(equations, ivals)
	}

	// fmt.Println(equations)

	Part1(equations)
	Part2(equations)
}

func Part1(equations [][]uint64) {
	ch := make(chan uint64)

	for _, nums := range equations {
		go DoCount(nums, ch)
	}

	valid_sum := uint64(0)

	for range equations {
		valid_sum += <-ch
	}

	fmt.Println("Part 1: ", valid_sum)
}

func DoCount(nums []uint64, ch chan uint64) {
	nums, result := Pop(nums)
	nums, first := Pop(nums)
	if !Count(result, first, nums) {
		result = 0
	}
	ch <- result
}

func Count(result uint64, tres uint64, nums []uint64) bool {
	if tres > result {
		return false
	}

	if len(nums) == 0 {
		return tres == result
	}
	nums, num := Pop(nums)
	if Count(result, tres*num, nums) {
		return true
	}
	return Count(result, tres+num, nums)
}

func Part2(equations [][]uint64) {
	ch := make(chan uint64)

	for _, nums := range equations {
		go DoCount2(nums, ch)
	}

	valid_sum := uint64(0)

	for range equations {
		valid_sum += <-ch
	}

	fmt.Println("Part 2: ", valid_sum)
}

func DoCount2(nums []uint64, ch chan uint64) {
	nums, result := Pop(nums)
	nums, first := Pop(nums)
	if !Count2(result, first, nums) {
		result = 0
	}
	ch <- result
}

func Count2(result uint64, tres uint64, nums []uint64) bool {
	if tres > result {
		return false
	}

	if len(nums) == 0 {
		return tres == result
	}
	nums, num := Pop(nums)

	if Count2(result, tres*num, nums) {
		return true
	}
	if Count2(result, tres+num, nums) {
		return true
	}

	multiplier := 1
	dnum := num
	for dnum > 0 {
		multiplier *= 10
		dnum /= 10
	}

	return Count2(result, tres*uint64(multiplier)+num, nums)
}

func Pop(stack []uint64) ([]uint64, uint64) {
	last := len(stack) - 1
	return stack[:last], stack[last]
}
