package kp

func makePolicyTable(n int, c int) [][]int {
    pt := make([][]int,n)
    for i:=0 ; i<n ; i++ {
        pt[i] = make([]int, c+1)
    }

    return pt
}

// Solve a knapsack problem with dynamic programming
func DynProg(kp *KnapsackProblemT) {
    var (
        v  []int			// value function for item i
	vv []int			// value function for item i+1
    )

    n := len(kp.Profit)			// n is the number of items we have
    kp.X = make([]int,n)		// X[i] = 0 for i=0,...,n-1

    policy := makePolicyTable(n, kp.Capacity)	// policy[i][s] stores the optimal decision
    						// for item i and rest capacity s
    vv = make([]int,kp.Capacity+1)		

    // Backward computation
    for i:=n-1 ; i>=0 ; i-- {			// for item=n-1,...,0
        v = make([]int, kp.Capacity+1)
        for s:=0 ; s<=kp.Capacity ; s++ {	// for rest capacity of s=0,...,Capacity
            v[s] = vv[s]			// not to select item i is always feasible
	   			// observe: the policy table represents this decision already

	    if s >= kp.Weight[i] {	// But if the rest capacity is large enough
	        if v[s] < kp.Profit[i] + vv[s-kp.Weight[i]] {	// we check wether selecting 
		    v[s] = kp.Profit[i] + vv[s-kp.Weight[i]]	// item i gives a better
		    policy[i][s] = 1				// solution
	        }
	    }
        }
        vv = v
    }

    // Forward computation
    kp.Obj = v[kp.Capacity]	// maximum 
    s := kp.Capacity		// go through the optimal decision starting with Capacity
    for i:=0 ; i<n ; i++ {
	kp.X[i] = policy[i][s]	// if the optimal decision is to select item i
	if policy[i][s] == 1 {	// we set X[i] to 1 and reduce the rest capacity by
	    s -= kp.Weight[i]	// weight[i]
	}
    }
}
