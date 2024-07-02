package main

import "errors"

func add(x int, y int) int {
	return x + y
}

func devide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by zero!")
	}
	return x / y, nil // nil - отсутствие (ошибки)
}
