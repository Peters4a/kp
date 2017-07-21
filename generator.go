package kp

import (
    "errors"
    "fmt"
    "math/rand"
    "sort"
    "time"
)


func Generate(gen KnapsackGeneratorT) (*KnapsackProblemT,error) {
    var (
	p []int		// profits
	w []int		// weights
        c int		// capacity
    )

    p = make([]int, gen.N)
    w = make([]int, gen.N)
    
    if gen.Seed != 0 {
        rand.Seed(gen.Seed)
    } else {
	rand.Seed(int64(time.Now().Nanosecond()))
    }

    for i:=0 ; i<gen.N ; i++ {		// weight are always uniformly distributed in
        w[i] = 1 + rand.Intn(gen.V)	// the interval[1,V]
    }

    if gen.CorrMode == "uncorrelated" {	// profit are also uniformly distributed in [1,V]
	for i:=0 ; i<gen.N ; i++ {
	    p[i] = 1 + rand.Intn(gen.V)
	}
    } else if gen.CorrMode == "weakly" {	// profit p[i] is uniformly distributed 
	for i:=0 ; i<gen.N ; i++ {		// in [w[i]-R,w[i]+R]
	    p[i] = w[i] + rand.Intn(2*gen.R+1) - gen.R
	    if p[i] <= 0 {			// if w[i] <= R the profit may become
	        p[i] += gen.R			// zero or negative: we have to prevent this
	    }
	}
    } else if gen.CorrMode == "strongly" {	// profit is w[i] + R
	for i:=0 ; i<gen.N ; i++ {
	    p[i] = w[i] + gen.R
	}
    } else {
        return nil, errors.New(fmt.Sprintf("unknown correlation mode: %s", gen.CorrMode))
    }
    

    if gen.CapMode == "halfwsum" {	// C = 0.5*\sum w[i]
        wsum := 0
	for i:=0 ; i<gen.N ; i++ {
	    wsum += w[i]
	}
	c = wsum/2
    } else if gen.CapMode == "doublev" { 	// C = 2*V
        c = 2*gen.V
    } else {
        return nil, errors.New(fmt.Sprintf("unknown capacity mode: %s", gen.CapMode))
    }

    perm := sortPerm(p, w)		// we sort the values according to decreasing p[i]/w[i]
    p = permute(p, perm)		// sortPerm() computes the permutation for sorting and
    w = permute(w, perm)		// we apply this permutation to p and w.

    kpp := &KnapsackProblemT{
        Profit   : p,
	Weight   : w,
	Capacity : c,
    }

    return kpp, nil
}

// Everything we need for sorting the items

type sortItem struct {
    index int		// original array index in p resp. w
    sprof float64	// p[i]/w[i] value
}

type bySpecProfit []sortItem

// Methods we need to implement Go's sorting interface
func (a bySpecProfit) Len()              int  { return len(a) }
func (a bySpecProfit) Swap(i int, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySpecProfit) Less(i int, j int) bool { return a[i].sprof > a[j].sprof }

// Compute the permutation that would sort the items according to
// decreasing specific profit (p[i]/w[i])
func sortPerm(p []int, w []int) []int {
    n := len(p)
    items := make([]sortItem, n)	// build and fill the array
    for i:=0 ; i<n ; i++ {
        items[i].index = i
	items[i].sprof = float64(p[i])/float64(w[i])
    }
    sort.Sort(bySpecProfit(items))	// sort the array
    perm := make([]int, n)		// the index values give as the permutation
    for i:=0 ; i<n ; i++ {
        perm[i] = items[i].index
    }
    return perm
}

// Apply a permutation perm to an array a
func permute(a []int, perm []int) []int {
    n := len(a)
    b := make([]int, n)
    for i:=0 ; i<n ; i++ {
        b[i] = a[perm[i]]
    }
    return b
}
