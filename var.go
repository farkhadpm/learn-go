package main

import (
	"fmt"
)

func print_var() {
	var intt int = 5
	var stringg string = "an apple"
	var floatt float64 = 9.865

	intt2 := 10 // :=
	const pi float64 = 3.14

	fmt.Println("Integer: ", intt, "\nString: ", stringg, "\nFloat: ", floatt, "\nInteger (:=): ", intt2, "\nConstant Pi = ", pi)
}
