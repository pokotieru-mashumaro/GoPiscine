package piscine

func SortParams(params []string) []string {
	for i := 0; i < len(params); i++ {
		for j := 0; j < len(params); j++ {
			if Compare(params[i], params[j]) == -1 {
				params[i], params[j] = params[j], params[i]
			}
		}
	}
	return params
}
