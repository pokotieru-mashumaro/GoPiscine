package piscine

func calc(ret int, nb int, power int) int {
	if power == 0 {
		return ret
	}
	return calc(ret*nb, nb, power-1)
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	return calc(1, nb, power)
}
