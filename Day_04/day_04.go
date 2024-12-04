package main

import (
	"bufio"
	"fmt"
	"os"
)

type V2 struct {
	x int
	y int
}

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

	dirs := []V2{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
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

func Check(rows [][]byte, startx int, starty int, dir V2) bool {
	height := len(rows)
	width := len(rows[0])

	xmas := []byte("XMAS")

	xend := startx + dir.x*len(xmas)
	yend := starty + dir.y*len(xmas)

	if xend < -1 || xend > width || yend < -1 || yend > height {
		return false
	}

	x := startx
	y := starty

	for i := range len(xmas) {
		if rows[y][x] != xmas[i] {
			return false
		}
		x += dir.x
		y += dir.y
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

	dirs := []V2{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	ms := [4]byte{}

	for i, d := range dirs {
		x := startx + d.x
		y := starty + d.y
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
