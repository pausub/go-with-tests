package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{15.0, 15.0})
	want := 60.0
	assertCorrectResult(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("test rectangle area", func(t *testing.T) {
		got := Rectangle{15.0, 15.0}.Area()
		want := 225.0
		assertCorrectResult(t, got, want)
	})
	t.Run("test circle area", func(t *testing.T) {
		got := Circle{10.0}.Area()
		want := 314.1592653589793
		assertCorrectResult(t, got, want)
	})
}

func assertCorrectResult(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
