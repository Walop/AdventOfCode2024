package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	disk_map := make([]uint16, 0, 20_000)

	file_index := uint16(0)

	for scanner.Scan() {
		len, _ := strconv.Atoi(scanner.Text())

		file := slices.Repeat([]uint16{file_index}, len)

		disk_map = slices.Concat(disk_map, file)

		file_index++

		scanner.Scan()
		len, _ = strconv.Atoi(scanner.Text())

		empty := slices.Repeat([]uint16{0xFFFF}, len)
		disk_map = slices.Concat(disk_map, empty)
	}

	Part1(slices.Clone(disk_map))
	Part2(slices.Clone(disk_map))
}

func Part1(disk_map []uint16) {
	//fmt.Println(disk_map)
	i := 0
	j := len(disk_map) - 1
	for {
		for disk_map[i] != 0xFFFF {
			i++
		}
		for disk_map[j] == 0xFFFF {
			j--
		}
		if i >= j {
			break
		}
		disk_map[i] = disk_map[j]
		disk_map[j] = 0xFFFF
	}

	sum := uint64(0)

	i = 0
	for disk_map[i] != 0xFFFF {
		sum += uint64(disk_map[i]) * uint64(i)
		i++
	}

	fmt.Println("Part 1: ", sum)
}

func Part2(disk_map []uint16) {
	i := len(disk_map) - 1
	first_empty := 0
	for disk_map[first_empty] != 0xFFFF {
		first_empty++
	}

	for disk_map[i] == 0xFFFF {
		i--
	}
	prev_id := disk_map[i]

	for i > first_empty {
		for disk_map[i] == 0xFFFF {
			i--
		}
		id := disk_map[i]
		for id > prev_id {
			for disk_map[i] == id {
				i--
			}
			id = disk_map[i]
		}
		prev_id = id
		file_len := 0
		// fmt.Println(id)
		for disk_map[i] == id {
			i--
			file_len++
		}
		j := first_empty
		empty_len := 0
		for j <= i && empty_len < file_len {
			if disk_map[j] != 0xFFFF {
				empty_len = 0
			} else {
				empty_len++
			}
			j++
		}

		if empty_len >= file_len {
			//fmt.Println(id, i+1, file_len, j-empty_len)
			for k := range file_len {
				disk_map[j-empty_len+k] = disk_map[i+1+k]
				disk_map[i+1+k] = 0xFFFF
			}
			if j-empty_len == first_empty {
				for disk_map[first_empty] != 0xFFFF {
					first_empty++
				}
			}
		}
	}

	sum := uint64(0)

	for i, id := range disk_map {
		if id != 0xFFFF {
			sum += uint64(id) * uint64(i)
		}
	}

	fmt.Println("Part 2: ", sum)
}
