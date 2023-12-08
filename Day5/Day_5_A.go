package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for _, s := range sd {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}
	sc.Scan()
	for i := 0; i < 7; i++ {
		stsl := getMap(sc)
		// fmt.Printf("stsl: %v\n", stsl)
		seeds = getNextList(stsl, seeds)
		// fmt.Printf("seeds: %v\n", seeds)
	}
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i] < seeds[j]
	})
	fmt.Println(seeds[0])
}

func getNextList(curMap [][]int, prevList []int) []int {
	nextListMap := map[int]int{}
	for _, mr := range curMap {
		for _, e := range prevList {
			if e >= mr[1] && e <= mr[1]+mr[2] {
				nextListMap[e] = e - mr[1] + mr[0]
			}
			// nextListMap[e] = e
			// nextList = append(nextList, e-mr[1]+mr[0])
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
