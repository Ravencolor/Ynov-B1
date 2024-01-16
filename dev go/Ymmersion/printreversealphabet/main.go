package main

import "github.com/01-edu/z01"

func main() {
	for char := (97 + 25); char >= (97); char-- {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}
