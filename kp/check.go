package kp

import (
    "errors"
)

func CheckSortedItems(kp KnapsackProblem) error {
    n := kp.N()
    for i:=1 ; i<n ; i++ {
        if float64(kp.Profit(i))/float64(kp.Weight(i)) >
           float64(kp.Profit(i-1))/float64(kp.Weight(i-1)) {
            return errors.New("wrong input: items are not sorted according to decreasing profit/weight")
        }
    }
    return nil
}
