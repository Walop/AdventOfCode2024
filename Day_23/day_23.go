package main

import (
	"AdventOfCode2024/util/timer"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type triplet struct {
	first  string
	second string
	third  string
}

func main() {
	defer timer.Timer("Main")()

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	connections := map[string][]string{}

	for scanner.Scan() {
		row := scanner.Text()
		split := strings.Split(row, "-")

		if _, exists := connections[split[0]]; !exists {
			connections[split[0]] = make([]string, 0, 13)
		}
		if _, exists := connections[split[1]]; !exists {
			connections[split[1]] = make([]string, 0, 13)
		}
		connections[split[0]] = append(connections[split[0]], split[1])
		connections[split[1]] = append(connections[split[1]], split[0])
	}

	//fmt.Println(connections)

	part1(connections)
	part2(connections)
}

func part1(connections map[string][]string) {
	triplets := make(map[triplet]struct{}, 50)

	for k := range connections {
		if k[0] == 't' {
			for _, v := range find_triplets(connections, k) {
				triplets[v] = struct{}{}
			}
		}
	}

	fmt.Println("Part 1: ", len(triplets))
}

func find_triplets(connections map[string][]string, start string) []triplet {
	triplets := make([]triplet, 0, 50)

	for _, n := range connections[start] {
		for _, n2 := range connections[n] {
			if slices.Contains(connections[n2], start) {
				path := []string{start, n, n2}
				slices.Sort(path)
				triplets = append(triplets, triplet{path[0], path[1], path[2]})
			}
		}
	}
	return triplets
}

func part2(connections map[string][]string) {
	triplets := make(map[triplet]struct{}, len(connections))
	for node := range connections {
		for _, v := range find_triplets(connections, node) {
			triplets[v] = struct{}{}
		}
	}

	most_common := ""
	most_common_count := 0
	start_nodes := make(map[string][]triplet, len(connections))
	for triplet := range triplets {
		tr := append(start_nodes[triplet.first], triplet)
		start_nodes[triplet.first] = tr
		if len(tr) > most_common_count {
			most_common_count = len(tr)
			most_common = triplet.first
		}
	}

	passwdparts := map[string]struct{}{}

	passwdparts[most_common] = struct{}{}

	for _, tr := range start_nodes[most_common] {
		passwdparts[tr.second] = struct{}{}
		passwdparts[tr.third] = struct{}{}
	}

	password := get_keys(passwdparts)
	slices.Sort(password)
	fmt.Println("Part 2: ", strings.Join(password, ","))
}

func get_keys(m map[string]struct{}) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
