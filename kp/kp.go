package kp

type KnapsackProblem interface {
    N()           int			// number of items, 0,...,n-1
    Profit(i int) int			// profit of item i
    Weight(i int) int			// weight of item i
    Capacity()    int			// capacity of the knapsack
}

// Knapsack problem data
type KnapsackData struct {
    Name    string `json:"name,omitempty"`	// problem name, optional
    Comment string `json:"comment,omitempty"`	// comment, optional
    Type    string `json:"type"`		// problem type, unused at the moment
    Dim     int    `json:"dimension"`		// problem size
    P       []int  `json:"profits"`		// profit values
    W       []int  `json:"weights"`		// weight values
    C       int    `json:"capacity"`		// capacity of the knapsack
    X       []int  `json:"x,omitempty"`	// binary decision variables
    Z       int    `json:"z,omitempty"`	// objective function value
    Xf      []float64 `json:"xf,omitempty"`
				// decision variables for solvers that may generate fractional
                                // values for the decision variables (e.g. LP relaxation)
}

func (kp KnapsackData) N() int {
    return kp.Dim
}

func (kp KnapsackData) Capacity() int {
    return kp.C
}

func (kp KnapsackData) Profit(i int) int {
    return kp.P[i]
}

func (kp KnapsackData) Weight(i int) int {
    return kp.W[i]
}

// Input for knapack problem generator
type KnapsackGenData struct {
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
