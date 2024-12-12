package main

import (
	"AdventOfCode2024/util/timer"
	"AdventOfCode2024/util/vec2"
	"bufio"
	"fmt"
	"os"
)

var (
	dirs = []vec2.Vec2{
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
		{X: -1, Y: -1},
	}
)

type searchState struct {
	area      int
	perimeter int
	corners   int
	garden    [][]byte
	visited   map[vec2.Vec2]bool
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	garden := make([][]byte, 0, 140)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		garden = append(garden, line)
	}

	visited := make(map[vec2.Vec2]bool)

	sum1 := 0
	sum2 := 0

	for y := range len(garden) {
		for x := range len(garden[0]) {
			state := searchState{
				0, 0, 0, garden, visited,
			}
			findAreaAndPerimeter(
				vec2.Vec2{X: x, Y: y},
				&state)
			sum1 += state.area * state.perimeter
			sum2 += state.area * state.corners
		}
	}

	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}

func findAreaAndPerimeter(p vec2.Vec2, state *searchState) {
	if state.visited[p] {
		return
	}
	state.visited[p] = true
	state.area++

	c := state.garden[p.Y][p.X]

	for i := 0; i < len(dirs); i += 2 {
		next := p.Add(&dirs[i])
		if isValid(&next, c, state.garden) {
			findAreaAndPerimeter(next, state)
		} else {
			state.perimeter++
		}
	}

	valid := make([]bool, 0, len(dirs))

	for _, d := range dirs {
		next := p.Add(&d)
		valid = append(valid, isValid(&next, c, state.garden))
	}

	for i := 0; i < len(valid); i += 2 {
		if !valid[i] && !valid[(i+2)%len(valid)] {
			state.corners++
			continue
		}
		if valid[i] && valid[(i+2)%len(valid)] && !valid[(i+1)%len(valid)] {
			state.corners++
		}
	}
}

func isValid(pos *vec2.Vec2, c byte, garden [][]byte) bool {
	switch {
	case pos.X < 0:
	case pos.X >= len(garden):
	case pos.Y < 0:
	case pos.Y >= len(garden[0]):
	case garden[pos.Y][pos.X] != c:
	default:
		return true
	}

	return false
}
