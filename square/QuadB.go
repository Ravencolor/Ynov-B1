package main

import "fmt"

func main() {
	QuadB(5,3)
}
func QuadB(x, y int) {
	if x > 0 && y > 0 {
		fmt.Println("/" + repeat("*", x-2) + "\\")
		for i := 1; i <= y-2; i++ {
			fmt.Print("*")
			for j := 1; j <= x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
			fmt.Println()
		}
		fmt.Println("\\" + repeat("*", x-2) + "/")
		fmt.Println()
	} else {
        fmt.Println("Les dimensions doivent être positives.")
        return
	}
}

func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
