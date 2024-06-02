package functions

import "testing"

func TestVariance(t *testing.T) {
    data := []int{1, 2, 3, 4, 5}
    expected := 2.0 // Expected variance of the data
    result := Variance(data)
    if result != expected {
        t.Errorf("Variance of %v was incorrect, got: %v, want: %v.", data, result, expected)
    }
}
