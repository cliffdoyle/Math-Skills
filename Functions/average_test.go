package functions

import "testing"

func TestAverage(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := 3.0 // Expected average of the data
	result := Average(data)
	if result != expected {
		t.Errorf("Average of %v was incorrect, got: %v, want: %v.", data, result, expected)
	}
}
