package kp

// Knapsack problem data
type KnapsackProblemT struct {
    Profit   []int `json:"profit"`		// profit values
    Weight   []int `json:"weight"`		// weight values
    Capacity int   `json:"capacity"`		// capacity of the knapsack
    X        []int `json:"x,omitempty"`		// binary decision variables
    Obj      int   `json:"obj,omitempty"`	// objective function value
    Xf       []float64 `json:"xf,omitempty"`
    				// decision variables for solvers that may generate fractional
                                // values for the decision variables (e.g. LP relaxation)
}

// Input for knapack problem generator
type KnapsackGeneratorT struct {
    N        int	// number of items
    V	     int	// maximal profit, weight value
                        // weights are uniformly random in [1,V]
    CorrMode string	// profit, weights may be "uncorrelated", "weakly" correlated or
                        // "strongly" correlated
    R        int        // weakly: p_j uniformly random in [w_j-R,w_j+R]
                        // strongly: p_j = w_j + r
    CapMode  string	// "doublev"  : Capacity = 2 * V
                        // "halfwsum" : Capacity = 0.5 * \sum_{j=1}^N w_j
    Seed     int64	// optional seed value for the random generator
}
