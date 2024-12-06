package main

import (
	"AdventOfCode2024/util/vec2"
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rows := make([][]byte, 0, 140)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		rows = append(rows, line)
	}

	Part1(rows)
	Part2(rows)
}

func Part1(rows [][]byte) {
	height := len(rows)
	width := len(rows[0])

	dirs := []vec2.Vec2{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	}

	count := 0

	for y := range height {
		for x := range width {
			for _, d := range dirs {
				if Check(rows, x, y, d) {
					count++
				}
			}
		}
	}

	fmt.Println("Part 1: ", count)
}

func Check(rows [][]byte, startx int, starty int, dir vec2.Vec2) bool {
	height := len(rows)
	width := len(rows[0])

	xmas := []byte("XMAS")

	xend := startx + dir.X*len(xmas)
	yend := starty + dir.Y*len(xmas)

	if xend < -1 || xend > width || yend < -1 || yend > height {
		return false
	}

	x := startx
	y := starty

	for i := range len(xmas) {
		if rows[y][x] != xmas[i] {
			return false
		}
		x += dir.X
		y += dir.Y
	}

	return true
}

func Part2(rows [][]byte) {
	height := len(rows)
	width := len(rows[0])
	count := 0

	for y := range height {
		for x := range width {
			if CheckMas(rows, x, y) {
				count++
			}
		}
	}

	fmt.Println("Part 2: ", count)
}

func CheckMas(rows [][]byte, startx int, starty int) bool {
	height := len(rows)
	width := len(rows[0])

	mas := []byte("MAS")

	if startx-1 < 0 || startx+1 >= width || starty-1 < 0 || starty+1 >= height {
		return false
	}

	if rows[starty][startx] != mas[1] {
		return false
	}

	dirs := []vec2.Vec2{
		{X: -1, Y: -1},
		{X: -1, Y: 1},
		{X: 1, Y: -1},
		{X: 1, Y: 1},
	}

	ms := [4]byte{}

	for i, d := range dirs {
		x := startx + d.X
		y := starty + d.Y
		if rows[y][x] != mas[0] && rows[y][x] != mas[2] {
			return false
		}
		ms[i] = rows[y][x]
	}

	if ms[0] == ms[3] || ms[1] == ms[2] {
		return false
	}

	return true
}
