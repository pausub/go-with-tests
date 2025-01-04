package main

type Stack[T any] struct {
	values[] T
}

// generic stack constructor
func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (t *Stack[T]) Push(value T) {
	t.values = append(t.values, value)
}

func (t *Stack[T]) IsEmpty() bool {
	return len(t.values) == 0
}

func (t *Stack[T]) Pop() (T, bool) {
	if t.IsEmpty() {
		var none T
		return none, false
	}
	
	index := len(t.values)-1
	value := t.values[index]
	t.values = t.values[:index]
	return value, true
}