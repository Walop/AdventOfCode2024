package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strconv"
)

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	seeds := make([]int64, 0, 2500)

	for scanner.Scan() {
		row := scanner.Text()
		num, _ := strconv.ParseInt(row, 10, 64)
		seeds = append(seeds, num)
	}

	part1(seeds)

	part2(seeds)
}

func part1(seeds []int64) {
	cpus := runtime.NumCPU()

	channels := make([]chan int64, cpus)
	for x := range channels {
		channels[x] = make(chan int64)
	}

	i := 0
	for batch := range slices.Chunk(seeds, len(seeds)/cpus+1) {
		go do_batch(batch, 2000, channels[i])
		i++
	}

	sum := int64(0)

	for j := range i {
		for n := range channels[j] {
			sum += n
		}
	}

	fmt.Println("Part 1: ", sum)
}

func do_batch(seeds []int64, rounds int, ch chan int64) {
	sum := int64(0)
	for i := range seeds {
		sum += do_rounds(seeds[i], rounds)
	}
	ch <- sum
	close(ch)
}

func do_rounds(seed int64, rounds int) int64 {
	current := seed
	for range rounds {
		current = next_number(current)
	}
	return current
}

func part2(seeds []int64) {
	cpus := runtime.NumCPU()

	channels := make([]chan map[int64]int64, cpus)
	for x := range channels {
		channels[x] = make(chan map[int64]int64)
	}

	i := 0
	for batch := range slices.Chunk(seeds, len(seeds)/cpus+1) {
		go do_batch2(batch, 2000, channels[i])
		i++
	}

	bananas := make(map[int64]int64)

	max := int64(0)
	for j := range i {
		for b := range channels[j] {
			for k, v := range b {
				new_val := bananas[k] + v
				bananas[k] = new_val
				if new_val > max {
					max = new_val
				}
			}
		}
	}

	fmt.Println("Part 2: ", max)
}

func do_batch2(seeds []int64, rounds int, ch chan map[int64]int64) {
	bananas := make(map[int64]int64)
	for i := range seeds {
		do_rounds2(seeds[i], rounds, bananas)
	}
	ch <- bananas
	close(ch)
}

func do_rounds2(seed int64, rounds int, bananas map[int64]int64) {
	current := seed

	seen := map[int64]struct{}{}

	prev_list := make([]int64, 0, rounds)
	prev := seed % 10
	for range rounds {
		current = next_number(current)
		last_number := current % 10
		diff := last_number - prev
		prev = last_number
		prev_list = append(prev_list, diff)
		if len(prev_list) > 4 {
			prev_list = prev_list[1:]
		}
		if len(prev_list) == 4 {
			key := prev_list[0]<<20 + prev_list[1]<<15 + prev_list[2]<<10 + prev_list[3]<<5
			_, found := seen[key]
			if !found {
				bananas[key] += last_number
				seen[key] = struct{}{}
			}
		}
	}
}

func next_number(seed int64) int64 {
	next1 := seed << 6
	next1 ^= seed
	next1 &= 16_777_215

	next2 := next1 >> 5
	next2 ^= next1

	next3 := next2 << 11
	next3 ^= next2
	next3 &= 16_777_215

	return next3
}
