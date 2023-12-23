package piscine

func UltimateDivMod(a *int, b *int) {
	quotient := *a / *b
	remainder := *a % *b
	*a = quotient
	*b = remainder
}
