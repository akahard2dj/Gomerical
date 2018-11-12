package optimize_test

import (
	"Gomerical/optimize"
	"math"
	"testing"
)

func targetBisectionFunction(x float64) float64 {
	return x*x*x + 4*x*x - 10.0
}

func targetNewtonFunction(x float64) float64 {
	return math.Cos(x) - x
}

func targetNewtonDerivFunction(x float64) float64 {
	return -math.Sin(x) - 1
}

func TestBisectionSolve(t *testing.T) {
	var sf optimize.SolveFunction
	sf = targetBisectionFunction
	result, _ := optimize.BisectionSolve(sf, 1, 2)

	if result != 1.36517333984375 {
		t.Error("Wrong result")
	}
}

func TestNewtonSolve(t *testing.T) {
	var sf optimize.SolveFunction
	var sdf optimize.SolveDerivFunction
	sf = targetNewtonFunction
	sdf = targetNewtonDerivFunction
	result, _ := optimize.NewtonSolve(sf, sdf, math.Pi/4.0)

	if result != 0.739085133215161 {
		t.Error("Wrong result")
	}
}
