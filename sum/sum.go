package sum

// array size is aprt of type. Its impossible to declare array
// type without size.
func Sum(a []int) (sum int) {
	for _, num := range a {
		sum += num
	}
	return
}

// Use append(slice, item) to add item to slice and not worry about length or capacity
func SumAll(slices ...[]int) (sums []int) {
	for _, nums := range slices {
		sums = append(sums, Sum(nums))
	}
	return
}

// Slice a slice with [start:end] operator. 
// Ommit numbers to capture everyting to that side
func SumAllTails(slices ...[]int) (sums []int) {
	for _, nums := range slices {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
