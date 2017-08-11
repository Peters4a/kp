package main

import (
    "fmt"

    "git.solver4all.com/Peter/kp/kp"
)

type item struct {
    p int
    w int
}

type kproblem struct {
    c int
    items []item
}

func (p kproblem) N()           int { return len(p.items) }
func (p kproblem) Profit(i int) int { return p.items[i].p }
func (p kproblem) Weight(i int) int { return p.items[i].w }
func (p kproblem) Capacity()    int { return p.c          }

func main() {
    p := kproblem{ c: 50,
                   items: []item{
		       item{ p: 70, w: 31 },
		       item{ p: 20, w: 10 },
		       item{ p: 39, w: 20 },
		       item{ p: 37, w: 19 },
		       item{ p:  7, w:  4 },
		       item{ p:  5, w:  3 },
		       item{ p: 10, w:  6 }, }, }

    x,z := kp.BranchAndBound(p)
    fmt.Printf("Optimal solution: %v\n", x)
    fmt.Printf("Overall profit: %v\n", z)
}
