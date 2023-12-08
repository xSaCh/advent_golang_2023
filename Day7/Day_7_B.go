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
		// hand = handleJoker(hand)
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
	// fmt.Printf("ranks: %v\n", ranks)
	// for _, r := range ranks {
	// 	fmt.Printf("%s %d\n", r.hand, r.bid)
	// }

	sum := 0
	for i, r := range ranks {
		sum += (i + 1) * r.bid
	}
	fmt.Printf("sum: %v\n", sum)
}

func getCardStrength(card rune) int {
	switch card {
	case 'J':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
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
	hand = handleJoker(hand)
	fre := make([]int, 14)
	freq := map[rune]int{}

	maxR := 0
	for _, c := range hand {
		fre[getCardStrength(c)] += 1
		freq[c] += 1
		if fre[getCardStrength(c)] > maxR {
			maxR = fre[getCardStrength(c)]
		}
	}

	// for i := 0; i < len(fre); i++ {
	// 	if fre[i] == maxR {
	// 		fre[i] += fre[0]
	// 		break
	// 	}
	// }
	// fre = fre[1:]
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

func handleJoker(hand string) string {
	freq := map[rune]int{}
	maxF := '0'
	for _, c := range hand {
		// if c == 'J' {
		// 	noJ++
		// }
		freq[c] = freq[c] + 1
		if c != 'J' && freq[maxF] < freq[c] {
			maxF = c
		}
	}
	h := strings.ReplaceAll(hand, "J", string(maxF))
	// fmt.Printf("maxF: %v\n", maxF)
	// fmt.Printf("freq: %v\n", freq)
	// fmt.Printf("h: %v\n", h)
	// cards := []rune(hand)
	// sort.Slice(cards, func(i, j int) bool {
	// 	return getCardStrength(cards[i]) > getCardStrength(cards[j])
	// })
	// fmt.Printf("string(cards): %v\n", string(cards))
	return h

}
