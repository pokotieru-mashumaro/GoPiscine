package piscine

func SplitWhiteSpaces(s string) []string {
	tmp := ""
	splits := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if tmp != "" {
				splits = append(splits, tmp)
				tmp = ""
			}
			continue
		}
		tmp += string(s[i])
		if i == len(s)-1 && tmp != "" {
			splits = append(splits, tmp)
		}
	}
	return splits
}
