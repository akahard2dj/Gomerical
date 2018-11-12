package main

import (
	"fmt"
	"gomerical/optimize"
)

func targetFunction(x float64) float64 {
	return x*x*x + 4*x*x - 10.0
}

func main() {
	var sf optimize.SolveFunction
	sf = targetFunction
	fmt.Println(optimize.BisectionSolve(sf, 1, 2))
}
