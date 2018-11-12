package optimize

import (
	"errors"
	"math"
)

// SolveFunction is target function for a solving
type SolveFunction func(float64) float64

// SolveDerivFunction is target first derive function for a solving
type SolveDerivFunction func(float64) float64

// BisectionSolve is solver using a bisection method
func BisectionSolve(f SolveFunction, start float64, end float64) (float64, error) {
	maxIter := 1000
	tol := 1e-4
	i := 1
	FA := f(start)
	for i < maxIter {
		p := start + (end-start)/2.0
		FP := f(p)
		if FP == 0 || (end-start)/2.0 < tol {
			return p, nil
		}
		i++
		if FA*FP > 0 {
			start = p
			FA = FP
		} else {
			end = p
		}
	}

	return -999, errors.New("Maximum number of itertions reached")
}

// NewtonSolve is solver using a newton's method
func NewtonSolve(f SolveFunction, df SolveDerivFunction, initialGeuess float64) (float64, error) {
	maxIter := 1000
	tol := 1e-5
	i := 1
	for i < maxIter {
		p := initialGeuess - f(initialGeuess)/df(initialGeuess)
		if math.Abs(p-initialGeuess) < tol {
			return p, nil
		}
		i++
		initialGeuess = p
	}
	return -999, errors.New("Maximum number of itertions reached")
}
