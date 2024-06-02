package functions

import "testing"

func TestStdDev(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := 1.4142135623730951// Expected standard deviation of the data
	result := StdDeviation(data)
	if result != expected {
		t.Errorf("Standard deviation of %v was incorrect, got: %v, want: %v.", data, result, expected)
	}
}
