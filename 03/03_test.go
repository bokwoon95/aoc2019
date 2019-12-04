package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalcManhattanDist(t *testing.T) {
	tests := []struct {
		curr     coord
		expected int
	}{
		{coord{1, 2}, 3},
		{coord{1, -2}, 3},
		{coord{-100, -999}, 1099},
	}
	for _, tt := range tests {
		got := calcManhattanDist(tt.curr)
		diff := cmp.Diff(tt.expected, got)
		if diff != "" {
			t.Errorf("-expected +got\n%s", diff)
		}
	}
}

func TestAddCoord(t *testing.T) {
	tests := []struct {
		d        *data1
		c        coord
		mask     int
		expected *data1
	}{
		{
			d: &data1{
				Coords:   map[coord]int{},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			c:    coord{X: 5, Y: 2},
			mask: 0b01,
			expected: &data1{
				Coords: map[coord]int{
					coord{X: 5, Y: 2}: 0b01,
				},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
		},
		{
			d: &data1{
				Coords: map[coord]int{
					coord{X: 99, Y: -99}: 0b00,
					coord{X: 5, Y: 2}:    0b10,
				},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			c:    coord{X: 5, Y: 2},
			mask: 0b01,
			expected: &data1{
				Coords: map[coord]int{
					coord{X: 99, Y: -99}: 0b00,
					coord{X: 5, Y: 2}:    0b11,
				},
				MinCoord: coord{X: 5, Y: 2},
				MinDist:  intOption{Valid: true, Value: 7},
			},
		},
	}
	for _, tt := range tests {
		tt.d.addCoordV1(tt.c, tt.mask)
		diff := cmp.Diff(tt.expected, tt.d)
		if diff != "" {
			t.Errorf("-expected +got\n%s", diff)
		}
	}
}

func TestAddCoords(t *testing.T) {
	tests := []struct {
		name     string
		d        *data1
		movement string
		mask     int
		expected *data1
	}{
		{
			name: "0,0 -> R003 (0b01)",
			d: &data1{
				Coords:    map[coord]int{},
				CurrCoord: coord{X: 0, Y: 0},
				MinCoord:  coord{},
				MinDist:   intOption{},
			},
			movement: "R003",
			mask:     0b01,
			expected: &data1{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
				},
				CurrCoord: coord{X: 3, Y: 0},
				MinCoord:  coord{},
				MinDist:   intOption{},
			},
		},
		{
			name: "3,0 -> U002 (0b01)",
			d: &data1{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
				},
				CurrCoord: coord{X: 3, Y: 0},
				MinCoord:  coord{},
				MinDist:   intOption{},
			},
			movement: "U002",
			mask:     0b01,
			expected: &data1{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,
				},
				CurrCoord: coord{X: 3, Y: 2},
				MinCoord:  coord{},
				MinDist:   intOption{},
			},
		},
		{
			name: "0,0 -> U002 (0b10)",
			d: &data1{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,
				},
				CurrCoord: coord{X: 0, Y: 0},
				MinCoord:  coord{},
				MinDist:   intOption{},
			},
			movement: "U002",
			mask:     0b10,
			expected: &data1{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,

					coord{X: 0, Y: 1}: 0b10,
					coord{X: 0, Y: 2}: 0b10,
				},
				CurrCoord: coord{X: 0, Y: 2},
				MinCoord:  coord{},
				MinDist:   intOption{},
			},
		},
		{
			name: "0,2 -> R003 (0b10)",
			d: &data1{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,

					coord{X: 0, Y: 1}: 0b10,
					coord{X: 0, Y: 2}: 0b10,
				},
				CurrCoord: coord{X: 0, Y: 2},
				MinCoord:  coord{},
				MinDist:   intOption{},
			},
			movement: "R003",
			mask:     0b10,
			expected: &data1{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b11,

					coord{X: 0, Y: 1}: 0b10,
					coord{X: 0, Y: 2}: 0b10,
					coord{X: 1, Y: 2}: 0b10,
					coord{X: 2, Y: 2}: 0b10,
				},
				CurrCoord: coord{X: 3, Y: 2},
				MinCoord:  coord{X: 3, Y: 2},
				MinDist:   intOption{Valid: true, Value: 5},
			},
		},
	}
	for _, tt := range tests {
		tt.d.addCoordsV1(tt.movement, tt.mask)
		diff := cmp.Diff(tt.expected, tt.d)
		if diff != "" {
			t.Errorf("name: %s -expected +got\n%s", tt.name, diff)
		}
	}
}

func TestCalcTotalDist(t *testing.T) {
	tests := []struct {
		record   map[int]int
		expected int
	}{
		{map[int]int{}, 0},
		{map[int]int{1: 27, 2: 56}, 83},
	}
	for _, tt := range tests {
		got := calcTotalDist(tt.record)
		diff := cmp.Diff(tt.expected, got)
		if diff != "" {
			t.Errorf("-expected +got\n%s", diff)
		}
	}
}

func TestAddCoordV2(t *testing.T) {
	tests := []struct {
		d2       *data2
		c        coord
		mask     int
		expected *data2
	}{
		{
			d2: &data2{
				Coords: map[coord]map[int]int{
					coord{X: 1, Y: 1}: map[int]int{
						0b01: 10,
					},
					coord{X: 5, Y: 2}: map[int]int{
						0b01: 5,
					},
				},
				CurrDist: 5,
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			c:    coord{X: 5, Y: 2},
			mask: 0b10,
			expected: &data2{
				Coords: map[coord]map[int]int{
					coord{X: 1, Y: 1}: map[int]int{
						0b01: 10,
					},
					coord{X: 5, Y: 2}: map[int]int{
						0b01: 5,
						0b10: 5,
					},
				},
				CurrDist: 5,
				MinCoord: coord{X: 5, Y: 2},
				MinDist:  intOption{Valid: true, Value: 10},
			},
		},
	}
	for _, tt := range tests {
		tt.d2.addCoordV2(tt.c, tt.mask)
		diff := cmp.Diff(tt.expected, tt.d2)
		if diff != "" {
			t.Errorf("-expected +got\n%s", diff)
		}
	}
}

func TestAddCoordsV2(t *testing.T) {
}
