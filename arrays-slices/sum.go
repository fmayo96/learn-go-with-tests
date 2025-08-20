package arraysslices

func Sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func SumAll(arr ...[]int) []int {
	out := make([]int, len(arr))
	for i, a := range arr {
		out[i] = Sum(a)
	}
	return out
}
