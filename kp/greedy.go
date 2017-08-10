package kp

// Primal greedy heuristic for the knapsack problem 
//
// Precondition: Profit[i]/Weight[i] >= Profit[i-1]/Weight[i-1] for i=1,...,n-1
//
// We do not check the precondition here!
// Greedy() computes a feasible solution even if the precondition is not valid,
// but this solution is usually considerably worse.
//
// To check the precondition, use kp.CheckSortedItems().
//
func Greedy(kp KnapsackProblem) ([]int,int) {
    n := kp.N()					// n is the number of items we have
    x := make([]int,n)				// X[i] = 0 for i=0,...,n-1
    c := kp.Capacity()
    z := 0					// at te beginning the knapsack is empty
    w := 0
    for i:=0 ; i<n ; i++ {			// we check each item
        if w+kp.Weight(i) <= c {		// if it fits into the knapsack
            w += kp.Weight(i)			// we take the item
            x[i] = 1
            z += kp.Profit(i)
        }
    }
    return x,z
}

// Dual greedy heuristic for the knapsack problem 
//
// We start with an infeasible solution (the knapsack contains all items) and
// remove items, starting with the item having the least profit/weight value,
// until the solution becomes feasible.
//
// Precondition: Profit[i]/Weight[i] >= Profit[i-1]/Weight[i-1] for i=1,...,n-1
//
// We do not check the precondition here!
// DualGreedy() computes a feasible solution even if the precondition is not valid,
// but this solution is usually considerably worse.
func DualGreedy(kp KnapsackProblem) ([]int,int) {
    n := kp.N()
    c := kp.Capacity()
    psum := 0
    wsum := 0
    x := make([]int,n)
    for i:=0 ; i<n ; i++ {	// put each item into the knapsack
        psum += kp.Profit(i)
        wsum += kp.Weight(i)
	x[i] = 1
    }

    for i:=n-1 ; i>=0 && wsum > c ; i-- {		// while solution is infeasible
        psum -= kp.Profit(i)				// remove item i
	wsum -= kp.Weight(i)
        x[i] = 0
    }
    return x,psum
}
