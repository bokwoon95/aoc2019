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
		if str == "" {
			continue
		}
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
	prog[2] = 3
	final := run(prog)
	fmt.Println(final)
	noun, verb := findNounVerb(prog, 19690720)
	fmt.Printf("noun: %d, verb: %d, answer: %d%d\n", noun, verb, noun, verb)
}

func run(prog []int) (p []int) {
	for _, num := range prog {
		p = append(p, num)
	}
	var pc int
Loop:
	for {
		nounIdx := p[pc+1]
		verbIdx := p[pc+2]
		outputIdx := p[pc+3]
		noun := p[nounIdx]
		verb := p[verbIdx]
		switch p[pc] {
		case OpAdd:
			p[outputIdx] = noun + verb
		case OpMul:
			p[outputIdx] = noun * verb
		case OpHalt:
			break Loop
		}
		pc += 4
	}
	return p
}

func findNounVerb(prog []int, target int) (noun, verb int) {
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			prog[1] = noun
			prog[2] = verb
			final := run(prog)
			if final[0] == target {
				return noun, verb
			}
		}
	}
	return -1, -1
}
