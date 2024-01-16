package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		i++
	}
	for i < len(s) && s[i] != ' ' {
		return 0
	}
	return result * sign
}
