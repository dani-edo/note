package main

import "fmt"

func main() {
	result := add(1, 2) // type is automatically inferred by the source data (1 and 2)
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
