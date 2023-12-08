package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	sd := strings.Split(strings.Split(readLine(sc), ": ")[1], " ")

	seeds := []int{}
	n := 0

	for i, s := range sd {
		if i%2 == 0 {
			nn, _ := strconv.Atoi(s)
			n = nn
		} else {
			nn, _ := strconv.Atoi(s)
			for j := n; j < n+nn; j++ {
				seeds = append(seeds, j)
			}
		}
	}
	sc.Scan()

	almanc := [][][]int{}
	for i := 0; i < 7; i++ {
		almanc = append(almanc, getMap(sc))
	}

	curList := seeds
	stsl := [][]int{}
	for i := 6; i >= 0; i-- {
		// stsl := getMap(sc)
		stsl = almanc[i]
		// fmt.Printf("stsl: %v\n", stsl)
		curList = getNextList(stsl, curList)
		fmt.Printf("curList: %v\n", curList)
	}

	s := getFList([][]int{{0, 79, 14}, {0, 55, 13}}, curList)
	fmt.Printf("s: %v\n", s)

	// ls := map[int]int{}
	// for _, v := range curList {
	// 	for i := 0; i < len(seeds)-1; i += 2 {
	// 		if v >= seeds[i] && v <= seeds[i]+seeds[i+1] {
	// 			ls[(seeds[i+1]+seeds[i])-v] = ls[(seeds[i+1]+seeds[i])-v] + 1
	// 			break
	// 		}
	// 	}
	// }
	// fmt.Printf("ls: %v\n", ls)
	// sort.Slice(seeds, func(i, j int) bool {
	// 	return seeds[i] < seeds[j]
	// })
	// fmt.Println(seeds[0])
}

func getFList(curMap [][]int, prevList []int) []int {
	nextListMap := map[int]int{}
	for _, mr := range curMap {
		for _, e := range prevList {
			if e >= mr[1] && e <= mr[1]+mr[2] {
				nextListMap[e] = e
			}
		}
	}
	nextList := []int{}
	// for _, e := range prevList {
	// 	_, ok := nextListMap[e]
	// 	if !ok {
	// 		nextList = append(nextList, e)
	// 	}
	// }
	for _, v := range nextListMap {
		nextList = append(nextList, v)
	}
	return nextList
}
func getNextList(curMap [][]int, prevList []int) []int {
	nextListMap := map[int]int{}
	for _, mr := range curMap {
		for _, e := range prevList {
			if e >= mr[1] && e <= mr[1]+mr[2] {
				nextListMap[e] = e - mr[1] + mr[0]
			}
		}
	}
	nextList := []int{}
	for _, e := range prevList {
		_, ok := nextListMap[e]
		if !ok {
			nextList = append(nextList, e)
		}
	}
	for _, v := range nextListMap {
		nextList = append(nextList, v)
	}
	return nextList
}

func getMap(sc *bufio.Scanner) [][]int {
	mp := [][]int{}
	sc.Scan()
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		// fmt.Println(sc.Text())
		ln := []int{}
		for _, n := range strings.Split(sc.Text(), " ") {
			in, _ := strconv.Atoi(n)
			ln = append(ln, in)
		}
		mp = append(mp, ln)
	}
	return mp
}
