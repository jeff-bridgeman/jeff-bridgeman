package maths

// Fibonnaci : Generates a list up to x terms
func Fibonnaci(x int) []int {

	sequence := make([]int, x)
	sequence[0] = 1
	sequence[1] = 2
	for i := 2; i < x; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}
	return sequence
}
