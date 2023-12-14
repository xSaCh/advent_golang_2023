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

		forAll := true
	mappLoop:
		for _, v := range mapp {

			for i, j := mp, mp-1; i < len(v) && j >= 0; i, j = i+1, j-1 {
				if v[i] != v[j] {
					forAll = false
					break mappLoop
				}
			}
		}
		if forAll {
			return mp
		}
	}
	return -1
}

func calcHoriMirror(mapp []string) int {
	for mp := 1; mp < len(mapp); mp++ {

		forAll := true
		for i, j := mp, mp-1; i < len(mapp) && j >= 0; i, j = i+1, j-1 {
			if mapp[i] != mapp[j] {
				forAll = false
				break
			}
		}

		if forAll {
			return mp
		}
	}
	return -1
}
