package piscine

func BasicAtoi2(s string) int {
	arr := []rune(s)
	n := len(s)
	v := 0
	for i := 0; i < n; i++ {
		if arr[i] < '0' || arr[i] > '9' {
			return 0
		} else {
			v *= 10
			v += int(arr[i]) - '0'
		}
	}
	return v
}
