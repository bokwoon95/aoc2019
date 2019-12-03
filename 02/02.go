package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bokwoon95/aoc2019/read"
)

// Ops denote an operation
const (
	OpAdd  = 1
	OpMul  = 2
	OpHalt = 99
)

type data []int

func (d *data) Append(line string) (err error) {
	strs := strings.Split(line, ",")
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*d = append(*d, num)
	}
	return nil
}

func main() {
	prog := data{}
	err := read.Input("input.txt", &prog)
	if err != nil {
		log.Fatal(err)
	}
	prog[1] = 12
	prog[2] = 2
	final := runV1(prog)
	fmt.Println(final)
}

func runV1(prog []int) (p []int) {
	for _, num := range prog {
		p = append(p, num)
	}
	var pc int
Loop:
	for {
		switch p[pc] {
		case OpAdd:
			p[p[pc+3]] = p[p[pc+2]] + p[p[pc+1]]
		case OpMul:
			p[p[pc+3]] = p[p[pc+2]] * p[p[pc+1]]
		case OpHalt:
			break Loop
		}
		pc += 4
	}
	return p
}
