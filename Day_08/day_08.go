package main

import (
	"AdventOfCode2024/util/vec2"
	"bufio"
	"fmt"
	"os"
)

type Layout struct {
	antennas map[rune][]vec2.Vec2
	height   int
	width    int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	layout := Layout{
		antennas: make(map[rune][]vec2.Vec2),
		height:   0,
		width:    1,
	}

	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			if i+1 > layout.width {
				layout.width = i + 1
			}
			if r != '.' {
				layout.antennas[r] = append(layout.antennas[r], vec2.Vec2{X: i, Y: layout.height})
			}
		}
		layout.height++
	}

	// fmt.Println(layout)

	Part1(&layout)
	Part2(&layout)
}

func Part1(layout *Layout) {
	antinodes := make(map[vec2.Vec2]bool)

	for _, v := range layout.antennas {
		for len(v) > 1 {
			head := v[0]
			v = v[1:]
			for _, pos := range v {
				dist := pos.Substract(head)
				antinode := head.Substract(dist)
				if CheckInBound(layout, antinode) {
					antinodes[antinode] = true
				}
				antinode = pos.Add(dist)
				if CheckInBound(layout, antinode) {
					antinodes[antinode] = true
				}
			}
		}
	}

	// PrintResult(layout, antinodes)
	fmt.Println("Part 1: ", len(antinodes))
}

func Part2(layout *Layout) {
	antinodes := make(map[vec2.Vec2]bool)

	for _, v := range layout.antennas {
		for len(v) > 1 {
			head := v[0]
			v = v[1:]
			antinodes[head] = true
			for _, pos := range v {
				dist := pos.Substract(head)
				antinode := head.Substract(dist)
				for CheckInBound(layout, antinode) {
					antinodes[antinode] = true
					antinode = antinode.Substract(dist)
				}
				antinode = head.Add(dist)
				for CheckInBound(layout, antinode) {
					antinodes[antinode] = true
					antinode = antinode.Add(dist)
				}
			}
		}
	}

	//PrintResult(layout, antinodes)
	fmt.Println("Part 2: ", len(antinodes))
}

func CheckInBound(layout *Layout, pos vec2.Vec2) bool {
	return pos.X >= 0 && pos.X < layout.width && pos.Y >= 0 && pos.Y < layout.height
}

func PrintResult(layout *Layout, antinodes map[vec2.Vec2]bool) {
	for y := range layout.height {
		for x := range layout.width {
			if antinodes[vec2.Vec2{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
