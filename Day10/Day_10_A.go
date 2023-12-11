package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Coord struct {
	x int
	y int
}

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)
	// sum := 0

	grid := [][]rune{}
	animalPos := Coord{}
	for sc.Scan() {
		xTile := []rune(sc.Text())

		xTile = slices.Insert(xTile, 0, '.')
		xTile = append(xTile, '.')

		grid = append(grid, xTile)
		i := strings.Index(sc.Text(), "S")
		if i >= 0 {
			animalPos = Coord{i + 1, len(grid)}
		}
	}

	grid = slices.Insert(grid, 0, []rune(strings.Repeat(".", len(grid[0]))))
	grid = append(grid, grid[0])

	fmt.Printf("animalPos: %v\n", animalPos)
	fmt.Printf(": %v\n", isValidTile(grid, animalPos, Coord{animalPos.x + 1, animalPos.y}))
	fmt.Printf(": %v\n", isValidTile(grid, animalPos, Coord{animalPos.x - 1, animalPos.y}))
	fmt.Printf(": %v\n", isValidTile(grid, animalPos, Coord{animalPos.x, animalPos.y + 1}))
	fmt.Printf(": %v\n", isValidTile(grid, animalPos, Coord{animalPos.x, animalPos.y - 1}))
	printGrid(grid)
	// return

	// weightMap := make([][]int, len(grid)-1)

	queue := []Coord{animalPos}
	visited := []Coord{}
	dis := 0
	// curPos := animalPos
	for len(queue) != 0 {
		curPos := pop(&queue)
		// fmt.Printf("curPos: %v\n", curPos)
		// printGridHigh(grid, curPos, maxDisX)
		// println("+++++++")

		nexCoord := Coord{curPos.x + 1, curPos.y}

		if isValidTile(grid, curPos, nexCoord) && !slices.Contains(visited, nexCoord) {
			push(&queue, nexCoord)
			visited = append(visited, nexCoord)
		}

		nexCoord = Coord{curPos.x - 1, curPos.y}
		if isValidTile(grid, curPos, nexCoord) && !slices.Contains(visited, nexCoord) {
			push(&queue, nexCoord)
			visited = append(visited, nexCoord)

		}

		nexCoord = Coord{curPos.x, curPos.y + 1}
		if isValidTile(grid, curPos, nexCoord) && !slices.Contains(visited, nexCoord) {
			push(&queue, nexCoord)
			visited = append(visited, nexCoord)

		}

		nexCoord = Coord{curPos.x, curPos.y - 1}
		if isValidTile(grid, curPos, nexCoord) && !slices.Contains(visited, nexCoord) {
			push(&queue, nexCoord)
			visited = append(visited, nexCoord)
		}
		dis++
	}
	fmt.Printf("maxDis: %v\n", dis/2)
	fmt.Printf("maxDis: %v\n", len(visited))
	// for i := 1; i < len(grid); i++ {
	// 	// x :=
	// 	weightMap[i-1] = make([]int, len(grid)-1)

	// 	for j := 1; j < len(grid[0]); j++ {

	// 	}
	// }

}

func pop(queue *[]Coord) Coord {
	top := (*queue)[0]
	*queue = slices.Delete(*queue, 0, 1)
	return top
}

func push(queue *[]Coord, val Coord) {
	*queue = append(*queue, val)
	// *queue = slices.Insert(*queue, 0, val)
}

func isValidTile(grid [][]rune, curPos, nextPos Coord) bool {
	cur := grid[curPos.y][curPos.x]
	next := grid[nextPos.y][nextPos.x]

	if next == '.' {
		return false
	}

	isRight := nextPos.x > curPos.x
	isTop := nextPos.y < curPos.y
	xEq := nextPos.x == curPos.x
	yEq := nextPos.y == curPos.y

	// x := nextPos.x - curPos.x
	// y := nextPos.y - curPos.y

	if cur == '-' || cur == 'S' {
		if yEq {
			if isRight {
				if strings.ContainsRune("-J7", next) {
					return true
				}
			} else {
				if strings.ContainsRune("-LF", next) {
					return true
				}
			}
		}
	}
	if cur == '|' || cur == 'S' {
		if xEq {
			if isTop {
				if strings.ContainsRune("|7F", next) {
					return true
				}
			} else {
				if strings.ContainsRune("|JL", next) {
					return true
				}
			}
		}
	}
	if cur == 'F' || cur == 'S' {

		if yEq && isRight && strings.ContainsRune("-J7", next) {
			return true
		}
		if xEq && !isTop && strings.ContainsRune("|JL", next) {
			return true
		}
	}
	if cur == '7' || cur == 'S' {

		if yEq && !isRight && strings.ContainsRune("-LF", next) {
			return true
		}
		if xEq && !isTop && strings.ContainsRune("|LJ", next) {
			return true
		}
	}
	if cur == 'L' || cur == 'S' {

		if yEq && isRight && strings.ContainsRune("-J7", next) {
			return true
		}
		if xEq && isTop && strings.ContainsRune("|F7", next) {
			return true
		}
	}
	if cur == 'J' || cur == 'S' {

		if yEq && !isRight && strings.ContainsRune("-LF", next) {
			return true
		}
		if xEq && isTop && strings.ContainsRune("|F7", next) {
			return true
		}

	}

	return false
}

func printGrid(g [][]rune) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			fmt.Printf("%c", g[i][j])
		}
		fmt.Println()
	}
}
func printGridHigh(g [][]rune, high Coord, x int) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if i == high.y && j == high.x {

				fmt.Printf("#")
			} else {
				fmt.Printf("%c", g[i][j])
			}
		}
		fmt.Println()
	}
}
