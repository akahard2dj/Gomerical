package optimize

import (
	"errors"
	"fmt"
)

type SolveFunction func(float64) float64

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

	return -999, errors.New("The number of iteration is reached by maxIter")
}
