package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/image/bmp"
)

var (
	width  = 101
	height = 103
)

type result struct {
	variance float64
	time     int
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	robots := make([][]int, 0, 500)

	robot := make([]int, 0, 4)
	buffer := make([]string, 0, 3)

	for scanner.Scan() {
		char := scanner.Text()

		if char != "-" && (char < "0" || char > "9") {
			if len(buffer) > 0 {
				num, _ := strconv.Atoi(strings.Join(buffer, ""))
				robot = append(slices.Clone(robot), num)
				buffer = buffer[:0]
			}
			if char == "\r" || char == "\n" {
				if len(robot) > 0 {
					robots = append(robots, robot)
					robot = robot[:0]
				}
			}
			continue
		}
		buffer = append(buffer, char)
	}

	final_positions := make(map[[2]int]int)

	for _, robot := range robots {
		final_pos := find_final_position(robot, 100)
		final_positions[final_pos] += 1
	}

	quadrants := []int{0, 0, 0, 0}

	for y := range height {
		for x := range width {
			count := final_positions[[2]int{x, y}]
			if count > 0 {
				if x < width/2 && y < height/2 {
					quadrants[0] += count
				} else if x > width/2 && y < height/2 {
					quadrants[1] += count
				} else if x < width/2 && y > height/2 {
					quadrants[2] += count
				} else if x > width/2 && y > height/2 {
					quadrants[3] += count
				}
				// fmt.Print(count)
			} else {
				// fmt.Print(".")
			}
		}
		// fmt.Print("\n")
	}

	// fmt.Println(final_positions)
	// fmt.Println(quadrants)

	fmt.Println("Part 1: ", quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3])

	ch := make(chan result)

	for i := range 10_000 {
		go calculate_variance(robots, i, ch)
	}

	min_variance := float64(1_000_000)
	min_i := 0

	for range 10_000 {
		res := <-ch
		if res.variance < min_variance {
			min_variance = res.variance
			min_i = res.time
		}
	}

	tree_positions := make(map[[2]int]bool)

	for _, robot := range robots {
		final_pos := find_final_position(robot, min_i)
		tree_positions[final_pos] = true
	}
	fmt.Println("Part 2: ", min_i)
	img := create_image(tree_positions)
	write_bmp(min_i, img)
}

func calculate_variance(robots [][]int, time int, ch chan result) {
	final_positions := make(map[[2]int]int)

	for _, robot := range robots {
		final_pos := find_final_position(robot, time)
		final_positions[final_pos] += 1
	}

	x_sum := 0
	y_sum := 0

	for pos := range final_positions {
		x_sum += pos[0]
		y_sum += pos[1]
	}

	average_x := float64(x_sum) / float64(len(final_positions))
	average_y := float64(y_sum) / float64(len(final_positions))

	variance := float64(0)

	for pos := range final_positions {
		variance += math.Abs(average_x-float64(pos[0])) + math.Abs(average_y-float64(pos[1]))
	}

	ch <- result{variance, time}
}

func find_final_position(robot []int, time int) [2]int {
	final_x := (robot[0] + robot[2]*time) % width
	if final_x < 0 {
		final_x += width
	}
	final_y := (robot[1] + robot[3]*time) % height
	if final_y < 0 {
		final_y += height
	}

	return [2]int{final_x, final_y}
}

func create_image(positions map[[2]int]bool) *image.Gray {
	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := range height {
		for x := range width {
			idx := y*img.Stride + x
			col := uint8(0)
			if positions[[2]int{x, y}] {
				col = uint8(255)
			}
			img.Pix[idx] = col
		}
	}

	return img
}

func write_bmp(i int, img *image.Gray) {
	out, err := os.Create(strconv.Itoa(i) + ".bmp")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	bmp.Encode(out, img)
}
