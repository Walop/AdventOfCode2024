package main

import (
	"AdventOfCode2024/util/timer"
	"AdventOfCode2024/util/vec2"
	"bufio"
	"fmt"
	"os"
)

type layout struct {
	robot vec2.Vec2
	walls map[vec2.Vec2]bool
	boxes map[vec2.Vec2]bool
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	warehouse := layout{
		walls: make(map[vec2.Vec2]bool),
		boxes: make(map[vec2.Vec2]bool),
	}

	y := 0

	for scanner.Scan() {
		row := scanner.Bytes()
		if len(row) == 0 {
			break
		}

		for x, b := range row {
			v := vec2.Vec2{X: x, Y: y}
			switch b {
			case '#':
				warehouse.walls[v] = true
			case 'O':
				warehouse.boxes[v] = true
			case '@':
				warehouse.robot = v
			}
		}
		y++
	}

	instructions := make([]byte, 0, 20040)

	for scanner.Scan() {
		row := scanner.Bytes()
		instructions = append(instructions, row...)
	}

	part1(&warehouse, instructions)

	part2(&warehouse, instructions)
}

func part1(initial_warehouse *layout, instructions []byte) {
	warehouse := layout{
		robot: initial_warehouse.robot,
		walls: copy_map(initial_warehouse.walls),
		boxes: copy_map(initial_warehouse.boxes),
	}

	for _, b := range instructions {
		switch b {
		case '<':
			move_robot(&warehouse, &vec2.Vec2{X: -1, Y: 0})
		case '^':
			move_robot(&warehouse, &vec2.Vec2{X: 0, Y: -1})
		case '>':
			move_robot(&warehouse, &vec2.Vec2{X: 1, Y: 0})
		default:
			move_robot(&warehouse, &vec2.Vec2{X: 0, Y: 1})
		}
	}

	sum := 0

	for b := range warehouse.boxes {
		sum += 100*b.Y + b.X
	}

	fmt.Println("Part 1: ", sum)
}

func move_robot(warehouse *layout, dir *vec2.Vec2) {
	move_to := warehouse.robot.Add(dir)
	for contains(warehouse.boxes, move_to) {
		move_to = move_to.Add(dir)
	}

	if contains(warehouse.walls, move_to) {
		return
	}

	warehouse.robot = warehouse.robot.Add(dir)

	warehouse.boxes[move_to] = true
	delete(warehouse.boxes, warehouse.robot)
}

func part2(initial_warehouse *layout, instructions []byte) {
	warehouse := layout{
		robot: initial_warehouse.robot,
		walls: copy_map(initial_warehouse.walls),
		boxes: copy_map(initial_warehouse.boxes),
	}

	warehouse.robot.X *= 2
	wide_walls := make(map[vec2.Vec2]bool)
	wx_max := 0
	wy_max := 0
	for w := range warehouse.walls {
		w.X *= 2
		wide_walls[w] = true
		w.X += 1
		wide_walls[w] = true
		if w.X > wx_max {
			wx_max = w.X
		}
		if w.Y > wy_max {
			wy_max = w.Y
		}
	}

	warehouse.walls = wide_walls

	wide_boxes := make(map[vec2.Vec2]bool)
	for b := range warehouse.boxes {
		b.X *= 2
		wide_boxes[b] = true
		b.X += 1
		wide_boxes[b] = false
	}

	warehouse.boxes = wide_boxes

	//print_warehouse(&warehouse, wx_max, wy_max)

	for _, b := range instructions {
		switch b {
		case '<':
			move_robot2(&warehouse, &vec2.Vec2{X: -1, Y: 0})
		case '^':
			move_robot2(&warehouse, &vec2.Vec2{X: 0, Y: -1})
		case '>':
			move_robot2(&warehouse, &vec2.Vec2{X: 1, Y: 0})
		default:
			move_robot2(&warehouse, &vec2.Vec2{X: 0, Y: 1})
		}
		//print_warehouse(&warehouse, wx_max, wy_max)
	}

	//print_warehouse(&warehouse, wx_max, wy_max)

	sum := 0

	for b, v := range warehouse.boxes {
		if v {
			sum += 100*b.Y + b.X
		}
	}

	fmt.Println("Part 2: ", sum)
}

func move_robot2(warehouse *layout, dir *vec2.Vec2) {
	boxes_to_move := make(map[vec2.Vec2]bool)

	move_to := warehouse.robot.Add(dir)

	if contains(warehouse.walls, move_to) {
		return
	}

	if !contains(warehouse.boxes, move_to) && !contains(warehouse.walls, move_to) {
		warehouse.robot = move_to
		return
	}

	if !move_boxes(warehouse, move_to, dir, boxes_to_move) {
		return
	}

	for b := range boxes_to_move {
		delete(warehouse.boxes, b)
	}

	for b, v := range boxes_to_move {
		next := b.Add(dir)
		warehouse.boxes[next] = v
	}

	warehouse.robot = move_to
}

func move_boxes(warehouse *layout, current vec2.Vec2, dir *vec2.Vec2, boxes_to_move map[vec2.Vec2]bool) bool {
	if contains(boxes_to_move, current) {
		return true
	}

	next := current.Add(dir)
	if contains(warehouse.walls, next) {
		return false
	}

	boxes_to_move[current] = warehouse.boxes[current]

	free := !contains(warehouse.boxes, next)

	move1_success := free || move_boxes(warehouse, next, dir, boxes_to_move)

	if !contains(warehouse.boxes, current) {
		fmt.Println(current)
		panic("Box error")
	}

	if warehouse.boxes[current] {
		current.X += 1
	} else {
		current.X -= 1
	}

	next = current.Add(dir)
	if contains(warehouse.walls, next) {
		return false
	}

	boxes_to_move[current] = warehouse.boxes[current]

	free = !contains(warehouse.boxes, next)

	move2_success := free || move_boxes(warehouse, next, dir, boxes_to_move)

	return move1_success && move2_success
}

func contains(m map[vec2.Vec2]bool, v vec2.Vec2) bool {
	_, ok := m[v]
	return ok
}

func copy_map(m map[vec2.Vec2]bool) map[vec2.Vec2]bool {
	m2 := make(map[vec2.Vec2]bool)
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

func print_warehouse(warehouse *layout, width int, height int) {
	for y := range height + 1 {
		for x := range width + 1 {
			p := vec2.Vec2{X: x, Y: y}
			if p == warehouse.robot {
				fmt.Print("@")
			} else if contains(warehouse.walls, p) {
				fmt.Print("#")
			} else if contains(warehouse.boxes, p) {
				if warehouse.boxes[p] {
					fmt.Print("[]")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
