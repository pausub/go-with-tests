package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{15.0, 15.0})
	want := 60.0
	assertCorrectResult(t, got, want)
}

func TestArea(t *testing.T) {

	// anonymous struct slice for Table Driven Tests
	// each line is table represents single test case
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, tableTest := range areaTests {
		// t.Run creates a separate test result line with given name
		// so its easier to see which exact shape succeeded and failed
		// instead of one failed TestArea test
		t.Run(tableTest.name, func(t *testing.T) {
			got := tableTest.shape.Area()
			assertCorrectResult(t, got, tableTest.want)
		})
	}
}

func assertCorrectResult(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
