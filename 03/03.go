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

type data1 struct {
	Coords    map[coord]int
	CurrCoord coord
	MinCoord  coord
	MinDist   intOption
}

type data2 struct {
	Coords    map[coord]map[int]int
	CurrCoord coord
	CurrDist  int
	MinCoord  coord
	MinDist   intOption
}

func main() {
	file, err := os.Open("03.txt")
	if err != nil {
		log.Fatal(err)
	}
	d1 := &data1{Coords: make(map[coord]int)}
	d2 := &data2{Coords: make(map[coord]map[int]int)}
	mask := 0b01 // First line's mask is 0b01, second line's mask is 0b10
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Before each iteration, reset the current coordinate to 0, 0
		d1.CurrCoord = coord{X: 0, Y: 0}
		d2.CurrCoord = coord{X: 0, Y: 0}
		d2.CurrDist = 0
		strs := strings.Split(scanner.Text(), ",")
		for _, str := range strs {
			d1.addCoordsV1(str, mask)
			d2.addCoordsV2(str, mask)
		}
		mask <<= 1 // Increment the bit mask
	}
	fmt.Println(d1.CurrCoord, d1.MinCoord, d1.MinDist, d1.Coords[d1.MinCoord])
	fmt.Println(d2.CurrCoord, d2.MinCoord, d2.MinDist, d2.Coords[d2.MinCoord])
}

// Given a movement e.g. R003, increments the X value of the current coordinate
// 3 times and calls d1.addCoordV1 withh the updated coordinate each time.
func (d1 *data1) addCoordsV1(movement string, mask int) {
	direction := movement[:1]
	offset, err := strconv.Atoi(movement[1:])
	if err != nil {
		panic(err)
	}
	for i := 0; i < offset; i++ {
		switch direction {
		case "U":
			d1.CurrCoord.Y++
		case "D":
			d1.CurrCoord.Y--
		case "L":
			d1.CurrCoord.X--
		case "R":
			d1.CurrCoord.X++
		} // increment or decrement the coordinate by 1 depending on the direction
		d1.addCoordV1(d1.CurrCoord, mask) // add the latest coordinate into the d1.Coords
	}
}

// Add a coordinate into d1.Coords. If the coordinate has been visited before,
// it is an intersection and the manhattan distance from the origin should be
// calculated. If the distance is lower than the previously recorded minimum
// distance, make this coordinate the new minimum coordinate + distance
func (d1 *data1) addCoordV1(c coord, mask int) {
	// Check if any non-mask bits are 1. If so, then some other wire has
	// previously crossed this coordinate i.e. INTERSECTION
	if d1.Coords[c]&^mask != 0 {
		dist := calcManhattanDist(c)
		if !d1.MinDist.Valid || dist < d1.MinDist.Value {
			d1.MinCoord = c
			d1.MinDist = intOption{Valid: true, Value: dist}
		}
	}
	// Set mask bits
	d1.Coords[c] |= mask
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

// Given a movement e.g. R003, increments the X value of the current coordinate
// 3 times and calls d2.addCoordV2 withh the updated coordinate each time.
func (d2 *data2) addCoordsV2(movement string, mask int) {
	direction := movement[:1]
	offset, err := strconv.Atoi(movement[1:])
	if err != nil {
		panic(err)
	}
	for i := 0; i < offset; i++ {
		switch direction {
		case "U":
			d2.CurrCoord.Y++
		case "D":
			d2.CurrCoord.Y--
		case "L":
			d2.CurrCoord.X--
		case "R":
			d2.CurrCoord.X++
		} // increment or decrement the coordinate by 1 depending on the direction
		d2.CurrDist++                     // increment the current distance travelled by 1
		d2.addCoordV2(d2.CurrCoord, mask) // add the latest coordinate into the d2.Coords
	}
}

// Add a coordinate into d2.Coords. If the coordinate has been visited before,
// it is an intersection and the total distance travelled by each wire up til
// the coordinate should be added up. If the sum of the distances travelled is
// lower than the previously recorded minimum distance travelled, make this the
// new minimum coordinate + distance
func (d2 *data2) addCoordV2(c coord, mask int) {
	if d2.Coords[c] == nil {
		d2.Coords[c] = make(map[int]int)
	}
	d2.Coords[c][mask] = d2.CurrDist
	if len(d2.Coords[c]) > 1 {
		totalDist := calcTotalDist(d2.Coords[c])
		if !d2.MinDist.Valid || totalDist < d2.MinDist.Value {
			d2.MinCoord = c
			d2.MinDist = intOption{Valid: true, Value: totalDist}
		}
	}
}

func calcTotalDist(record map[int]int) (dist int) {
	for _, v := range record {
		dist += v
	}
	return dist
}
