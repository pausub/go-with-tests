package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of numbers in array of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectNumber(t, got, want, numbers)
	})
}

func assertCorrectNumber(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d, numbers: %v", got, want, numbers)
	}
}
