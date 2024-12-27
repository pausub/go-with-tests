package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{15.0, 15.0})
	want := 60.0
	assertCorrectResult(t, got, want)
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectResult(t, got, want)
	}

	t.Run("test rectangle area", func(t *testing.T) {
		checkArea(t, Rectangle{15.0, 15.0}, 225.0)
	})
	t.Run("test circle area", func(t *testing.T) {
		checkArea(t,  Circle{10.0}, 314.1592653589793)
	})
}

func assertCorrectResult(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
