package piscine

func IsDuplicate(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return true
			}
		}
	}
	return false
}

func IsExist(str1 string, str2 string) bool {
	count := 0
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				count++
				break
			}
		}
	}

	if count == len(str1) {
		return true
	} else {
		return false
	}
}

func IsAlNum(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	} else if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func IsAlNumStr(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsAlNum(c) {
			return false
		}
	}
	return true
}

func Calc(s string, base string, length int) int {
	if length == 1 {
		return Index(base, string(s[length-1]))
	}

	tmp := len(base) * Calc(s, base, length-1)
	tmp += Index(base, string(s[length-1]))
	return tmp
}

func AtoiBase(s string, base string) int {
	if len(base) < 2 || IsDuplicate(base) || !IsAlNumStr(base) || !IsExist(s, base) {
		return 0
	}
	return Calc(s, base, len(s))
}
