package main

import (
    "fmt"
    "log"

    "git.solver4all.com/Peter/kp/kp"
)

func main() {
    g := kp.KnapsackGenData{ N: 40, V: 200, CorrMode: "strongly",
                             R: 15, CapMode: "halfwsum" }
    p,err := kp.Generate(g)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Problem:\nP: %v\nW: %v\nC: %v\n\n", p.P, p.W, p.C)

    _,zz := kp.UpperBound(p)
    fmt.Printf("Upper Bound: %v\n", zz)

    x,z := kp.BranchAndBound(p)
    fmt.Printf("BaB solution: %v\n", x)
    fmt.Printf("BaB overall profit: %v\n", z)

    x,z = kp.Greedy(p)
    fmt.Printf("GRD solution: %v\n", x)
    fmt.Printf("GRD overall profit: %v\n", z)
}
