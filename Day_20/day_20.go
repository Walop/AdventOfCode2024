package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"os"
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

type pathpoint struct {
	point
	length int
}

var dirs = [4]point{
	{x: 0, y: -1},
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: -1, y: 0},
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	maze := make([][]byte, 0, 141)
	start := point{}
	end := point{}

	scanner := bufio.NewScanner(f)

	y := 0

	for scanner.Scan() {
		row := scanner.Bytes()
		copy_row := make([]byte, 0, len(row))
		copy_row = append(copy_row, row...)
		maze = append(maze, copy_row)
		for x := range row {
			switch row[x] {
			case 'S':
				start.x = x
				start.y = y
			case 'E':
				end.x = x
				end.y = y
			}
		}
		y++
	}

	to_end := without_cheat(maze, start, end)

	cheats := count_cheats(maze, end, start, to_end, 100)

	fmt.Println("Part 1: ", cheats)

	long_cheats := count_long_cheats(maze, end, start, to_end, 100)

	fmt.Println("Part 2: ", long_cheats)
}

func without_cheat(maze [][]byte, start point, end point) map[point]int {
	visited := make(map[point]int)

	to_visit := make([]pathpoint, 0, len(maze)*len(maze[0]))
	to_visit = append(to_visit, pathpoint{start, 0})

	for len(to_visit) > 0 {
		current := to_visit[0]
		to_visit = to_visit[1:]

		visited[current.point] = current.length

		if current.point == end {
			continue
		}

		for d := range dirs {
			next := *current.point.add(&dirs[d])
			_, v := visited[next]
			if !v && maze[next.y][next.x] != '#' {
				to_visit = append(to_visit, pathpoint{next, current.length + 1})
			}
		}
	}

	return visited
}

func count_cheats(maze [][]byte, start point, end point, to_end map[point]int, time_save int) int {
	visited := make(map[point]int)

	to_visit := make([]pathpoint, 0, len(maze)*len(maze[0]))
	to_visit = append(to_visit, pathpoint{start, 0})

	cheats := 0

	for len(to_visit) > 0 {
		current := to_visit[0]
		to_visit = to_visit[1:]

		visited[current.point] = current.length

		if current.point == end {
			continue
		}

		for d := range dirs {
			next := *current.point.add(&dirs[d])
			_, v := visited[next]
			if !v && maze[next.y][next.x] != '#' {
				to_visit = append(to_visit, pathpoint{next, current.length + 1})
			}
			cheat_next := *current.point.add(&dirs[d]).add(&dirs[d])
			cheat_start := to_end[current.point]
			cheat_end, found := to_end[cheat_next]
			if found && cheat_start-cheat_end-2 >= time_save {
				cheats++
			}
		}
	}

	return cheats
}

func count_long_cheats(maze [][]byte, start point, end point, to_end map[point]int, time_save int) int {
	visited := make(map[point]int)

	to_visit := make([]pathpoint, 0, len(maze)*len(maze[0]))
	to_visit = append(to_visit, pathpoint{start, 0})

	in_bounds := make_in_bounds(len(maze[0]), len(maze))

	cheats := 0

	visited_count := 0

	for len(to_visit) > 0 {
		current := to_visit[0]
		to_visit = to_visit[1:]

		visited[current.point] = current.length

		if current.point == end {
			continue
		}

		visited_count++

		cheats += long_cheat(maze, current.point, to_end, time_save, in_bounds)

		for d := range dirs {
			next := *current.point.add(&dirs[d])
			_, v := visited[next]
			if !v && maze[next.y][next.x] != '#' {
				to_visit = append(to_visit, pathpoint{next, current.length + 1})
			}
		}
	}

	return cheats
}

func long_cheat(maze [][]byte, start point, to_end map[point]int, time_save int, in_bounds func(*point) bool) int {
	visited := make(map[point]int)

	to_visit := make([]pathpoint, 0, len(maze)*len(maze[0]))
	to_visit = append(to_visit, pathpoint{start, 0})
	to_visit_set := make(map[point]bool)

	cheats := 0

	cheat_start := to_end[start]

	for len(to_visit) > 0 {
		current := to_visit[0]
		to_visit = to_visit[1:]
		delete(to_visit_set, current.point)

		visited[current.point] = current.length

		cheat_end, found := to_end[current.point]
		if found && cheat_start-cheat_end-current.length >= time_save {
			cheats++
		}

		if current.length == 20 {
			continue
		}

		for d := range dirs {
			next := *current.point.add(&dirs[d])
			_, v := visited[next]
			if !to_visit_set[next] && !v && in_bounds(&next) {
				to_visit = append(to_visit, pathpoint{next, current.length + 1})
				to_visit_set[next] = true
			}
		}
	}

	return cheats
}

func make_in_bounds(max_x int, max_y int) func(*point) bool {
	return func(p *point) bool {
		switch {
		case p.x < 0:
		case p.x >= max_x:
		case p.y < 0:
		case p.y >= max_y:
		default:
			return true
		}

		return false
	}
}
