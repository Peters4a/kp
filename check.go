package kp

import (
    "errors"
)

func CheckEqualLength(kp *KnapsackProblemT) error {
    if len(kp.Profit) != len(kp.Weight) {
        return errors.New("wrong input: profit and weight arrays have unequal length")
    }
    return nil
}

func CheckSortedItems(kp *KnapsackProblemT) error {
    n := len(kp.Profit)
    for i:=1 ; i<n ; i++ {
        if float64(kp.Profit[i])/float64(kp.Weight[i]) >
           float64(kp.Profit[i-1])/float64(kp.Weight[i-1]) {
            return errors.New("wrong input: items are not sorted according to decreasing profit/weight")
        }
    }
    return nil
}
