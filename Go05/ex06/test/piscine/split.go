package piscine

func IsSep(s string, sep string, i int) bool {
	for j := 0; j < len(sep); j++ {
		if s[i+j] != sep[j] {
			return false
		}
	}
	return true
}

func Split(s, sep string) []string {
	tmp := ""
	splits := []string{}
	if len(sep) == 0 {
		return append(splits, s)
	}
	for i := 0; i < len(s); i++ {
		if IsSep(s, sep, i) {
			if tmp != "" {
				splits = append(splits, tmp)
				tmp = ""
			}
			i += len(sep) - 1
			continue
		}
		tmp += string(s[i])
		if i == len(s)-1 && tmp != "" {
			splits = append(splits, tmp)
		}
	}
	return splits
}
