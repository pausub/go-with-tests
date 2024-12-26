package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of numbers in array of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectNumbers(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of all integers in slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 0, 1})
		want := []int{6, 2}
		assertCorrectResult(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum of all tails of integers in slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 0, 1})
		want := []int{5, 1}
		assertCorrectResult(t, got, want)
	})
	t.Run("Sum of all tails of integers of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{9, 1})
		want := []int{0, 1}
		assertCorrectResult(t, got, want)
	})
}

func assertCorrectNumbers(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d, numbers: %v", got, want, numbers)
	}
}

func assertCorrectResult(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
