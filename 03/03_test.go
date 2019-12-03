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
		d        *data
		c        coord
		mask     int
		expected *data
	}{
		{
			d: &data{
				Coords:   map[coord]int{},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			c:    coord{X: 5, Y: 2},
			mask: 0b01,
			expected: &data{
				Coords: map[coord]int{
					coord{X: 5, Y: 2}: 0b01,
				},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
		},
		{
			d: &data{
				Coords: map[coord]int{
					coord{X: 99, Y: -99}: 0b00,
					coord{X: 5, Y: 2}:    0b10,
				},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			c:    coord{X: 5, Y: 2},
			mask: 0b01,
			expected: &data{
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
		tt.d.addCoord(tt.c, tt.mask)
		diff := cmp.Diff(tt.expected, tt.d)
		if diff != "" {
			t.Errorf("-expected +got\n%s", diff)
		}
	}
}

func TestAddCoords(t *testing.T) {
	tests := []struct {
		name     string
		d        *data
		movement string
		mask     int
		expected *data
	}{
		{
			name: "0,0 -> R003 (0b01)",
			d: &data{
				Coords:   map[coord]int{},
				Curr:     coord{X: 0, Y: 0},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			movement: "R003",
			mask:     0b01,
			expected: &data{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
				},
				Curr:     coord{X: 3, Y: 0},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
		},
		{
			name: "3,0 -> U002 (0b01)",
			d: &data{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
				},
				Curr:     coord{X: 3, Y: 0},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			movement: "U002",
			mask:     0b01,
			expected: &data{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,
				},
				Curr:     coord{X: 3, Y: 2},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
		},
		{
			name: "0,0 -> U002 (0b10)",
			d: &data{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,
				},
				Curr:     coord{X: 0, Y: 0},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			movement: "U002",
			mask:     0b10,
			expected: &data{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,

					coord{X: 0, Y: 1}: 0b10,
					coord{X: 0, Y: 2}: 0b10,
				},
				Curr:     coord{X: 0, Y: 2},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
		},
		{
			name: "0,2 -> R003 (0b10)",
			d: &data{
				Coords: map[coord]int{
					coord{X: 1, Y: 0}: 0b01,
					coord{X: 2, Y: 0}: 0b01,
					coord{X: 3, Y: 0}: 0b01,
					coord{X: 3, Y: 1}: 0b01,
					coord{X: 3, Y: 2}: 0b01,

					coord{X: 0, Y: 1}: 0b10,
					coord{X: 0, Y: 2}: 0b10,
				},
				Curr:     coord{X: 0, Y: 2},
				MinCoord: coord{},
				MinDist:  intOption{},
			},
			movement: "R003",
			mask:     0b10,
			expected: &data{
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
				Curr:     coord{X: 3, Y: 2},
				MinCoord: coord{X: 3, Y: 2},
				MinDist:  intOption{Valid: true, Value: 5},
			},
		},
	}
	for _, tt := range tests {
		tt.d.addCoords(tt.movement, tt.mask)
		diff := cmp.Diff(tt.expected, tt.d)
		if diff != "" {
			t.Errorf("name: %s -expected +got\n%s", tt.name, diff)
		}
	}
}
