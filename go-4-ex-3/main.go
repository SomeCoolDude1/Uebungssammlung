package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	var results []float64
	d := math.Pow(b, 2) - 4*a*c

	if d < 0 {
		return results
	} else if d == 0 {
		results = append(results, (-b / (2 * a)))
		return results
	} else {
		results = append(results, ((-b + math.Sqrt(d)) / (2 * a)))
		results = append(results, ((-b - math.Sqrt(d)) / (2 * a)))
		return results
	}
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(3, 4, 1)) //[-0.333333333333333, -1]
	fmt.Println(computeQuadraticFormula(2, 4, 2)) //[-1]
	fmt.Println(computeQuadraticFormula(3, 4, 2)) //[]
}
