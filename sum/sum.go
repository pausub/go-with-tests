package sum

// array size is aprt of type. Its impossible to declare array 
// type without size.
func Sum(a []int) (sum int) {
	for _, num := range a {
		sum += num
	}
	return
}
