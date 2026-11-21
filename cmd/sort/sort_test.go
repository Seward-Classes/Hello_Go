package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{name: "empty", input: []int{}, expect: []int{}},
		{name: "single value", input: []int{4}, expect: []int{4}},
		{name: "duplicates and negatives", input: []int{3, -1, 3, 0, -5}, expect: []int{-5, -1, 0, 3, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quickSort(test.input)
			if !reflect.DeepEqual(test.input, test.expect) {
				t.Fatalf("quickSort() = %v, want %v", test.input, test.expect)
			}
		})
	}
}
