package main

import (
    "fmt"

    "git.solver4all.com/Peter/kp/kp"
)

func main() {
    p := kp.KnapsackData{ Dim: 7,
                          P: []int{ 70, 20, 39, 37, 7, 5, 10 },
		          W: []int{ 31, 10, 20, 19, 4, 3, 6 },
		          C: 50 }
    x,z := kp.BranchAndBound(p)
    fmt.Printf("Problem: P: %v, W: %v, C: %v\n", p.P, p.W, p.C)
    fmt.Printf("Optimal solution: %v\n", x)
    fmt.Printf("Overall profit: %v\n", z)
}
