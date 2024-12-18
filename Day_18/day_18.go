package main

import (
	"AdventOfCode2024/util/timer"
	"AdventOfCode2024/util/vec2"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	pos  vec2.Vec2
	path []vec2.Vec2
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	corrupted := make(map[vec2.Vec2]int)

	time := 0

	for scanner.Scan() {
		row := scanner.Text()
		split := strings.Split(row, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		point := vec2.Vec2{
			X: x,
			Y: y,
		}
		corrupted[point] = time
		time++
	}

	start := vec2.Vec2{X: 0, Y: 0}
	end := vec2.Vec2{X: 70, Y: 70}

	// for y := range end.Y {
	// 	for x := range end.Y {
	// 		c_time, ok := corrupted[vec2.Vec2{X: x, Y: y}]
	// 		if ok && c_time < 12 {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	shortest := find_path(12, start, end, corrupted)

	fmt.Println("Part 1: ", len(shortest.path))

	current_time := 0

	for shortest.pos.X == end.X {
		current_time++
		shortest = find_path(current_time, start, end, corrupted)
	}

	var block vec2.Vec2

	for k, v := range corrupted {
		if v == current_time {
			block = k
		}
	}

	fmt.Println("Part 2:", current_time, block)
}

func find_path(current_time int, start vec2.Vec2, end vec2.Vec2, corrupted map[vec2.Vec2]int) node {
	dirs := [4]vec2.Vec2{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}

	to_visit := make([]node, 0, end.X*end.Y)
	to_visit = append(to_visit, node{start, make([]vec2.Vec2, 0, 0)})

	visited := make(map[vec2.Vec2]bool)

	for len(to_visit) > 0 {
		current := to_visit[0]
		to_visit = to_visit[1:]

		visited[current.pos] = true

		if current.pos == end {
			return current
		}

		for d := range dirs {
			next := current.pos.Add(&dirs[d])
			c_time, ok := corrupted[next]
			if inBounds(next, start, end) && !visited[next] && !(ok && c_time <= current_time) && !contains_pos(to_visit, next) {
				next_path := make([]vec2.Vec2, 0, len(current.path)+1)
				next_path = append(next_path, current.path...)
				next_path = append(next_path, next)
				to_visit = append(to_visit, node{next, next_path})
			}
		}
	}

	return node{}
}

func inBounds(p vec2.Vec2, start vec2.Vec2, end vec2.Vec2) bool {
	switch {
	case p.X < start.X:
	case p.X > end.X:
	case p.Y < start.Y:
	case p.Y > end.Y:
	default:
		return true
	}

	return false
}

func contains_pos(nodes []node, pos vec2.Vec2) bool {
	for i := range nodes {
		if nodes[i].pos == pos {
			return true
		}
	}
	return false
}
