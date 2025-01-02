package piscine

func CalcConvert(n int, base string) string {
	if n <= 0 {
		return ""
	}
	l := len(base)
	tmp := CalcConvert(n/l, base)
	tmp += string(base[n%l])
	return tmp
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AtoiBase(nbr, baseFrom)
	return CalcConvert(n, baseTo)
}
