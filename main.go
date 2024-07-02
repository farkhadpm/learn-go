package main

import (
	"fmt"
)

func main() {
	fmt.Print("Assalamu Aleykum, brother!")

	fmt.Println("\n\nVariables:")
	print_var()

	fmt.Println("\n\nSumm:")
	add_result := add(1, 2)
	fmt.Println("x + y = ", add_result)

	fmt.Println("\n\nDivision: ")
	devide_res, err := devide(32, 0)
	if err != nil {
		fmt.Println("Erorr:", err)
	}
	fmt.Println("x / y = ", devide_res)

}
