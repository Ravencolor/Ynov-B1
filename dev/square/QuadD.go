package main

import "fmt"

func main() {
	QuadD(5,3)
}
func QuadD(x, y int) {
	if x > 0 && y > 0 {
		fmt.Println("A" + repeat("B", x-2) + "C")
		for i := 1; i <= y-2; i++ {
			fmt.Print("B")
			for j := 1; j <= x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("B")
			fmt.Println()
		}
		fmt.Println("A" + repeat("B", x-2) + "C")
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
