package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	img := [][]rune{}
	for sc.Scan() {
		img = append(img, []rune(sc.Text()))
	}

	rows := []int{}
	cols := []int{}
	for i := 0; i < len(img); i++ {
		if !strings.ContainsRune(string(img[i]), '#') {
			rows = append(rows, i)
		}
		isEmpty := true
		for j := 0; j < len(img[0]); j++ {
			if img[j][i] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			cols = append(cols, i)
		}
	}
	newRows, newCols := make([]int, len(img)), make([]int, len(img))

	for i := 0; i < len(img); i++ {
		newRows[i] = 1000000
		newCols[i] = 1000000
	}

	galaxiesPos := []Coord{}
	for i, r := range img {
		for j, c := range r {
			if c == '#' {
				galaxiesPos = append(galaxiesPos, Coord{j, i})
				newCols[j] = 1
				newRows[i] = 1
			}

		}
	}

	for i := 1; i < len(img); i++ {
		newRows[i] += newRows[i-1]
		newCols[i] += newCols[i-1]
	}

	sum := 0
	for i := 0; i < len(galaxiesPos); i++ {
		for j := i + 1; j < len(galaxiesPos); j++ {
			sum += manhanttenDis(Coord{newCols[galaxiesPos[i].x], newRows[galaxiesPos[i].y]}, Coord{newCols[galaxiesPos[j].x], newRows[galaxiesPos[j].y]})
		}
	}
	fmt.Printf("sum: %v\n", sum)

}

func manhanttenDis(p1, p2 Coord) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x)) + math.Abs(float64(p1.y)-float64(p2.y)))
}
