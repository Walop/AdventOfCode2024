package main

import (
	"AdventOfCode2024/util/vec2"
	"bufio"
	"fmt"
	"os"
)

var (
	DIRS [4]vec2.Vec2
)

type Layout struct {
	width     int
	height    int
	obstacles map[vec2.Vec2]bool
	guard     vec2.Vec2
}

type VD struct {
	v vec2.Vec2
	d int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	layout := new(Layout)

	layout.obstacles = make(map[vec2.Vec2]bool)

	for scanner.Scan() {
		line := scanner.Text()
		for xpos, r := range line {
			if r == '#' {
				layout.obstacles[vec2.Vec2{X: xpos, Y: layout.height}] = true
			}
			if r == '^' {
				layout.guard.X = xpos
				layout.guard.Y = layout.height
			}
			if xpos > layout.width {
				layout.width = xpos
			}
		}
		layout.height++
	}

	layout.width++

	DIRS = [4]vec2.Vec2{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}

	Part1(layout)
	Part2(layout)
}

func Part1(layout *Layout) {
	current_dir := 0

	visited := make(map[vec2.Vec2]bool)

	guard := layout.guard

	for {
		visited[guard] = true

		next := vec2.Add(guard, DIRS[current_dir])

		if layout.obstacles[next] {
			current_dir = (current_dir + 1) % len(DIRS)
			continue
		}

		if CheckOutOfBounds(layout, next) {
			break
		}
		guard = next
	}

	fmt.Println("Part 1: ", len(visited))
}

func Part2(layout *Layout) {
	current_dir := 0

	loop_obstacles := make(map[vec2.Vec2]bool)

	guard := layout.guard

	for {
		next := vec2.Add(guard, DIRS[current_dir])

		if layout.obstacles[next] {
			current_dir = (current_dir + 1) % len(DIRS)
			continue
		}

		if CheckOutOfBounds(layout, next) {
			break
		}

		if CheckLoop(layout, layout.guard, next, 0) {
			loop_obstacles[next] = true
		}
		guard = next
	}

	fmt.Println("Part 2: ", len(loop_obstacles))
}

func CheckLoop(layout *Layout, current vec2.Vec2, loop_obstacle vec2.Vec2, dir int) bool {
	visited := make(map[VD]bool)

	for {
		next := vec2.Add(current, DIRS[dir])

		for next == loop_obstacle || layout.obstacles[next] {
			dir = (dir + 1) % len(DIRS)
			next = vec2.Add(current, DIRS[dir])
		}

		if CheckOutOfBounds(layout, next) {
			break
		}

		if visited[VD{current, dir}] {
			return true
		}

		visited[VD{current, dir}] = true

		current = next
	}

	return false
}

func CheckOutOfBounds(layout *Layout, p vec2.Vec2) bool {
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
