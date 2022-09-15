package main

import "fmt"

func main() {
	var num int
	z := 0

	fmt.Printf("Enter a number:")
	fmt.Scanf("%d", &num)

	for x := 0; x < num; x++ {
		if num/2 > z {
			for y := 0; y < num; y++ {

				if y+1 < num-x {
					fmt.Printf(" ")

				} else {
					fmt.Printf("# ")

				}
			}
			fmt.Println()
			z++
		} else {
			for y := 0; y < num; y++ {
				if y < x {
					fmt.Printf(" ")
				} else {
					fmt.Printf("# ")
				}
			}
			fmt.Println()
		}
	}
}
