package calculator

import (
	"math"
)

type Result struct {
	Result float64 `json:"result"`
}

func Add(first, second float64) Result {
	return Result{Result: first + second}
}

func Sub(first, second float64) Result {
	return Result{Result: first - second}
}

func Mul(first, second float64) Result {
	return Result{Result: first * second}
}

func Div(first, second float64) Result {
	return Result{Result: first / second}
}

func Pow(first, second float64) Result {
	return Result{Result: math.Pow(first, second)}
}
