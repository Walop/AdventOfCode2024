package main

import (
	"AdventOfCode2024/util/timer"
	"AdventOfCode2024/util/vec2"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Layout struct {
	height_map [][]uint8
	height     int
	width      int
}

type Result struct {
	routes    int
	reachable int
}

var (
	DIRS [4]vec2.Vec2 = [4]vec2.Vec2{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}
)

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	layout := Layout{
		height_map: make([][]uint8, 0, 43),
		height:     0,
		width:      1,
	}

	start_pos := make([]vec2.Vec2, 0, 150)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]uint8, 0, 43)
		for i, r := range line {
			if i+1 > layout.width {
				layout.width = i + 1
			}
			num, _ := strconv.ParseUint(string(r), 10, 8)
			row = append(row, uint8(num))
			if num == 0 {
				start_pos = append(start_pos, vec2.Vec2{X: i, Y: layout.height})
			}
		}
		layout.height_map = append(layout.height_map, row)
		layout.height++
	}

	// for _, row := range layout.heightMap {
	// 	for _, val := range row {
	// 		fmt.Print(strconv.FormatUint(uint64(val), 10))
	// 	}
	// 	fmt.Print("\n")
	// }

	ch := make(chan Result)

	for _, pos := range start_pos {
		go FindRoutes(&layout, pos, ch)
	}

	sum := 0
	uq_sum := 0

	for range start_pos {
		res := <-ch
		sum += res.routes
		uq_sum += res.reachable
	}

	fmt.Println("Part 1: ", uq_sum)
	fmt.Println("Part 2: ", sum)
}

func FindRoutes(layout *Layout, start_pos vec2.Vec2, ch chan Result) {
	results := FindRoute(layout, start_pos)
	uq := make(map[vec2.Vec2]bool)
	for _, r := range results {
		uq[r] = true
	}

	ch <- Result{
		routes:    len(results),
		reachable: len(uq),
	}
}

func FindRoute(layout *Layout, pos vec2.Vec2) []vec2.Vec2 {
	current := layout.height_map[pos.Y][pos.X]
	if current == 9 {
		return []vec2.Vec2{pos}
	}
	results := make([]vec2.Vec2, 0, 20)
	for _, d := range DIRS {
		next_pos := pos.Add(&d)
		if IsOutOfBounds(layout, &next_pos) {
			continue
		}
		next := layout.height_map[next_pos.Y][next_pos.X]
		if next-current == 1 {
			results = append(results, FindRoute(layout, next_pos)...)
		}
	}
	return results
}

func IsOutOfBounds(layout *Layout, p *vec2.Vec2) bool {
	switch {
	case p.X < 0:
	case p.X == layout.width:
	case p.Y < 0:
	case p.Y == layout.height:
	default:
		return false
	}

	return true
}
