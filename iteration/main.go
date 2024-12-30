package main

func Repeat(s string, times int) (result string) {
	for i := 0; i < times; i++ {
		result += s
	}
	return
}
