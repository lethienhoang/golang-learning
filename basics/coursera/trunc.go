package main

import (
	"math"
)

func roundToEven(x float64) float64 {
	t := math.Trunc(x)
	odd := math.Remainder(t, 2) != 0

	if d := math.Abs(x - t); d > 0.5 || (d == 0.5 && odd) {
		return t + math.Copysign(1, x)
	}

	return t
}

func main() {
	var x int
	var y *int
	z := 3
	y = &z
	x = &y
}
