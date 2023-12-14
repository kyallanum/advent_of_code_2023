package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func slicesMap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func mapsClear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	totalPaths := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		pipes, hints, _ := strings.Cut(currentLine, " ")
		pipesTimes5, hintsTimes5 := "", ""

		for i := 0; i < 5; i++ {
			pipesTimes5, hintsTimes5 = pipesTimes5+pipes+"?", hintsTimes5+hints+","
		}

		pipes, hints = strings.TrimSuffix(pipesTimes5, "?"), strings.TrimSuffix(hintsTimes5, ",")

		bytePipes := []byte(pipes)
		sliceMap := slicesMap(strings.Split(hints, ","), atoi)
		currentPaths := countPossible(bytePipes, sliceMap)
		totalPaths += currentPaths
	}
	fmt.Println(totalPaths)
}

func countPossible(s []byte, c []int) int {
	position := 0

	cstates := map[[4]int]int{{0, 0, 0, 0}: 1}
	nstates := map[[4]int]int{}

	for len(cstates) > 0 {
		for state, num := range cstates {
			si, ci, cc, expdot := state[0], state[1], state[2], state[3]
			if si == len(s) {
				if ci == len(c) {
					position += num
				}
				continue
			}
			switch {
			case (s[si] == '#' || s[si] == '?') && ci < len(c) && expdot == 0:
				if s[si] == '?' && cc == 0 {
					nstates[[4]int{si + 1, ci, cc, expdot}] += num
				}
				cc++
				if cc == c[ci] {
					ci++
					cc = 0
					expdot = 1
				}
				nstates[[4]int{si + 1, ci, cc, expdot}] += num
			case (s[si] == '.' || s[si] == '?') && cc == 0:
				expdot = 0
				nstates[[4]int{si + 1, ci, cc, expdot}] += num
			}
		}
		cstates, nstates = nstates, cstates
		mapsClear(nstates)
	}
	return position
}
