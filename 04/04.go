package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var count1, count2 int
	for n := 136760; n <= 595730; n++ {
		if isValidV1(n) {
			count1++
		}
		if isValidV2(n) {
			count2++
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
}

func isValidV1(n int) bool {
	num := strconv.Itoa(n)
	return twoAdjacentSameV1(num) && increasingDigits(num)
}

func isValidV2(n int) bool {
	num := strconv.Itoa(n)
	return twoAdjacentSameV2(num) && increasingDigits(num)
}

func twoAdjacentSameV1(num string) bool {
	lc := utf8.RuneCountInString(num)
	runes := []rune(num)
	for i, j := 0, 1; i < lc-1 && j < lc; i, j = i+1, j+1 {
		if runes[i] == runes[j] {
			return true
		}
	}
	return false
}

func increasingDigits(num string) bool {
	lc := utf8.RuneCountInString(num)
	runes := []rune(num)
	for i, j := 0, 1; i < lc-1 && j < lc; i, j = i+1, j+1 {
		a, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return false
		}
		b, err := strconv.Atoi(string(runes[j]))
		if err != nil {
			return false
		}
		if b < a {
			return false
		}
	}
	return true
}

func twoAdjacentSameV2(num string) bool {
	type byteOption struct {
		Valid bool
		Value byte
	}
	var prev byteOption
	for len(num) > 0 {
		head := num[0]
		tail := num[1:]
		if prev.Valid && prev.Value == head {
			if len(tail) == 0 || len(tail) > 0 && head != tail[0] {
				return true
			}
			for len(tail) > 0 && head == tail[0] {
				tail = tail[1:]
			}
		}
		prev = byteOption{Valid: true, Value: head}
		num = tail
	}
	return false
}
