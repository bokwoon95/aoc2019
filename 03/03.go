package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct{ X, Y int }

type intOption struct {
	Valid bool
	Value int
}

type data struct {
	Coords   map[coord]int
	Curr     coord
	MinCoord coord
	MinDist  intOption
}

func main() {
	file, err := os.Open("03.txt")
	if err != nil {
		log.Fatal(err)
	}
	d := &data{Coords: make(map[coord]int)} // Initialize data
	mask := 0b01 // First line's mask is 0b01, second line's mask is 0b10
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d.Curr = coord{X: 0, Y: 0} // reset the current coord to 0, 0
		strs := strings.Split(scanner.Text(), ",")
		for _, str := range strs {
			d.addCoords(str, mask)
		}
		mask <<= 1 // Increment the bit mask
	}
	fmt.Println(d.Curr, d.MinCoord, d.MinDist)
}

// Given a movement e.g. R003, increments the X value of the current coordinate
// 3 times and calls (*.data).addCoord withh the updated coordinate each time.
func (d *data) addCoords(movement string, mask int) {
	direction := movement[:1]
	offset, err := strconv.Atoi(movement[1:])
	if err != nil {
		panic(err)
	}
	for i := 0; i < offset; i++ {
		switch direction {
		case "U":
			d.Curr.Y++
		case "D":
			d.Curr.Y--
		case "L":
			d.Curr.X--
		case "R":
			d.Curr.X++
		}
		d.addCoord(d.Curr, mask)
	}
}

// Given a coord c and an int mask, adds the map[coord]int entry to d.Coords.
func (d *data) addCoord(c coord, mask int) {
	// Check if any non-mask bits are 1. If so, then some other wire has
	// previously crossed this coordinate i.e. INTERSECTION
	if d.Coords[c]&^mask != 0 {
		dist := calcManhattanDist(c)
		if !d.MinDist.Valid || dist < d.MinDist.Value {
			d.MinCoord = c
			d.MinDist = intOption{Valid: true, Value: dist}
		}
	}
	// Set mask bits
	d.Coords[c] |= mask
}

func calcManhattanDist(curr coord) (dist int) {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	return abs(curr.X) + abs(curr.Y)
}
