package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	tmp := 0
	for i := 1; i <= nb; i++ {
		tmp = result * i
		if tmp < result {
			return 0
		}
		result = tmp
	}
	return result
}
