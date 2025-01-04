package main

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("assert integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("assert strings", func(t *testing.T) {
		AssertEqual(t, "this", "this")
		AssertNotEqual(t, "this", "that")
	})
}

func TestStack(t *testing.T) {
	t.Run("assert int stack", func(t *testing.T) {
		
		intStack := new(Stack[int])

		// or using constructor:
		intStack = NewStack[int]()

		AssertTrue(t, intStack.IsEmpty())

		value, found := intStack.Pop()
		AssertEqual(t, value, 0)
		AssertFalse(t, found)

		intStack.Push(123)
		AssertFalse(t, intStack.IsEmpty())

		value, found = intStack.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, found)
		AssertTrue(t, intStack.IsEmpty())

		intStack.Push(2)
		intStack.Push(1)

		value, found = intStack.Pop()
		AssertEqual(t, value, 1)
		AssertTrue(t, found)
		AssertFalse(t, intStack.IsEmpty())

		value, found = intStack.Pop()
		AssertEqual(t, value, 2)
		AssertTrue(t, found)
		AssertTrue(t, intStack.IsEmpty())
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Error("expected true, got false")
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Error("expected false, got true")
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

// accepts any argument in got and want, like java Object
func AssertEqualInterface(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

// any constraints bot got and want to be the same type
func AssertEqualAny[T any](t *testing.T, got, want T) {
	t.Helper()

	// compilation error. any can not be compared.
	// if got != want {
	// 	 t.Errorf("got: %v, want: %v", got, want)
	// }
}
