package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	maps := [][]string{}

	tmp := []string{}
	for sc.Scan() {
		s := sc.Text()
		if s == "" {
			maps = append(maps, tmp)
			tmp = []string{}
		} else {
			tmp = append(tmp, s)
		}
	}
	maps = append(maps, tmp)

	v, h := 0, 0
	for _, m := range maps {
		s := calcVertMirror(m)
		if s != -1 {
			v += s
		} else {
			h += calcHoriMirror(m)
		}
	}
	fmt.Printf("ans: %v\n", (h*100)+v)

}

func calcVertMirror(mapp []string) int {
	for mp := 1; mp < len(mapp[0]); mp++ {

		count := 0
	mappLoop:
		for _, v := range mapp {

			for i, j := mp, mp-1; i < len(v) && j >= 0; i, j = i+1, j-1 {
				if v[i] != v[j] {
					count++
				}
				if count > 1 {
					break mappLoop
				}
			}
		}
		if count == 1 {
			return mp
		}
	}
	return -1
}

func calcHoriMirror(mapp []string) int {
	for mp := 1; mp < len(mapp); mp++ {

		smutCnt := 0
		for i, j := mp, mp-1; i < len(mapp) && j >= 0; i, j = i+1, j-1 {
			smutCnt += getDiffCount(mapp[i], mapp[j])
			if smutCnt > 1 {
				break
			}
		}

		if smutCnt == 1 {
			return mp
		}
	}
	return -1
}

func getDiffCount(a, b string) int {
	x, y := min(a, b), max(a, b)
	count := 0

	for i := range x {
		if x[i] != y[i] {
			count++
		}
	}
	return count
}
