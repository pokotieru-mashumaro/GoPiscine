package piscine

func Join(strs []string, sep string) string {
	ret := ""
	for i, elem := range strs {
		ret += elem
		if i != len(strs)-1 {
			ret += sep
		}
	}
	return ret
}
