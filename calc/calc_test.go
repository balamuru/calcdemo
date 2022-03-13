package calc_test

import (
	"github.com/balamuru/calcdemo/calc"
	"testing"
)

//simple test
func TestAdd(t *testing.T) {
	got := calc.Add(1, 2)
	expected := 3
	if got != expected {
		t.Fatalf("Got %v, Expected %v\n", got, expected)
	}
}

//simple test
func TestSub(t *testing.T) {
	got := calc.Sub(3, 2)
	expected := 1
	if got != expected {
		t.Fatalf("Got %v, Expected %v\n", got, expected)
	}
}

//test data set
func TestAddTable(t *testing.T) {
	data := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 2, y: 1, expected: 3},
		{1, 2, 3},
		{1, 5, 6},
	}

	for _, v := range data {
		got := calc.Add(v.x, v.y)
		expected := v.expected
		if got != expected {
			t.Fatalf("Got %v, Expected %v\n", got, expected)
		}
	}
}
