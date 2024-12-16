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

type pathpoint struct {
	point
	dir    int
	points int
	path   []point
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

	// for y := range len(maze) {
	// 	for x := range len(maze[0]) {
	// 		fmt.Print(string(maze[y][x]))
	// 	}
	// 	fmt.Print("\n")
	// }

	to_visit := make([]pathpoint, 0, 11000)
	visited := make([]pathpoint, 0, 11000)

	to_visit = append(to_visit, pathpoint{
		dir:    1,
		points: 0,
		path:   []point{},
		point: point{
			x: start.x,
			y: start.y,
		},
	})

	dirs := [4]point{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	best := 1000000000

	seen := make([]point, 0, 5000)

	// visited_count := 0

	for len(to_visit) > 0 {
		i := find_smallest(to_visit)
		current := to_visit[i]

		to_visit = remove_item(to_visit, i)

		if has_visited(visited, current.point, current.dir) > -1 {
			continue
		}

		next := point{current.x + dirs[current.dir].x, current.y + dirs[current.dir].y}

		if next.x == end.x && next.y == end.y {
			if best >= current.points+1 {
				best = current.points + 1
				seen = append(seen, current.path...)
			}
			continue
		}

		next_path := make([]point, 0, len(current.path)+1)
		next_path = append(next_path, current.path...)
		next_path = append(next_path, next)

		vi := to_be_visited(to_visit, next, current.dir)
		if vi != -1 {
			if to_visit[vi].points >= current.points+1 {
				if to_visit[vi].points == current.points+1 {
					to_visit[vi].path = append(to_visit[vi].path, next_path...)
				} else {
					to_visit[vi].path = next_path
					to_visit[vi].points = current.points + 1
				}
			}
		} else {
			if maze[next.y][next.x] == '.' {
				to_visit = append(to_visit, pathpoint{
					point: point{
						x: next.x,
						y: next.y,
					},
					dir:    current.dir,
					points: current.points + 1,
					path:   next_path,
				})
			}
		}

		new_dir := (current.dir + len(dirs) + 1) % len(dirs)

		vi = to_be_visited(to_visit, current.point, new_dir)
		if vi != -1 {
			if to_visit[vi].points >= current.points+1000 {
				if to_visit[vi].points == current.points+1000 {
					to_visit[vi].path = append(to_visit[vi].path, current.path...)
				} else {
					to_visit[vi].path = current.path
					to_visit[vi].points = current.points + 1000
				}
			}
		} else {
			to_visit = append(to_visit, pathpoint{
				point: point{
					x: current.x,
					y: current.y,
				},
				dir:    new_dir,
				points: current.points + 1000,
				path:   current.path,
			})
		}

		new_dir = (current.dir + len(dirs) - 1) % len(dirs)

		vi = to_be_visited(to_visit, current.point, new_dir)
		if vi != -1 {
			if to_visit[vi].points >= current.points+1000 {
				if to_visit[vi].points == current.points+1000 {
					to_visit[vi].path = append(to_visit[vi].path, current.path...)
				} else {
					to_visit[vi].path = current.path
					to_visit[vi].points = current.points + 1000
				}
			}
		} else {
			to_visit = append(to_visit, pathpoint{
				point: point{
					x: current.x,
					y: current.y,
				},
				dir:    new_dir,
				points: current.points + 1000,
				path:   current.path,
			})
		}

		visited = append(visited, current)

		// visited_count++
		// if visited_count%100 == 0 {
		// 	fmt.Println(visited_count)
		// }
	}

	fmt.Println("Best ", best)
	uq := make(map[point]struct{})
	for _, s := range seen {
		uq[s] = struct{}{}
	}

	// for y := range len(maze) {
	// 	for x := range len(maze[0]) {
	// 		if point_visited(visited, point{x, y}) {
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(string(maze[y][x]))
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	fmt.Println("Seen unique ", len(uq)+2)
}

func find_smallest(to_visit []pathpoint) int {
	min_i := 0
	sm := 1_000_000
	for i := range to_visit {
		if to_visit[i].points < sm {
			min_i = i
			sm = to_visit[i].points
		}
	}

	return min_i
}

func remove_item(to_visit []pathpoint, i int) []pathpoint {
	start := to_visit[:i]
	end := to_visit[i+1:]
	return append(start, end...)
}

func to_be_visited(to_visit []pathpoint, p point, dir int) int {
	for i, v := range to_visit {
		if v.x == p.x && v.y == p.y && v.dir == dir {
			return i
		}
	}
	return -1
}

func has_visited(visited []pathpoint, p point, dir int) int {
	for i, v := range visited {
		if v.x == p.x && v.y == p.y && v.dir == dir {
			return i
		}
	}
	return -1
}

func point_visited(visited []pathpoint, p point) bool {
	for _, v := range visited {
		if v.x == p.x && v.y == p.y {
			return true
		}
	}
	return false
}
