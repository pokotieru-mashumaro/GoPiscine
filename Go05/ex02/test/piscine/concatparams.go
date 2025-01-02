package piscine

func ConcatParams(args []string) string {
	tmp := ""
	for i, s := range args {
		tmp += s
		if i != len(args)-1 {
			tmp += "\n"
		}
	}
	return tmp
}
