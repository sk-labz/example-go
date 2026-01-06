package calculator


import "errors"

func Add(x, y int) (int, error) {
	return x + y, nil
}

func Subtract(x, y int) (int, error) {
	return x - y, nil
}

func Multiply(x, y int) (int, error) {
	return x * y, nil
}

