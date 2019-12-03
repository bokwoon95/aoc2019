package main

import "testing"

func TestGetFuelV1(t *testing.T) {
	tests := []struct {
		mass     int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		got := getFuelV1(tt.mass)
		if tt.expected != got {
			t.Errorf("mass:%d expected:%d got:%d", tt.mass, tt.expected, got)
		}
	}
}

func TestGetFuelV2(t *testing.T) {
	tests := []struct {
		mass     int
		expected int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		got := getFuelV2(tt.mass)
		if tt.expected != got {
			t.Errorf("mass:%d expected:%d got:%d", tt.mass, tt.expected, got)
		}
	}
}
