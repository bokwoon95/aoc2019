package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bokwoon95/aoc2019/read"
)

type data []int

func (d *data) Append(line string) (err error) {
	mass, err := strconv.Atoi(line)
	*d = append(*d, mass)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	masses := data{}
	err := read.Input("input.txt", &masses)
	if err != nil {
		log.Fatal(err)
	}

	totalFuel := 0
	for _, mass := range masses {
		fuel := getFuelV2(mass)
		fmt.Printf("mass:%d fuel:%d\n", mass, fuel)
		totalFuel += fuel
	}
	fmt.Println(totalFuel)
}

func getFuelV1(mass int) (fuel int) {
	fuel = (mass / 3) - 2
	return fuel
}

func getFuelV2(mass int) (fuel int) {
	for mass > 0 {
		mass = mass/3 - 2
		if mass > 0 {
			fuel += mass
		}
	}
	return fuel
}
