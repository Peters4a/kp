package kp

import (
    "math"
)

// Upper Bound Procedure for the knapsack Problem
//
// Precondition: Profit[i]/Weight[i] >= Profit[i-1]/Weight[i-1] for i=1,...,n-1
//
// We do not check the precondition here!
// If the preconditioin is violated, the upper bound value will be probably wrong.
// Use kp.CheckSortedItems() to check the precondition.
func UpperBound(kp *KnapsackProblemT) {
    n := len(kp.Profit)
    kp.Xf = make([]float64,n)
    ub := 0
    c := kp.Capacity
    i := 0
    for ; i<n && kp.Weight[i]<=c ; i++ {
        ub += kp.Profit[i]
	c -= kp.Weight[i]
	kp.Xf[i] = 1.0
    }
    if i<n {
        kp.Xf[i] = float64(c) / float64(kp.Weight[i])
	ub += int(math.Floor(float64(kp.Profit[i])*kp.Xf[i]))
    }
    kp.Obj = ub
}

// Upper bound procedure for use within branch and bound algorithms.
// Uses the same algorithm as UpperBound(), but only returns the upper bound value.
// Moreover it is applicable to a subset of the items starting with index istart and
// a residual capacity c.
func uBound1P(p []int, w []int, c int, istart int) int {
    n := len(p)
    ub := 0
    i := istart
    for ; i<n && w[i]<=c ; i++ {
        ub += p[i]
        c -= w[i]
    }
    if i<n {
        ub += int(math.Floor(float64(p[i]) * float64(c) / float64(w[i])))
    }
    return ub
}
