package piscine

func AppendRange(min, max int) []int {
	array := []int{}
	if min >= max {
		return array
	}
	for i := min; i < max; i++ {
		array = append(array, i)
	}
	return array
}
