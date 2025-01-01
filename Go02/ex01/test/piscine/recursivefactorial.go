package piscine

func calc(ret int, nb int) int {
	if nb == 0 {
		return ret
	}
	tmp := ret * nb
	if tmp < ret {
		return 0
	}
	return calc(tmp, nb-1)
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	return calc(1, nb)
}
