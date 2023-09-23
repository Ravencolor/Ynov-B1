package piscine

func Split(str, charset string) []string {
	ln := 0
	for i := range charset {
		ln = i + 1
	}
	ln2 := 0
	for i := range str {
		ln2 = i + 1
	}
	for i := 0; i < ln2-ln; i++ {
		if str[i:i+ln] == charset {
			str = str[:i] + " " + str[i+ln:]
			ln2 -= ln
		}
	}
	return SplitWhiteSpaces(str)
}

func SplitWhiteSpaces(str string) []string {
	TextToString := ""
	t := []string{}
	for i, v := range str {
		if i == lent2(str)-1 && string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
			t = append(t, TextToString)
		} else if string(v) != " " && string(v) != "\t" && string(v) != "\n" {
			TextToString += string(v)
		} else {
			if len(TextToString) >= 1 {
				t = append(t, TextToString)
			}
			TextToString = ""
		}
	}
	return t
}

func lent2(d string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}
