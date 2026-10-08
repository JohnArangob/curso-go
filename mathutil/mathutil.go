package mathutil

import (
	"errors"
	"fmt"
)

func Add(a, b int) (error, int) {
	if a > b {
		return nil, a + b
	}

	fmt.Println("El primer número debe ser mayor que el segundo.")
	return errors.New("el primer número debe ser mayor que el segundo"), 0
}

func Subtract(a, b int) int {
	return a - b
}
