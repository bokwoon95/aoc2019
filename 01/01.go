package main

import (
	"log"

	"github.com/bokwoon95/aoc2019/read"
)

type t []string

func (d *t) Append(line string) (err error) {
	return err
}

func main() {
	data := t{}
	err := read.Input("input.txt", &data)
	if err != nil {
		log.Fatal(err)
	}
}
