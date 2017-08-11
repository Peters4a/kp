package main

import (
    "fmt"
    "log"

    "git.solver4all.com/Peter/kp/kp"
)

func main() {
    g := kp.KnapsackGenData{ N: 50, V: 100, CorrMode: "weakly",
                             R: 20, CapMode: "halfwsum" }
    p,err := kp.Generate(g)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Problem:\nP: %v\nW: %v\nC: %v\n\n", p.P, p.W, p.C)

    x,z := kp.BranchAndBound(p)
    fmt.Printf("BaB solution: %v\n", x)
    fmt.Printf("BaB overall profit: %v\n", z)

    x,z = kp.BranchAndBoundHS(p)
    fmt.Printf("HS solution:  %v\n", x)
    fmt.Printf("HS overall profit:  %v\n", z)

    x,z = kp.DynProg(p)
    fmt.Printf("DP solution:  %v\n", x)
    fmt.Printf("DP overall profit:  %v\n", z)

    x,z = kp.Greedy(p)
    fmt.Printf("GRD solution: %v\n", x)
    fmt.Printf("GRD overall profit: %v\n", z)
}
