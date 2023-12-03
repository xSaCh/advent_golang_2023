package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, _ := os.Open("./day3.txt")
	defer f.Close()
	sc := bufio.NewScanner(f)
	scheme := [][]rune{}
	sum := 0
	for sc.Scan() {
		scheme = append(scheme, []rune(sc.Text()))
	}

	for i := 0; i < len(scheme); i++ {
		for j := 0; j < len(scheme[i]); j++ {
			if scheme[i][j] != '.' && !unicode.IsDigit(scheme[i][j]) {
				sum += getNums(scheme, i, j)
			}
		}

	}

	println(sum)
}

func getNums(sh [][]rune, i, j int) int {

	nums := make([]int, 0)
	if i != 0 {
		nums = append(nums[:], getHoriNums(sh, i-1, j)[:]...)
	}
	nums = append(nums[:], getHoriNums(sh, i, j)[:]...)
	if i+1 < len(sh) {
		nums = append(nums[:], getHoriNums(sh, i+1, j)[:]...)
	}

	if len(nums) == 2 {
		return nums[0] * nums[1]
	}
	return 0
}

func getHoriNums(sh [][]rune, i int, j int) []int {
	ln, rn := "", ""

	for k := j + 1; k < len(sh[i]); k++ {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		rn += string(sh[i][k])
	}
	for k := j - 1; k >= 0; k-- {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		ln = string(sh[i][k]) + ln
	}

	if unicode.IsDigit(sh[i][j]) {
		n, _ := strconv.Atoi(ln + string(sh[i][j]) + rn)
		return []int{n}
	}
	num := []int{}
	lnn, _ := strconv.Atoi(ln)
	rnn, _ := strconv.Atoi(rn)
	if lnn != 0 {
		num = append(num, lnn)
	}
	if rnn != 0 {
		num = append(num, rnn)
	}
	return num
}
