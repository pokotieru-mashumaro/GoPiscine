package piscine

func BasicJoin(elems []string) string {
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	return ret
}
