package functions

import "testing"

func TestMedian(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := 3.0 // Expected median of the data
	result := Median(data)
	if result != expected {
		t.Errorf("Median of %v was incorrect, got: %v, want: %v.", data, result, expected)
	}
}
