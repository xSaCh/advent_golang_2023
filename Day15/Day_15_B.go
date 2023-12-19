package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	seqs := []string{}

	for sc.Scan() {
		str := sc.Text()
		seqs = strings.Split(str, ",")
	}

	hashmap := [256]*orderedmap.OrderedMap[string, int]{}
	isInit := [256]bool{false}
	// sum := 0
	for _, s := range seqs {
		if s[len(s)-1] == '-' {
			h := hash(s[:len(s)-1])

			if isInit[h] {
				hashmap[h].Delete(s[:len(s)-1])
			}
		} else {
			seqS := strings.Split(s, "=")
			h := hash(seqS[0])
			n, _ := strconv.Atoi(seqS[1])

			if !isInit[h] {
				isInit[h] = true
				hashmap[h] = orderedmap.NewOrderedMap[string, int]()
			}
			hashmap[h].Set(seqS[0], n)
		}
	}

	sum := 0
	for i, l := range hashmap {
		if isInit[i] && l.Len() > 0 {
			c := 1
			for el := l.Front(); el != nil; el = el.Next() {
				a := (i + 1) * el.Value * c
				sum += a
				c++
			}

		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func hash(str string) int {
	n := 0
	for _, c := range str {
		n += int(c)
		n *= 17
		n %= 256
	}
	return n
}
