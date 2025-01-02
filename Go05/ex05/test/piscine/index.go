package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				break
			}
			if j == len(toFind)-1 {
				return i
			}
		}
	}
	return -1
}
