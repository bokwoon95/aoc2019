package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsValidV1(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{111111, true},
		{122345, true},
		{223450, false},
		{123789, false},
	}
	for i, tt := range tests {
		got := isValidV1(tt.n)
		diff := cmp.Diff(tt.expected, got)
		if diff != "" {
			t.Errorf("%d: -expected +got\n%s", i, diff)
		}
	}
}

func TestIsValidV2(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{111111, false},
		{123444, false},
		{111122, true},
		{111123, false},
	}
	for i, tt := range tests {
		if i+1 == 3 {
			if i == i {
			}
		}
		got := isValidV2(tt.n)
		diff := cmp.Diff(tt.expected, got)
		if diff != "" {
			t.Errorf("test %d: -expected +got\n%s", i+1, diff)
		}
	}
}
