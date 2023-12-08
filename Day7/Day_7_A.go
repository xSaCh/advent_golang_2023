package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Round struct {
	hand string
	bid  int
	rank int
}

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)

	ranks := []Round{}
	for sc.Scan() {
		s := strings.Split(sc.Text(), " ")
		hand := s[0]
		bid, _ := strconv.Atoi(s[1])

		rank := getTypeRank(hand)
		ranks = append(ranks, Round{hand, bid, rank})
	}
	sort.Slice(ranks, func(i, j int) bool {
		if ranks[i].rank == ranks[j].rank {
			return !isHandGreater(ranks[i].hand, ranks[j].hand)
		}
		return ranks[i].rank < ranks[j].rank
	})

	sum := 0
	for i, r := range ranks {
		sum += (i + 1) * r.bid
	}
	fmt.Printf("sum: %v\n", sum)
}

func getCardStrength(card rune) int {
	switch card {
	case '2':
		return 1
	case '3':
		return 2
	case '4':
		return 3
	case '5':
		return 4
	case '6':
		return 5
	case '7':
		return 6
	case '8':
		return 7
	case '9':
		return 8
	case 'T':
		return 9
	case 'J':
		return 10
	case 'Q':
		return 11
	case 'K':
		return 12
	case 'A':
		return 13
	}
	return 0
}

func getTypeRank(hand string) int {
	fre := make([]int, 14)

	for _, c := range hand {
		fre[getCardStrength(c)] += 1
	}
	sort.Slice(fre, func(i, j int) bool {
		return fre[i] > fre[j]
	})
	// fmt.Printf("freq: %v\n", fre)

	if fre[0] == 5 {
		return 7
	}
	if fre[0] == 4 {
		return 6
	}
	if fre[0] == 3 && fre[1] == 2 {
		return 5
	}
	if fre[0] == 3 && fre[1] == 1 {
		return 4
	}
	if fre[0] == 2 && fre[1] == 2 {
		return 3
	}
	if fre[0] == 2 && fre[1] == 1 {
		return 2
	}

	return 1

}

func isHandGreater(c1 string, c2 string) bool {
	for i := 0; i < len(c1); i++ {
		if c1[i] == c2[i] {
			continue
		}
		if getCardStrength(rune(c1[i])) > getCardStrength(rune(c2[i])) {
			return true
		} else {
			return false
		}

	}
	return false
}
