package kp

// A state in the branch and bound process.
type stateT struct {
    decision int	// decision(=1 or =0)
    nitems   int	// number of items considered
    psum     int	// profit sum of this state
    capacity int	// residual capacity
    ubound   int	// maximal profit for residual capacity
    phi      int	// psum + ubound
    father   *stateT	// preceeding state
}

// Solve a knapsack problem by Branch and Bound.
// Here we use a best upper bound strategy which leads to an A*-algorithm.
// This is simply achieved by using a priority queue as agenda.
func BranchAndBound(kp KnapsackProblem) ([]int,int) {
    var (
        state1 *stateT
        state2 *stateT
    )

    n := kp.N()					// number of items
    agenda := []*stateT{ initialState(kp) }	// initial state of our agenda

    for {
        state := agenda[0]		// get the first element of the agenda (priority queue)
	if state.nitems == n {		// goal state: optimal solution found
	    return optSol(kp, state)	// store it in kp
					// and we are done.
	}
						// no goal state: nitems < n
	if state.capacity >= kp.Weight(state.nitems) {	// is X[item]=1 feasible? if yes:
	    state1 = successor1(kp,state)	// successor for X[item] = 1
	} else {
	    state1 = nil
	}
	state2 = successor0(kp,state)		// successor for X[item] = 0
	agenda = pqUpdate(agenda,state1,state2)		// update the agenda
    }
}

// Solve a knapsack problem by Branch and Bound.
// Here we use a depth first strategy which leads to an algorithm similiar to the
// classical branch and bound algorithm by Horowitz and Sahni.
// We simply achieve the depth first strategy by using a stack as agenda.
// The garbage collector should keep the used memory small, because the agenda
// contains only one path (with sibling nodes, the size of the agenda is bounded by 2n+1).
func BranchAndBoundHS(kp KnapsackProblem) ([]int,int) {
    var (
	stateB *stateT				// actual best solution
    )

    n := kp.N()					// number of items
    pmax := 0					// actual best solution value
    agenda := []*stateT{ initialState(kp) }	// initial state of our agenda

    for {
        if len(agenda) == 0 {			// if the agenda is empty we are done.
	    return optSol(kp, stateB)		// we store the best solution we found
	}
	state := agenda[len(agenda)-1]		// take the top of the stack
	agenda = agenda[0:len(agenda)-1]	// pop
	if state.nitems == n {			// popped state is a goal state
	    if state.psum > pmax {		// new best solution? if yes
		pmax = state.psum		// store new best solution value
	        stateB = state			// and pointer to this solution
	    }
	} else if state.phi > pmax {		// not a goal state but upper bound larger
	    agenda = append(agenda,successor0(kp,state))	// push for decision = 0
	    if state.capacity >= kp.Weight(state.nitems) {// if residual capacity is large enough
		agenda = append(agenda,successor1(kp,state))	// push for decision = 1
	    }
	}
    }
}

func initialState(kp KnapsackProblem) *stateT {
    state := &stateT{			// initial state
	decision : -1,			// no decision
        nitems   : 0,			// no item considered yet
	psum     : 0,			// no profit yet
	capacity : kp.Capacity(),	// knapsack is empty
	ubound   : uBound1P(kp, kp.Capacity(), 0),
	father   : nil,			// root of the search tree has no father
    }
    state.phi = state.psum + state.ubound

    return state
}

func successor0(kp KnapsackProblem, state *stateT) *stateT {
    state2 := &stateT{			// X[item]=0 is always feasible
	decision : 0,
	nitems   : state.nitems + 1,
	psum     : state.psum,		// psum and capacity remain unchanged
	capacity : state.capacity,	// but the upper bound may change
	ubound   : uBound1P(kp, state.capacity, state.nitems+1),
	father   : state,
    }
    state2.phi = state2.psum + state2.ubound

    return state2
}

func successor1(kp KnapsackProblem, state *stateT) *stateT {
    item := state.nitems
    state1 := &stateT{			// construct a state with X[item] = 1
	decision : 1,
	nitems   : state.nitems + 1,
	psum     : state.psum + kp.Profit(item),	// additional profit
	capacity : state.capacity - kp.Weight(item),	// less residual capacity
	ubound   : state.ubound - kp.Profit(item),	// trick: on X[item]=1
	father   : state,				// psum+ubound doesn't change
    }
    state1.phi = state1.psum + state1.ubound

    return state1
}

func optSol(kp KnapsackProblem, state *stateT) ([]int,int) {
    n := kp.N()
    x := make([]int, n)
    z := state.psum		// copy the decisions into decision vector x
    for ; state.nitems>0 ; state=state.father {
	x[state.nitems-1] = state.decision
    }
    return x,z
}

// Update the agenda, which is organized as a max-heap.
// s1 and s2 are the successor states of pq[0].
// s1 results from decision = 1 and may be nil (if decision = 1 is infeasible).
// s2 results from decision = 0 and is alwas != nil.
// The max-heap is a left fully binary tree organized in an array.
// The root (largest element) is at index 0.
// A state at index i has its left and right son at index 2i+1 resp. 2i+2.
func pqUpdate(pq []*stateT, s1 *stateT, s2 *stateT) []*stateT {
    var (
        s *stateT
    )

    if s1 != nil {		// s1 (decision = 1) maybe nil
        s = s1			// s2 (decision = 0) is always != nil
    } else {
        s = s2
    }
    pq[0] = s			// the first state != nil substitutes the agenda head
    reheapTop(pq)		// heap property maybe violated ==> reconstitute the heap property

    if s1 != nil {		// an eventually second state is appended
        pq = append(pq, s2)
	reheapBottom(pq)	// reconstitute the heap property
    }

    return pq
}

// Reconstitute the heap property for a new root.
// The root (largest element) is at index 0.
// A state at index i has its left and right son at index 2i+1 resp. 2i+2.
func reheapTop(pq []*stateT) {
    l := len(pq)
    i := 0			// at the root we start
    for {			// index 2i+1 is the left, index 2i+2 the right son of i
	if 2*i+1 >= l {		// no left son? ==> leaf: we are done.
	    return
	}
	if 2*i+2 >= l {		// only a left son
	    if pq[i].phi >= pq[2*i+1].phi {	// left son is not greater? done.
	        return
	    }
	    j := 2*i+1				// the left son is greater
	    pq[i], pq[j] = pq[j], pq[i]		// we swap and proceed
	    i = j
	} else {		// left and right son
	    if pq[i].phi >= pq[2*i+1].phi && pq[i].phi >= pq[2*i+2].phi {
	        return		// greater or equal than both sons? done.
	    }
	    j := 2*i+1		// at least one son is greater, maybe the left
	    if pq[2*i+2].phi > pq[2*i+1].phi {	// but eventually the right
	        j = 2*i+2
	    }
	    pq[i], pq[j] = pq[j], pq[i]		// we swap with the larger son
	    i = j				// and proceed
	}
    }
}

// Reconstitute the heap property for an appended element.
// The father of a node at index i is at index (i-1)/2.
// The root (largest element) is at index 0.
func reheapBottom(pq []*stateT) {
    l := len(pq)
    i := l-1			// at the last element we start
    for {
        if i==0 {		// We have reached the root? done.
	    return
	}
	ifath := (i-1)/2	// index of father
	if pq[ifath].phi >= pq[i].phi {	// son is not greater? done.
	    return
	}
	pq[i], pq[ifath] = pq[ifath], pq[i]	// son is greater: we swap
	i = ifath				// and proceed
    }
}
