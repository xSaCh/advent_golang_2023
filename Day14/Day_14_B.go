package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	dish := [][]rune{}

	for sc.Scan() {
		dish = append(dish, []rune(sc.Text()))
	}
	// printD(dish)
	// prev := copyGrid(dish)

	cycle := func(g [][]rune) {
		tilt(g) // North

		rotateClock(g)
		tilt(g) // West

		rotateClock(g)
		tilt(g) // South

		rotateClock(g)
		tilt(g) // East
		rotateClock(g)
	}

	cycleRepeatAfter, st := floyd(cycle, dish)
	left := 1e9 % cycleRepeatAfter

	for k := 0; k < cycleRepeatAfter+left; k++ {
		cycle(st)
	}
	sum := 0
	for i, v := range st {
		sum += strings.Count(string(v), "O") * (len(st) - i)

	}
	fmt.Printf("sum: %v\n", sum)
}

func copyGrid(grid [][]rune) [][]rune {
	prev := make([][]rune, len(grid))
	for i := range grid {
		prev[i] = make([]rune, len(grid[i]))
		copy(prev[i], grid[i])
	}
	return prev
}

func tilt(dish [][]rune) {
	for i := 0; i < len(dish); i++ {
		for j := 0; j < len(dish[i]); j++ {
			if dish[j][i] != 'O' {
				continue
			}

			newI := getEmptySpot(dish, j, i)
			if newI == -1 {
				continue
			}
			dish[newI][i], dish[j][i] = dish[j][i], dish[newI][i]
		}

	}
}

func getEmptySpot(d [][]rune, minH, cln int) int {
	lastEm := 0
	for i := minH - 1; i >= 0; i-- {
		if d[i][cln] == '.' {
			lastEm = i
		}
		if d[i][cln] == '#' || d[i][cln] == 'O' {
			return i + 1
		}
	}
	return lastEm
}

func rotateClock(a [][]rune) {
	N := len(a)
	for i := 0; i < N/2; i++ {
		for j := i; j < N-i-1; j++ {

			temp := a[i][j]
			a[i][j] = a[N-1-j][i]
			a[N-1-j][i] = a[N-1-i][N-1-j]
			a[N-1-i][N-1-j] = a[j][N-1-i]
			a[j][N-1-i] = temp
		}
	}
}

func printD(a [][]rune) {
	for _, v := range a {
		fmt.Println(string(v))
	}
}

func isEqualGrid(a, b [][]rune) bool {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func floyd(f func([][]rune), x0 [][]rune) (int, [][]rune) {
	//"""Floyd's cycle detection algorithm."""
	// Main phase of algorithm: finding a repetition x_i = x_2i.
	// The hare moves twice as quickly as the tortoise and
	// the distance between them increases by 1 at each step.
	// Eventually they will both be inside the cycle and then,
	// at some point, the distance between them will be
	// divisible by the period λ.
	tortoise, hare := copyGrid(x0), copyGrid(x0)
	f(tortoise) // f(x0) is the element/node next to x0.
	f(hare)
	f(hare)
	for !isEqualGrid(tortoise, hare) {
		f(tortoise)
		f(hare)
		f(hare)
	}
	// At this point the tortoise position, ν, which is also equal
	// to the distance between hare and tortoise, is divisible by
	// the period λ. So hare moving in cycle one step at a time,
	// and tortoise (reset to x0) moving towards the cycle, will
	// intersect at the beginning of the cycle. Because the
	// distance between them is constant at 2ν, a multiple of λ,
	// they will agree as soon as the tortoise reaches index μ.

	// Find the position μ of first repetition.
	//mu = 0
	//tortoise = x0
	// for (isEqualGrid(tortoise ,hare)):
	//     tortoise = f(tortoise)
	//     hare = f(hare)   // Hare and tortoise move at same speed
	//     mu += 1

	// // Find the length of the shortest cycle starting from x_μ
	// // The hare moves one step at a time while tortoise is still.
	// // lam is incremented until λ is found.
	lam := 1
	f(hare)
	for !isEqualGrid(tortoise, hare) {
		f(hare)
		lam += 1
	}

	return lam, hare
}
